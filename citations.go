package goieeeapi

func GetCitations(client HTTPClient, id int) (*GetCitationsResponse, error) {
	resp := &GetCitationsResponse{}
	err := getEndp(client, id, "citations", resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

type GetCitationsResponse struct {
	GetProgramTermsAccepted     bool              `json:"getProgramTermsAccepted"`
	FormulaStrippedArticleTitle *string           `json:"formulaStrippedArticleTitle,omitempty"`
	PaperCitations              PaperCitations    `json:"paperCitations"`
	IsReadingRoomArticle        bool              `json:"isReadingRoomArticle"`
	IsGetArticle                bool              `json:"isGetArticle"`
	IsGetAddressInfoCaptured    bool              `json:"isGetAddressInfoCaptured"`
	IsMarketingOptIn            bool              `json:"isMarketingOptIn"`
	IsProduct                   bool              `json:"isProduct"`
	IsChapter                   bool              `json:"isChapter"`
	IsEarlyAccess               bool              `json:"isEarlyAccess"`
	NonIeeeCitationCount        *string           `json:"nonIeeeCitationCount,omitempty"`
	ContentTypeDisplay          *string           `json:"contentTypeDisplay,omitempty"`
	PatentCitationCount         string            `json:"patentCitationCount"`
	PatentCitations             *[]PatentCitation `json:"patentCitations,omitempty"`
	MlTime                      string            `json:"mlTime"`
	LastUpdate                  *string           `json:"lastupdate,omitempty"`
	MediaPath                   *string           `json:"mediaPath,omitempty"`
	Title                       *string           `json:"title,omitempty"`
	ContentType                 *string           `json:"contentType,omitempty"`
	IEEECitationCount           *string           `json:"ieeeCitationCount,omitempty"`
	PublicationNumber           *string           `json:"publicationNumber,omitempty"`
	HTMLFlagLegacy              *string           `json:"htmlFlag,omitempty"`
}

type PaperCitations struct {
	IEEE    *[]Reference `json:"ieee,omitempty"`
	NonIEEE *[]Reference `json:"nonIeee,omitempty"`
}

type PatentCitation struct {
	AppNum            string    `json:"appNum"`
	Assignees         []string  `json:"assignees"`
	FilingDate        string    `json:"filingDate"`
	GoogleScholarLink string    `json:"googleScholarLink"`
	GrantDate         string    `json:"grantDate"`
	ID                string    `json:"id"`
	Inventors         string    `json:"inventors"`
	IPCClassList      string    `json:"ipcClassList"`
	IPCClasses        []string  `json:"ipcClasses"`
	Order             string    `json:"order"`
	PatentAbstract    string    `json:"patentAbstract"`
	PatentLink        string    `json:"patentLink"`
	PatentNumber      string    `json:"patentNumber"`
	PdfLink           *string   `json:"pdfLink,omitempty"`
	PocClassList      *string   `json:"pocClassList,omitempty"`
	PocClasses        *[]string `json:"pocClasses,omitempty"`
	Title             string    `json:"title"`
}
