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
	GetProgramTermsAccepted     bool           `json:"getProgramTermsAccepted"`
	FormulaStrippedArticleTitle string         `json:"formulaStrippedArticleTitle"`
	PaperCitations              PaperCitations `json:"paperCitations"`
	IsReadingRoomArticle        bool           `json:"isReadingRoomArticle"`
	IsGetArticle                bool           `json:"isGetArticle"`
	IsGetAddressInfoCaptured    bool           `json:"isGetAddressInfoCaptured"`
	IsMarketingOptIn            bool           `json:"isMarketingOptIn"`
	IsProduct                   bool           `json:"isProduct"`
	IsChapter                   bool           `json:"isChapter"`
	IsEarlyAccess               bool           `json:"isEarlyAccess"`
	NonIeeeCitationCount        string         `json:"nonIeeeCitationCount"`
	ContentTypeDisplay          string         `json:"contentTypeDisplay"`
	PatentCitationCount         string         `json:"patentCitationCount"`
	MlTime                      string         `json:"mlTime"`
	LastUpdate                  string         `json:"lastupdate"`
	MediaPath                   string         `json:"mediaPath"`
	Title                       string         `json:"title"`
	ContentType                 string         `json:"contentType"`
	IEEECitationCount           string         `json:"ieeeCitationCount"`
	PublicationNumber           string         `json:"publicationNumber"`
}

type PaperCitations struct {
	IEEE []Reference `json:"ieee"`
}
