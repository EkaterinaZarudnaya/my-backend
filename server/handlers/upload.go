package handlers

import (
	"fmt"
	"html/template"
	"net/http"
)

func Upload(w http.ResponseWriter, req *http.Request) {

	templateFile := template.Must(template.ParseFiles("templates/index.html"))

	if req.Method == http.MethodPost {
		handleUpload(w, req)
		return
	}

	templateFile.ExecuteTemplate(w, "index.html", nil)
}

func handleUpload(w http.ResponseWriter, req *http.Request) {

	req.ParseMultipartForm(10 << 20)

	file, fileHeader, err := req.FormFile("file")
	if err != nil {
		fmt.Printf("File retrieval error: %v\n", err)
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}
	fmt.Println("Retrieving the file successfully.")
	defer file.Close()

	FileHandler(fileHeader.Filename, file, w)

	http.Redirect(w, req, "/?success=true", http.StatusSeeOther)
}
