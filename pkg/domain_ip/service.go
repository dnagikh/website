package domain_ip

import (
	"fmt"
	"net"
)

func DomainIP(domain string) ([]string, error) {
	ips, err := net.LookupIP(domain)
	if err != nil {
		return []string{}, fmt.Errorf("failed to lookup IP for domain %q", domain)
	}

	var list []string
	for _, ip := range ips {
		list = append(list, ip.String())
	}

	if len(list) == 0 {
		return []string{}, fmt.Errorf("IP address not found for domain %q", domain)
	}

	return list, nil
}
