package goieeeapi_test

import (
	"encoding/json"
	"net/http"
	"reflect"
	"testing"

	goieeeapi "github.com/pandatix/go-ieee-api"
)

func TestGetCitations(t *testing.T) {
	t.Parallel()

	var tests = map[string]struct {
		Client            goieeeapi.HTTPClient
		ID                int
		ExpectedCitations *goieeeapi.GetCitationsResponse
		ExpectedErr       error
	}{
		"nil-client": {
			Client:            nil,
			ID:                0,
			ExpectedCitations: nil,
			ExpectedErr:       goieeeapi.ErrNilClient,
		},
		"failing-client": {
			Client:            newFakeHTTPClient(``, 0, errFake),
			ID:                0,
			ExpectedCitations: nil,
			ExpectedErr:       errFake,
		},
		"unexpected-statuscode": {
			Client:            newFakeHTTPClient(``, 0, nil),
			ID:                0,
			ExpectedCitations: nil,
			ExpectedErr: &goieeeapi.ErrUnexpectedStatus{
				StatusCode: 0,
				Body:       []byte(``),
			},
		},
		"failing-unmarshal": {
			Client:            newFakeHTTPClient(`{[}]`, http.StatusOK, nil),
			ID:                0,
			ExpectedCitations: nil,
			ExpectedErr: &json.SyntaxError{
				Offset: 2,
			},
		},
		"valid-call": {
			Client: newFakeHTTPClient(`{"getProgramTermsAccepted":false,"formulaStrippedArticleTitle":"Deep Learning Solutions for Tandem-X-Based Forest Classification","paperCitations":{"ieee":[{"order":"1","displayText":"S. Vitale, G. Ferraioli, V. Pascazio, \"Edge Preserving Cnn Sar Despeckling Algorithm\", <i>Remote Sensing Conference (LAGIRS) 2020 IEEE Latin American GRSS & ISPRS</i>, pp. 12-15, 2020.","links":{"documentLink":"/document/9165559","pdfLink":"/stamp/stamp.jsp?tp=&arnumber=9165559","openUrlImgLoc":"/assets/img/btn.find-in-library.png","pdfSize":"202KB"},"googleScholarLink":"https://scholar.google.com/scholar?as_q=Edge+Preserving+Cnn+Sar+Despeckling+Algorithm&as_occt=title&hl=en&as_sdt=0%2C31","title":"Edge Preserving Cnn Sar Despeckling Algorithm"}]},"isReadingRoomArticle":false,"isGetArticle":false,"isGetAddressInfoCaptured":false,"isMarketingOptIn":false,"isProduct":false,"isChapter":false,"isEarlyAccess":false,"nonIeeeCitationCount":"0","contentTypeDisplay":"Conferences","patentCitationCount":"0","mlTime":"PT0.033714S","lastupdate":"2021-08-21","mediaPath":"/mediastore_new/IEEE/content/media/8891871/8897702/8900441","title":"Deep Learning Solutions for Tandem-X-Based Forest Classification","contentType":"conferences","ieeeCitationCount":"1","publicationNumber":"8891871"}`, http.StatusOK, nil),
			ID:     0,
			ExpectedCitations: &goieeeapi.GetCitationsResponse{
				GetProgramTermsAccepted:     false,
				FormulaStrippedArticleTitle: str("Deep Learning Solutions for Tandem-X-Based Forest Classification"),
				PaperCitations: goieeeapi.PaperCitations{
					IEEE: &[]goieeeapi.Reference{
						{
							Order:       "1",
							DisplayText: str("S. Vitale, G. Ferraioli, V. Pascazio, \"Edge Preserving Cnn Sar Despeckling Algorithm\", <i>Remote Sensing Conference (LAGIRS) 2020 IEEE Latin American GRSS & ISPRS</i>, pp. 12-15, 2020."),
							Links: &goieeeapi.Links{
								DocumentLink:  str("/document/9165559"),
								PdfLink:       str("/stamp/stamp.jsp?tp=&arnumber=9165559"),
								OpenURLImgLoc: "/assets/img/btn.find-in-library.png",
								PdfSize:       str("202KB"),
							},
							GoogleScholarLink: str("https://scholar.google.com/scholar?as_q=Edge+Preserving+Cnn+Sar+Despeckling+Algorithm&as_occt=title&hl=en&as_sdt=0%2C31"),
							Title:             str("Edge Preserving Cnn Sar Despeckling Algorithm"),
						},
					},
				},
				IsReadingRoomArticle:     false,
				IsGetArticle:             false,
				IsGetAddressInfoCaptured: false,
				IsMarketingOptIn:         false,
				IsProduct:                false,
				IsChapter:                false,
				IsEarlyAccess:            false,
				NonIeeeCitationCount:     str("0"),
				ContentTypeDisplay:       str("Conferences"),
				PatentCitationCount:      "0",
				MlTime:                   "PT0.033714S",
				LastUpdate:               str("2021-08-21"),
				MediaPath:                str("/mediastore_new/IEEE/content/media/8891871/8897702/8900441"),
				Title:                    str("Deep Learning Solutions for Tandem-X-Based Forest Classification"),
				ContentType:              str("conferences"),
				IEEECitationCount:        str("1"),
				PublicationNumber:        str("8891871"),
			},
			ExpectedErr: nil,
		},
	}

	for testname, tt := range tests {
		t.Run(testname, func(t *testing.T) {
			cts, err := goieeeapi.GetCitations(tt.Client, tt.ID)

			if !reflect.DeepEqual(cts, tt.ExpectedCitations) {
				t.Errorf("Failed to get expected citations: got \"%v\" instead of \"%v\".", cts, tt.ExpectedCitations)
			}
			checkErr(err, tt.ExpectedErr, t)
		})
	}
}
