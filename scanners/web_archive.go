package scanners

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"sync"
	"time"
)

type webArchive struct {
	mu   sync.Mutex
	Name string
}

func NewWebArchive() *webArchive {
	return &webArchive{
		Name: "Web.archive.org scan",
	}
}

func (wa *webArchive) GetName() string {
	return wa.Name
}

func (wa *webArchive) GetSubdomains(target string) ([]string, error) {
	log.Printf("%v scanning", wa.GetName())

	urlT := fmt.Sprintf("https://web.archive.org/cdx/search/cdx?matchType=domain&fl=original&output=json&collapse=urlkey&url=%s", target)
	client := &http.Client{}

	ctx, cancel := context.WithTimeout(context.Background(), 10 * time.Second)
	defer cancel()
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, urlT, nil)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	resp, err := client.Do(req)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status: got %v", resp.Status)
	}

	var waResp [][]string
	if err := json.NewDecoder(resp.Body).Decode(&waResp); err != nil {
		return nil, err
	}

	var subdomain []string
	var wg sync.WaitGroup
	for _, urlSlc := range waResp {
		wg.Add(1)
		go func(urlSlc []string) {
			defer wg.Done()

			for _, urlStr := range urlSlc {
				u, err := url.Parse(urlStr)
				if err != nil {
					log.Printf("Got an error while parsing url: %s", err)
					continue
				}

				host := u.Hostname()
				if host != "" {
					wa.mu.Lock()
					subdomain = append(subdomain, host)
					wa.mu.Unlock()
				}
			}
		}(urlSlc)
	}

	wg.Wait()
	return subdomain, nil
}
