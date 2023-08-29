package scanners

import (
	"log"
	"sync"
	"time"

	"github.com/amosehiguese/subdscanner/subdomain"
)

// Scan scans subdomains concurrently using various scanners.
func Scan(url string) (*[]subdomain.Subdomain, error) {
	var mu sync.Mutex
	var result []subdomain.Subdomain

	var wg sync.WaitGroup
	startTime := time.Now()
	scanners := GetAllScanners()

	for _, scanner := range scanners {
		wg.Add(1)
		go func(sc SubdomainScanner) {
			defer wg.Done()

			subdomains, err := sc.GetSubdomains(url)
			if err != nil {
				log.Printf("using %v, unable to get subdomain err ->%v", sc.GetName(), err)
				return
			}

			// remove duplicates
			subdomainMap := make(map[string]struct{})
			for _, s := range subdomains {
				subdomainMap[s] = struct{}{}
			}

			var subDomainSet []string
			for sMapKey := range subdomainMap {
				subDomainSet = append(subDomainSet, sMapKey)
			}

			// Filter domain that do not resolve
			var resolvedSubdomains []string
			for _, addr := range subDomainSet {
				if resolveDNS(addr) {
					resolvedSubdomains = append(resolvedSubdomains, addr)
				}
			}

			// Start port scanning
			for _, subdom := range resolvedSubdomains {
				sd := subdomain.ScanPorts(subdom)
				mu.Lock()
				result = append(result, sd)
				mu.Unlock()
			}
			log.Println("step done")
		}(scanner)
	}

	wg.Wait()
	endTime := time.Now()
	totalScanTime := endTime.Sub(startTime)

	log.Println("Scan completed in ->", totalScanTime)

	return &result, nil
}

func resolveDNS(domain string) bool {
	rs := subdomain.NewResolver()
	_, err := rs.ResolveDNS(domain)
	return err == nil
}
