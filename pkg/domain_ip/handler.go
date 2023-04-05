package domain_ip

import (
	"github.com/dnagikh/website/pkg/utils"
	"html/template"
	"net/http"
)

type ViewData struct {
	Domain string
	Trace  []string
}

func Index(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/domain_ip" {
		utils.ErrorHandler(w, r, http.StatusNotFound)
		return
	}

	data := ViewData{}
	err := r.ParseForm()
	if r.Method == http.MethodPost && r.Body != nil {
		data.Domain = r.Form.Get("domain")
		data.Trace, err = DomainIP(data.Domain)
		if err != nil {
			data.Trace = append(data.Trace, err.Error())
		}
	}

	files := []string{
		"./ui/templates/layout.html",
		"./ui/templates/domain_ip/index.html",
	}

	templates := template.Must(template.ParseFiles(files...))
	err = templates.ExecuteTemplate(w, "layout", data)
	if err != nil {
		utils.ErrorHandler(w, r, http.StatusInternalServerError)
		return
	}
}
