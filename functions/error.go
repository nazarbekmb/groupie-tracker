package functions

import (
	"html/template"
	"log"
	"net/http"
)

func ErrorHandler(w http.ResponseWriter, status int) {
	tmpl, err := template.ParseFiles("ui/templates/error.html")
	if err != nil {
		http.Error(w, http.StatusText(status), status)
		return
	}
	w.WriteHeader(status)
	ErrorMessage := ErrorStatus{status, http.StatusText(status)}
	err = tmpl.Execute(w, ErrorMessage)
	if err != nil {
		log.Println(err)
		http.Error(w, http.StatusText(status), status)
		return
	}
}

type ErrorStatus struct {
	Status  int
	Message string
}
