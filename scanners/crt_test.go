package scanners

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCrtSubdomainScanner(t *testing.T) {
	crt := NewCrt()
	target := "vwrm.com"

	subdomains, err := crt.GetSubdomains(target)
	require.NoError(t, err)

	if len(subdomains) == 0 {
		t.Fatalf("No subdomains found")
	}

	t.Logf("Found %d subdomains:", len(subdomains))
	for _, subdomain := range subdomains {
		t.Logf("-> %s", subdomain)
	}
}
