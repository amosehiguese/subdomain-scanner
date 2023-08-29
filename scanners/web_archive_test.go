package scanners

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestWebArchiveSubdomainScanner(t *testing.T) {
	wa := NewWebArchive()

	target := "vwrm.com" // Replace with the target domain you want to test

	subdomains, err := wa.GetSubdomains(target)
	require.NoError(t, err)

	if len(subdomains) == 0 {
		t.Fatalf("No subdomains found")
	}

	t.Logf("Found %d subdomains:", len(subdomains))
	for _, subdomain := range subdomains {
		t.Logf("-> %s", subdomain)
	}
}
