package dkim

import (
	"github.com/dnagikh/website/pkg/utils"
	"html/template"
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
		utils.ErrorHandler(w, r, http.StatusNotFound)
		return
	}

	files := []string{
		"./ui/templates/layout.html",
		"./ui/templates/dkim/index.html",
	}

	templates := template.Must(template.ParseFiles(files...))
	err := templates.ExecuteTemplate(w, "layout", nil)
	if err != nil {
		utils.ErrorHandler(w, r, http.StatusInternalServerError)
		return
	}
}

func Result(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/dkim/result" {
		utils.ErrorHandler(w, r, http.StatusNotFound)
		return
	}

	err := r.ParseForm()
	if err != nil {
		utils.ErrorHandler(w, r, http.StatusBadRequest)
		return
	}

	length, err := strconv.Atoi(r.Form.Get("length"))
	if err != nil || length <= 0 {
		utils.ErrorHandler(w, r, http.StatusBadRequest)
		return
	}

	data := &ViewData{
		Domain:   r.Form.Get("domain"),
		Selector: r.Form.Get("selector"),
		Length:   length,
	}

	data.Private, data.Public, err = GenerateRSA(data.Domain, data.Selector, data.Length)
	if err != nil {
		utils.ErrorHandler(w, r, http.StatusInternalServerError)
		return
	}

	files := []string{
		"./ui/templates/layout.html",
		"./ui/templates/dkim/result.html",
	}

	templates := template.Must(template.ParseFiles(files...))
	err = templates.ExecuteTemplate(w, "layout", *data)
	if err != nil {
		utils.ErrorHandler(w, r, http.StatusInternalServerError)
		return
	}
}
