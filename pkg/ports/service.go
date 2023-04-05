package ports

import (
	"fmt"
	"github.com/dnagikh/website/pkg/utils"
	"net"
	"strconv"
	"strings"
)

func CheckPorts(domain string, ports []string) (map[string]string, error) {
	if len(ports) == 0 {
		return nil, fmt.Errorf("specify ports to check domain %q", domain)
	}

	if !utils.IsValidDomain(domain) {
		return nil, fmt.Errorf("invalid domain name %q", domain)
	}

	list := make(map[string]string)
	for _, port := range ports {
		conn, err := net.Dial("tcp", domain+":"+port)
		if err != nil {
			list[port] = "closed"
			continue
		}
		conn.Close()
		list[port] = "open"
	}

	return list, nil
}

func ValidatePortList(portList string) (bool, error) {
	for _, ch := range portList {
		if ch != ',' && (ch < '0' || ch > '9') {
			return false, fmt.Errorf("invalid characters in port list")
		}
	}

	portStrings := strings.Split(portList, ",")
	for _, portString := range portStrings {
		port, err := strconv.Atoi(portString)
		if err != nil {
			return false, err
		}
		if port < 0 || port > 65535 {
			return false, fmt.Errorf("port %d is out of range (0-65535)", port)
		}
	}

	return true, nil
}
