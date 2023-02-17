package MapsClient

import (
	"context"
	"log"
	"googlemaps.github.io/maps"
	"os"
)


func GetTransitDirection(origin string, destination string) []maps.Route {
	apiKey, exists := os.LookupEnv("API_KEY")
	log.Println(apiKey)
	if !exists {
		log.Println("File .env not found")
	}
	
	c, err := maps.NewClient(maps.WithAPIKey(apiKey), maps.WithRateLimit(10))
	if err != nil {
		log.Fatalf("fatal error: %s", err)
	}

	r := &maps.DirectionsRequest{
		Origin:      origin,
		Destination: destination,
		Mode: "transit",
	}
	route, _, err := c.Directions(context.Background(), r)
	if err != nil {
		log.Fatalf("fatal error: %s", err)
	}

	return route
}