package utils

import (
	"fmt"
	"html/template"
	"net/http"
	"regexp"
)

func ErrorHandler(w http.ResponseWriter, r *http.Request, status int) {
	w.WriteHeader(status)
	files := []string{
		"./ui/templates/layout.html",
		fmt.Sprintf("./ui/templates/errors/%d.html", status),
	}
	templates := template.Must(template.ParseFiles(files...))
	templates.ExecuteTemplate(w, "layout", nil)
}

func IsValidDomain(domain string) bool {
	return regexp.MustCompile(`^(?i)[a-z0-9-]+(\.[a-z0-9-]+)+\.?$`).MatchString(domain)
}
