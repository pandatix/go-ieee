package integration_test

import (
	"bytes"
	"encoding/json"
	"testing"

	goieeeapi "github.com/pandatix/go-ieee-api"
	"github.com/stretchr/testify/assert"
)

func TestIntegrationPostSearch(t *testing.T) {
	var tests = map[string]struct {
		Params goieeeapi.PostSearchParams
	}{
		"QSD": {
			Params: goieeeapi.PostSearchParams{
				Highlight: true,
				MatchPubs: true,
				NewSearch: b(true),
				QueryText: "qsd",
				ReturnFacets: []string{
					"ALL",
				},
				ReturnType: "SEARCH",
			},
		},
		"space-biology-advanced": {
			Params: goieeeapi.PostSearchParams{
				Action:       str("search"),
				Highlight:    true,
				MatchBoolean: b(true),
				MatchPubs:    true,
				NewSearch:    b(true),
				QueryText:    `("Full Text Only":space) AND ("Authors":biology) NOT ("All Metadata":it)`,
				Ranges: &[]string{
					"1884_1997_Year",
				},
				ReturnFacets: []string{
					"ALL",
				},
				ReturnType: "SEARCH",
			},
		},
		"botnet-detection-2012_2016-Magazines-Early_Access_Articles-Web_sites": {
			Params: goieeeapi.PostSearchParams{
				Highlight: true,
				MatchPubs: true,
				QueryText: "Botnet detection",
				Ranges: &[]string{
					"2012_2016_Year",
				},
				Refinements: &[]string{
					"ContentType:Magazines",
					"ContentType:Early Access Articles",
					"ControlledTerms:Web sites",
				},
				ReturnFacets: []string{
					"ALL",
				},
				ReturnType: "SEARCH",
			},
		},
	}

	for testname, tt := range tests {
		t.Run(testname, func(t *testing.T) {
			assert := assert.New(t)

			client := &MdwClient{}
			rslts, err := goieeeapi.PostSearch(client, tt.Params)

			// Ensure no error
			if !assert.Nil(err) {
				t.Errorf("Last body [%s]\n", client.LastBody)
			}

			// Reencode to JSON
			buf := &bytes.Buffer{}
			_ = json.NewEncoder(buf).Encode(rslts)

			// Decode both to interfaces
			var expected interface{}
			var actual interface{}
			_ = json.Unmarshal(client.LastBody, &expected)
			_ = json.Unmarshal(buf.Bytes(), &actual)

			// Compares both to check valid API (and not nil)
			assert.NotNil(expected)
			assert.Equal(expected, actual)
		})
	}
}

func str(str string) *string {
	return &str
}

func b(b bool) *bool {
	return &b
}
