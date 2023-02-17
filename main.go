package main

import (
	"encoding/json"
	"fmt"
	"github.com/joho/godotenv"
	"net/http"
	"mapsenhancement/MapsClient"
	"log"
)


func singleDirection(w http.ResponseWriter, r *http.Request)  {
	origin := r.URL.Query().Get("origin")
	destination := r.URL.Query().Get("destination")

	route := MapsClient.GetTransitDirection(origin, destination)
	response, err := json.Marshal(route);
	if err != nil {
		fmt.Println("Error in serializing response")
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	w.Write(response)

}


func main() {
	godotenv.Load()
	http.HandleFunc("/api/getSingleDestination", singleDirection)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
