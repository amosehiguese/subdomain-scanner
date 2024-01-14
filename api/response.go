package api

import "github.com/amosehiguese/subdscanner/subdomain"

type Subdomains struct {
	Subdomains []subdomain.Subdomain
}

type errorResp struct {
	Success   bool `json:"success"`
	ErrorCode int  `json:"errCode"`
	Message   any  `json:"message"`
}

func NewError(errCode int, msg any) *errorResp {
	return &errorResp{
		Success:   false,
		ErrorCode: errCode,
		Message:   msg,
	}
}
