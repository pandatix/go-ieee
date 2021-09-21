# Go IEEE API

[![reference](https://godoc.org/github.com/pandatix/go-ieee-api/v5?status.svg=)](https://pkg.go.dev/github.com/pandatix/go-ieee-api)
[![go report](https://goreportcard.com/badge/github.com/PandatiX/go-ieee-api)](https://goreportcard.com/report/github.com/PandatiX/go-ieee-api)
[![codecov](https://codecov.io/gh/PandatiX/go-ieee-api/branch/master/graph/badge.svg)](https://codecov.io/gh/PandatiX/go-ieee-api)
[![CI](https://github.com/PandatiX/go-ieee-api/actions/workflows/ci.yaml/badge.svg)](https://github.com/PandatiX/go-ieee-api/actions?query=workflow%3Aci+)

Go IEEE API wraps the REST IEEE API, for the following endpoints:
 - search
 - document/abstract
 - document/authors
 - document/citations
 - document/disclaimer
 - document/figures
 - document/footnotes
 - document/keywords
 - document/metrics
 - document/multimedia
 - document/references

## How to use

In case you want to fetch references for a given ID, you could do the following.

```go
import (
	"log"
	"net/http"

	goieeeapi "github.com/pandatix/go-ieee-api"
)

func main() {
	client := &http.Client{}
	refs, err := goieeeapi.GetReferences(client, 1234567)
	if err != nil {
		log.Fatal(err)
	}

	for i, ref := range *refs.References {
		log.Printf("Ref #%d: %s\n", i, *ref.Title)
	}
}
```
