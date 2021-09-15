package goieeeapi

func GetFigures(client HTTPClient, id int) (*GetFiguresResponse, error) {
	resp := &GetFiguresResponse{}
	err := getEndp(client, id, "figures", resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

type GetFiguresResponse struct {
	UserInfo                    UserInfo  `json:"userInfo"`
	Figures                     *[]Figure `json:"figures,omitempty"`
	ArticleNumber               string    `json:"articleNumber"`
	GetProgramTermsAccepted     bool      `json:"getProgramTermsAccepted"`
	FormulaStrippedArticleTitle *string   `json:"formulaStrippedArticleTitle,omitempty"`
	PubLink                     *string   `json:"pubLink,omitempty"`
	AllowComments               bool      `json:"allowComments"`
	IssueLink                   string    `json:"issueLink"`
	IsReadingRoomArticle        bool      `json:"isReadingRoomArticle"`
	IsGetArticle                bool      `json:"isGetArticle"`
	IsGetAddressInfoCaptured    bool      `json:"isGetAddressInfoCaptured"`
	IsMarketingOptIn            bool      `json:"isMarketingOptIn"`
	Publisher                   *string   `json:"publisher,omitempty"`
	XploreDocumentType          *string   `json:"xploreDocumentType,omitempty"`
	IsPromo                     bool      `json:"isPromo"`
	IsNotDynamicOrStatic        bool      `json:"isNotDynamicOrStatic"`
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
	PersistentLink              *string   `json:"persistentLink,omitempty"`
	Title                       *string   `json:"title,omitempty"`
	ContentTypeDisplay          *string   `json:"contentTypeDisplay,omitempty"`
	MlTime                      string    `json:"mlTime"`
	LastUpdate                  *string   `json:"lastupdate,omitempty"`
	MediaPath                   *string   `json:"mediaPath,omitempty"`
	ContentType                 *string   `json:"contentType,omitempty"`
	PublicationNumber           *string   `json:"publicationNumber,omitempty"`
	HTMLFlagLegacy              *string   `json:"htmlFlag,omitempty"`
}

type Figure struct {
	ID          string  `json:"id"`
	Label       string  `json:"label"`
	Graphic     Graphic `json:"graphic"`
	Caption     string  `json:"caption"`
	Part        *string `json:"part,omitempty"`
	Orientation *string `json:"orientation,omitempty"`
	Position    *string `json:"position,omitempty"`
}

type Graphic struct {
	Href        string  `json:"href"`
	Small       string  `json:"small"`
	Large       string  `json:"large"`
	Hires       string  `json:"hires"`
	Glance      string  `json:"glance"`
	Orientation *string `json:"orientation,omitempty"`
	Position    *string `json:"position,omitempty"`
	Type        *string `json:"type,omitempty"`
}
