package main

import (
	"context"
	"fmt"
	"log"

	"github.com/cloudflare/cloudflare-go"
)

func main() {
	api, err := cloudflare.New("apiKey", "admin@stocktonsoft.com")
	if err != nil {
		log.Fatal(err)
	}

	zoneID, err := api.ZoneIDByName("domain.com")
	if err != nil {
		log.Fatal(err)
	}

	pageOpts := cloudflare.PaginationOptions{
		PerPage: 5,
		Page:    1,
	}
	rateLimits, _, err := api.ListRateLimits(context.Background(), zoneID, pageOpts)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%+v\n", rateLimits)
	for _, r := range rateLimits {
		fmt.Printf("%+v\n", r)
	}
}
