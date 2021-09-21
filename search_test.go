package goieeeapi_test

import (
	"encoding/json"
	"net/http"
	"reflect"
	"testing"

	goieeeapi "github.com/pandatix/go-ieee-api"
)

func TestPostSearch(t *testing.T) {
	t.Parallel()

	var tests = map[string]struct {
		Client          goieeeapi.HTTPClient
		Params          goieeeapi.PostSearchParams
		ExpectedResults *goieeeapi.PostSearchResponse
		ExpectedErr     error
	}{
		"nil-client": {
			Client:          nil,
			Params:          goieeeapi.PostSearchParams{},
			ExpectedResults: nil,
			ExpectedErr:     goieeeapi.ErrNilClient,
		},
		"failing-client": {
			Client:          newFakeHTTPClient(``, 0, errFake),
			Params:          goieeeapi.PostSearchParams{},
			ExpectedResults: nil,
			ExpectedErr:     errFake,
		},
		"unexpected-statuscode": {
			Client:          newFakeHTTPClient(``, 0, nil),
			Params:          goieeeapi.PostSearchParams{},
			ExpectedResults: nil,
			ExpectedErr: &goieeeapi.ErrUnexpectedStatus{
				StatusCode: 0,
				Body:       []byte(``),
			},
		},
		"failing-unmarshal": {
			Client:          newFakeHTTPClient(`{[}]`, http.StatusOK, nil),
			Params:          goieeeapi.PostSearchParams{},
			ExpectedResults: nil,
			ExpectedErr: &json.SyntaxError{
				Offset: 2,
			},
		},
		"valid-call": {
			Client: newFakeHTTPClient(`{"userInfo":{"institute":false,"member":false,"individual":false,"guest":false,"subscribedContent":false,"fileCabinetContent":false,"fileCabinetUser":false,"institutionalFileCabinetUser":false,"ip":"255.255.255.255","showPatentCitations":true,"showGet802Link":false},"records":[{"authors":[{"preferredName":"Zou Futai","normalizedName":"Z. Futai","firstName":"Zou","lastName":"Futai","searchablePreferredName":"Zou Futai"},{"preferredName":"Zhang Siyu","normalizedName":"Z. Siyu","firstName":"Zhang","lastName":"Siyu","searchablePreferredName":"Zhang Siyu"},{"preferredName":"Rao Weixiong","normalizedName":"R. Weixiong","firstName":"Rao","lastName":"Weixiong","searchablePreferredName":"Rao Weixiong"}],"patentCitationCount":0,"accessType":{"type":"locked","message":"Full text access may be available. Click article title to sign in or learn about subscription options."},"publicationYear":"2013","volume":"10","issue":"11","publicationNumber":"6245522","documentLink":"/document/6674213/","articleNumber":"6674213","doi":"10.1109/CC.2013.6674213","citationCount":4,"isNumber":"6674198","articleTitle":"Hybrid detection and tracking of fast-flux botnet on domain name system traffic","pdfSize":"10699","rightslinkFlag":true,"rightsLink":"http://s100.copyright.com/AppDispatchServlet?publisherName=ieee&publication=6245522&title=Hybrid+detection+and+tracking+of+fast-flux+botnet+on+domain+name+system+traffic&isbn=null&publicationDate=Nov.+2013&author=Zou+Futai%3B+Zhang+Siyu%3B+Rao+Weixiong&ContentID=10.1109/CC.2013.6674213&orderBeanReset=true&startPage=81&endPage=94&volumeNum=10&issueNum=11","startPage":"81","endPage":"94","publicationDate":"Nov. 2013","showDataset":false,"showVideo":false,"ephemera":false,"vj":false,"showAlgorithm":false,"publicationLink":"/xpl/RecentIssue.jsp?punumber=6245522","downloadCount":824,"htmlLink":"/document/6674213/","citationsLink":"/document/6674213/citations?tabFilter=papers","showHtml":false,"publisher":"IEEE","redline":false,"showCheckbox":true,"handleProduct":false,"contentType":"IEEE Magazines","abstract":"Fast-flux is a Domain Name System (DNS) technique used by [::botnets::] to organise compromised hosts into a high-availability, load-balancing network that is similar to Content Delivery Networks (CDNs). Fast-Flux Service Networks (FFSNs) are usually used as proxies of phishing websites and malwares, and hide upstream servers that host actual content. In this paper, by analysing recursive DNS traf...","articleContentType":"Magazines","pdfLink":"/stamp/stamp.jsp?tp=&arnumber=6674213","highlightedTitle":"Hybrid [::detection::] and tracking of fast-flux [::botnet::] on domain name system traffic","publicationTitle":"China Communications","displayPublicationTitle":"China Communications","isStandard":false,"isConference":false,"isJournalAndMagazine":true,"isMagazine":false,"isJournal":false,"isBook":false,"course":false,"isBookWithoutChapters":false,"displayContentType":"Magazine Article","docIdentifier":"IEEE Magazines","isEarlyAccess":false},{"authors":[{"preferredName":"Aditya K. Sood","normalizedName":"A. K. Sood","firstName":"Aditya K.","lastName":"Sood","searchablePreferredName":"Aditya K. Sood","id":38548373600},{"preferredName":"Sherali Zeadally","normalizedName":"S. Zeadally","firstName":"Sherali","lastName":"Zeadally","searchablePreferredName":"Sherali Zeadally","id":37271608700}],"patentCitationCount":0,"accessType":{"type":"locked","message":"Full text access may be available. Click article title to sign in or learn about subscription options."},"publicationYear":"2016","volume":"14","issue":"4","publicationNumber":"8013","documentLink":"/document/7535098/","articleNumber":"7535098","doi":"10.1109/MSP.2016.76","citationCount":39,"isNumber":"7535064","articleTitle":"A Taxonomy of Domain-Generation Algorithms","pdfSize":"4526","rightslinkFlag":true,"rightsLink":"http://s100.copyright.com/AppDispatchServlet?publisherName=ieee&publication=8013&title=A+Taxonomy+of+Domain-Generation+Algorithms&isbn=null&publicationDate=July-Aug.+2016&author=Aditya+K.+Sood%3B+Sherali+Zeadally&ContentID=10.1109/MSP.2016.76&orderBeanReset=true&startPage=46&endPage=53&volumeNum=14&issueNum=4","startPage":"46","endPage":"53","publicationDate":"July-Aug. 2016","showDataset":false,"showVideo":false,"ephemera":false,"vj":false,"showAlgorithm":false,"publicationLink":"/xpl/RecentIssue.jsp?punumber=8013","downloadCount":1023,"htmlLink":"/document/7535098/","citationsLink":"/document/7535098/citations?tabFilter=papers","showHtml":true,"publisher":"IEEE","redline":false,"showCheckbox":true,"handleProduct":false,"contentType":"IEEE Magazines","majorTopic":"Security Smorgasbord","abstract":"Domain-generation algorithms (DGAs) allow attackers to manage infection-spreading websites and command-and-control (C&amp;C) deployments by altering domain names on a timely basis. DGAs have made the infection and C&amp;C architecture more robust and supportive for attackers. This detailed taxonomy of DGAs highlights the problem and offers solutions to combat DGAs through [::detection::] of drive-...","articleContentType":"Magazines","pdfLink":"/stamp/stamp.jsp?tp=&arnumber=7535098","highlightedTitle":"A Taxonomy of Domain-Generation Algorithms","publicationTitle":"IEEE Security & Privacy","displayPublicationTitle":"IEEE Security & Privacy","isStandard":false,"isConference":false,"isJournalAndMagazine":true,"isMagazine":false,"isJournal":false,"isBook":false,"course":false,"isBookWithoutChapters":false,"displayContentType":"Magazine Article","docIdentifier":"IEEE Magazines","isEarlyAccess":false}],"breadCrumbs":[{"type":"search","children":[{"value":"Botnet detection","reference":"Botnet detection","display":"Botnet detection"}]},{"type":"refinement","key":"Content Type","value":"ContentType:","children":[{"type":"refinement","key":"Magazines","value":"ContentType:Magazines"},{"type":"refinement","key":"Early Access Articles","value":"ContentType:Early Access Articles"}]},{"type":"refinement","key":"Publication Topics","value":"ControlledTerms:","children":[{"type":"refinement","key":"Web sites","value":"ControlledTerms:Web sites"}]},{"type":"range","key":"Year","reference":"2012_2016_Year","start":"2012","end":"2016"}],"showStandardDictionary":false,"searchType":"basic","totalRecords":2,"subscribedContentApplied":false,"promoApplied":false,"startRecord":0,"facets":[{"id":"ContentType","name":"Content Type","numRecords":-1,"children":[{"id":"ContentType:Conferences","name":"Conferences","numRecords":6,"active":"true"},{"id":"ContentType:Journals","name":"Journals","numRecords":1,"active":"true"}],"active":"true"},{"id":"PublicationYear","name":"Year","numRecords":-1,"children":[{"id":"PublicationYear:2013","name":"2013","numRecords":1,"active":"true"},{"id":"PublicationYear:2016","name":"2016","numRecords":1,"active":"true"}],"active":"true"},{"id":"Author","name":"Author","numRecords":-1,"children":[{"id":"Author:Aditya K. Sood","name":"Aditya K. Sood","numRecords":1,"active":"true"},{"id":"Author:Sherali Zeadally","name":"Sherali Zeadally","numRecords":1,"active":"true"},{"id":"Author:Zou Futai","name":"Zou Futai","numRecords":1,"active":"true"},{"id":"Author:Zhang Siyu","name":"Zhang Siyu","numRecords":1,"active":"true"},{"id":"Author:Rao Weixiong","name":"Rao Weixiong","numRecords":1,"active":"true"}],"active":"true"},{"id":"Affiliation","name":"Affiliation","numRecords":-1,"children":[{"id":"Affiliation:School of Software Engineering, Tongji University, Shanghai 200092, China","name":"School of Software Engineering, Tongji University, Shanghai 200092, China","numRecords":1,"active":"true"},{"id":"Affiliation:Elastica","name":"Elastica","numRecords":1,"active":"true"},{"id":"Affiliation:School of Information Security Engineering, Shanghai Jiao Tong University, Shanghai 200240, China","name":"School of Information Security Engineering, Shanghai Jiao Tong University, Shanghai 200240, China","numRecords":1,"active":"true"},{"id":"Affiliation:University of Kentucky","name":"University of Kentucky","numRecords":1,"active":"true"}],"active":"true"},{"id":"PublicationTitle","name":"Publication Title","numRecords":-1,"children":[{"id":"PublicationTitle:China Communications","name":"China Communications","numRecords":1,"active":"true"},{"id":"PublicationTitle:IEEE Security .AND. Privacy","name":"IEEE Security & Privacy","numRecords":1,"active":"true"}],"active":"true"},{"id":"Publisher","name":"Publisher","numRecords":-1,"children":[{"id":"Publisher:IEEE","name":"IEEE","numRecords":2,"active":"true"}],"active":"true"},{"id":"ControlledTerms","name":"Publication Topics","numRecords":-1,"children":[{"id":"ControlledTerms:computer network security","name":"computer network security","numRecords":3,"active":"true"},{"id":"ControlledTerms:invasive software","name":"invasive software","numRecords":3,"active":"true"},{"id":"ControlledTerms:Internet","name":"Internet","numRecords":2,"active":"true"},{"id":"ControlledTerms:security of data","name":"security of data","numRecords":2,"active":"true"},{"id":"ControlledTerms:artificial intelligence","name":"artificial intelligence","numRecords":1,"active":"true"},{"id":"ControlledTerms:cloud computing","name":"cloud computing","numRecords":1,"active":"true"},{"id":"ControlledTerms:command and control systems","name":"command and control systems","numRecords":1,"active":"true"},{"id":"ControlledTerms:digital forensics","name":"digital forensics","numRecords":1,"active":"true"},{"id":"ControlledTerms:distributed sensors","name":"distributed sensors","numRecords":1,"active":"true"},{"id":"ControlledTerms:fractals","name":"fractals","numRecords":1,"active":"true"},{"id":"ControlledTerms:genetic algorithms","name":"genetic algorithms","numRecords":1,"active":"true"},{"id":"ControlledTerms:inference mechanisms","name":"inference mechanisms","numRecords":1,"active":"true"},{"id":"ControlledTerms:military computing","name":"military computing","numRecords":1,"active":"true"},{"id":"ControlledTerms:network servers","name":"network servers","numRecords":1,"active":"true"},{"id":"ControlledTerms:neural nets","name":"neural nets","numRecords":1,"active":"true"},{"id":"ControlledTerms:peer-to-peer computing","name":"peer-to-peer computing","numRecords":1,"active":"true"},{"id":"ControlledTerms:resource allocation","name":"resource allocation","numRecords":1,"active":"true"},{"id":"ControlledTerms:telecommunication traffic","name":"telecommunication traffic","numRecords":1,"active":"true"},{"id":"ControlledTerms:uncertainty handling","name":"uncertainty handling","numRecords":1,"active":"true"}],"active":"true"}],"endRecord":2,"totalPages":1}`, 200, nil),
			Params: goieeeapi.PostSearchParams{
				Highlight: true,
				MatchPubs: true,
				QueryText: "Botnet detection",
				Ranges: &[]string{
					"2012_2016_Year",
				},
				Refinements: &[]string{
					"ContentType:Magazines",
					"ContentType:Early Access Articles",
					"ControlledTerms:Web sites",
				},
				ReturnFacets: []string{
					"ALL",
				},
				ReturnType: "SEARCH",
			},
			ExpectedResults: &goieeeapi.PostSearchResponse{
				UserInfo: goieeeapi.UserInfo{
					Institute:                    false,
					Member:                       false,
					Individual:                   false,
					Guest:                        false,
					SubscribedContent:            false,
					FileCabinetContent:           false,
					FileCabinetUser:              false,
					InstitutionalFileCabinetUser: false,
					IP:                           str("255.255.255.255"),
					ShowPatentCitations:          true,
					ShowGet802Link:               false,
				},
				Records: &[]goieeeapi.Record{
					{
						Authors: []goieeeapi.RecordAuthor{
							{
								PreferredName:           "Zou Futai",
								NormalizedName:          "Z. Futai",
								FirstName:               str("Zou"),
								LastName:                "Futai",
								SearchablePreferredName: "Zou Futai",
							}, {
								PreferredName:           "Zhang Siyu",
								NormalizedName:          "Z. Siyu",
								FirstName:               str("Zhang"),
								LastName:                "Siyu",
								SearchablePreferredName: "Zhang Siyu",
							}, {
								PreferredName:           "Rao Weixiong",
								NormalizedName:          "R. Weixiong",
								FirstName:               str("Rao"),
								LastName:                "Weixiong",
								SearchablePreferredName: "Rao Weixiong",
							},
						},
						PatentCitationCount: 0,
						AccessType: goieeeapi.AccessType{
							Type:    "locked",
							Message: "Full text access may be available. Click article title to sign in or learn about subscription options.",
						},
						PublicationYear:         "2013",
						Volume:                  str("10"),
						Issue:                   str("11"),
						PublicationNumber:       "6245522",
						DocumentLink:            "/document/6674213/",
						ArticleNumber:           "6674213",
						DOI:                     str("10.1109/CC.2013.6674213"),
						CitationCount:           4,
						IsNumber:                "6674198",
						ArticleTitle:            "Hybrid detection and tracking of fast-flux botnet on domain name system traffic",
						PdfSize:                 "10699",
						RightsLinkFlag:          true,
						RightsLink:              str("http://s100.copyright.com/AppDispatchServlet?publisherName=ieee&publication=6245522&title=Hybrid+detection+and+tracking+of+fast-flux+botnet+on+domain+name+system+traffic&isbn=null&publicationDate=Nov.+2013&author=Zou+Futai%3B+Zhang+Siyu%3B+Rao+Weixiong&ContentID=10.1109/CC.2013.6674213&orderBeanReset=true&startPage=81&endPage=94&volumeNum=10&issueNum=11"),
						StartPage:               "81",
						EndPage:                 "94",
						PublicationDate:         str("Nov. 2013"),
						ShowDataset:             false,
						ShowVideo:               false,
						Ephemera:                false,
						VJ:                      false,
						ShowAlgorithm:           false,
						PublicationLink:         "/xpl/RecentIssue.jsp?punumber=6245522",
						DownloadCount:           824,
						HTMLLink:                "/document/6674213/",
						CitationsLink:           str("/document/6674213/citations?tabFilter=papers"),
						ShowHTML:                false,
						Publisher:               "IEEE",
						Redline:                 false,
						ShowCheckbox:            true,
						HandleProduct:           false,
						ContentType:             "IEEE Magazines",
						Abstract:                "Fast-flux is a Domain Name System (DNS) technique used by [::botnets::] to organise compromised hosts into a high-availability, load-balancing network that is similar to Content Delivery Networks (CDNs). Fast-Flux Service Networks (FFSNs) are usually used as proxies of phishing websites and malwares, and hide upstream servers that host actual content. In this paper, by analysing recursive DNS traf...",
						ArticleContentType:      "Magazines",
						PdfLink:                 "/stamp/stamp.jsp?tp=&arnumber=6674213",
						HighlightedTitle:        "Hybrid [::detection::] and tracking of fast-flux [::botnet::] on domain name system traffic",
						PublicationTitle:        "China Communications",
						DisplayPublicationTitle: "China Communications",
						IsStandard:              false,
						IsConference:            false,
						IsJournalAndMagazine:    true,
						IsMagazine:              false,
						IsJournal:               false,
						IsBook:                  false,
						Course:                  false,
						IsBookWithoutChapters:   false,
						DisplayContentType:      "Magazine Article",
						DocIdentifier:           "IEEE Magazines",
						IsEarlyAccess:           false,
					}, {
						Authors: []goieeeapi.RecordAuthor{
							{
								PreferredName:           "Aditya K. Sood",
								NormalizedName:          "A. K. Sood",
								FirstName:               str("Aditya K."),
								LastName:                "Sood",
								SearchablePreferredName: "Aditya K. Sood",
								ID:                      i(38548373600),
							}, {
								PreferredName:           "Sherali Zeadally",
								NormalizedName:          "S. Zeadally",
								FirstName:               str("Sherali"),
								LastName:                "Zeadally",
								SearchablePreferredName: "Sherali Zeadally",
								ID:                      i(37271608700),
							},
						},
						PatentCitationCount: 0,
						AccessType: goieeeapi.AccessType{
							Type:    "locked",
							Message: "Full text access may be available. Click article title to sign in or learn about subscription options.",
						},
						PublicationYear:         "2016",
						Volume:                  str("14"),
						Issue:                   str("4"),
						PublicationNumber:       "8013",
						DocumentLink:            "/document/7535098/",
						ArticleNumber:           "7535098",
						DOI:                     str("10.1109/MSP.2016.76"),
						CitationCount:           39,
						IsNumber:                "7535064",
						ArticleTitle:            "A Taxonomy of Domain-Generation Algorithms",
						PdfSize:                 "4526",
						RightsLinkFlag:          true,
						RightsLink:              str("http://s100.copyright.com/AppDispatchServlet?publisherName=ieee&publication=8013&title=A+Taxonomy+of+Domain-Generation+Algorithms&isbn=null&publicationDate=July-Aug.+2016&author=Aditya+K.+Sood%3B+Sherali+Zeadally&ContentID=10.1109/MSP.2016.76&orderBeanReset=true&startPage=46&endPage=53&volumeNum=14&issueNum=4"),
						StartPage:               "46",
						EndPage:                 "53",
						PublicationDate:         str("July-Aug. 2016"),
						ShowDataset:             false,
						ShowVideo:               false,
						Ephemera:                false,
						VJ:                      false,
						ShowAlgorithm:           false,
						PublicationLink:         "/xpl/RecentIssue.jsp?punumber=8013",
						DownloadCount:           1023,
						HTMLLink:                "/document/7535098/",
						CitationsLink:           str("/document/7535098/citations?tabFilter=papers"),
						ShowHTML:                true,
						Publisher:               "IEEE",
						Redline:                 false,
						ShowCheckbox:            true,
						HandleProduct:           false,
						ContentType:             "IEEE Magazines",
						MajorTopic:              str("Security Smorgasbord"),
						Abstract:                "Domain-generation algorithms (DGAs) allow attackers to manage infection-spreading websites and command-and-control (C&amp;C) deployments by altering domain names on a timely basis. DGAs have made the infection and C&amp;C architecture more robust and supportive for attackers. This detailed taxonomy of DGAs highlights the problem and offers solutions to combat DGAs through [::detection::] of drive-...",
						ArticleContentType:      "Magazines",
						PdfLink:                 "/stamp/stamp.jsp?tp=&arnumber=7535098",
						HighlightedTitle:        "A Taxonomy of Domain-Generation Algorithms",
						PublicationTitle:        "IEEE Security & Privacy",
						DisplayPublicationTitle: "IEEE Security & Privacy",
						IsStandard:              false,
						IsConference:            false,
						IsJournalAndMagazine:    true,
						IsMagazine:              false,
						IsJournal:               false,
						IsBook:                  false,
						Course:                  false,
						IsBookWithoutChapters:   false,
						DisplayContentType:      "Magazine Article",
						DocIdentifier:           "IEEE Magazines",
						IsEarlyAccess:           false,
					},
				},
				BreadCrumbs: []goieeeapi.BreadCrumb{
					{
						Type: "search",
						Children: &[]goieeeapi.BreadCrumbChild{
							{
								Value:     "Botnet detection",
								Reference: str("Botnet detection"),
								Display:   str("Botnet detection"),
							},
						},
					}, {
						Type:  "refinement",
						Key:   str("Content Type"),
						Value: str("ContentType:"),
						Children: &[]goieeeapi.BreadCrumbChild{
							{
								Type:  str("refinement"),
								Key:   str("Magazines"),
								Value: "ContentType:Magazines",
							}, {
								Type:  str("refinement"),
								Key:   str("Early Access Articles"),
								Value: "ContentType:Early Access Articles",
							},
						},
					}, {
						Type:  "refinement",
						Key:   str("Publication Topics"),
						Value: str("ControlledTerms:"),
						Children: &[]goieeeapi.BreadCrumbChild{
							{
								Type:  str("refinement"),
								Key:   str("Web sites"),
								Value: "ControlledTerms:Web sites",
							},
						},
					}, {
						Type:      "range",
						Key:       str("Year"),
						Reference: str("2012_2016_Year"),
						Start:     str("2012"),
						End:       str("2016"),
					},
				},
				ShowStandardDictionary:   false,
				SearchType:               "basic",
				TotalRecords:             2,
				SubscribedContentApplied: false,
				PromoApplied:             false,
				StartRecord:              0,
				Facets: &[]goieeeapi.Facet{
					{
						ID:         "ContentType",
						Name:       "Content Type",
						NumRecords: -1,
						Children: &[]goieeeapi.Facet{
							{
								ID:         "ContentType:Conferences",
								Name:       "Conferences",
								NumRecords: 6,
								Active:     "true",
							}, {
								ID:         "ContentType:Journals",
								Name:       "Journals",
								NumRecords: 1,
								Active:     "true",
							},
						},
						Active: "true",
					}, {
						ID:         "PublicationYear",
						Name:       "Year",
						NumRecords: -1,
						Children: &[]goieeeapi.Facet{
							{
								ID:         "PublicationYear:2013",
								Name:       "2013",
								NumRecords: 1,
								Active:     "true",
							}, {
								ID:         "PublicationYear:2016",
								Name:       "2016",
								NumRecords: 1,
								Active:     "true",
							},
						},
						Active: "true",
					}, {
						ID:         "Author",
						Name:       "Author",
						NumRecords: -1,
						Children: &[]goieeeapi.Facet{
							{
								ID:         "Author:Aditya K. Sood",
								Name:       "Aditya K. Sood",
								NumRecords: 1,
								Active:     "true",
							}, {
								ID:         "Author:Sherali Zeadally",
								Name:       "Sherali Zeadally",
								NumRecords: 1,
								Active:     "true",
							}, {
								ID:         "Author:Zou Futai",
								Name:       "Zou Futai",
								NumRecords: 1,
								Active:     "true",
							}, {
								ID:         "Author:Zhang Siyu",
								Name:       "Zhang Siyu",
								NumRecords: 1,
								Active:     "true",
							}, {
								ID:         "Author:Rao Weixiong",
								Name:       "Rao Weixiong",
								NumRecords: 1,
								Active:     "true",
							},
						},
						Active: "true",
					}, {
						ID:         "Affiliation",
						Name:       "Affiliation",
						NumRecords: -1,
						Children: &[]goieeeapi.Facet{
							{
								ID:         "Affiliation:School of Software Engineering, Tongji University, Shanghai 200092, China",
								Name:       "School of Software Engineering, Tongji University, Shanghai 200092, China",
								NumRecords: 1,
								Active:     "true",
							}, {
								ID:         "Affiliation:Elastica",
								Name:       "Elastica",
								NumRecords: 1,
								Active:     "true",
							}, {
								ID:         "Affiliation:School of Information Security Engineering, Shanghai Jiao Tong University, Shanghai 200240, China",
								Name:       "School of Information Security Engineering, Shanghai Jiao Tong University, Shanghai 200240, China",
								NumRecords: 1,
								Active:     "true",
							}, {
								ID:         "Affiliation:University of Kentucky",
								Name:       "University of Kentucky",
								NumRecords: 1,
								Active:     "true",
							},
						},
						Active: "true",
					}, {
						ID:         "PublicationTitle",
						Name:       "Publication Title",
						NumRecords: -1,
						Children: &[]goieeeapi.Facet{
							{
								ID:         "PublicationTitle:China Communications",
								Name:       "China Communications",
								NumRecords: 1,
								Active:     "true",
							}, {
								ID:         "PublicationTitle:IEEE Security .AND. Privacy",
								Name:       "IEEE Security & Privacy",
								NumRecords: 1,
								Active:     "true",
							},
						},
						Active: "true",
					}, {
						ID:         "Publisher",
						Name:       "Publisher",
						NumRecords: -1,
						Children: &[]goieeeapi.Facet{
							{
								ID:         "Publisher:IEEE",
								Name:       "IEEE",
								NumRecords: 2,
								Active:     "true",
							},
						},
						Active: "true",
					}, {
						ID:         "ControlledTerms",
						Name:       "Publication Topics",
						NumRecords: -1,
						Children: &[]goieeeapi.Facet{
							{
								ID:         "ControlledTerms:computer network security",
								Name:       "computer network security",
								NumRecords: 3,
								Active:     "true",
							}, {
								ID:         "ControlledTerms:invasive software",
								Name:       "invasive software",
								NumRecords: 3,
								Active:     "true",
							}, {
								ID:         "ControlledTerms:Internet",
								Name:       "Internet",
								NumRecords: 2,
								Active:     "true",
							}, {
								ID:         "ControlledTerms:security of data",
								Name:       "security of data",
								NumRecords: 2,
								Active:     "true",
							}, {
								ID:         "ControlledTerms:artificial intelligence",
								Name:       "artificial intelligence",
								NumRecords: 1,
								Active:     "true",
							}, {
								ID:         "ControlledTerms:cloud computing",
								Name:       "cloud computing",
								NumRecords: 1,
								Active:     "true",
							}, {
								ID:         "ControlledTerms:command and control systems",
								Name:       "command and control systems",
								NumRecords: 1,
								Active:     "true",
							}, {
								ID:         "ControlledTerms:digital forensics",
								Name:       "digital forensics",
								NumRecords: 1,
								Active:     "true",
							}, {
								ID:         "ControlledTerms:distributed sensors",
								Name:       "distributed sensors",
								NumRecords: 1,
								Active:     "true",
							}, {
								ID:         "ControlledTerms:fractals",
								Name:       "fractals",
								NumRecords: 1,
								Active:     "true",
							}, {
								ID:         "ControlledTerms:genetic algorithms",
								Name:       "genetic algorithms",
								NumRecords: 1,
								Active:     "true",
							}, {
								ID:         "ControlledTerms:inference mechanisms",
								Name:       "inference mechanisms",
								NumRecords: 1,
								Active:     "true",
							}, {
								ID:         "ControlledTerms:military computing",
								Name:       "military computing",
								NumRecords: 1,
								Active:     "true",
							}, {
								ID:         "ControlledTerms:network servers",
								Name:       "network servers",
								NumRecords: 1,
								Active:     "true",
							}, {
								ID:         "ControlledTerms:neural nets",
								Name:       "neural nets",
								NumRecords: 1,
								Active:     "true",
							}, {
								ID:         "ControlledTerms:peer-to-peer computing",
								Name:       "peer-to-peer computing",
								NumRecords: 1,
								Active:     "true",
							}, {
								ID:         "ControlledTerms:resource allocation",
								Name:       "resource allocation",
								NumRecords: 1,
								Active:     "true",
							}, {
								ID:         "ControlledTerms:telecommunication traffic",
								Name:       "telecommunication traffic",
								NumRecords: 1,
								Active:     "true",
							}, {
								ID:         "ControlledTerms:uncertainty handling",
								Name:       "uncertainty handling",
								NumRecords: 1,
								Active:     "true",
							},
						},
						Active: "true",
					},
				},
				EndRecord:  2,
				TotalPages: 1,
			},
			ExpectedErr: nil,
		},
	}

	for testname, tt := range tests {
		t.Run(testname, func(t *testing.T) {
			rslts, err := goieeeapi.PostSearch(tt.Client, tt.Params)

			if !reflect.DeepEqual(rslts, tt.ExpectedResults) {
				t.Errorf("Failed to get expected search results: got \"%v\" instead of \"%v\".", rslts, tt.ExpectedResults)
			}
			checkErr(err, tt.ExpectedErr, t)
		})
	}
}
