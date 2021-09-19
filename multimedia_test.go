package goieeeapi_test

import (
	"encoding/json"
	"net/http"
	"reflect"
	"testing"

	goieeeapi "github.com/pandatix/go-ieee-api"
)

func TestGetMultimedia(t *testing.T) {
	t.Parallel()

	var tests = map[string]struct {
		Client          goieeeapi.HTTPClient
		ID              int
		ExpectedFigures *goieeeapi.GetMultimediaResponse
		ExpectedErr     error
	}{
		"nil-client": {
			Client:          nil,
			ID:              0,
			ExpectedFigures: nil,
			ExpectedErr:     goieeeapi.ErrNilClient,
		},
		"failing-client": {
			Client:          newFakeHTTPClient(``, 0, errFake),
			ID:              0,
			ExpectedFigures: nil,
			ExpectedErr:     errFake,
		},
		"unexpected-statuscode": {
			Client:          newFakeHTTPClient(``, 0, nil),
			ID:              0,
			ExpectedFigures: nil,
			ExpectedErr: &goieeeapi.ErrUnexpectedStatus{
				StatusCode: 0,
				Body:       []byte(``),
			},
		},
		"failing-unmarshal": {
			Client:          newFakeHTTPClient(`{[}]`, http.StatusOK, nil),
			ID:              0,
			ExpectedFigures: nil,
			ExpectedErr: &json.SyntaxError{
				Offset: 2,
			},
		},
		"valid-call": {
			Client: newFakeHTTPClient(`{"userInfo":{"institute":false,"member":false,"individual":false,"guest":false,"subscribedContent":false,"fileCabinetContent":false,"fileCabinetUser":false,"institutionalFileCabinetUser":false,"showPatentCitations":true,"showGet802Link":false,"showOpenUrlLink":false,"tracked":false,"delegatedAdmin":false,"desktop":false,"isInstitutionDashboardEnabled":false,"isInstitutionProfileEnabled":false,"isRoamingEnabled":false,"isDelegatedAdmin":false,"isMdl":false,"isCwg":false},"articleNumber":"8900441","getProgramTermsAccepted":false,"allowComments":false,"pubLink":"/xpl/conhome/8891871/proceeding","issueLink":"/xpl/tocresult.jsp?isnumber=null","formulaStrippedArticleTitle":"Deep Learning Solutions for Tandem-X-Based Forest Classification","isReadingRoomArticle":false,"isGetArticle":false,"isGetAddressInfoCaptured":false,"isMarketingOptIn":false,"publisher":"IEEE","isDynamicHtml":false,"isFreeDocument":false,"displayDocTitle":"Deep Learning Solutions for Tandem-X-Based Forest Classification","isStandard":false,"isMorganClaypool":false,"isConference":true,"isProduct":false,"isPromo":false,"isBookWithoutChapters":false,"persistentLink":"https://ieeexplore.ieee.org/servlet/opac?punumber=8891871","isEarlyAccess":false,"isJournal":false,"isBook":false,"isChapter":false,"isStaticHtml":false,"isOpenAccess":false,"isEphemera":false,"htmlAbstractLink":"/document/8900441/","isSMPTE":false,"isOUP":false,"isSAE":false,"isNow":false,"isCustomDenial":false,"isNotDynamicOrStatic":true,"xploreDocumentType":"Conference Publication","contentTypeDisplay":"Conferences","mlTime":"PT0.027462S","lastupdate":"2021-08-21","mediaPath":"/mediastore_new/IEEE/content/media/8891871/8897702/8900441","title":"Deep Learning Solutions for Tandem-X-Based Forest Classification","contentType":"conferences","publicationNumber":"8891871"}`, http.StatusOK, nil),
			ID:     0,
			ExpectedFigures: &goieeeapi.GetMultimediaResponse{
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
				AllowComments:               false,
				PubLink:                     str("/xpl/conhome/8891871/proceeding"),
				IssueLink:                   "/xpl/tocresult.jsp?isnumber=null",
				FormulaStrippedArticleTitle: str("Deep Learning Solutions for Tandem-X-Based Forest Classification"),
				IsReadingRoomArticle:        false,
				IsGetArticle:                false,
				IsGetAddressInfoCaptured:    false,
				IsMarketingOptIn:            false,
				Publisher:                   str("IEEE"),
				IsDynamicHTML:               false,
				IsFreeDocument:              false,
				DisplayDocTitle:             str("Deep Learning Solutions for Tandem-X-Based Forest Classification"),
				IsStandard:                  false,
				IsMorganClaypool:            false,
				IsConference:                true,
				IsProduct:                   false,
				IsPromo:                     false,
				IsBookWithoutChapters:       false,
				PersistentLink:              str("https://ieeexplore.ieee.org/servlet/opac?punumber=8891871"),
				IsEarlyAccess:               false,
				IsJournal:                   false,
				IsBook:                      false,
				IsChapter:                   false,
				IsStaticHTML:                false,
				IsOpenAccess:                false,
				IsEphemera:                  false,
				HTMLAbstractLink:            "/document/8900441/",
				IsSMPTE:                     false,
				IsOUP:                       false,
				IsSAE:                       false,
				IsNow:                       false,
				IsCustomDenial:              false,
				IsNotDynamicOrStatic:        true,
				XploreDocumentType:          str("Conference Publication"),
				ContentTypeDisplay:          str("Conferences"),
				MlTime:                      "PT0.027462S",
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
			mltmda, err := goieeeapi.GetMultimedia(tt.Client, tt.ID)

			if !reflect.DeepEqual(mltmda, tt.ExpectedFigures) {
				t.Errorf("Failed to get expected multimedia: got \"%v\" instead of \"%v\".", mltmda, tt.ExpectedFigures)
			}
			checkErr(err, tt.ExpectedErr, t)
		})
	}
}
