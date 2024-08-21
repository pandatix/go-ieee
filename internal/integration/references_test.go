package integration_test

import (
	"bytes"
	"encoding/json"
	"testing"

	goieeeapi "github.com/pandatix/go-ieee-api"
	"github.com/stretchr/testify/assert"
)

func TestIntegrationGetReferences(t *testing.T) {
	for testname, id := range tests {
		t.Run(testname, func(t *testing.T) {
			assert := assert.New(t)

			client := &MdwClient{}
			authors, err := goieeeapi.GetReferences(client, id, nil)

			// Ensure no error
			if !assert.Nil(err) {
				t.Errorf("Last body [%s]\n", client.LastBody)
			}

			// Reencode to JSON
			buf := &bytes.Buffer{}
			_ = json.NewEncoder(buf).Encode(authors)

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
