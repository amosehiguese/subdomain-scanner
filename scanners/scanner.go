package scanners

type Scanner interface {
	GetName() string
}

type SubdomainScanner interface {
	Scanner
	GetSubdomains(url string) ([]string, error)
}

func GetAllScanners() []SubdomainScanner {
	return []SubdomainScanner{
		NewCrt(),
		NewWebArchive(),
	}
}
