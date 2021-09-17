package goieeeapi_test

import (
	"encoding/json"
	"net/http"
	"reflect"
	"testing"

	goieeeapi "github.com/pandatix/go-ieee-api"
)

func TestGetAbstract(t *testing.T) {
	t.Parallel()

	var tests = map[string]struct {
		Client           goieeeapi.HTTPClient
		ID               int
		ExpectedAbstract *goieeeapi.GetAbstractResponse
		ExpectedErr      error
	}{
		"nil-client": {
			Client:           nil,
			ID:               0,
			ExpectedAbstract: nil,
			ExpectedErr:      goieeeapi.ErrNilClient,
		},
		"failing-client": {
			Client:           newFakeHTTPClient(``, 0, errFake),
			ID:               0,
			ExpectedAbstract: nil,
			ExpectedErr:      errFake,
		},
		"unexpected-statuscode": {
			Client:           newFakeHTTPClient(``, 0, nil),
			ID:               0,
			ExpectedAbstract: nil,
			ExpectedErr: &goieeeapi.ErrUnexpectedStatus{
				StatusCode: 0,
				Body:       []byte(``),
			},
		},
		"failing-unmarshal": {
			Client:           newFakeHTTPClient(`{[}]`, http.StatusOK, nil),
			ID:               0,
			ExpectedAbstract: nil,
			ExpectedErr: &json.SyntaxError{
				Offset: 2,
			},
		},
		"valid-call": {
			Client: newFakeHTTPClient(`{"userInfo":{"institute":false,"member":false,"individual":false,"guest":false,"subscribedContent":false,"fileCabinetContent":false,"fileCabinetUser":false,"institutionalFileCabinetUser":false,"showPatentCitations":true,"showGet802Link":false,"showOpenUrlLink":false,"tracked":false,"delegatedAdmin":false,"desktop":false,"isInstitutionDashboardEnabled":false,"isInstitutionProfileEnabled":false,"isRoamingEnabled":false,"isDelegatedAdmin":false,"isMdl":false,"isCwg":false},"authors":[{"name":"Marco Wagler","affiliation":["Institute of Power Transmission Systems, Technische Universität, München, Munich, Germany"],"ringgoldId":"9184","firstName":"Marco","lastName":"Wagler","id":"37085668708"},{"name":"Rolf Witzmann","affiliation":["Institute of Power Transmission Systems, Technische Universität, München, Munich, Germany"],"ringgoldId":"9184","firstName":"Rolf","lastName":"Witzmann","id":"37568408200"}],"isbn":[{"format":"Electronic ISBN","value":"978-1-5090-2157-4","isbnType":""},{"format":"Print on Demand(PoD) ISBN","value":"978-1-5090-2158-1","isbnType":""}],"issn":[{"format":"Electronic ISSN","value":"2160-8563"}],"articleNumber":"7520000","getProgramTermsAccepted":false,"sections":{"abstract":"true","authors":"true","figures":"true","multimedia":"false","references":"true","citedby":"false","keywords":"true","definitions":"false","algorithm":"false","dataset":"false","cadmore":"false","footnotes":"false","disclaimer":"false","relatedContent":"false"},"pdfUrl":"/stamp/stamp.jsp?tp=&arnumber=7520000","keywords":[{"type":"IEEE Keywords","kwd":["Sensitivity analysis","Low voltage","Load flow","Reactive power","Loading","Jacobian matrices"]},{"type":"INSPEC: Controlled Indexing","kwd":["distribution networks","load flow","power grids","power transformers"]},{"type":"INSPEC: Non-Controlled Indexing","kwd":["sensitivity analysis","flexible distribution grid","real low voltage grids","load flow modification","node voltages","line currents","transformer loadings"]},{"type":"Author Keywords ","kwd":["Sensitivity","Smart Grid","Distribution Grid","Intelligent Grid Operation"]}],"pubLink":"/xpl/conhome/7513513/proceeding","allowComments":false,"abstract":"In this paper a sensitivity analysis for several real low voltage grids is performed. The term sensitivity hereby expresses the impact of a load flow modification in a particular node with respect to other node voltages, line currents and transformer loadings. The question of how many nodes have a sufficient effect on problem areas as well as the size of such a sensitive region is answered for low voltage grids.","doi":"10.1109/TDC.2016.7520000","doiLink":"https://doi.org/10.1109/TDC.2016.7520000","rightsLink":"http://s100.copyright.com/AppDispatchServlet?publisherName=ieee&publication=proceedings&title=Sensitivity+analysis+-+a+key+element+for+the+operation+of+a+flexible+distribution+grid&isbn=978-1-5090-2157-4&publicationDate=May+2016&author=Marco+Wagler&ContentID=10.1109/TDC.2016.7520000&orderBeanReset=true&startPage=1&endPage=5&proceedingName=2016+IEEE%2FPES+Transmission+and+Distribution+Conference+and+Exposition+%28T%26D%29","startPage":"1","endPage":"5","publicationTitle":"2016 IEEE/PES Transmission and Distribution Conference and Exposition (T&D)","displayPublicationTitle":"2016 IEEE/PES Transmission and Distribution Conference and Exposition (T&D)","pdfPath":"/iel7/7513513/7519846/07520000.pdf","issueLink":"/xpl/tocresult.jsp?isnumber=7519846","formulaStrippedArticleTitle":"Sensitivity analysis - a key element for the operation of a flexible distribution grid","isReadingRoomArticle":false,"isGetArticle":false,"isGetAddressInfoCaptured":false,"isMarketingOptIn":false,"pubTopics":[{"name":"Engineering Profession"},{"name":"General Topics for Engineers"},{"name":"Power, Energy and Industry Applications"}],"publisher":"IEEE","isDynamicHtml":true,"isFreeDocument":false,"htmlAbstractLink":"/document/7520000/","isCustomDenial":false,"isSMPTE":false,"isOUP":false,"conferenceDate":"3-5 May 2016","isNotDynamicOrStatic":false,"chronOrPublicationDate":"May 2016","isNow":false,"isSAE":false,"isPromo":false,"displayDocTitle":"Sensitivity analysis - a key element for the operation of a flexible distribution grid","isStandard":false,"isConference":true,"isProduct":false,"isMorganClaypool":false,"isJournal":false,"isBook":false,"isBookWithoutChapters":false,"isEarlyAccess":false,"accessionNumber":"16160184","isOpenAccess":false,"publicationDate":"May 2016","isEphemera":false,"htmlLink":"/document/7520000/","persistentLink":"https://ieeexplore.ieee.org/servlet/opac?punumber=7513513","isChapter":false,"isStaticHtml":true,"xploreDocumentType":"Conference Publication","startPage":"1","openAccessFlag":"F","ephemeraFlag":"false","title":"Sensitivity analysis - a key element for the operation of a flexible distribution grid","confLoc":"Dallas, TX, USA","accessionNumber":"16160184","html_flag":"true","ml_html_flag":"true","sourcePdf":"TD2016-000106.pdf","mlTime":"PT0.030537S","xplore-pub-id":"7513513","pdfPath":"/iel7/7513513/7519846/07520000.pdf","isNumber":"7519846","rightsLinkFlag":"1","contentType":"conferences","publicationDate":"May 2016","publicationNumber":"7513513","xplore-issue":"7519846","articleId":"7520000","publicationTitle":"2016 IEEE/PES Transmission and Distribution Conference and Exposition (T&D)","sections":{"abstract":"true","authors":"true","figures":"true","multimedia":"false","references":"true","citedby":"false","keywords":"true","definitions":"false","algorithm":"false","dataset":"false","cadmore":"false","footnotes":"false","disclaimer":"false","relatedContent":"false"},"contentTypeDisplay":"Conferences","conferenceDate":"3-5 May 2016","subType":"IEEE Conference","_value":"IEEE","lastupdate":"2021-08-15","mediaPath":"/mediastore_new/IEEE/content/media/7513513/7519846/7520000","endPage":"5","displayPublicationTitle":"2016 IEEE/PES Transmission and Distribution Conference and Exposition (T&D)","doi":"10.1109/TDC.2016.7520000"}`, http.StatusOK, nil),
			ID:     0,
			ExpectedAbstract: &goieeeapi.GetAbstractResponse{
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
				Authors: &[]goieeeapi.Author{
					{
						Name: "Marco Wagler",
						Affiliation: &[]string{
							"Institute of Power Transmission Systems, Technische Universität, München, Munich, Germany",
						},
						RingGoldID: str("9184"),
						FirstName:  str("Marco"),
						LastName:   "Wagler",
						ID:         str("37085668708"),
					}, {
						Name: "Rolf Witzmann",
						Affiliation: &[]string{
							"Institute of Power Transmission Systems, Technische Universität, München, Munich, Germany",
						},
						RingGoldID: str("9184"),
						FirstName:  str("Rolf"),
						LastName:   "Witzmann",
						ID:         str("37568408200"),
					},
				},
				ISBN: &[]goieeeapi.ISBN{
					{
						Format:   "Electronic ISBN",
						Value:    "978-1-5090-2157-4",
						ISBNType: "",
					}, {
						Format:   "Print on Demand(PoD) ISBN",
						Value:    "978-1-5090-2158-1",
						ISBNType: "",
					},
				},
				ISSN: &[]goieeeapi.ISSN{
					{
						Format: "Electronic ISSN",
						Value:  "2160-8563",
					},
				},
				ArticleNumber:           "7520000",
				GetProgramTermsAccepted: false,
				Sections: &goieeeapi.Sections{
					Abstract:       "true",
					Authors:        "true",
					Figures:        "true",
					Multimedia:     "false",
					References:     "true",
					CitedBy:        "false",
					Keywords:       "true",
					Definitions:    "false",
					Algorithm:      "false",
					Dataset:        "false",
					Cadmore:        "false",
					Footnotes:      "false",
					Disclaimer:     "false",
					RelatedContent: "false",
				},
				PdfURL: str("/stamp/stamp.jsp?tp=&arnumber=7520000"),
				Keywords: &[]goieeeapi.Keyword{
					{
						Type: str("IEEE Keywords"),
						Kwd: []string{
							"Sensitivity analysis",
							"Low voltage",
							"Load flow",
							"Reactive power",
							"Loading",
							"Jacobian matrices",
						},
					}, {
						Type: str("INSPEC: Controlled Indexing"),
						Kwd: []string{
							"distribution networks",
							"load flow",
							"power grids",
							"power transformers",
						},
					}, {
						Type: str("INSPEC: Non-Controlled Indexing"),
						Kwd: []string{
							"sensitivity analysis",
							"flexible distribution grid",
							"real low voltage grids",
							"load flow modification",
							"node voltages",
							"line currents",
							"transformer loadings",
						},
					}, {
						Type: str("Author Keywords "),
						Kwd: []string{
							"Sensitivity",
							"Smart Grid",
							"Distribution Grid",
							"Intelligent Grid Operation",
						},
					},
				},
				PubLink:                     str("/xpl/conhome/7513513/proceeding"),
				AllowComments:               false,
				Abstract:                    str("In this paper a sensitivity analysis for several real low voltage grids is performed. The term sensitivity hereby expresses the impact of a load flow modification in a particular node with respect to other node voltages, line currents and transformer loadings. The question of how many nodes have a sufficient effect on problem areas as well as the size of such a sensitive region is answered for low voltage grids."),
				DOI:                         str("10.1109/TDC.2016.7520000"),
				DOILink:                     str("https://doi.org/10.1109/TDC.2016.7520000"),
				RightsLink:                  str("http://s100.copyright.com/AppDispatchServlet?publisherName=ieee&publication=proceedings&title=Sensitivity+analysis+-+a+key+element+for+the+operation+of+a+flexible+distribution+grid&isbn=978-1-5090-2157-4&publicationDate=May+2016&author=Marco+Wagler&ContentID=10.1109/TDC.2016.7520000&orderBeanReset=true&startPage=1&endPage=5&proceedingName=2016+IEEE%2FPES+Transmission+and+Distribution+Conference+and+Exposition+%28T%26D%29"),
				StartPage:                   str("1"),
				EndPage:                     str("5"),
				PublicationTitle:            str("2016 IEEE/PES Transmission and Distribution Conference and Exposition (T&D)"),
				DisplayPublicationTitle:     str("2016 IEEE/PES Transmission and Distribution Conference and Exposition (T&D)"),
				PdfPath:                     str("/iel7/7513513/7519846/07520000.pdf"),
				IssueLink:                   "/xpl/tocresult.jsp?isnumber=7519846",
				FormulaStrippedArticleTitle: str("Sensitivity analysis - a key element for the operation of a flexible distribution grid"),
				IsReadingRoomArticle:        false,
				IsGetArticle:                false,
				IsGetAddressInfoCaptured:    false,
				IsMarketingOptIn:            false,
				PubTopics: &[]goieeeapi.PubTopic{
					{
						Name: "Engineering Profession",
					}, {
						Name: "General Topics for Engineers",
					}, {
						Name: "Power, Energy and Industry Applications",
					},
				},
				Publisher:              str("IEEE"),
				IsDynamicHTML:          true,
				IsFreeDocument:         false,
				HTMLAbstractLink:       "/document/7520000/",
				IsCustomDenial:         false,
				IsSMPTE:                false,
				IsOUP:                  false,
				ConferenceDate:         str("3-5 May 2016"),
				IsNotDynamicOrStatic:   false,
				ChronOrPublicationDate: str("May 2016"),
				IsNow:                  false,
				IsSAE:                  false,
				IsPromo:                false,
				DisplayDocTitle:        str("Sensitivity analysis - a key element for the operation of a flexible distribution grid"),
				IsStandard:             false,
				IsConference:           true,
				IsProduct:              false,
				IsMorganClaypool:       false,
				IsJournal:              false,
				IsBook:                 false,
				IsBookWithoutChapters:  false,
				IsEarlyAccess:          false,
				AccessionNumber:        str("16160184"),
				IsOpenAccess:           false,
				PublicationDate:        str("May 2016"),
				IsEphemera:             false,
				HTMLLink:               str("/document/7520000/"),
				PersistentLink:         str("https://ieeexplore.ieee.org/servlet/opac?punumber=7513513"),
				IsChapter:              false,
				IsStaticHTML:           true,
				XploreDocumentType:     str("Conference Publication"),
				OpenAccessFlag:         str("F"),
				EphemeraFlag:           str("false"),
				Title:                  str("Sensitivity analysis - a key element for the operation of a flexible distribution grid"),
				ConfLoc:                str("Dallas, TX, USA"),
				HTMLFlag:               str("true"),
				MlHTMLFlag:             str("true"),
				SourcePdf:              str("TD2016-000106.pdf"),
				MlTime:                 "PT0.030537S",
				XplorePubID:            str("7513513"),
				IsNumber:               str("7519846"),
				RightsLinkFlag:         str("1"),
				ContentType:            str("conferences"),
				PublicationNumber:      str("7513513"),
				XploreIssue:            str("7519846"),
				ArticleID:              str("7520000"),
				ContentTypeDisplay:     str("Conferences"),
				SubType:                str("IEEE Conference"),
				Value:                  str("IEEE"),
				LastUpdate:             str("2021-08-15"),
				MediaPath:              str("/mediastore_new/IEEE/content/media/7513513/7519846/7520000"),
			},
			ExpectedErr: nil,
		},
	}

	for testname, tt := range tests {
		t.Run(testname, func(t *testing.T) {
			abs, err := goieeeapi.GetAbstract(tt.Client, tt.ID)

			if !reflect.DeepEqual(abs, tt.ExpectedAbstract) {
				t.Errorf("Failed to get expected abs: got \"%v\" instead of \"%v\".", abs, tt.ExpectedAbstract)
			}
			checkErr(err, tt.ExpectedErr, t)
		})
	}
}
