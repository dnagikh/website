package dns

import (
	"github.com/dnagikh/website/pkg/utils"
	"html/template"
	"net/http"
)

type ViewData struct {
	Domain     string
	Type       string
	Trace      []string
	Types      *[]string
	DnsServers *map[string]string
	Dns        string
}

func Index(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/dns" {
		utils.ErrorHandler(w, r, http.StatusNotFound)
		return
	}

	data := ViewData{}
	data.Types = &[]string{
		"A",
		"NS",
		"CNAME",
		"MX",
		"TXT",
	}
	data.DnsServers = &map[string]string{
		"#1 Google (USA)":    "8.8.8.8",
		"#2 Yandex (Russia)": "77.88.8.8",
		"#3 114DNS (China)":  "114.114.114.114",
	}

	err := r.ParseForm()
	if r.Method == http.MethodPost && r.Body != nil {
		data.Domain = r.Form.Get("domain")
		data.Type = r.Form.Get("type")
		data.Dns = r.Form.Get("dns")
		data.Trace, err = checkDNS(data.Domain, data.Type, data.Dns)
		if err != nil {
			data.Trace = append(data.Trace, err.Error())
		}
	}

	files := []string{
		"./ui/templates/layout.html",
		"./ui/templates/dns/index.html",
	}

	templates := template.Must(template.ParseFiles(files...))
	err = templates.ExecuteTemplate(w, "layout", data)
	if err != nil {
		utils.ErrorHandler(w, r, http.StatusInternalServerError)
		return
	}
}
