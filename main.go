package main

import (
	"fmt"
	"net/http"
	u "pdf-creator-example/pdfGenerator"

	"github.com/gorilla/mux"
)

func GeneratePDF(w http.ResponseWriter, r *http.Request) {
	rpdf := u.NewRequestPdf("")

	//html template path
	templatePath := "templates/correios.html"

	//path for download pdf
	outputPath := "storage/example.pdf"

	//html template data
	templateData := struct {
		Title        string
		Description  string
		TrackingCode string
		Address      string
		Contact      string
	}{
		Title:        "Correios label example",
		Description:  "Corra imprimir a etiqueta.",
		TrackingCode: "ASD123IJA231",
		Address:      "Higino Aguiar, 114, SC",
		Contact:      "Leoni",
	}

	if err := rpdf.ParseTemplate(templatePath, templateData); err == nil {
		pdfbytes, _ := rpdf.GeneratePDF(outputPath)
		w.Header().Set("Content-type", "application/pdf")
		w.Write(pdfbytes)
	} else {
		fmt.Println(err)
	}
}

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/", GeneratePDF)

	fmt.Println("Listening server at port 8080")
	http.ListenAndServe(":8080", r)
}
