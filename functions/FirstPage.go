package functions

import (
	"encoding/json"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
)

func Index(w http.ResponseWriter, r *http.Request) { // с помощью  "w" обращаемся к конкретной страничке, "r" - отслеживает подключение к конкретной странице (POST)
	if r.URL.Path != "/" {
		ErrorHandler(w, http.StatusNotFound)
		return
	}

	url := "https://groupietrackers.herokuapp.com/api/artists"
	response, err := http.Get(url)
	if err != nil {
		ErrorHandler(w, http.StatusInternalServerError)
		return
	}
	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Println("Error:88", err)
		return
	}
	var data []artists
	err = json.Unmarshal(body, &data)
	if err != nil {
		ErrorHandler(w, http.StatusInternalServerError)
		return
	}

	if r.Method != "GET" {
		ErrorHandler(w, http.StatusInternalServerError)
		return
	}
	tmpl, err := template.ParseFiles("ui/templates/mainindex.html")
	err = tmpl.Execute(w, &data)
}
