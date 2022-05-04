package ddns_test

import (
	"encoding/base64"
	"fmt"
	"net"
	"net/http"
	"os"
	"testing"

	"github.com/gruntwork-io/terratest/modules/terraform"
)

func TestDDNS(t *testing.T) {
	domain, ok := os.LookupEnv("DDNS_DOMAIN")
	if !ok {
		t.Fatalf("expected DDNS_DOMAIN env var to be set")
	}
	subdomain, ok := os.LookupEnv("DDNS_SUBDOMAIN")
	if !ok {
		t.Fatalf("expected DDNS_SUBDOMAIN env var to be set")
	}
	opts := &terraform.Options{
		TerraformDir: "../_examples",
		Vars: map[string]interface{}{
			"domain":    domain,
			"subdomain": subdomain,
		},
	}
	defer terraform.Destroy(t, opts)
	terraform.InitAndApply(t, opts)

	fqdn := fmt.Sprintf("%s.%s", subdomain, domain)

	checkHostnameIP(t, fqdn, "0.0.0.0")

	apiGatewayURL := terraform.OutputRequired(t, opts, "url")
	username := terraform.OutputRequired(t, opts, "username")
	password := terraform.OutputRequired(t, opts, "password")

	fmt.Println(username, password, apiGatewayURL)

	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("%s/nic/update?myip=1.1.1.1", apiGatewayURL), nil)
	if err != nil {
		t.Fatalf("could not create the request: %v", err)
	}
	req.Header.Set("Authorization", fmt.Sprintf("Basic %s", encode(username, password)))
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		t.Fatalf("could not execute the request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		t.Fatalf("expected status code of %d, got %d", http.StatusOK, resp.StatusCode)
	}
	checkHostnameIP(t, fqdn, "1.1.1.1")
}

func checkHostnameIP(t *testing.T, fqdn string, ip string) {
	ips, err := net.LookupIP(fqdn)
	if err != nil {
		t.Fatalf("expected IP lookup not to be nil got: %v", err)
	}

	ipsString := Map(ips, func(ip net.IP) string { return ip.String() })

	if !Includes(ipsString, ip) {
		t.Fatalf("expected DNS lookup %v to include %v value", ipsString, ip)
	}
}

func encode(username, password string) string {
	pass := fmt.Sprintf("%s:%s", username, password)
	return base64.StdEncoding.EncodeToString([]byte(pass))
}
