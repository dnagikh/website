package myip

import (
	"html/template"
	"log"
	"net/http"
)

type ViewData struct {
	Ip string
}

func Index(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/my_ip" {
		http.NotFound(w, r)
		return
	}

	files := []string{
		"./ui/templates/layout.html",
		"./ui/templates/myip/index.html",
	}

	templates := template.Must(template.ParseFiles(files...))
	err := templates.ExecuteTemplate(w, "layout", ViewData{
		Ip: ReadUserIP(r),
	})
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Internal Server Error", 500)
		return
	}
}
