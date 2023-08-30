package scanners

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
)

type crtSh struct {
	Name string
}

func NewCrt() *crtSh {
	return &crtSh{
		Name: "crtsh subdomains scan",
	}
}

type CrtResp struct {
	NameValue string `json:"name_value"`
}

func (c *crtSh) GetName() string {
	return c.Name
}

func (c *crtSh) GetSubdomains(target string) ([]string, error) {
	log.Printf("%v scanning", c.GetName())
	url := fmt.Sprintf("https://crt.sh/?q=%%25.%s&output=json", target)

	resp, err := http.Get(url)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status: got %v", resp.Status)
	}

	var newCrtResp []CrtResp
	if err := json.NewDecoder(resp.Body).Decode(&newCrtResp); err != nil {
		return nil, err
	}

	var subdomains []string
	for _, resp := range newCrtResp {
		nameValue := strings.TrimSpace(resp.NameValue)
		subdomains = append(subdomains, nameValue)
	}

	return subdomains, nil

}
