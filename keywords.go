package goieeeapi

func GetKeywords(client HTTPClient, id int) (*GetKeywordsResponse, error) {
	resp := &GetKeywordsResponse{}
	err := getEndp(client, id, "keywords", resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

type GetKeywordsResponse struct {
	UserInfo                    UserInfo  `json:"userInfo"`
	ArticleNumber               string    `json:"articleNumber"`
	GetProgramTermsAccepted     bool      `json:"getProgramTermsAccepted"`
	FormulaStrippedArticleTitle string    `json:"formulaStrippedArticleTitle"`
	PubLink                     string    `json:"pubLink"`
	Keywords                    []Keyword `json:"keywords"`
	AllowComments               bool      `json:"allowComments"`
	IssueLink                   string    `json:"issueLink"`
	IsReadingRoomArticle        bool      `json:"isReadingRoomArticle"`
	IsGetArticle                bool      `json:"isGetArticle"`
	IsGetAddressInfoCaptured    bool      `json:"isGetAddressInfoCaptured"`
	IsMarketingOptIn            bool      `json:"isMarketingOptIn"`
	Publisher                   string    `json:"publisher"`
	XploreDocumentType          string    `json:"xploreDocumentType"`
	IsPromo                     bool      `json:"isPromo"`
	IsNotDynamicOrStatic        bool      `json:"isNotDynamicOrStatic"`
	HTMLAbstractLink            string    `json:"htmlAbstractLink"`
	IsCustomDenial              bool      `json:"isCustomDenial"`
	IsSAE                       bool      `json:"isSAE"`
	IsDynamicHTML               bool      `json:"isDynamicHtml"`
	IsFreeDocument              bool      `json:"isFreeDocument"`
	DisplayDocTitle             string    `json:"displayDocTitle"`
	IsStandard                  bool      `json:"isStandard"`
	IsSMPTE                     bool      `json:"isSMPTE"`
	IsOUP                       bool      `json:"isOUP"`
	IsNow                       bool      `json:"isNow"`
	IsProduct                   bool      `json:"isProduct"`
	IsMorganClaypool            bool      `json:"isMorganClaypool"`
	IsJournal                   bool      `json:"isJournal"`
	IsBook                      bool      `json:"isBook"`
	IsBookWithoutChapters       bool      `json:"isBookWithoutChapters"`
	IsOpenAccess                bool      `json:"isOpenAccess"`
	IsEphemera                  bool      `json:"isEphemera"`
	IsConference                bool      `json:"isConference"`
	IsChapter                   bool      `json:"isChapter"`
	IsStaticHTML                bool      `json:"isStaticHtml"`
	IsEarlyAccess               bool      `json:"isEarlyAccess"`
	PersistentLink              string    `json:"persistentLink"`
	ContentTypeDisplay          string    `json:"contentTypeDisplay"`
	MlTime                      string    `json:"mlTime"`
	LastUpdate                  string    `json:"lastupdate"`
	MediaPath                   string    `json:"mediaPath"`
	Title                       string    `json:"title"`
	ContentType                 string    `json:"contentType"`
	PublicationNumber           string    `json:"publicationNumber"`
}

type Keyword struct {
	Type string   `json:"type"`
	Kwd  []string `json:"kwd"`
}
