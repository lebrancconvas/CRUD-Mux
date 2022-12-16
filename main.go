// Link: https://www.youtube.com/watch?v=gIe543ufGmE  
package main 

import (
	_"fmt" 
	"net/http"  
	"encoding/json" 
	"github.com/gorilla/mux" 
	"log" 
)

type Track struct {
	ID int `json: "id"`
	Title string `json: "title"`
	Artist *Artist `json: "artist"` 
	Album string `json: "album"`
	Year int `json: "year"` 
}

type Artist struct { 
	ID int `json: "id"`
	FirstName string `json: "firstname"`
	LastName string `json: "lastname"` 
}

var tracks []Track; 

// Get All Tracks. 
func getTracks(write http.ResponseWriter, read *http.Request) {
	write.Header().Set("Content-Type", "application/json") 
	json.NewEncoder(write).Encode(tracks)
}

// Get Single Track. 
func getTrack(write http.ResponseWriter, read *http.Request) {
	write.Header().Set("Content-Type", "application/json")
	params := mux.Vars(read) 
	// Not finished. 
}

func main() {
	router := mux.NewRouter()

	tracks = append(tracks, Track{
		ID: 1,
		Title: "Shanghai Beach", 
		Artist: &Artist{
			ID: 1,
			FirstName: "Frances",
			LastName: "Yip",
		},
		Album: "Recollecting Frances",
		Year: 2003,  
	})

	router.HandleFunc("/api/tracks", getTracks).Methods("GET")
	log.Fatal(http.ListenAndServe(":8080", router)) 
}