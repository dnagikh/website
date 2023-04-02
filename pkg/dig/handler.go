package dig

import (
	"github.com/dnagikh/website/pkg/utils"
	"html/template"
	"net/http"
	"strconv"
)

type ViewData struct {
	Domain string
	Type   uint16
	Trace  string
	Types  *map[uint16]string
}

func Index(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/dig" {
		utils.ErrorHandler(w, r, http.StatusNotFound)
		return
	}

	data := ViewData{}
	data.Types = &map[uint16]string{
		1:  "A",
		2:  "NS",
		5:  "CNAME",
		15: "MX",
		16: "TXT",
	}

	err := r.ParseForm()
	if err == nil {
		dnsType, err := strconv.ParseUint(r.Form.Get("type"), 10, 16)
		if err != nil {
			data.Trace = "Choose the correct type"
		}

		data.Domain = r.Form.Get("domain")
		data.Type = uint16(dnsType)
		data.Trace = GetTrace(r, &data)
	}

	files := []string{
		"./ui/templates/layout.html",
		"./ui/templates/dig/index.html",
	}

	templates := template.Must(template.ParseFiles(files...))
	err = templates.ExecuteTemplate(w, "layout", data)
	if err != nil {
		utils.ErrorHandler(w, r, http.StatusInternalServerError)
		return
	}
}
