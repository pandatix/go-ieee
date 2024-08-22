package api

import (
	"path"
	"strconv"
)

// GetDocumentDisclaimer fetch the /document/<id>/disclaimer REST endpoint and
// returns the document disclaimer information.
func (client *IEEEClient) GetDocumentDisclaimer(id int, opts ...Option) (*GetDisclaimerResponse, error) {
	resp := &GetDisclaimerResponse{}
	err := client.get(path.Join("document", strconv.Itoa(id), "disclaimer"), nil, resp, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

type GetDisclaimerResponse struct {
	UserInfo                    UserInfo `json:"userInfo"`
	ArticleNumber               string   `json:"articleNumber"`
	GetProgramTermsAccepted     bool     `json:"getProgramTermsAccepted"`
	PubLink                     *string  `json:"pubLink,omitempty"`
	AllowComments               bool     `json:"allowComments"`
	IssueLink                   string   `json:"issueLink"`
	FormulaStrippedArticleTitle *string  `json:"formulaStrippedArticleTitle,omitempty"`
	IsGetAddressInfoCaptured    bool     `json:"isGetAddressInfoCaptured"`
	IsMarketingOptIn            bool     `json:"isMarketingOptIn"`
	Publisher                   *string  `json:"publisher,omitempty"`
	IsPromo                     bool     `json:"isPromo"`
	IsSMPTE                     bool     `json:"isSMPTE"`
	IsOUP                       bool     `json:"isOUP"`
	IsSAE                       bool     `json:"isSAE"`
	IsNow                       bool     `json:"isNow"`
	IsCustomDenial              bool     `json:"isCustomDenial"`
	DisplayDocTitle             *string  `json:"displayDocTitle,omitempty"`
	IsStandard                  bool     `json:"isStandard"`
	HTMLAbstractLink            string   `json:"htmlAbstractLink"`
	IsProduct                   bool     `json:"isProduct"`
	IsOpenAccess                bool     `json:"isOpenAccess"`
	IsEphemera                  bool     `json:"isEphemera"`
	IsConference                bool     `json:"isConference"`
	IsEarlyAccess               bool     `json:"isEarlyAccess"`
	IsJournal                   bool     `json:"isJournal"`
	IsBook                      bool     `json:"isBook"`
	IsBookWithoutChapters       bool     `json:"isBookWithoutChapters"`
	IsChapter                   bool     `json:"isChapter"`
	IsGiveaway                  bool     `json:"isGiveaway"`
	IsLatestStandard            bool     `json:"isLatestStandard"`
	IsOnlineOnly                bool     `json:"isOnlineOnly"`
	IsSpringer                  bool     `json:"isSpringer"`
	IsTranslation               bool     `json:"isTranslation"`
	PersistentLink              *string  `json:"persistentLink,omitempty"`
	XploreDocumentType          *string  `json:"xploreDocumentType,omitempty"`
	IsFreeDocument              bool     `json:"isFreeDocument"`
	IsDynamicHTML               bool     `json:"isDynamicHtml"`
	ContentTypeDisplay          *string  `json:"contentTypeDisplay,omitempty"`
	MlTime                      string   `json:"mlTime"`
	LastUpdate                  *string  `json:"lastupdate,omitempty"`
	MediaPath                   *string  `json:"mediaPath,omitempty"`
	Title                       *string  `json:"title,omitempty"`
	ContentType                 *string  `json:"contentType,omitempty"`
	PublicationNumber           *string  `json:"publicationNumber,omitempty"`
	HTMLFlagLegacy              *string  `json:"htmlFlag,omitempty"`
	HasStandardVersions         bool     `json:"hasStandardVersions"`
}
