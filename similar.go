package goieeeapi

func GetSimilar(client HTTPClient, id int) (*GetSimilarResponse, error) {
	resp := &GetSimilarResponse{}
	err := getEndp(client, id, "similar", resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

type GetSimilarResponse struct {
	GetProgramTermsAccepted     bool       `json:"getProgramTermsAccepted"`
	Similar                     *[]Similar `json:"similar,omitempty"`
	FormulaStrippedArticleTitle *string    `json:"formulaStrippedArticleTitle,omitempty"`
	IsReadingRoomArticle        bool       `json:"isReadingRoomArticle"`
	IsGetArticle                bool       `json:"isGetArticle"`
	IsGetAddressInfoCaptured    bool       `json:"isGetAddressInfoCaptured"`
	IsMarketingOptIn            bool       `json:"isMarketingOptIn"`
	Publisher                   *string    `json:"publisher,omitempty"`
	IsCustomDenial              bool       `json:"isCustomDenial"`
	IsSMPTE                     bool       `json:"isSMPTE"`
	IsOUP                       bool       `json:"isOUP"`
	IsNow                       bool       `json:"isNow"`
	IsSAE                       bool       `json:"isSAE"`
	IsStandard                  bool       `json:"isStandard"`
	IsConference                bool       `json:"isConference"`
	IsProduct                   bool       `json:"isProduct"`
	IsMorganClaypool            bool       `json:"isMorganClaypool"`
	IsJournal                   bool       `json:"isJournal"`
	IsBook                      bool       `json:"isBook"`
	IsEarlyAccess               bool       `json:"isEarlyAccess"`
	PersistentLink              *string    `json:"persistentLink,omitempty"`
	IsChapter                   bool       `json:"isChapter"`
	ContentTypeDisplay          *string    `json:"contentTypeDisplay,omitempty"`
	MlTime                      *string    `json:"mlTime,omitempty"`
	LastUpdate                  *string    `json:"lastupdate,omitempty"`
	MediaPath                   *string    `json:"mediaPath,omitempty"`
	Title                       *string    `json:"title,omitempty"`
	ContentType                 *string    `json:"contentType,omitempty"`
	PublicationNumber           *string    `json:"publicationNumber,omitempty"`
	HTMLFlagLegacy              *string    `json:"htmlFlag,omitempty"`
}

type Similar struct {
	ArticleNumber     string    `json:"articleNumber"`
	Title             string    `json:"title"`
	PublicationNumber string    `json:"publicationNumber"`
	PublicationTitle  string    `json:"publicationTitle"`
	PublicationYear   string    `json:"publicationYear"`
	StandardNumber    *string   `json:"standardNumber,omitempty"`
	Author            *[]string `json:"author,omitempty"`
	ContentType       string    `json:"contentType"`
	DocumentLink      string    `json:"documentLink"`
	IsArticle         string    `json:"isArticle"`
}
