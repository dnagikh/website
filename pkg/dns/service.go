package dns

import (
	"context"
	"fmt"
	"github.com/dnagikh/website/pkg/utils"
	"net"
)

func checkDNS(domain string, recordType string, dnsServer string) ([]string, error) {
	if !utils.IsValidDomain(domain) {
		return nil, fmt.Errorf("bad domain name %q", domain)
	}

	resolver := &net.Resolver{
		PreferGo: true,
		Dial: func(ctx context.Context, network, address string) (net.Conn, error) {
			d := net.Dialer{}
			return d.DialContext(ctx, "udp", dnsServer+":53")
		},
	}

	var records []string
	switch recordType {
	case "A":
		ips, err := resolver.LookupIPAddr(context.Background(), domain)
		if err != nil {
			return nil, fmt.Errorf("A records not found for %q", domain)
		}
		for _, ip := range ips {
			records = append(records, ip.String())
		}
	case "TXT":
		txts, err := resolver.LookupTXT(context.Background(), domain)
		if err != nil {
			return nil, fmt.Errorf("TXT not found for %q", domain)
		}
		for _, txt := range txts {
			records = append(records, txt)
		}
	case "NS":
		nss, err := resolver.LookupNS(context.Background(), domain)
		if err != nil {
			return nil, fmt.Errorf("NS not found for %q", domain)
		}
		for _, ns := range nss {
			records = append(records, ns.Host)
		}
	case "CNAME":
		cname, err := resolver.LookupCNAME(context.Background(), domain)
		if err != nil {
			return nil, fmt.Errorf("CNAME not found for %q", domain)
		}
		records = append(records, cname)
	case "MX":
		mxs, err := resolver.LookupMX(context.Background(), domain)
		if err != nil {
			return nil, fmt.Errorf("MX not nothing found for %q", domain)
		}
		for _, mx := range mxs {
			records = append(records, fmt.Sprintf("%v %v", mx.Pref, mx.Host))
		}
	default:
		return nil, fmt.Errorf("unsupported record type %q", recordType)
	}

	if len(records) == 0 {
		return nil, fmt.Errorf("no records found %q", domain)
	}

	return records, nil
}
