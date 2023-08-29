package subdomain

type Subdomain struct {
	Domain    string `json:"domain"`
	OpenPorts []Port `json:"open_ports"`
}
