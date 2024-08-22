package api

import (
	"path"
	"strconv"
)

// GetDocumentMetrics fetch the /document/<id>/metrics REST endpoint and
// returns the document metrics information.
func (client *IEEEClient) GetDocumentMetrics(id int, opts ...Option) (*GetMetricsResponse, error) {
	resp := &GetMetricsResponse{}
	err := client.get(path.Join("document", strconv.Itoa(id), "metrics"), nil, resp, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

type GetMetricsResponse struct {
	ArticleNumber            string  `json:"articleNumber"`
	Metrics                  Metrics `json:"metrics"`
	GetProgramTermsAccepted  bool    `json:"getProgramTermsAccepted"`
	AllowComments            bool    `json:"allowComments"`
	IssueLink                string  `json:"issueLink"`
	IsGetAddressInfoCaptured bool    `json:"isGetAddressInfoCaptured"`
	IsMarketingOptIn         bool    `json:"isMarketingOptIn"`
	IsPromo                  bool    `json:"isPromo"`
	HTMLAbstractLink         string  `json:"htmlAbstractLink"`
	IsCustomDenial           bool    `json:"isCustomDenial"`
	IsSAE                    bool    `json:"isSAE"`
	IsDynamicHTML            bool    `json:"isDynamicHtml"`
	IsFreeDocument           bool    `json:"isFreeDocument"`
	IsStandard               bool    `json:"isStandard"`
	IsSMPTE                  bool    `json:"isSMPTE"`
	IsOUP                    bool    `json:"isOUP"`
	IsNow                    bool    `json:"isNow"`
	IsProduct                bool    `json:"isProduct"`
	IsJournal                bool    `json:"isJournal"`
	IsBook                   bool    `json:"isBook"`
	IsBookWithoutChapters    bool    `json:"isBookWithoutChapters"`
	IsOpenAccess             bool    `json:"isOpenAccess"`
	IsEphemera               bool    `json:"isEphemera"`
	IsConference             bool    `json:"isConference"`
	IsChapter                bool    `json:"isChapter"`
	IsEarlyAccess            bool    `json:"isEarlyAccess"`
	HasStandardVersions      bool    `json:"hasStandardVersions"`
	IsGiveaway               bool    `json:"isGiveaway"`
	IsLatestStandard         bool    `json:"isLatestStandard"`
	IsOnlineOnly             bool    `json:"isOnlineOnly"`
	IsSpringer               bool    `json:"isSpringer"`
	IsTranslation            bool    `json:"isTranslation"`
}

type Metrics struct {
	CitationCountPaper  int      `json:"citationCountPaper"`
	CitationCountPatent int      `json:"citationCountPatent"`
	TotalDownloads      int      `json:"totalDownloads"`
	Biblio              []Biblio `json:"biblio"`
	DOI                 *string  `json:"doi,omitempty"`
	WosCount            *string  `json:"wos_count,omitempty"`
	WosURL              *string  `json:"wos_url,omitempty"`
	WosCitationCount    int      `json:"wosCitationCount"`
	ScopusCount         *string  `json:"scopus_count,omitempty"`
	ScopusURL           *string  `json:"scopus_url,omitempty"`
}

type Biblio struct {
	Year                     string `json:"year"`
	BestMonthInYear          int    `json:"bestMonthInYear"`
	BestMonthInYearString    string `json:"bestMonthInYearString"`
	YearToDateDownloads      int    `json:"yearToDateDownloads"`
	TotalArticleDownloads    int    `json:"totalArticleDownloads"`
	TotalArticleDownadsSince string `json:"totalArticleDownloadsSince"`
	Jan                      string `json:"Jan"`
	Feb                      string `json:"Feb"`
	Mar                      string `json:"Mar"`
	Apr                      string `json:"Apr"`
	May                      string `json:"May"`
	Jun                      string `json:"Jun"`
	Jul                      string `json:"Jul"`
	Aug                      string `json:"Aug"`
	Sep                      string `json:"Sep"`
	Oct                      string `json:"Oct"`
	Nov                      string `json:"Nov"`
	Dec                      string `json:"Dec"`
}
