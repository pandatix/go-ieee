package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/pandatix/go-ieee/v2/api"
)

func main() {
	// Create an IEEEClient.
	client, _ := api.NewIEEEClient(&http.Client{})

	// Manipulate your context if necessary.
	// Here we set up a timeout to make sure we don't wait indefinitely on IEEE
	// if the API service is downgraded.
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	// Then we issue our search request.
	// Values are copied from the IEEExplore frontend calls, there is no worry
	// to have considering this API is not documented thus you can't be
	// omniscient on their use.
	res, err := client.PostSearch(&api.PostSearchParams{
		Highlight: true,
		MatchPubs: true,
		QueryText: "White Rabbit",
		Refinements: []string{
			"ContentType:Conferences",
			"ContentType:Journals",
		},
		ReturnFacets: []string{
			"ALL",
		},
		ReturnType: "SEARCH",
	}, api.WithContext(ctx))
	if err != nil {
		log.Fatal(err)
	}

	// Finally, we manipulate our results as we need.
	fmt.Printf("Got %d records !\n", res.TotalRecords)
}
