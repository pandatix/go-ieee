package api

import (
	"path"
	"strconv"
)

// GetDocumentKeywords fetch the /document/<id>/keywords REST endpoint and
// returns the document keywords information.
func (client *IEEEClient) GetDocumentKeywords(id int, opts ...Option) (*GetKeywordsResponse, error) {
	resp := &GetKeywordsResponse{}
	err := client.get(path.Join("document", strconv.Itoa(id), "keywords"), nil, resp, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

type GetKeywordsResponse struct {
	UserInfo                    UserInfo  `json:"userInfo"`
	ArticleNumber               string    `json:"articleNumber"`
	GetProgramTermsAccepted     bool      `json:"getProgramTermsAccepted"`
	FormulaStrippedArticleTitle *string   `json:"formulaStrippedArticleTitle,omitempty"`
	PubLink                     *string   `json:"pubLink,omitempty"`
	Keywords                    []Keyword `json:"keywords,omitempty"`
	AllowComments               bool      `json:"allowComments"`
	IssueLink                   string    `json:"issueLink"`
	IsGetAddressInfoCaptured    bool      `json:"isGetAddressInfoCaptured"`
	IsMarketingOptIn            bool      `json:"isMarketingOptIn"`
	Publisher                   *string   `json:"publisher,omitempty"`
	XploreDocumentType          *string   `json:"xploreDocumentType,omitempty"`
	IsPromo                     bool      `json:"isPromo"`
	HTMLAbstractLink            string    `json:"htmlAbstractLink"`
	IsCustomDenial              bool      `json:"isCustomDenial"`
	IsSAE                       bool      `json:"isSAE"`
	IsDynamicHTML               bool      `json:"isDynamicHtml"`
	IsFreeDocument              bool      `json:"isFreeDocument"`
	DisplayDocTitle             *string   `json:"displayDocTitle,omitempty"`
	IsStandard                  bool      `json:"isStandard"`
	IsSMPTE                     bool      `json:"isSMPTE"`
	IsOUP                       bool      `json:"isOUP"`
	IsNow                       bool      `json:"isNow"`
	IsProduct                   bool      `json:"isProduct"`
	IsJournal                   bool      `json:"isJournal"`
	IsBook                      bool      `json:"isBook"`
	IsBookWithoutChapters       bool      `json:"isBookWithoutChapters"`
	IsOpenAccess                bool      `json:"isOpenAccess"`
	IsEphemera                  bool      `json:"isEphemera"`
	IsConference                bool      `json:"isConference"`
	IsChapter                   bool      `json:"isChapter"`
	IsEarlyAccess               bool      `json:"isEarlyAccess"`
	PersistentLink              *string   `json:"persistentLink,omitempty"`
	ContentTypeDisplay          *string   `json:"contentTypeDisplay,omitempty"`
	MlTime                      string    `json:"mlTime"`
	LastUpdate                  *string   `json:"lastupdate,omitempty"`
	MediaPath                   *string   `json:"mediaPath,omitempty"`
	Title                       *string   `json:"title,omitempty"`
	ContentType                 *string   `json:"contentType,omitempty"`
	PublicationNumber           *string   `json:"publicationNumber,omitempty"`
	HTMLFlagLegacy              *string   `json:"htmlFlag,omitempty"`
	HasStandardVersions         bool      `json:"hasStandardVersions"`
	IsGiveaway                  bool      `json:"isGiveaway"`
	IsLatestStandard            bool      `json:"isLatestStandard"`
	IsOnlineOnly                bool      `json:"isOnlineOnly"`
	IsSpringer                  bool      `json:"isSpringer"`
	IsTranslation               bool      `json:"isTranslation"`
}

type Keyword struct {
	Type *string  `json:"type,omitempty"`
	Kwd  []string `json:"kwd"`
}
