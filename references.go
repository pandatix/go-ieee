package goieeeapi

func GetReferences(client HTTPClient, id int) (*GetReferencesResponse, error) {
	resp := &GetReferencesResponse{}
	err := getEndp(client, id, "references", resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

type GetReferencesResponse struct {
	UserInfo                    UserInfo    `json:"userInfo"`
	References                  []Reference `json:"references"`
	ArticleNumber               string      `json:"articleNumber"`
	GetProgramTermsAccepted     bool        `json:"getProgramTermsAccepted"`
	FormulaStrippedArticleTitle string      `json:"formulaStrippedArticleTitle"`
	PubLink                     string      `json:"pubLink"`
	AllowComments               bool        `json:"allowComments"`
	IssueLink                   string      `json:"issueLink"`
	IsReadingRoomArticle        bool        `json:"isReadingRoomArticle"`
	IsGetArticle                bool        `json:"isGetArticle"`
	IsGetAddressInfoCaptured    bool        `json:"isGetAddressInfoCaptured"`
	IsMarketingOptIn            bool        `json:"isMarketingOptIn"`
	Publisher                   string      `json:"publisher"`
	XploreDocumentType          string      `json:"xploreDocumentType"`
	IsPromo                     bool        `json:"isPromo"`
	IsNotDynamicOrStatic        bool        `json:"isNotDynamicOrStatic"`
	HTMLAbstractLink            string      `json:"htmlAbstractLink"`
	IsCustomDenial              bool        `json:"isCustomDenial"`
	IsSAE                       bool        `json:"isSAE"`
	IsDynamicHTML               bool        `json:"isDynamicHtml"`
	IsFreeDocument              bool        `json:"isFreeDocument"`
	DisplayDocTitle             string      `json:"displayDocTitle"`
	IsStandard                  bool        `json:"isStandard"`
	IsSMPTE                     bool        `json:"isSMPTE"`
	IsOUP                       bool        `json:"isOUP"`
	IsNow                       bool        `json:"isNow"`
	IsProduct                   bool        `json:"isProduct"`
	IsMorganClaypool            bool        `json:"isMorganClaypool"`
	IsJournal                   bool        `json:"isJournal"`
	IsBook                      bool        `json:"isBook"`
	IsBookWithoutChapters       bool        `json:"isBookWithoutChapters"`
	IsOpenAccess                bool        `json:"isOpenAccess"`
	IsEphemera                  bool        `json:"isEphemera"`
	IsConference                bool        `json:"isConference"`
	HTMLLink                    string      `json:"htmlLink"`
	IsChapter                   bool        `json:"isChapter"`
	IsStaticHTML                bool        `json:"isStaticHtml"`
	IsEarlyAccess               bool        `json:"isEarlyAccess"`
	PersistentLink              string      `json:"persistentLink"`
	ArticleId                   string      `json:"articleId"`
	OpenAccessFlag              string      `json:"openAccessFlag"`
	EphemeraFlag                string      `json:"ephemeraFlag"`
	Title                       string      `json:"title"`
	ContentTypeDisplay          string      `json:"contentTypeDisplay"`
	HTMLFlag                    string      `json:"html_flag"`
	MlHTMLFlag                  string      `json:"ml_html_flag"`
	MlTime                      string      `json:"mlTime"`
	LastUpdate                  string      `json:"lastupdate"`
	MediaPath                   string      `json:"mediaPath"`
	ContentType                 string      `json:"contentType"`
	Definitions                 string      `json:"definitions"`
	PublicationNumber           string      `json:"publicationNumber"`
}

type Reference struct {
	Order             string     `json:"order"`
	Text              *string    `json:"text,omitempty"`
	DisplayText       *string    `json:"displayText"`
	Title             string     `json:"title"`
	Context           *[]Context `json:"context,omitempty"`
	Links             *Links     `json:"links,omitempty"`
	GoogleScholarLink string     `json:"googleScholarLink"`
	RefType           *string    `json:"refType,omitempty"`
	ID                *string    `json:"id,omitempty"`
}

type Context struct {
	Sec  string `json:"sec"`
	Text string `json:"text"`
	Part string `json:"part"`
}

type Links struct {
	DocumentLink  *string `json:"documentLink,omitempty"`
	PdfLink       *string `json:"pdfLink,omitempty"`
	Abstract      *string `json:"abstract,omitempty"`
	CrossRefLink  *string `json:"crossRefLink,omitempty"`
	OpenURLImgLoc string  `json:"openUrlImgLoc"`
	PdfSize       *string `json:"pdfSize,omitempty"`
}
