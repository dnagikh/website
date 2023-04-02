package dkim

import (
	"html/template"
	"log"
	"net/http"
	"strconv"
)

type ViewData struct {
	Domain   string
	Selector string
	Length   int

	Public  string
	Private string
}

func Index(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/dkim" {
		http.NotFound(w, r)
		return
	}

	files := []string{
		"./ui/templates/layout.html",
		"./ui/templates/dkim/index.html",
	}

	templates := template.Must(template.ParseFiles(files...))
	err := templates.ExecuteTemplate(w, "layout", nil)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Internal Server Error", 500)
		return
	}
}

func Result(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/dkim/result" {
		http.NotFound(w, r)
		return
	}

	err := r.ParseForm()
	if err != nil {
		log.Println(err.Error())
		return
	}

	length, err := strconv.Atoi(r.Form.Get("length"))
	if err != nil || length <= 0 {
		log.Println("Invalid key length")
		return
	}

	data := &ViewData{
		Domain:   r.Form.Get("domain"),
		Selector: r.Form.Get("selector"),
		Length:   length,
	}

	data.Private, data.Public, err = GenerateRSA(data.Domain, data.Selector, data.Length)
	if err != nil {
		log.Println(err.Error())
		return
	}

	files := []string{
		"./ui/templates/layout.html",
		"./ui/templates/dkim/result.html",
	}

	templates := template.Must(template.ParseFiles(files...))
	err = templates.ExecuteTemplate(w, "layout", *data)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Internal Server Error", 500)
		return
	}
}
