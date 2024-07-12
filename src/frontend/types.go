package main

type Port struct {
	ConnOpen bool
	OpenPort uint32
}

type Subdomain struct {
	Domain string `json:"domain"`
	Ports  any    `json:"open_ports"`
}
