package dig

import (
	"fmt"
	"github.com/lixiangzhong/dnsutil"
	"net/http"
	"regexp"
)

func GetTrace(r *http.Request, data *ViewData) string {
	if len(data.Domain) <= 0 {
		return ""
	}

	if !isValidDomain(data.Domain) {
		return "Not valid URL"
	}

	var dig dnsutil.Dig
	msg, err := dig.GetMsg(data.Type, data.Domain)
	if err != nil {
		fmt.Println(err)
		return err.Error()
	}

	return msg.String()
}

func isValidDomain(domain string) bool {
	return regexp.MustCompile(`^(?i)[a-z0-9-]+(\.[a-z0-9-]+)+\.?$`).MatchString(domain)
}
