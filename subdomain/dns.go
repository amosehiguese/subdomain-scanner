package subdomain

import (
	"log"
	"net"
)

type resolver struct{}

func NewResolver() *resolver {
	log.Println("Generating dns resolver...")
	return &resolver{}
}

// ResolveDNS checks to see if a subdomain resolves according to the Domain Naming System
func (r *resolver) ResolveDNS(host string) (*Subdomain, error) {
	_, err := net.LookupHost(host)
	if err != nil {
		log.Println("Got an error while trying to resolve DNS", err)
		return nil, err
	}

	return &Subdomain{}, nil
}
