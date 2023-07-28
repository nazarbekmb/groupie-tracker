package functions

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func GetApi2(url string) (rel Relations, err error) {
	response, err := http.Get(url)
	if err != nil {
		err = fmt.Errorf("Error:33 %v", err)
		return
	}
	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		err = fmt.Errorf("Error:741 %v", err)
		return
	}
	err = json.Unmarshal(body, &rel)
	if err != nil {
		err = fmt.Errorf("Error: Unmarshaling failed %v", err)
		return
	}
	return
}

type artists struct {
	ID            int                 `json:"id"`
	Image         string              `json:"image"`
	Name          string              `json:"name"`
	Members       []string            `json:"members"`
	CreationDate  int                 `json:"creationDate"`
	FirstAlbum    string              `json:"firstAlbum"`
	RelationsID   string              `json:"relations"`
	LocationsDate map[string][]string `json:"datesLocations"`
}

type Locations struct {
	ID        int      `json:"id"`
	Locations []string `json:"locations"`
}

type Relations struct {
	ID        int                 `json:"id"`
	Relations map[string][]string `json:"datesLocations"`
}
