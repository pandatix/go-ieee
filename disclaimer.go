package goieeeapi

func GetDisclaimer(client HTTPClient, id int) (*GetDisclaimerResponse, error) {
	resp := &GetDisclaimerResponse{}
	err := getEndp(client, id, "disclaimer", resp)
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
	IsReadingRoomArticle        bool     `json:"isReadingRoomArticle"`
	IsGetArticle                bool     `json:"isGetArticle"`
	IsGetAddressInfoCaptured    bool     `json:"isGetAddressInfoCaptured"`
	IsMarketingOptIn            bool     `json:"isMarketingOptIn"`
	Publisher                   *string  `json:"publisher,omitempty"`
	IsPromo                     bool     `json:"isPromo"`
	IsSMPTE                     bool     `json:"isSMPTE"`
	IsOUP                       bool     `json:"isOUP"`
	IsSAE                       bool     `json:"isSAE"`
	IsNow                       bool     `json:"isNow"`
	IsCustomDenial              bool     `json:"isCustomDenial"`
	IsNotDynamicOrStatic        bool     `json:"isNotDynamicOrStatic"`
	DisplayDocTitle             *string  `json:"displayDocTitle,omitempty"`
	IsStandard                  bool     `json:"isStandard"`
	HTMLAbstractLink            string   `json:"htmlAbstractLink"`
	IsProduct                   bool     `json:"isProduct"`
	IsMorganClaypool            bool     `json:"isMorganClaypool"`
	IsOpenAccess                bool     `json:"isOpenAccess"`
	IsEphemera                  bool     `json:"isEphemera"`
	IsConference                bool     `json:"isConference"`
	IsEarlyAccess               bool     `json:"isEarlyAccess"`
	IsJournal                   bool     `json:"isJournal"`
	IsBook                      bool     `json:"isBook"`
	IsBookWithoutChapters       bool     `json:"isBookWithoutChapters"`
	IsChapter                   bool     `json:"isChapter"`
	IsStaticHTML                bool     `json:"isStaticHtml"`
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
}
