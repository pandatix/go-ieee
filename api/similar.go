package api

import (
	"path"
	"strconv"
)

// GetDocumentSimilar fetch the /document/<id>/similar REST endpoint and
// returns the document similar information.
func (client *IEEEClient) GetDocumentSimilar(id int, opts ...Option) (*GetDocumentSimilarResponse, error) {
	resp := &GetDocumentSimilarResponse{}
	err := client.get(path.Join("document", strconv.Itoa(id), "similar"), nil, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

type GetDocumentSimilarResponse struct {
	GetProgramTermsAccepted     bool      `json:"getProgramTermsAccepted"`
	Similar                     []Similar `json:"similar,omitempty"`
	FormulaStrippedArticleTitle *string   `json:"formulaStrippedArticleTitle,omitempty"`
	IsGetAddressInfoCaptured    bool      `json:"isGetAddressInfoCaptured"`
	IsMarketingOptIn            bool      `json:"isMarketingOptIn"`
	Publisher                   *string   `json:"publisher,omitempty"`
	IsCustomDenial              bool      `json:"isCustomDenial"`
	IsSMPTE                     bool      `json:"isSMPTE"`
	IsOUP                       bool      `json:"isOUP"`
	IsNow                       bool      `json:"isNow"`
	IsSAE                       bool      `json:"isSAE"`
	IsStandard                  bool      `json:"isStandard"`
	IsConference                bool      `json:"isConference"`
	IsProduct                   bool      `json:"isProduct"`
	IsJournal                   bool      `json:"isJournal"`
	IsBook                      bool      `json:"isBook"`
	IsEarlyAccess               bool      `json:"isEarlyAccess"`
	PersistentLink              *string   `json:"persistentLink,omitempty"`
	IsChapter                   bool      `json:"isChapter"`
	ContentTypeDisplay          *string   `json:"contentTypeDisplay,omitempty"`
	MlTime                      *string   `json:"mlTime,omitempty"`
	LastUpdate                  *string   `json:"lastupdate,omitempty"`
	MediaPath                   *string   `json:"mediaPath,omitempty"`
	Title                       *string   `json:"title,omitempty"`
	ContentType                 *string   `json:"contentType,omitempty"`
	PublicationNumber           *string   `json:"publicationNumber,omitempty"`
	HTMLFlagLegacy              *string   `json:"htmlFlag,omitempty"`
	IsSpringer                  bool      `json:"isSpringer"`
}

type Similar struct {
	ArticleNumber     string   `json:"articleNumber"`
	Title             string   `json:"title"`
	PublicationNumber string   `json:"publicationNumber"`
	PublicationTitle  string   `json:"publicationTitle"`
	PublicationYear   string   `json:"publicationYear"`
	StandardNumber    *string  `json:"standardNumber,omitempty"`
	Author            []string `json:"author,omitempty"`
	ContentType       string   `json:"contentType"`
	DocumentLink      string   `json:"documentLink"`
	IsArticle         string   `json:"isArticle"`
}
