// api/handler.go
package api

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

type Music struct {
	Id     string `json:"id"`
	Name   string `json:"name"`
	Artist string `json:"artist"`
	Time   string `json:"time"`
}

var album = []Music{
	{Id: "1", Name: "Look what you've done", Artist: "Drake", Time: "5:02"},
	{Id: "2", Name: "Margaret", Artist: "Lana Del Rey", Time: "5:40"},
	{Id: "3", Name: "Bigudi", Artist: "Ivan Dorn", Time: "4:53"},
	{Id: "4", Name: "Santorini", Artist: "Mot", Time: "3:02"},
}

func HealthCheck(w http.ResponseWriter, r *http.Request) {
	responseText := "Welcome to My GO TSIS\nAuthor: Gaziza"
	w.Header().Set("Content-Type", "text/plain")

	_, err := w.Write([]byte(responseText))
	if err != nil {
		http.Error(w, "Error writing response", http.StatusInternalServerError)
		return
	}
}

func GetAllMusics(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(album)
}

func GetMusic(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)
	musicID := params["id"]

	var foundMusic *Music
	for _, m := range album {
		if m.Id == musicID {
			foundMusic = &m
			break
		}
	}
	json.NewEncoder(w).Encode(foundMusic)
}
