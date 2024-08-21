package goieeeapi

func GetMultimedia(client HTTPClient, id int) (*GetMultimediaResponse, error) {
	resp := &GetMultimediaResponse{}
	err := getEndp(client, id, "multimedia", nil, resp)
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
	IsReadingRoomArticle        bool     `json:"isReadingRoomArticle"`
	IsGetArticle                bool     `json:"isGetArticle"`
	IsGetAddressInfoCaptured    bool     `json:"isGetAddressInfoCaptured"`
	IsMarketingOptIn            bool     `json:"isMarketingOptIn"`
	Publisher                   *string  `json:"publisher,omitempty"`
	IsDynamicHTML               bool     `json:"isDynamicHtml"`
	IsFreeDocument              bool     `json:"isFreeDocument"`
	DisplayDocTitle             *string  `json:"displayDocTitle,omitempty"`
	IsStandard                  bool     `json:"isStandard"`
	IsMorganClaypool            bool     `json:"isMorganClaypool"`
	IsConference                bool     `json:"isConference"`
	IsProduct                   bool     `json:"isProduct"`
	IsPromo                     bool     `json:"isPromo"`
	IsBookWithoutChapters       bool     `json:"isBookWithoutChapters"`
	PersistentLink              *string  `json:"persistentLink,omitempty"`
	IsEarlyAccess               bool     `json:"isEarlyAccess"`
	IsJournal                   bool     `json:"isJournal"`
	IsBook                      bool     `json:"isBook"`
	IsChapter                   bool     `json:"isChapter"`
	IsStaticHTML                bool     `json:"isStaticHtml"`
	IsOpenAccess                bool     `json:"isOpenAccess"`
	IsEphemera                  bool     `json:"isEphemera"`
	HTMLAbstractLink            string   `json:"htmlAbstractLink"`
	IsSMPTE                     bool     `json:"isSMPTE"`
	IsOUP                       bool     `json:"isOUP"`
	IsSAE                       bool     `json:"isSAE"`
	IsNow                       bool     `json:"isNow"`
	IsCustomDenial              bool     `json:"isCustomDenial"`
	IsNotDynamicOrStatic        bool     `json:"isNotDynamicOrStatic"`
	XploreDocumentType          *string  `json:"xploreDocumentType,omitempty"`
	ContentTypeDisplay          *string  `json:"contentTypeDisplay,omitempty"`
	MlTime                      string   `json:"mlTime"`
	LastUpdate                  *string  `json:"lastupdate,omitempty"`
	MediaPath                   *string  `json:"mediaPath,omitempty"`
	Title                       *string  `json:"title,omitempty"`
	ContentType                 *string  `json:"contentType,omitempty"`
	PublicationNumber           *string  `json:"publicationNumber,omitempty"`
	HTMLFlagLegacy              *string  `json:"htmlFlag,omitempty"`
}
