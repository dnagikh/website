package pkg

import (
	"github.com/dnagikh/website/pkg/utils"
	"html/template"
	"net/http"
)

func Index(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		utils.ErrorHandler(w, r, http.StatusNotFound)
		return
	}

	files := []string{
		"./ui/templates/layout.html",
		"./ui/templates/index.html",
	}

	templates := template.Must(template.ParseFiles(files...))
	err := templates.ExecuteTemplate(w, "layout", nil)
	if err != nil {
		utils.ErrorHandler(w, r, http.StatusInternalServerError)
		return
	}
}
