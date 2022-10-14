package handlers

import (
	"html/template"
	"my-backend/templates"
	"net/http"
)

var IndexHtml string

func Home(w http.ResponseWriter, req *http.Request) {
	temps := templates.GetTemp()
	templateFile := template.Must(template.New("index.html").Parse(temps["index"]))
	templateFile.ExecuteTemplate(w, "index.html", nil)
}
