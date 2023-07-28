package functions

import (
	"encoding/json"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"regexp"
	"strings"
)

func Artists(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		ErrorHandler(w, http.StatusMethodNotAllowed)
		return
	}

	id := r.URL.Query().Get("id")
	var data artists
	url := "https://groupietrackers.herokuapp.com/api/artists/" + id
	response, err := http.Get(url)
	if err != nil {
		log.Println("Error:", err)
		return
	}
	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Println("Error:", err)
		return
	}
	err = json.Unmarshal(body, &data)

	relations, err := GetApi2(data.RelationsID)
	if err != nil {
		ErrorHandler(w, http.StatusNotFound)
		return
	}

	response, err = http.Get(data.RelationsID)
	if err != nil {
		ErrorHandler(w, http.StatusNotFound)
		return
	}
	defer response.Body.Close()
	body, err = ioutil.ReadAll(response.Body)
	if err != nil {
		log.Println("Error:", err)
		return
	}
	err = json.Unmarshal(body, &data)
	tmpl, err := template.ParseFiles("ui/templates/info.html")
	newRelations := make(map[string][]string)
	for k, v := range relations.Relations {
		k = strings.Title(k)
		k = regexp.MustCompile("-").ReplaceAllString(k, ", ")
		k = strings.Title(strings.ReplaceAll(k, "_", " "))
		k = regexp.MustCompile("Usa").ReplaceAllString(k, "USA")
		k = regexp.MustCompile("Uk").ReplaceAllString(k, "UK")
		newRelations[k] = v
	}
	data.LocationsDate = newRelations
	err = tmpl.Execute(w, data)
}
