package handlers

import (
	"html/template"
	"net/http"
)

var IndexHtml string

func Home(w http.ResponseWriter, req *http.Request) {
	templateFile := template.Must(template.New("index.html").Parse(IndexHtml))
	templateFile.ExecuteTemplate(w, "index.html", nil)
}
