package dig

import (
	"fmt"
	"github.com/dnagikh/website/pkg/utils"
	"github.com/lixiangzhong/dnsutil"
	"net/http"
)

func GetTrace(r *http.Request, data *ViewData) string {
	if len(data.Domain) <= 0 {
		return ""
	}

	if !utils.IsValidDomain(data.Domain) {
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
