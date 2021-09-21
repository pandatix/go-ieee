package goieeeapi

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
)

func PostSearch(client HTTPClient, params PostSearchParams) (*PostSearchResponse, error) {
	if client == nil {
		return nil, ErrNilClient
	}

	// Encode payload
	buf := &bytes.Buffer{}
	_ = json.NewEncoder(buf).Encode(params)

	// Issue request
	req, _ := http.NewRequest(http.MethodPost, "https://ieeexplore.ieee.org/rest/search", buf)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Referer", "https://ieeexplore.ieee.org")
	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	// Extract the response body
	body, _ := io.ReadAll(res.Body)

	// Check status code
	if res.StatusCode != http.StatusOK {
		return nil, &ErrUnexpectedStatus{
			StatusCode: res.StatusCode,
			Body:       body,
		}
	}

	// Decode response body
	resp := &PostSearchResponse{}
	err = json.Unmarshal(body, resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

type PostSearchParams struct {
	Action       *string   `json:"action,omitempty"`
	Highlight    bool      `json:"highlight"`
	MatchBoolean *bool     `json:"matchBoolean,omitempty"`
	MatchPubs    bool      `json:"matchPubs"`
	NewSearch    *bool     `json:"newsearch,omitempty"`
	QueryText    string    `json:"queryText"`
	Ranges       *[]string `json:"ranges,omitempty"`
	Refinements  *[]string `json:"refinements,omitempty"`
	ReturnFacets []string  `json:"returnFacets"`
	ReturnType   string    `json:"returnType"`
}

type PostSearchResponse struct {
	UserInfo                 UserInfo     `json:"userInfo"`
	Records                  *[]Record    `json:"records,omitempty"`
	BreadCrumbs              []BreadCrumb `json:"breadCrumbs"`
	ShowStandardDictionary   bool         `json:"showStandardDictionary"`
	SearchType               string       `json:"searchType"`
	TotalRecords             int          `json:"totalRecords"`
	SubscribedContentApplied bool         `json:"subscribedContentApplied"`
	PromoApplied             bool         `json:"promoApplied"`
	StartRecord              int          `json:"startRecord"`
	Facets                   *[]Facet     `json:"facets,omitempty"`
	EndRecord                int          `json:"endRecord"`
	TotalPages               int          `json:"totalPages"`
}

type Record struct {
	Authors                 []RecordAuthor `json:"authors"`
	PatentCitationCount     int            `json:"patentCitationCount"`
	PatentCitationsLink     *string        `json:"patentCitationsLink,omitempty"`
	AccessType              AccessType     `json:"accessType"`
	Volume                  *string        `json:"volume,omitempty"`
	Issue                   *string        `json:"issue,omitempty"`
	PublicationYear         string         `json:"publicationYear"`
	PublicationNumber       string         `json:"publicationNumber"`
	DocumentLink            string         `json:"documentLink"`
	ArticleNumber           string         `json:"articleNumber"`
	DOI                     *string        `json:"doi,omitempty"`
	CitationCount           int            `json:"citationCount"`
	IsNumber                string         `json:"isNumber"`
	ArticleTitle            string         `json:"articleTitle"`
	PdfSize                 string         `json:"pdfSize"`
	RightsLinkFlag          bool           `json:"rightslinkFlag"`
	RightsLink              *string        `json:"rightsLink,omitempty"`
	StartPage               string         `json:"startPage"`
	EndPage                 string         `json:"endPage"`
	PublicationDate         *string        `json:"publicationDate,omitempty"`
	ShowDataset             bool           `json:"showDataset"`
	ShowVideo               bool           `json:"showVideo"`
	Ephemera                bool           `json:"ephemera"`
	VJ                      bool           `json:"vj"`
	ShowAlgorithm           bool           `json:"showAlgorithm"`
	PublicationLink         string         `json:"publicationLink"`
	DownloadCount           int            `json:"downloadCount"`
	HTMLLink                string         `json:"htmlLink"`
	CitationsLink           *string        `json:"citationsLink,omitempty"`
	ShowHTML                bool           `json:"showHtml"`
	Publisher               string         `json:"publisher"`
	Redline                 bool           `json:"redline"`
	ShowCheckbox            bool           `json:"showCheckbox"`
	HandleProduct           bool           `json:"handleProduct"`
	ContentType             string         `json:"contentType"`
	MajorTopic              *string        `json:"majorTopic,omitempty"`
	Abstract                string         `json:"abstract"`
	ArticleContentType      string         `json:"articleContentType"`
	PdfLink                 string         `json:"pdfLink"`
	HighlightedTitle        string         `json:"highlightedTitle"`
	PublicationTitle        string         `json:"publicationTitle"`
	DisplayPublicationTitle string         `json:"displayPublicationTitle"`
	IsStandard              bool           `json:"isStandard"`
	IsConference            bool           `json:"isConference"`
	IsJournalAndMagazine    bool           `json:"isJournalAndMagazine"`
	IsMagazine              bool           `json:"isMagazine"`
	IsJournal               bool           `json:"isJournal"`
	IsBook                  bool           `json:"isBook"`
	Course                  bool           `json:"course"`
	IsBookWithoutChapters   bool           `json:"isBookWithoutChapters"`
	DisplayContentType      string         `json:"displayContentType"`
	DocIdentifier           string         `json:"docIdentifier"`
	IsEarlyAccess           bool           `json:"isEarlyAccess"`
}

type RecordAuthor struct {
	PreferredName           string  `json:"preferredName"`
	NormalizedName          string  `json:"normalizedName"`
	FirstName               *string `json:"firstName,omitempty"`
	LastName                string  `json:"lastName"`
	SearchablePreferredName string  `json:"searchablePreferredName"`
	ID                      *int    `json:"id,omitempty"`
}

type AccessType struct {
	Type    string `json:"type"`
	Message string `json:"message"`
}

type BreadCrumb struct {
	Type      string             `json:"type"`
	Key       *string            `json:"key,omitempty"`
	Value     *string            `json:"value,omitempty"`
	Reference *string            `json:"reference,omitempty"`
	Start     *string            `json:"start,omitempty"`
	End       *string            `json:"end,omitempty"`
	Children  *[]BreadCrumbChild `json:"children,omitempty"`
}

type BreadCrumbChild struct {
	Type      *string `json:"type,omitempty"`
	Key       *string `json:"key,omitempty"`
	Value     string  `json:"value"`
	Reference *string `json:"reference,omitempty"`
	Display   *string `json:"display,omitempty"`
}

type Facet struct {
	ID         string   `json:"id"`
	Name       string   `json:"name"`
	NumRecords int      `json:"numRecords"`
	Children   *[]Facet `json:"children,omitempty"`
	Active     string   `json:"active"`
}
