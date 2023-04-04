package utils

import (
	"fmt"
	"html/template"
	"net/http"
	"regexp"
	"strings"
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
	re := regexp.MustCompile(`^[a-zA-Z0-9._-]+\.[a-zA-Z]{2,}$`)
	return re.MatchString(domain)
}

func CleanDomain(domain string) (string, error) {
	if strings.HasPrefix(domain, "http://") {
		domain = strings.Replace(domain, "http://", "", 1)
	} else if strings.HasPrefix(domain, "https://") {
		domain = strings.Replace(domain, "https://", "", 1)
	}

	if strings.Contains(domain, "?") {
		domain = domain[:strings.Index(domain, "?")]
	}

	if strings.Contains(domain, "/") {
		domain = domain[:strings.Index(domain, "/")]
	}

	if !IsValidDomain(domain) {
		return "", fmt.Errorf("bad domain: %q", domain)
	}

	return domain, nil
}
