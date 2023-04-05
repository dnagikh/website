package ports

import (
	"github.com/dnagikh/website/pkg/utils"
	"html/template"
	"net/http"
	"strings"
)

type ViewData struct {
	Domain   string
	Ports    string
	PortList []string
	Trace    map[string]string
}

func Index(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/ports" {
		utils.ErrorHandler(w, r, http.StatusNotFound)
		return
	}

	data := ViewData{}
	err := r.ParseForm()
	if r.Method == http.MethodPost && r.Body != nil {
		data.Domain = r.Form.Get("domain")
		data.Ports = r.Form.Get("ports")
		correctPorts, err := ValidatePortList(data.Ports)
		if !correctPorts {
			data.Trace = map[string]string{
				"*": err.Error(),
			}
		} else {
			data.PortList = strings.Split(data.Ports, ",")
			data.Trace, err = CheckPorts(data.Domain, data.PortList)
			if err != nil {
				data.Trace = map[string]string{
					"*": err.Error(),
				}
			}
		}
	} else {
		data.Ports = "80,443,22"
	}

	files := []string{
		"./ui/templates/layout.html",
		"./ui/templates/ports/index.html",
	}

	templates := template.Must(template.ParseFiles(files...))
	err = templates.ExecuteTemplate(w, "layout", data)
	if err != nil {
		utils.ErrorHandler(w, r, http.StatusInternalServerError)
		return
	}
}
