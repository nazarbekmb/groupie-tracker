package main

import (
	"groupie-tracker/functions"
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", functions.Index)
	mux.HandleFunc("/info", functions.Artists)
	mux.Handle("/ui/", http.StripPrefix("/ui/", http.FileServer(http.Dir("ui")))) // каждый раз когда идёт обращение нач-ся на "ui" ui убирается по оставшемуся пути ищется соответствующий файл.
	log.Println("Server is listening...http://127.0.0.1:7777/")
	err := http.ListenAndServe(":7777", mux)
	log.Println(err)
}
