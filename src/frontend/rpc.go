package main

import (
	"context"
	"io"
	"net/http"
	"time"

	pb "github.com/amosehiguese/subdomain-scanner/src/frontend/genproto/subdomain_v1"
	"go.uber.org/zap"
)

func (fe *frontendServer) getSubdomainsByApiQuery(ctx context.Context, target string) ([]string, error) {
	response, err := pb.NewApiQueryServiceClient(fe.apiQuerySvcConn).GetSubdomainsByApiQuery(ctx, &pb.ApiQueryRequest{Target: target})
	if err != nil {
		return nil, err
	}
	return response.Subdomains, nil
}

func (fe *frontendServer) getSubdomainsByBruteForce(ctx context.Context, target string) ([]string, error) {
	response, err := pb.NewBruteServiceClient(fe.bruteForceSvcConn).GetSubdomainsByBruteForce(ctx, &pb.BruteForceRequest{Target: target})
	if err != nil {
		return nil, err
	}

	return response.Subdomains, nil
}

func (fe *frontendServer) resolveDNS(ctx context.Context, hosts []string) ([]string, error) {
	response, err := pb.NewResolveDnsServiceClient(fe.bruteForceSvcConn).ResolveDns(ctx, &pb.ResolveDnsRequest{Hosts: hosts})
	if err != nil {
		return nil, err
	}

	var result []string
	for {
		resolveDnsResponse, err := response.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, err
		}

		result = append(result, resolveDnsResponse.Subdomain)
	}
	return result, nil
}

func (fe *frontendServer) portScan(ctx context.Context, hosts []string) ([]Subdomain, error) {
	stream, err := pb.NewPortScanServiceClient(fe.portScanSvcConn).ScanForOpenPorts(ctx)
	if err != nil {
		return nil, err
	}

	waitc := make(chan struct{})
	result := make(chan *pb.Subdomain)

	go func() {
		for {
			in, err := stream.Recv()
			if err == io.EOF {
				close(waitc)
				return
			}
			if err != nil {
				return
			}
			result <- in
		}
	}()

	for _, host := range hosts {
		if err := stream.Send(&pb.PortScanRequest{Host: host}); err != nil {
			return nil, err
		}
	}

	var subdomains []Subdomain
	for subdomain := range result {
		subd := Subdomain{
			Domain: subdomain.Domain,
			Ports:  subdomain.Ports,
		}
		subdomains = append(subdomains, subd)
	}

	return subdomains, nil
}

func (fe *frontendServer) scan(r *http.Request, domain string) ([]Subdomain, error) {
	zapLog := r.Context().Value(ctxKeyLog{}).(*zap.Logger)
	var result []string

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*20)
	defer cancel()
	subdomains, err := fe.getSubdomainsByApiQuery(ctx, domain)
	if err != nil {
		zapLog.With(
			zap.Field(
				zap.Error(err),
			),
		).Error("Failed to get subdomain by API queries.")
		return nil, err
	}

	result = append(result, subdomains...)

	subdomains, err = fe.getSubdomainsByBruteForce(ctx, domain)
	if err != nil {
		zapLog.With(
			zap.Field(
				zap.Error(err),
			),
		).Error("Failed to get subdomain by brute force.")
	}

	result = append(result, subdomains...)

	result, err = fe.resolveDNS(ctx, result)
	if err != nil {
		zapLog.With(
			zap.Field(
				zap.Error(err),
			),
		).Error("Failed to resolve dns.")

		return nil, err
	}

	var subds []Subdomain
	subds, err = fe.portScan(ctx, result)
	if err != nil {
		zapLog.With(
			zap.Field(
				zap.Error(err),
			),
		).Error("Port scanning failed.")

		return nil, err
	}

	return subds, nil
}
