package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/pandatix/go-ieee/v2/api"
)

const docID = 5070278

func main() {
	// Create an IEEEClient.
	client, _ := api.NewIEEEClient(&http.Client{})

	// Manipulate your context if necessary.
	// Here we set up a timeout to make sure we don't wait indefinitely on IEEE
	// if the API service is downgraded.
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	// The we issue our document references request.
	res, err := client.GetDocumentReferences(docID, &api.GetReferencesParams{
		// No parameters needed here...
	}, api.WithContext(ctx))
	if err != nil {
		log.Fatal(err)
	}

	// Then print their title and look if the document contains
	// itself in its own references.
	containsItself := false
	for _, ref := range res.References {
		if ref.Title == nil {
			continue
		}
		fmt.Printf("%s\n", *ref.Title)
		if ref.Links != nil && ref.Links.ArticleNumber != nil && *ref.Links.ArticleNumber == strconv.Itoa(docID) {
			containsItself = true
		}
	}
	if containsItself {
		log.Fatalf("Document %d contains itself", docID)
	}
}
