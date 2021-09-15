package goieeeapi_test

import (
	"encoding/json"
	"net/http"
	"reflect"
	"testing"

	goieeeapi "github.com/pandatix/go-ieee-api"
)

func TestGetFootnotes(t *testing.T) {
	t.Parallel()

	var tests = map[string]struct {
		Client            goieeeapi.HTTPClient
		ID                int
		ExpectedFootnotes *goieeeapi.GetFootnotesResponse
		ExpectedErr       error
	}{
		"nil-client": {
			Client:            nil,
			ID:                0,
			ExpectedFootnotes: nil,
			ExpectedErr:       goieeeapi.ErrNilClient,
		},
		"failing-client": {
			Client:            newFakeHTTPClient(``, 0, errFake),
			ID:                0,
			ExpectedFootnotes: nil,
			ExpectedErr:       errFake,
		},
		"unexpected-statuscode": {
			Client:            newFakeHTTPClient(``, 0, nil),
			ID:                0,
			ExpectedFootnotes: nil,
			ExpectedErr: &goieeeapi.ErrUnexpectedStatus{
				StatusCode: 0,
				Body:       []byte(``),
			},
		},
		"failing-unmarshal": {
			Client:            newFakeHTTPClient(`{[}]`, http.StatusOK, nil),
			ID:                0,
			ExpectedFootnotes: nil,
			ExpectedErr: &json.SyntaxError{
				Offset: 2,
			},
		},
		"valid-call": {
			Client: newFakeHTTPClient(`{"userInfo":{"institute":false,"member":false,"individual":false,"guest":false,"subscribedContent":false,"fileCabinetContent":false,"fileCabinetUser":false,"institutionalFileCabinetUser":false,"showPatentCitations":true,"showGet802Link":false,"showOpenUrlLink":false,"tracked":false,"delegatedAdmin":false,"desktop":false,"isInstitutionDashboardEnabled":false,"isInstitutionProfileEnabled":false,"isRoamingEnabled":false,"isDelegatedAdmin":false,"isMdl":false,"isCwg":false},"articleNumber":"8900441","getProgramTermsAccepted":false,"formulaStrippedArticleTitle":"Deep Learning Solutions for Tandem-X-Based Forest Classification","pubLink":"/xpl/conhome/8891871/proceeding","allowComments":false,"issueLink":"/xpl/tocresult.jsp?isnumber=null","footnote":[{"id":"fn1","label":"","text":"We assume the ground-truth density map as membership degree."},{"id":"fn2","label":"","text":"A pass on the whole training dataset."}],"isReadingRoomArticle":false,"isGetArticle":false,"isGetAddressInfoCaptured":false,"isMarketingOptIn":false,"publisher":"IEEE","xploreDocumentType":"Conference Publication","isPromo":false,"isNotDynamicOrStatic":true,"htmlAbstractLink":"/document/8900441/","isCustomDenial":false,"isSAE":false,"isDynamicHtml":false,"isFreeDocument":false,"displayDocTitle":"Deep Learning Solutions for Tandem-X-Based Forest Classification","isStandard":false,"isSMPTE":false,"isOUP":false,"isNow":false,"isProduct":false,"isMorganClaypool":false,"isJournal":false,"isBook":false,"isBookWithoutChapters":false,"isOpenAccess":false,"isEphemera":false,"isConference":true,"isChapter":false,"isStaticHtml":false,"isEarlyAccess":false,"persistentLink":"https://ieeexplore.ieee.org/servlet/opac?punumber=8891871","contentTypeDisplay":"Conferences","mlTime":"PT0.02496S","lastupdate":"2021-08-21","mediaPath":"/mediastore_new/IEEE/content/media/8891871/8897702/8900441","title":"Deep Learning Solutions for Tandem-X-Based Forest Classification","contentType":"conferences","publicationNumber":"8891871"}`, http.StatusOK, nil),
			ID:     0,
			ExpectedFootnotes: &goieeeapi.GetFootnotesResponse{
				UserInfo: goieeeapi.UserInfo{
					Institute:                     false,
					Member:                        false,
					Individual:                    false,
					Guest:                         false,
					SubscribedContent:             false,
					FileCabinetContent:            false,
					FileCabinetUser:               false,
					InstitutionalFileCabinetUser:  false,
					ShowPatentCitations:           true,
					ShowGet802Link:                false,
					ShowOpenURLLink:               false,
					Tracked:                       false,
					DelegatedAdmin:                false,
					Desktop:                       false,
					IsInstitutionDashboardEnabled: false,
					IsInstitutionProfileEnabled:   false,
					IsRoamingEnabled:              false,
					IsDelegatedAdmin:              false,
					IsMdl:                         false,
					IsCwg:                         false,
				},
				ArticleNumber:               "8900441",
				GetProgramTermsAccepted:     false,
				FormulaStrippedArticleTitle: str("Deep Learning Solutions for Tandem-X-Based Forest Classification"),
				PubLink:                     str("/xpl/conhome/8891871/proceeding"),
				AllowComments:               false,
				IssueLink:                   "/xpl/tocresult.jsp?isnumber=null",
				Footnote: &[]goieeeapi.Footnote{
					{
						ID:    "fn1",
						Label: "",
						Text:  "We assume the ground-truth density map as membership degree.",
					}, {
						ID:    "fn2",
						Label: "",
						Text:  "A pass on the whole training dataset.",
					},
				},
				IsReadingRoomArticle:     false,
				IsGetArticle:             false,
				IsGetAddressInfoCaptured: false,
				IsMarketingOptIn:         false,
				Publisher:                str("IEEE"),
				XploreDocumentType:       str("Conference Publication"),
				IsPromo:                  false,
				IsNotDynamicOrStatic:     true,
				HTMLAbstractLink:         "/document/8900441/",
				IsCustomDenial:           false,
				IsSAE:                    false,
				IsDynamicHTML:            false,
				IsFreeDocument:           false,
				DisplayDocTitle:          str("Deep Learning Solutions for Tandem-X-Based Forest Classification"),
				IsStandard:               false,
				IsSMPTE:                  false,
				IsOUP:                    false,
				IsNow:                    false,
				IsProduct:                false,
				IsMorganClaypool:         false,
				IsJournal:                false,
				IsBook:                   false,
				IsBookWithoutChapters:    false,
				IsOpenAccess:             false,
				IsEphemera:               false,
				IsConference:             true,
				IsChapter:                false,
				IsStaticHTML:             false,
				IsEarlyAccess:            false,
				PersistentLink:           str("https://ieeexplore.ieee.org/servlet/opac?punumber=8891871"),
				Title:                    str("Deep Learning Solutions for Tandem-X-Based Forest Classification"),
				ContentTypeDisplay:       str("Conferences"),
				MlTime:                   "PT0.02496S",
				LastUpdate:               str("2021-08-21"),
				MediaPath:                str("/mediastore_new/IEEE/content/media/8891871/8897702/8900441"),
				ContentType:              str("conferences"),
				PublicationNumber:        str("8891871"),
			},
			ExpectedErr: nil,
		},
	}

	for testname, tt := range tests {
		t.Run(testname, func(t *testing.T) {
			ftns, err := goieeeapi.GetFootnotes(tt.Client, tt.ID)

			if !reflect.DeepEqual(ftns, tt.ExpectedFootnotes) {
				t.Errorf("Failed to get expected footnotes: got \"%v\" instead of \"%v\".", ftns, tt.ExpectedFootnotes)
			}
			checkErr(err, tt.ExpectedErr, t)
		})
	}
}
