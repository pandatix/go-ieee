package goieeeapi_test

import (
	"encoding/json"
	"net/http"
	"reflect"
	"testing"

	goieeeapi "github.com/pandatix/go-ieee-api"
)

func TestGetSimilar(t *testing.T) {
	t.Parallel()

	var tests = map[string]struct {
		Client          goieeeapi.HTTPClient
		ID              int
		ExpectedSimilar *goieeeapi.GetSimilarResponse
		ExpectedErr     error
	}{
		"nil-client": {
			Client:          nil,
			ID:              0,
			ExpectedSimilar: nil,
			ExpectedErr:     goieeeapi.ErrNilClient,
		},
		"failing-client": {
			Client:          newFakeHTTPClient(``, 0, errFake),
			ID:              0,
			ExpectedSimilar: nil,
			ExpectedErr:     errFake,
		},
		"unexpected-statuscode": {
			Client:          newFakeHTTPClient(``, 0, nil),
			ID:              0,
			ExpectedSimilar: nil,
			ExpectedErr: &goieeeapi.ErrUnexpectedStatus{
				StatusCode: 0,
				Body:       []byte(``),
			},
		},
		"failing-unmarshal": {
			Client:          newFakeHTTPClient(`{[}]`, http.StatusOK, nil),
			ID:              0,
			ExpectedSimilar: nil,
			ExpectedErr: &json.SyntaxError{
				Offset: 2,
			},
		},
		"valid-call": {
			Client: newFakeHTTPClient(`{"getProgramTermsAccepted":false,"similar":[{"articleNumber":"8378314","title":"A comprehensive three-phase load flow method for integrated MV and LV distribution networks","publicationNumber":"8370759","publicationTitle":"2017 IEEE Innovative Smart Grid Technologies - Asia (ISGT-Asia)","publicationYear":"2017","author":["Yang Fu","Xiaoming Zhou","Xiangjing Su","Peter Wolfs","Mohammad A.S. Masoum"],"contentType":"Conferences","documentLink":"/document/8378314/","isArticle":"true"},{"articleNumber":"9321393","title":"Optimal Battery Energy Storage System Scheduling Based on Mutation-Improved Grey Wolf Optimizer Using GPU-Accelerated Load Flow in Active Distribution Networks","publicationNumber":"6287639","publicationTitle":"IEEE Access","publicationYear":"2021","author":["Dorian O. Sidea","Irina I. Picioroaga","Constantin Bulac"],"contentType":"Journals & Magazines","documentLink":"/document/9321393/","isArticle":"true"},{"articleNumber":"8076277","title":"Power grid configuration influence on the geomagnetically induced currents value in power transformers","publicationNumber":"8053490","publicationTitle":"2017 International Conference on Industrial Engineering, Applications and Manufacturing (ICIEAM)","publicationYear":"2017","author":["V. V. Vakhnina","A. A. Kuvshinov","A. A. Chernenko","D. A. Kretov"],"contentType":"Conferences","documentLink":"/document/8076277/","isArticle":"true"}],"formulaStrippedArticleTitle":"Sensitivity analysis - a key element for the operation of a flexible distribution grid","isReadingRoomArticle":false,"isGetArticle":false,"isGetAddressInfoCaptured":false,"isMarketingOptIn":false,"publisher":"IEEE","isCustomDenial":false,"isSMPTE":false,"isOUP":false,"isNow":false,"isSAE":false,"isStandard":false,"isConference":true,"isProduct":false,"isMorganClaypool":false,"isJournal":false,"isBook":false,"isEarlyAccess":false,"persistentLink":"https://ieeexplore.ieee.org/servlet/opac?punumber=7513513","isChapter":false,"contentTypeDisplay":"Conferences","mlTime":"PT0.034521S","lastupdate":"2021-08-15","mediaPath":"/mediastore_new/IEEE/content/media/7513513/7519846/7520000","title":"Sensitivity analysis - a key element for the operation of a flexible distribution grid","contentType":"conferences","publicationNumber":"7513513"}`, http.StatusOK, nil),
			ID:     0,
			ExpectedSimilar: &goieeeapi.GetSimilarResponse{
				GetProgramTermsAccepted: false,
				Similar: &[]goieeeapi.Similar{
					{
						ArticleNumber:     "8378314",
						Title:             "A comprehensive three-phase load flow method for integrated MV and LV distribution networks",
						PublicationNumber: "8370759",
						PublicationTitle:  "2017 IEEE Innovative Smart Grid Technologies - Asia (ISGT-Asia)",
						PublicationYear:   "2017",
						Author: &[]string{
							"Yang Fu",
							"Xiaoming Zhou",
							"Xiangjing Su",
							"Peter Wolfs",
							"Mohammad A.S. Masoum",
						},
						ContentType:  "Conferences",
						DocumentLink: "/document/8378314/",
						IsArticle:    "true",
					}, {
						ArticleNumber:     "9321393",
						Title:             "Optimal Battery Energy Storage System Scheduling Based on Mutation-Improved Grey Wolf Optimizer Using GPU-Accelerated Load Flow in Active Distribution Networks",
						PublicationNumber: "6287639",
						PublicationTitle:  "IEEE Access",
						PublicationYear:   "2021",
						Author: &[]string{
							"Dorian O. Sidea",
							"Irina I. Picioroaga",
							"Constantin Bulac",
						},
						ContentType:  "Journals & Magazines",
						DocumentLink: "/document/9321393/",
						IsArticle:    "true",
					}, {
						ArticleNumber:     "8076277",
						Title:             "Power grid configuration influence on the geomagnetically induced currents value in power transformers",
						PublicationNumber: "8053490",
						PublicationTitle:  "2017 International Conference on Industrial Engineering, Applications and Manufacturing (ICIEAM)",
						PublicationYear:   "2017",
						Author: &[]string{
							"V. V. Vakhnina",
							"A. A. Kuvshinov",
							"A. A. Chernenko",
							"D. A. Kretov",
						},
						ContentType:  "Conferences",
						DocumentLink: "/document/8076277/",
						IsArticle:    "true",
					},
				},
				FormulaStrippedArticleTitle: str("Sensitivity analysis - a key element for the operation of a flexible distribution grid"),
				IsReadingRoomArticle:        false,
				IsGetArticle:                false,
				IsGetAddressInfoCaptured:    false,
				IsMarketingOptIn:            false,
				Publisher:                   str("IEEE"),
				IsCustomDenial:              false,
				IsSMPTE:                     false,
				IsOUP:                       false,
				IsNow:                       false,
				IsSAE:                       false,
				IsStandard:                  false,
				IsConference:                true,
				IsProduct:                   false,
				IsMorganClaypool:            false,
				IsJournal:                   false,
				IsBook:                      false,
				IsEarlyAccess:               false,
				PersistentLink:              str("https://ieeexplore.ieee.org/servlet/opac?punumber=7513513"),
				IsChapter:                   false,
				ContentTypeDisplay:          str("Conferences"),
				MlTime:                      str("PT0.034521S"),
				LastUpdate:                  str("2021-08-15"),
				MediaPath:                   str("/mediastore_new/IEEE/content/media/7513513/7519846/7520000"),
				Title:                       str("Sensitivity analysis - a key element for the operation of a flexible distribution grid"),
				ContentType:                 str("conferences"),
				PublicationNumber:           str("7513513"),
			},
			ExpectedErr: nil,
		},
	}

	for testname, tt := range tests {
		t.Run(testname, func(t *testing.T) {
			sim, err := goieeeapi.GetSimilar(tt.Client, tt.ID)

			if !reflect.DeepEqual(sim, tt.ExpectedSimilar) {
				t.Errorf("Failed to get expected similar: got \"%v\" instead of \"%v\".", sim, tt.ExpectedSimilar)
			}
			checkErr(err, tt.ExpectedErr, t)
		})
	}
}
