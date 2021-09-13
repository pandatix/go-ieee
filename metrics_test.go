package goieeeapi_test

import (
	"encoding/json"
	"net/http"
	"reflect"
	"testing"

	goieeeapi "github.com/pandatix/go-ieee-api"
)

func TestGetMetrics(t *testing.T) {
	t.Parallel()

	var tests = map[string]struct {
		Client          goieeeapi.HTTPClient
		ID              int
		ExpectedMetrics *goieeeapi.GetMetricsResponse
		ExpectedErr     error
	}{
		"nil-client": {
			Client:          nil,
			ID:              0,
			ExpectedMetrics: nil,
			ExpectedErr:     goieeeapi.ErrNilClient,
		},
		"failing-client": {
			Client:          newFakeHTTPClient(``, 0, errFake),
			ID:              0,
			ExpectedMetrics: nil,
			ExpectedErr:     errFake,
		},
		"unexpected-statuscode": {
			Client:          newFakeHTTPClient(``, 0, nil),
			ID:              0,
			ExpectedMetrics: nil,
			ExpectedErr: &goieeeapi.ErrUnexpectedStatus{
				StatusCode: 0,
				Body:       []byte(``),
			},
		},
		"failing-unmarshal": {
			Client:          newFakeHTTPClient(`{[}]`, http.StatusOK, nil),
			ID:              0,
			ExpectedMetrics: nil,
			ExpectedErr: &json.SyntaxError{
				Offset: 2,
			},
		},
		"valid-call": {
			Client: newFakeHTTPClient(`{"articleNumber":"8900441","metrics":{"citationCountPaper":0,"citationCountPatent":0,"totalDownloads":54,"biblio":[{"year":"2021","bestMonthInYear":5,"bestMonthInYearString":"Jun","yearToDateDownloads":13,"totalArticleDownloads":54,"totalArticleDownloadsSince":"Nov 2019","Jul":"2","Oct":"-","Feb":"-","Apr":"1","Jun":"5","Aug":"-","Dec":"-","May":"3","Nov":"-","Jan":"2","Mar":"-","Sep":"-"},{"year":"2020","bestMonthInYear":5,"bestMonthInYearString":"Jun","yearToDateDownloads":27,"totalArticleDownloads":54,"totalArticleDownloadsSince":"Nov 2019","Jul":"3","Oct":"2","Feb":"3","Apr":"1","Jun":"6","Aug":"1","Dec":"-","May":"3","Nov":"2","Jan":"2","Mar":"1","Sep":"3"},{"year":"2019","bestMonthInYear":10,"bestMonthInYearString":"Nov","yearToDateDownloads":14,"totalArticleDownloads":54,"totalArticleDownloadsSince":"Nov 2019","Jul":"-","Oct":"-","Feb":"-","Apr":"-","Jun":"-","Aug":"-","Dec":"5","May":"-","Nov":"9","Jan":"-","Mar":"-","Sep":"-"}],"doi":"10.1109/IGARSS.2019.8900441"},"getProgramTermsAccepted":false,"allowComments":false,"issueLink":"/xpl/tocresult.jsp?isnumber=null","isReadingRoomArticle":false,"isGetArticle":false,"isGetAddressInfoCaptured":false,"isMarketingOptIn":false,"isPromo":false,"isNotDynamicOrStatic":true,"htmlAbstractLink":"/document/8900441/","isCustomDenial":false,"isSAE":false,"isDynamicHtml":false,"isFreeDocument":false,"isStandard":false,"isSMPTE":false,"isOUP":false,"isNow":false,"isProduct":false,"isMorganClaypool":false,"isJournal":false,"isBook":false,"isBookWithoutChapters":false,"isOpenAccess":false,"isEphemera":false,"isConference":false,"isChapter":false,"isStaticHtml":false,"isEarlyAccess":false}`, http.StatusOK, nil),
			ID:     0,
			ExpectedMetrics: &goieeeapi.GetMetricsResponse{
				ArticleNumber: "8900441",
				Metrics: goieeeapi.Metrics{
					CitationCountPaper:  0,
					CitationCountPatent: 0,
					TotalDownloads:      54,
					Biblio: []goieeeapi.Biblio{
						{
							Year:                     "2021",
							BestMonthInYear:          5,
							BestMonthInYearString:    "Jun",
							YearToDateDownloads:      13,
							TotalArticleDownloads:    54,
							TotalArticleDownadsSince: "Nov 2019",
							Jan:                      "2",
							Feb:                      "-",
							Mar:                      "-",
							Apr:                      "1",
							May:                      "3",
							Jun:                      "5",
							Jul:                      "2",
							Aug:                      "-",
							Sep:                      "-",
							Oct:                      "-",
							Nov:                      "-",
							Dec:                      "-",
						}, {
							Year:                     "2020",
							BestMonthInYear:          5,
							BestMonthInYearString:    "Jun",
							YearToDateDownloads:      27,
							TotalArticleDownloads:    54,
							TotalArticleDownadsSince: "Nov 2019",
							Jan:                      "2",
							Feb:                      "3",
							Mar:                      "1",
							Apr:                      "1",
							May:                      "3",
							Jun:                      "6",
							Jul:                      "3",
							Aug:                      "1",
							Sep:                      "3",
							Oct:                      "2",
							Nov:                      "2",
							Dec:                      "-",
						}, {
							Year:                     "2019",
							BestMonthInYear:          10,
							BestMonthInYearString:    "Nov",
							YearToDateDownloads:      14,
							TotalArticleDownloads:    54,
							TotalArticleDownadsSince: "Nov 2019",
							Jan:                      "-",
							Feb:                      "-",
							Mar:                      "-",
							Apr:                      "-",
							May:                      "-",
							Jun:                      "-",
							Jul:                      "-",
							Aug:                      "-",
							Sep:                      "-",
							Oct:                      "-",
							Nov:                      "9",
							Dec:                      "5",
						},
					},
					DOI: "10.1109/IGARSS.2019.8900441",
				},
				GetProgramTermsAccepted:  false,
				AllowComments:            false,
				IssueLink:                "/xpl/tocresult.jsp?isnumber=null",
				IsReadingRoomArticle:     false,
				IsGetArticle:             false,
				IsGetAddressInfoCaptured: false,
				IsMarketingOptIn:         false,
				IsPromo:                  false,
				IsNotDynamicOrStatic:     true,
				HTMLAbstractLink:         "/document/8900441/",
				IsCustomDenial:           false,
				IsSAE:                    false,
				IsDynamicHTML:            false,
				IsFreeDocument:           false,
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
				IsConference:             false,
				IsChapter:                false,
				IsStaticHTML:             false,
				IsEarlyAccess:            false,
			},
			ExpectedErr: nil,
		},
	}

	for testname, tt := range tests {
		t.Run(testname, func(t *testing.T) {
			metrics, err := goieeeapi.GetMetrics(tt.Client, tt.ID)

			if !reflect.DeepEqual(metrics, tt.ExpectedMetrics) {
				t.Errorf("Failed to get expected metrics: got \"%v\" instead of \"%v\".", metrics, tt.ExpectedMetrics)
			}
			checkErr(err, tt.ExpectedErr, t)
		})
	}
}
