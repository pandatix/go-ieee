package api

import (
	"path"
	"strconv"
)

// GetDocumentMultimedia fetch the /document/<id>/mutimedia REST endpoint and
// returns the document mutimedia information.
func (client *IEEEClient) GetDocumentMultimedia(id int, opts ...Option) (*GetMultimediaResponse, error) {
	resp := &GetMultimediaResponse{}
	err := client.get(path.Join("document", strconv.Itoa(id), "multimedia"), nil, resp, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

type GetMultimediaResponse struct {
	UserInfo                    UserInfo `json:"userInfo"`
	ArticleNumber               string   `json:"articleNumber"`
	GetProgramTermsAccepted     bool     `json:"getProgramTermsAccepted"`
	AllowComments               bool     `json:"allowComments"`
	PubLink                     *string  `json:"pubLink,omitempty"`
	IssueLink                   string   `json:"issueLink"`
	FormulaStrippedArticleTitle *string  `json:"formulaStrippedArticleTitle,omitempty"`
	IsGetAddressInfoCaptured    bool     `json:"isGetAddressInfoCaptured"`
	IsMarketingOptIn            bool     `json:"isMarketingOptIn"`
	Publisher                   *string  `json:"publisher,omitempty"`
	IsDynamicHTML               bool     `json:"isDynamicHtml"`
	IsFreeDocument              bool     `json:"isFreeDocument"`
	DisplayDocTitle             *string  `json:"displayDocTitle,omitempty"`
	IsStandard                  bool     `json:"isStandard"`
	IsConference                bool     `json:"isConference"`
	IsProduct                   bool     `json:"isProduct"`
	IsPromo                     bool     `json:"isPromo"`
	IsBookWithoutChapters       bool     `json:"isBookWithoutChapters"`
	PersistentLink              *string  `json:"persistentLink,omitempty"`
	IsEarlyAccess               bool     `json:"isEarlyAccess"`
	IsJournal                   bool     `json:"isJournal"`
	IsBook                      bool     `json:"isBook"`
	IsChapter                   bool     `json:"isChapter"`
	IsOpenAccess                bool     `json:"isOpenAccess"`
	IsEphemera                  bool     `json:"isEphemera"`
	HTMLAbstractLink            string   `json:"htmlAbstractLink"`
	IsSMPTE                     bool     `json:"isSMPTE"`
	IsOUP                       bool     `json:"isOUP"`
	IsSAE                       bool     `json:"isSAE"`
	IsNow                       bool     `json:"isNow"`
	IsCustomDenial              bool     `json:"isCustomDenial"`
	XploreDocumentType          *string  `json:"xploreDocumentType,omitempty"`
	ContentTypeDisplay          *string  `json:"contentTypeDisplay,omitempty"`
	MlTime                      string   `json:"mlTime"`
	LastUpdate                  *string  `json:"lastupdate,omitempty"`
	MediaPath                   *string  `json:"mediaPath,omitempty"`
	Title                       *string  `json:"title,omitempty"`
	ContentType                 *string  `json:"contentType,omitempty"`
	PublicationNumber           *string  `json:"publicationNumber,omitempty"`
	HTMLFlagLegacy              *string  `json:"htmlFlag,omitempty"`
	HasStandardVersions         bool     `json:"hasStandardVersions"`
	IsGiveaway                  bool     `json:"isGiveaway"`
	IsLatestStandard            bool     `json:"isLatestStandard"`
	IsOnlineOnly                bool     `json:"isOnlineOnly"`
	IsSpringer                  bool     `json:"isSpringer"`
	IsTranslation               bool     `json:"isTranslation"`
}
