package goieeeapi_test

import (
	"encoding/json"
	"net/http"
	"reflect"
	"testing"

	goieeeapi "github.com/pandatix/go-ieee-api"
)

func TestGetAuthors(t *testing.T) {
	t.Parallel()

	var tests = map[string]struct {
		Client          goieeeapi.HTTPClient
		ID              int
		ExpectedAuthors *goieeeapi.GetAuthorsResponse
		ExpectedErr     error
	}{
		"nil-client": {
			Client:          nil,
			ID:              0,
			ExpectedAuthors: nil,
			ExpectedErr:     goieeeapi.ErrNilClient,
		},
		"failing-client": {
			Client:          newFakeHTTPClient(``, 0, errFake),
			ID:              0,
			ExpectedAuthors: nil,
			ExpectedErr:     errFake,
		},
		"unexpected-statuscode": {
			Client:          newFakeHTTPClient(``, 0, nil),
			ID:              0,
			ExpectedAuthors: nil,
			ExpectedErr: &goieeeapi.ErrUnexpectedStatus{
				StatusCode: 0,
				Body:       []byte(``),
			},
		},
		"failing-unmarshal": {
			Client:          newFakeHTTPClient(`{[}]`, http.StatusOK, nil),
			ID:              0,
			ExpectedAuthors: nil,
			ExpectedErr: &json.SyntaxError{
				Offset: 2,
			},
		},
		"valid-call": {
			Client: newFakeHTTPClient(`{"userInfo":{"institute":false,"member":false,"individual":false,"guest":false,"subscribedContent":false,"fileCabinetContent":false,"fileCabinetUser":false,"institutionalFileCabinetUser":false,"showPatentCitations":true,"showGet802Link":false,"showOpenUrlLink":false,"tracked":false,"delegatedAdmin":false,"desktop":false,"isInstitutionDashboardEnabled":false,"isInstitutionProfileEnabled":false,"isRoamingEnabled":false,"isDelegatedAdmin":false,"isMdl":false,"isCwg":false},"authors":[{"name":"Antonio Mazza","affiliation":["DIETI, University Federico II, Naples, Italy"],"ringgoldId":"9307","firstName":"Antonio","lastName":"Mazza","id":"37086504809"},{"name":"Francescopaolo Sica","affiliation":["Microwaves and Radar Institute, DLR, Weßling, Germany"],"firstName":"Francescopaolo","lastName":"Sica","id":"38529295500"}],"articleNumber":"8900441","getProgramTermsAccepted":false,"formulaStrippedArticleTitle":"Deep Learning Solutions for Tandem-X-Based Forest Classification","pubLink":"/xpl/conhome/8891871/proceeding","allowComments":false,"issueLink":"/xpl/tocresult.jsp?isnumber=null","isReadingRoomArticle":false,"isGetArticle":false,"isGetAddressInfoCaptured":false,"isMarketingOptIn":false,"publisher":"IEEE","xploreDocumentType":"Conference Publication","isPromo":false,"isNotDynamicOrStatic":true,"htmlAbstractLink":"/document/8900441/","isCustomDenial":false,"isSAE":false,"isDynamicHtml":false,"isFreeDocument":false,"displayDocTitle":"Deep Learning Solutions for Tandem-X-Based Forest Classification","isStandard":false,"isSMPTE":false,"isOUP":false,"isNow":false,"isProduct":false,"isMorganClaypool":false,"isJournal":false,"isBook":false,"isBookWithoutChapters":false,"isOpenAccess":false,"isEphemera":false,"isConference":true,"isChapter":false,"isStaticHtml":false,"isEarlyAccess":false,"persistentLink":"https://ieeexplore.ieee.org/servlet/opac?punumber=8891871","contentTypeDisplay":"Conferences","mlTime":"PT0.004696S","lastupdate":"2021-08-21","mediaPath":"/mediastore_new/IEEE/content/media/8891871/8897702/8900441","title":"Deep Learning Solutions for Tandem-X-Based Forest Classification","contentType":"conferences","publicationNumber":"8891871"}`, http.StatusOK, nil),
			ID:     0,
			ExpectedAuthors: &goieeeapi.GetAuthorsResponse{
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
				Authors: []goieeeapi.Author{
					{
						Name:        "Antonio Mazza",
						Affiliation: []string{"DIETI, University Federico II, Naples, Italy"},
						RingGoldID:  str("9307"),
						FirstName:   "Antonio",
						LastName:    "Mazza",
						ID:          "37086504809",
					}, {
						Name:        "Francescopaolo Sica",
						Affiliation: []string{"Microwaves and Radar Institute, DLR, Weßling, Germany"},
						FirstName:   "Francescopaolo",
						LastName:    "Sica",
						ID:          "38529295500",
					},
				},
				ArticleNumber:               "8900441",
				GetProgramTermsAccepted:     false,
				FormulaStrippedArticleTitle: "Deep Learning Solutions for Tandem-X-Based Forest Classification",
				PubLink:                     "/xpl/conhome/8891871/proceeding",
				AllowComments:               false,
				IssueLink:                   "/xpl/tocresult.jsp?isnumber=null",
				IsReadingRoomArticle:        false,
				IsGetArticle:                false,
				IsGetAddressInfoCaptured:    false,
				IsMarketingOptIn:            false,
				Publisher:                   "IEEE",
				XploreDocumentType:          "Conference Publication",
				IsPromo:                     false,
				IsNotDynamicOrStatic:        true,
				HTMLAbstractLink:            "/document/8900441/",
				IsCustomDenial:              false,
				IsSAE:                       false,
				IsDynamicHTML:               false,
				IsFreeDocument:              false,
				DisplayDocTitle:             "Deep Learning Solutions for Tandem-X-Based Forest Classification",
				IsStandard:                  false,
				IsSMPTE:                     false,
				IsOUP:                       false,
				IsNow:                       false,
				IsProduct:                   false,
				IsMorganClaypool:            false,
				IsJournal:                   false,
				IsBook:                      false,
				IsBookWithoutChapters:       false,
				IsOpenAccess:                false,
				IsEphemera:                  false,
				IsConference:                true,
				IsChapter:                   false,
				IsStaticHTML:                false,
				IsEarlyAccess:               false,
				PersistentLink:              "https://ieeexplore.ieee.org/servlet/opac?punumber=8891871",
				Title:                       "Deep Learning Solutions for Tandem-X-Based Forest Classification",
				ContentTypeDisplay:          "Conferences",
				MlTime:                      "PT0.004696S",
				LastUpdate:                  "2021-08-21",
				MediaPath:                   "/mediastore_new/IEEE/content/media/8891871/8897702/8900441",
				ContentType:                 "conferences",
				PublicationNumber:           "8891871",
			},
			ExpectedErr: nil,
		},
	}

	for testname, tt := range tests {
		t.Run(testname, func(t *testing.T) {
			authors, err := goieeeapi.GetAuthors(tt.Client, tt.ID)

			if !reflect.DeepEqual(authors, tt.ExpectedAuthors) {
				t.Errorf("Failed to get expected authors: got \"%v\" instead of \"%v\".", authors, tt.ExpectedAuthors)
			}
			checkErr(err, tt.ExpectedErr, t)
		})
	}
}
