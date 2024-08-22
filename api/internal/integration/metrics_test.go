package integration_test

import (
	"bytes"
	"encoding/json"
	"testing"

	"github.com/pandatix/go-ieee/v2/api"
	"github.com/stretchr/testify/assert"
)

func Test_I_GetDocumentMetrics(t *testing.T) {
	for testname, id := range tests {
		t.Run(testname, func(t *testing.T) {
			assert := assert.New(t)

			mdw := &MdwClient{}
			client, _ := api.NewIEEEClient(mdw)
			mtrcs, err := client.GetDocumentMetrics(id)

			// Ensure no error
			if !assert.Nil(err) {
				t.Errorf("Last body [%s]\n", mdw.LastBody)
			}

			// Reencode to JSON
			buf := &bytes.Buffer{}
			_ = json.NewEncoder(buf).Encode(mtrcs)

			// Decode both to interfaces
			var expected interface{}
			var actual interface{}
			_ = json.Unmarshal(mdw.LastBody, &expected)
			_ = json.Unmarshal(buf.Bytes(), &actual)

			// Compares both to check valid API (and not nil)
			assert.NotNil(expected)
			assert.Equal(expected, actual)
		})
	}
}
