package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"

	_ "github.com/joho/godotenv/autoload"
	"googlemaps.github.io/maps"
)

func main() {
	c, err := maps.NewClient(maps.WithAPIKey(os.Getenv("GOOGLE_API_KEY")))
	if err != nil {
		log.Fatalf("starting maps client: %s", err)
	}

	res, err := c.DistanceMatrix(context.Background(),
		&maps.DistanceMatrixRequest{
			Origins: []string{
				"32.90967,-97.09656",
				"32.90842,-97.09633",
				"32.91025,-97.09617",
				"32.90858,-97.0978",
			},
			Destinations: []string{
				"32.90842,-97.09633",
				"32.91025,-97.09617",
				"32.90858,-97.0978",
				"32.9086,-97.09785",
			},
			Mode:  maps.TravelModeBicycling,
			Units: maps.UnitsMetric,
		},
	)

	if err != nil {
		log.Fatalf("requesting distance matrix: %s", err)
	}

	b, err := json.MarshalIndent(res, "", "  ")
	if err != nil {
		log.Fatalf("trying to parse the response: %s", err)
	}

	fmt.Printf("%s", b)
}
