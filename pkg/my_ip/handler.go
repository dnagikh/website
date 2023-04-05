package my_ip

import (
	"github.com/dnagikh/website/pkg/utils"
	"html/template"
	"net/http"
)

type ViewData struct {
	Ip string
}

func Index(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/my_ip" {
		utils.ErrorHandler(w, r, http.StatusNotFound)
		return
	}

	files := []string{
		"./ui/templates/layout.html",
		"./ui/templates/my_ip/index.html",
	}

	templates := template.Must(template.ParseFiles(files...))
	err := templates.ExecuteTemplate(w, "layout", ViewData{
		Ip: ReadUserIP(r),
	})
	if err != nil {
		utils.ErrorHandler(w, r, http.StatusInternalServerError)
		return
	}
}
