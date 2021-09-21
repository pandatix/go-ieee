package goieeeapi_test

import (
	"encoding/json"
	"net/http"
	"reflect"
	"testing"

	goieeeapi "github.com/pandatix/go-ieee-api"
)

func TestGetFigures(t *testing.T) {
	t.Parallel()

	var tests = map[string]struct {
		Client          goieeeapi.HTTPClient
		ID              int
		ExpectedFigures *goieeeapi.GetFiguresResponse
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
			Client: newFakeHTTPClient(`{"userInfo":{"institute":false,"member":false,"individual":false,"guest":false,"subscribedContent":false,"fileCabinetContent":false,"fileCabinetUser":false,"institutionalFileCabinetUser":false,"showPatentCitations":true,"showGet802Link":false,"showOpenUrlLink":false,"tracked":false,"delegatedAdmin":false,"desktop":false,"isInstitutionDashboardEnabled":false,"isInstitutionProfileEnabled":false,"isRoamingEnabled":false,"isDelegatedAdmin":false,"isMdl":false,"isCwg":false},"figures":[{"id":"fig1","label":"Fig. 1.","graphic":{"href":"mazza1-p4-mazza.tif","small":"mazza1-p4-mazza-small.gif","large":"mazza1-p4-mazza-large.gif","hires":"mazza1-p4-mazza-hires.gif","glance":"mazza1-p4-mazza-glance.gif"},"caption":"\n\n<p>The mask produced by deep learning approaches are clean compared to the baseline.</p>\n\n","part":"1"},{"id":"fig2","label":"Fig. 2.","graphic":{"href":"mazza2-p4-mazza.tif","small":"mazza2-p4-mazza-small.gif","large":"mazza2-p4-mazza-large.gif","hires":"mazza2-p4-mazza-hires.gif","glance":"mazza2-p4-mazza-glance.gif"},"caption":"\n\n<p>False positives and oversmoothing for DL solutions and faliure of the baseline.</p>\n\n","part":"1"},{"id":"fig3","label":"Fig. 3.","graphic":{"href":"mazza3-p4-mazza.tif","small":"mazza3-p4-mazza-small.gif","large":"mazza3-p4-mazza-large.gif","hires":"mazza3-p4-mazza-hires.gif","glance":"mazza3-p4-mazza-glance.gif"},"caption":"\n\n<p>False negatives for all.</p>\n\n","part":"1"}],"articleNumber":"8900441","getProgramTermsAccepted":false,"formulaStrippedArticleTitle":"Deep Learning Solutions for Tandem-X-Based Forest Classification","pubLink":"/xpl/conhome/8891871/proceeding","allowComments":false,"issueLink":"/xpl/tocresult.jsp?isnumber=null","isReadingRoomArticle":false,"isGetArticle":false,"isGetAddressInfoCaptured":false,"isMarketingOptIn":false,"publisher":"IEEE","xploreDocumentType":"Conference Publication","isPromo":false,"isNotDynamicOrStatic":true,"htmlAbstractLink":"/document/8900441/","isCustomDenial":false,"isSAE":false,"isDynamicHtml":false,"isFreeDocument":false,"displayDocTitle":"Deep Learning Solutions for Tandem-X-Based Forest Classification","isStandard":false,"isSMPTE":false,"isOUP":false,"isNow":false,"isProduct":false,"isMorganClaypool":false,"isJournal":false,"isBook":false,"isBookWithoutChapters":false,"isOpenAccess":false,"isEphemera":false,"isConference":true,"isChapter":false,"isStaticHtml":false,"isEarlyAccess":false,"persistentLink":"https://ieeexplore.ieee.org/servlet/opac?punumber=8891871","contentTypeDisplay":"Conferences","mlTime":"PT0.025641S","lastupdate":"2021-08-21","mediaPath":"/mediastore_new/IEEE/content/media/8891871/8897702/8900441","title":"Deep Learning Solutions for Tandem-X-Based Forest Classification","contentType":"conferences","publicationNumber":"8891871"}`, http.StatusOK, nil),
			ID:     0,
			ExpectedFigures: &goieeeapi.GetFiguresResponse{
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
				Figures: &[]goieeeapi.Figure{
					{
						ID:    "fig1",
						Label: "Fig. 1.",
						Graphic: goieeeapi.Graphic{
							Href:   "mazza1-p4-mazza.tif",
							Small:  "mazza1-p4-mazza-small.gif",
							Large:  "mazza1-p4-mazza-large.gif",
							Hires:  "mazza1-p4-mazza-hires.gif",
							Glance: "mazza1-p4-mazza-glance.gif",
						},
						Caption: "\n\n<p>The mask produced by deep learning approaches are clean compared to the baseline.</p>\n\n",
						Part:    str("1"),
					}, {
						ID:    "fig2",
						Label: "Fig. 2.",
						Graphic: goieeeapi.Graphic{
							Href:   "mazza2-p4-mazza.tif",
							Small:  "mazza2-p4-mazza-small.gif",
							Large:  "mazza2-p4-mazza-large.gif",
							Hires:  "mazza2-p4-mazza-hires.gif",
							Glance: "mazza2-p4-mazza-glance.gif",
						},
						Caption: "\n\n<p>False positives and oversmoothing for DL solutions and faliure of the baseline.</p>\n\n",
						Part:    str("1"),
					}, {
						ID:    "fig3",
						Label: "Fig. 3.",
						Graphic: goieeeapi.Graphic{
							Href:   "mazza3-p4-mazza.tif",
							Small:  "mazza3-p4-mazza-small.gif",
							Large:  "mazza3-p4-mazza-large.gif",
							Hires:  "mazza3-p4-mazza-hires.gif",
							Glance: "mazza3-p4-mazza-glance.gif",
						},
						Caption: "\n\n<p>False negatives for all.</p>\n\n",
						Part:    str("1"),
					},
				},
				ArticleNumber:               "8900441",
				GetProgramTermsAccepted:     false,
				FormulaStrippedArticleTitle: str("Deep Learning Solutions for Tandem-X-Based Forest Classification"),
				PubLink:                     str("/xpl/conhome/8891871/proceeding"),
				AllowComments:               false,
				IssueLink:                   "/xpl/tocresult.jsp?isnumber=null",
				IsReadingRoomArticle:        false,
				IsGetArticle:                false,
				IsGetAddressInfoCaptured:    false,
				IsMarketingOptIn:            false,
				Publisher:                   str("IEEE"),
				XploreDocumentType:          str("Conference Publication"),
				IsPromo:                     false,
				IsNotDynamicOrStatic:        true,
				HTMLAbstractLink:            "/document/8900441/",
				IsCustomDenial:              false,
				IsSAE:                       false,
				IsDynamicHTML:               false,
				IsFreeDocument:              false,
				DisplayDocTitle:             str("Deep Learning Solutions for Tandem-X-Based Forest Classification"),
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
				PersistentLink:              str("https://ieeexplore.ieee.org/servlet/opac?punumber=8891871"),
				Title:                       str("Deep Learning Solutions for Tandem-X-Based Forest Classification"),
				ContentTypeDisplay:          str("Conferences"),
				MlTime:                      "PT0.025641S",
				LastUpdate:                  str("2021-08-21"),
				MediaPath:                   str("/mediastore_new/IEEE/content/media/8891871/8897702/8900441"),
				ContentType:                 str("conferences"),
				PublicationNumber:           str("8891871"),
			},
			ExpectedErr: nil,
		},
	}

	for testname, tt := range tests {
		t.Run(testname, func(t *testing.T) {
			cts, err := goieeeapi.GetFigures(tt.Client, tt.ID)

			if !reflect.DeepEqual(cts, tt.ExpectedFigures) {
				t.Errorf("Failed to get expected figures: got \"%v\" instead of \"%v\".", cts, tt.ExpectedFigures)
			}
			checkErr(err, tt.ExpectedErr, t)
		})
	}
}
