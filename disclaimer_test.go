package goieeeapi_test

import (
	"encoding/json"
	"net/http"
	"reflect"
	"testing"

	goieeeapi "github.com/pandatix/go-ieee-api"
)

func TestGetDisclaimer(t *testing.T) {
	t.Parallel()

	var tests = map[string]struct {
		Client             goieeeapi.HTTPClient
		ID                 int
		ExpectedDisclaimer *goieeeapi.GetDisclaimerResponse
		ExpectedErr        error
	}{
		"nil-client": {
			Client:             nil,
			ID:                 0,
			ExpectedDisclaimer: nil,
			ExpectedErr:        goieeeapi.ErrNilClient,
		},
		"failing-client": {
			Client:             newFakeHTTPClient(``, 0, errFake),
			ID:                 0,
			ExpectedDisclaimer: nil,
			ExpectedErr:        errFake,
		},
		"unexpected-statuscode": {
			Client:             newFakeHTTPClient(``, 0, nil),
			ID:                 0,
			ExpectedDisclaimer: nil,
			ExpectedErr: &goieeeapi.ErrUnexpectedStatus{
				StatusCode: 0,
				Body:       []byte(``),
			},
		},
		"failing-unmarshal": {
			Client:             newFakeHTTPClient(`{[}]`, http.StatusOK, nil),
			ID:                 0,
			ExpectedDisclaimer: nil,
			ExpectedErr: &json.SyntaxError{
				Offset: 2,
			},
		},
		"valid-call": {
			Client: newFakeHTTPClient(`{"userInfo":{"institute":false,"member":false,"individual":false,"guest":false,"subscribedContent":false,"fileCabinetContent":false,"fileCabinetUser":false,"institutionalFileCabinetUser":false,"showPatentCitations":true,"showGet802Link":false,"showOpenUrlLink":false,"tracked":false,"desktop":false,"delegatedAdmin":false,"isInstitutionDashboardEnabled":false,"isInstitutionProfileEnabled":false,"isRoamingEnabled":false,"isDelegatedAdmin":false,"isMdl":false,"isCwg":false},"articleNumber":"8900441","getProgramTermsAccepted":false,"pubLink":"/xpl/conhome/8891871/proceeding","allowComments":false,"issueLink":"/xpl/tocresult.jsp?isnumber=null","formulaStrippedArticleTitle":"Deep Learning Solutions for Tandem-X-Based Forest Classification","isReadingRoomArticle":false,"isGetArticle":false,"isGetAddressInfoCaptured":false,"isMarketingOptIn":false,"publisher":"IEEE","isPromo":false,"isSMPTE":false,"isOUP":false,"isSAE":false,"isNow":false,"isCustomDenial":false,"isNotDynamicOrStatic":true,"displayDocTitle":"Deep Learning Solutions for Tandem-X-Based Forest Classification","isStandard":false,"htmlAbstractLink":"/document/8900441/","isProduct":false,"isMorganClaypool":false,"isOpenAccess":false,"isEphemera":false,"isConference":true,"isEarlyAccess":false,"isJournal":false,"isBook":false,"isBookWithoutChapters":false,"isChapter":false,"isStaticHtml":false,"persistentLink":"https://ieeexplore.ieee.org/servlet/opac?punumber=8891871","xploreDocumentType":"Conference Publication","isFreeDocument":false,"isDynamicHtml":false,"contentTypeDisplay":"Conferences","mlTime":"PT0.006651S","lastupdate":"2021-08-21","mediaPath":"/mediastore_new/IEEE/content/media/8891871/8897702/8900441","title":"Deep Learning Solutions for Tandem-X-Based Forest Classification","contentType":"conferences","publicationNumber":"8891871"}`, http.StatusOK, nil),
			ID:     0,
			ExpectedDisclaimer: &goieeeapi.GetDisclaimerResponse{
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
					ShowOpenURLLink:               b(false),
					Tracked:                       b(false),
					DelegatedAdmin:                b(false),
					Desktop:                       b(false),
					IsInstitutionDashboardEnabled: b(false),
					IsInstitutionProfileEnabled:   b(false),
					IsRoamingEnabled:              b(false),
					IsDelegatedAdmin:              b(false),
					IsMdl:                         b(false),
					IsCwg:                         b(false),
				},
				ArticleNumber:               "8900441",
				GetProgramTermsAccepted:     false,
				PubLink:                     str("/xpl/conhome/8891871/proceeding"),
				AllowComments:               false,
				IssueLink:                   "/xpl/tocresult.jsp?isnumber=null",
				FormulaStrippedArticleTitle: str("Deep Learning Solutions for Tandem-X-Based Forest Classification"),
				IsReadingRoomArticle:        false,
				IsGetArticle:                false,
				IsGetAddressInfoCaptured:    false,
				IsMarketingOptIn:            false,
				Publisher:                   str("IEEE"),
				IsPromo:                     false,
				IsSMPTE:                     false,
				IsOUP:                       false,
				IsSAE:                       false,
				IsNow:                       false,
				IsCustomDenial:              false,
				IsNotDynamicOrStatic:        true,
				DisplayDocTitle:             str("Deep Learning Solutions for Tandem-X-Based Forest Classification"),
				IsStandard:                  false,
				HTMLAbstractLink:            "/document/8900441/",
				IsProduct:                   false,
				IsMorganClaypool:            false,
				IsOpenAccess:                false,
				IsEphemera:                  false,
				IsConference:                true,
				IsEarlyAccess:               false,
				IsJournal:                   false,
				IsBook:                      false,
				IsBookWithoutChapters:       false,
				IsChapter:                   false,
				IsStaticHTML:                false,
				PersistentLink:              str("https://ieeexplore.ieee.org/servlet/opac?punumber=8891871"),
				XploreDocumentType:          str("Conference Publication"),
				IsFreeDocument:              false,
				IsDynamicHTML:               false,
				ContentTypeDisplay:          str("Conferences"),
				MlTime:                      "PT0.006651S",
				LastUpdate:                  str("2021-08-21"),
				MediaPath:                   str("/mediastore_new/IEEE/content/media/8891871/8897702/8900441"),
				Title:                       str("Deep Learning Solutions for Tandem-X-Based Forest Classification"),
				ContentType:                 str("conferences"),
				PublicationNumber:           str("8891871"),
			},
			ExpectedErr: nil,
		},
	}

	for testname, tt := range tests {
		t.Run(testname, func(t *testing.T) {
			disc, err := goieeeapi.GetDisclaimer(tt.Client, tt.ID)

			if !reflect.DeepEqual(disc, tt.ExpectedDisclaimer) {
				t.Errorf("Failed to get expected disclaimer: got \"%v\" instead of \"%v\".", disc, tt.ExpectedDisclaimer)
			}
			checkErr(err, tt.ExpectedErr, t)
		})
	}
}
