package goieeeapi_test

import (
	"encoding/json"
	"net/http"
	"reflect"
	"testing"

	goieeeapi "github.com/pandatix/go-ieee-api"
)

func TestGetReferences(t *testing.T) {
	t.Parallel()

	var tests = map[string]struct {
		Client             goieeeapi.HTTPClient
		ID                 int
		ExpectedReferences *goieeeapi.GetReferencesResponse
		ExpectedErr        error
	}{
		"nil-client": {
			Client:             nil,
			ID:                 0,
			ExpectedReferences: nil,
			ExpectedErr:        goieeeapi.ErrNilClient,
		},
		"failing-client": {
			Client:             newFakeHTTPClient(``, 0, errFake),
			ID:                 0,
			ExpectedReferences: nil,
			ExpectedErr:        errFake,
		},
		"unexpected-statuscode": {
			Client:             newFakeHTTPClient(``, 0, nil),
			ID:                 0,
			ExpectedReferences: nil,
			ExpectedErr: &goieeeapi.ErrUnexpectedStatus{
				StatusCode: 0,
				Body:       []byte(``),
			},
		},
		"failing-unmarshal": {
			Client:             newFakeHTTPClient(`{[}]`, http.StatusOK, nil),
			ID:                 0,
			ExpectedReferences: nil,
			ExpectedErr: &json.SyntaxError{
				Offset: 2,
			},
		},
		"valid-call": {
			Client: newFakeHTTPClient(`{"userInfo":{"institute":false,"member":false,"individual":false,"guest":false,"subscribedContent":false,"fileCabinetContent":false,"fileCabinetUser":false,"institutionalFileCabinetUser":false,"showPatentCitations":true,"showGet802Link":false,"showOpenUrlLink":false,"tracked":false,"delegatedAdmin":false,"desktop":false,"isInstitutionDashboardEnabled":false,"isInstitutionProfileEnabled":false,"isRoamingEnabled":false,"isDelegatedAdmin":false,"isMdl":false,"isCwg":false},"references":[{"order":"1","text":"A. Chakraborty, M. V. R. Seshasai, C. Sudhakar Reddy and V. K. Dadhwal, \"Persistent negative changes in seasonal greenness over different forest types of india using modis time series ndvi data (20012014)\", <em>Ecological Indicators</em>, vol. 85, pp. 887-903, 2018.","title":"Persistent negative changes in seasonal greenness over different forest types of india using modis time series ndvi data (20012014)","context":[{"sec":"sec1","text":" The Normalized Difference Vegetation Index (NDVI) is a notable example of a standard and simple vegetation indicator that is extracted through a straightforward combination of spectral bands [1].","part":"1"}],"links":{"crossRefLink":"https://doi.org/10.1016/j.ecolind.2017.11.032","openUrlImgLoc":"/assets/img/btn.find-in-library.png"},"googleScholarLink":"https://scholar.google.com/scholar?as_q=Persistent+negative+changes+in+seasonal+greenness+over+different+forest+types+of+india+using+modis+time+series+ndvi+data+%2820012014%29&as_occt=title&hl=en&as_sdt=0%2C31","refType":"biblio","id":"ref1"},{"order":"2","text":"V. J. Pasquarella, C. E. Holden and C. E. Woodcock, \"Improved mapping of forest type using spectral-temporal landsat features\", <em>Remote Sensing of Environment</em>, vol. 210, pp. 193-207, 2018.","title":"Improved mapping of forest type using spectral-temporal landsat features","context":[{"sec":"sec1","text":" More specific indicators can also be derived from multispectral images, like the Enhanced Vegetation Index (EVI), more suited to discriminate canopy [2].","part":"1"}],"links":{"crossRefLink":"https://doi.org/10.1016/j.rse.2018.02.064","openUrlImgLoc":"/assets/img/btn.find-in-library.png"},"googleScholarLink":"https://scholar.google.com/scholar?as_q=Improved+mapping+of+forest+type+using+spectral-temporal+landsat+features&as_occt=title&hl=en&as_sdt=0%2C31","refType":"biblio","id":"ref2"},{"order":"3","text":"J. Inglada et al., \"Assessment of an operational system for crop type map production using high temporal and spatial resolution satellite optical imagery\", <em>Remote Sensing</em>, vol. 7, no. 9, pp. 12356-12379, 2015.","title":"Assessment of an operational system for crop type map production using high temporal and spatial resolution satellite optical imagery","context":[{"sec":"sec1","text":" However, the use of optical data is severely undermined by their dependence on the weather conditions, which can be only partially mitigated through multitemporal processing and data fusion techniques [3], [4], [5].","part":"1"}],"googleScholarLink":"https://scholar.google.com/scholar?as_q=Assessment+of+an+operational+system+for+crop+type+map+production+using+high+temporal+and+spatial+resolution+satellite+optical+imagery&as_occt=title&hl=en&as_sdt=0%2C31","refType":"biblio","id":"ref3"},{"order":"4","text":"G. Scarpa, M. Gargiulo, A. Mazza and R. Gaetano, \"A CNN-Based Fusion Method for Feature Extraction from Sentinel Data\", <em>Remote Sensing</em>, vol. 10, no. 2, 2018.","title":"A CNN-Based Fusion Method for Feature Extraction from Sentinel Data","context":[{"sec":"sec1","text":" However, the use of optical data is severely undermined by their dependence on the weather conditions, which can be only partially mitigated through multitemporal processing and data fusion techniques [3], [4], [5].","part":"1"}],"googleScholarLink":"https://scholar.google.com/scholar?as_q=A+CNN-Based+Fusion+Method+for+Feature+Extraction+from+Sentinel+Data&as_occt=title&hl=en&as_sdt=0%2C31","refType":"biblio","id":"ref4"},{"order":"5","text":"A. Errico, C. V. Angelino, L. Cicala, D. P. Podobinski, G. Persechino, C. Ferrara, et al., \"SAR/multispectral image fusion for the detection of environmental hazards with a gis\", <em>Proceedings of SPIE - The International Society for Optical Engineering</em>, vol. 9245, 2014.","title":"SAR/multispectral image fusion for the detection of environmental hazards with a gis","context":[{"sec":"sec1","text":" However, the use of optical data is severely undermined by their dependence on the weather conditions, which can be only partially mitigated through multitemporal processing and data fusion techniques [3], [4], [5].","part":"1"}],"googleScholarLink":"https://scholar.google.com/scholar?as_q=SAR%2Fmultispectral+image+fusion+for+the+detection+of+environmental+hazards+with+a+gis&as_occt=title&hl=en&as_sdt=0%2C31","refType":"biblio","id":"ref5"},{"order":"6","text":"R. Gaetano, D. Amitrano, G. Masi, G. Poggi, G. Ruello, L. Verdoliva, et al., \"Exploration of multitemporal COSMO-skymed data via interactive tree-structured MRF segmentation\", <em>IEEE Journal of Selected Topics in Applied Earth Observations and Remote Sensing</em>, vol. 7, no. 7, pp. 2763-2775, 2014.","title":"Exploration of multitemporal COSMO-skymed data via interactive tree-structured MRF segmentation","context":[{"sec":"sec1","text":" On the contrary, SAR data are almost weather insensitive and carry precious information related to ground geometry and electromagnetic propagation [6].","part":"1"}],"links":{"documentLink":"/document/6881820","pdfLink":"/stamp/stamp.jsp?tp=&arnumber=6881820","abstract":"We propose a new approach for remote sensing data exploration, based on a tight human-machine interaction. The analyst uses a number of powerful and user-friendly image classification/segmentation tools to obtain a satisfactory thematic map, based only on visual assessment and expertise. All processing tools are in the framework of the tree-structured MRF model, which allows for a flexible and spatially adaptive description of the data. We test the proposed approach for the exploration of multit...","openUrlImgLoc":"/assets/img/btn.find-in-library.png","pdfSize":"3176KB"},"googleScholarLink":"https://scholar.google.com/scholar?as_q=Exploration+of+multitemporal+COSMO-skymed+data+via+interactive+tree-structured+MRF+segmentation&as_occt=title&hl=en&as_sdt=0%2C31","refType":"biblio","id":"ref6"},{"order":"7","text":"R. Hagensieker and B. Waske, \"Evaluation of multi-frequency sar images for tropical land cover mapping\", <em>Remote Sensing</em>, vol. 10, no. 2, 2018.","title":"Evaluation of multi-frequency sar images for tropical land cover mapping","context":[{"sec":"sec1","text":" In [7] SAR images obtained in different bands are combined for land cover classification.","part":"1"}],"links":{"crossRefLink":"https://doi.org/10.3390/rs10020257","openUrlImgLoc":"/assets/img/btn.find-in-library.png"},"googleScholarLink":"https://scholar.google.com/scholar?as_q=Evaluation+of+multi-frequency+sar+images+for+tropical+land+cover+mapping&as_occt=title&hl=en&as_sdt=0%2C31","refType":"biblio","id":"ref7"},{"order":"8","text":"M. Martone, P. Rizzoli, C. Wecklich, C. Gonzlez, J.-L. Bueso-Bello, P. Valdo, et al., \"The global forest/non-forest map from tandem-x interferometric SAR data\", <em>Remote Sensing of Environment</em>, vol. 205, pp. 352-373, 2018.","title":"The global forest/non-forest map from tandem-x interferometric SAR data","context":[{"sec":"sec1","text":" In [8], [9], the TanDEM-X forest/non-forest map is generated from TanDEM-X bistatic interferometric images, by linking the presence of vegetation to the retrieved InSAR volume decorrelation.","part":"1"},{"sec":"sec2a","text":"In this work we exploit the same dataset used in [8], described in Section 3, including the ground-truth reference which is given in terms of density of forest in a squared area of 6 × 6 meters.","part":"1"},{"sec":"sec3","text":" Following the methodology used in [8] we have chosen the accuray indicator ACC, a widespread quality index for binary classification problems, defined as\n\nwhere TP, TN, FP, and FN count true positive, true negative, false positive, and false negative pixels, respectively.","part":"1"},{"sec":"sec3","text":" To this end we followed the same criterion used in [8], maximizing the Pearson coefficient ϕ with respect to the threshold pair (one for prediction, one for reference).","part":"1"},{"sec":"sec3","text":" The threshold pair that corresponds to the maximum value of ϕ is the optimal choice according to this criterion (see [8] for a deeper discussion).","part":"1"},{"sec":"sec3","text":" It has to be remarked that for a fair comparison with the proposed methods, the baseline solution of [8] was used without masking any class, contrarily to what is done in the original formulation.","part":"1"},{"sec":"sec3","text":" Specifically, in [8] city and water classes are excluded by means of available masks, because forests, cities and water classes all exhibit a low volume correlation, the core feature proposed to classify forests.","part":"1"}],"links":{"crossRefLink":"https://doi.org/10.1016/j.rse.2017.12.002","openUrlImgLoc":"/assets/img/btn.find-in-library.png"},"googleScholarLink":"https://scholar.google.com/scholar?as_q=The+global+forest%2Fnon-forest+map+from+tandem-x+interferometric+SAR+data&as_occt=title&hl=en&as_sdt=0%2C31","refType":"biblio","id":"ref8"},{"order":"9","text":"M. Martone, F. Sica, C. Gonzlez, J.-L. Bueso-Bello, P. Valdo and P. Rizzoli, \"High-resolution forest mapping from tandem-x interferometric data exploiting nonlocal filtering\", <em>Remote Sensing</em>, vol. 10, pp. 1477, 2018.","title":"High-resolution forest mapping from tandem-x interferometric data exploiting nonlocal filtering","context":[{"sec":"sec1","text":" In [8], [9], the TanDEM-X forest/non-forest map is generated from TanDEM-X bistatic interferometric images, by linking the presence of vegetation to the retrieved InSAR volume decorrelation.","part":"1"}],"googleScholarLink":"https://scholar.google.com/scholar?as_q=High-resolution+forest+mapping+from+tandem-x+interferometric+data+exploiting+nonlocal+filtering&as_occt=title&hl=en&as_sdt=0%2C31","refType":"biblio","id":"ref9"},{"order":"10","text":"K. He, X. Zhang, S. Ren and J. Sun, \"Deep residual learning for image recognition\", <em>CVPR</em>, June 2016.","title":"Deep residual learning for image recognition","context":[{"sec":"sec1","text":" Specifically, we define and train from scratch two DL architectures following the ResNet [10] and the DenseNet [11] models, respectively.","part":"1"},{"sec":"sec2","text":" In this work, we consider two state-of-the-art CNN models, ResNet [10] and DenseNet [11], which are particularly appealing as they can be reach a considerable depth avoiding vanishing gradient problems during training.","part":"1"},{"sec":"sec2","text":" For both models, we describe here only the main functional aspects of interest for the present work, referring to the original papers [10], [11] for a thorough description of the network architecture.","part":"1"},{"sec":"sec2","text":" Although functionally unnecessary, skip connections have proven to speed-up the training [10], [12].","part":"1"}],"links":{"documentLink":"/document/7780459","pdfLink":"/stamp/stamp.jsp?tp=&arnumber=7780459","abstract":"Deeper neural networks are more difficult to train. We present a residual learning framework to ease the training of networks that are substantially deeper than those used previously. We explicitly reformulate the layers as learning residual functions with reference to the layer inputs, instead of learning unreferenced functions. We provide comprehensive empirical evidence showing that these residual networks are easier to optimize, and can gain accuracy from considerably increased depth. On the...","openUrlImgLoc":"/assets/img/btn.find-in-library.png","pdfSize":"282KB"},"googleScholarLink":"https://scholar.google.com/scholar?as_q=Deep+residual+learning+for+image+recognition&as_occt=title&hl=en&as_sdt=0%2C31","refType":"biblio","id":"ref10"},{"order":"11","text":"G. Huang, Z. Liu, L. Van Der Maaten and K. Q. Weinberger, \"Densely connected convolutional networks\", <em>CVPR</em>, 2017.","title":"Densely connected convolutional networks","context":[{"sec":"sec1","text":" Specifically, we define and train from scratch two DL architectures following the ResNet [10] and the DenseNet [11] models, respectively.","part":"1"},{"sec":"sec2","text":" In this work, we consider two state-of-the-art CNN models, ResNet [10] and DenseNet [11], which are particularly appealing as they can be reach a considerable depth avoiding vanishing gradient problems during training.","part":"1"},{"sec":"sec2","text":" For both models, we describe here only the main functional aspects of interest for the present work, referring to the original papers [10], [11] for a thorough description of the network architecture.","part":"1"}],"links":{"documentLink":"/document/8099726","pdfLink":"/stamp/stamp.jsp?tp=&arnumber=8099726","abstract":"Recent work has shown that convolutional networks can be substantially deeper, more accurate, and efficient to train if they contain shorter connections between layers close to the input and those close to the output. In this paper, we embrace this observation and introduce the Dense Convolutional Network (DenseNet), which connects each layer to every other layer in a feed-forward fashion. Whereas traditional convolutional networks with L layers have L connections-one between each layer and its ...","openUrlImgLoc":"/assets/img/btn.find-in-library.png","pdfSize":"395KB"},"googleScholarLink":"https://scholar.google.com/scholar?as_q=Densely+connected+convolutional+networks&as_occt=title&hl=en&as_sdt=0%2C31","refType":"biblio","id":"ref11"},{"order":"12","text":"G. Scarpa, S. Vitale and D. Cozzolino, \"Target-adaptive cnn-based pansharpening\", <em>IEEE Transactions on Geoscience and Remote Sensing</em>, vol. 56, no. 9, pp. 5443-5457, Sep. 2018.","title":"Target-adaptive cnn-based pansharpening","context":[{"sec":"sec2","text":" Although functionally unnecessary, skip connections have proven to speed-up the training [10], [12].","part":"1"}],"links":{"documentLink":"/document/8334206","pdfLink":"/stamp/stamp.jsp?tp=&arnumber=8334206","abstract":"We recently proposed a convolutional neural network (CNN) for remote sensing image pansharpening obtaining a significant performance gain over the state of the art. In this paper, we explore a number of architectural and training variations to this baseline, achieving further performance gains with a lightweight network that trains very fast. Leveraging on this latter property, we propose a target-adaptive usage modality that ensures a very good performance also in the presence of a mismatch wit...","openUrlImgLoc":"/assets/img/btn.find-in-library.png","pdfSize":"4862KB"},"googleScholarLink":"https://scholar.google.com/scholar?as_q=Target-adaptive+cnn-based+pansharpening&as_occt=title&hl=en&as_sdt=0%2C31","refType":"biblio","id":"ref12"},{"order":"13","text":"A. Krizhevsky, I. Sutskever and G. E Hinton, \"Imagenet classification with deep convolutional neural networks\", <em>Advances in neural information processing systems</em>, pp. 1097-1105, 2012.","title":"Imagenet classification with deep convolutional neural networks","context":[{"sec":"sec2","text":"Here we propose for both models a cascade architecture with six convolutional layers with 3 × 3 kernels interleaved by ReLU (Rectified Linear Unit) activation functions [13].","part":"1"}],"googleScholarLink":"https://scholar.google.com/scholar?as_q=Imagenet+classification+with+deep+convolutional+neural+networks&as_occt=title&hl=en&as_sdt=0%2C31","refType":"biblio","id":"ref13"},{"order":"14","text":"G. Csurka, D. Larlus, F. Perronnin and F. Meylan, \"What is a good evaluation measure for semantic segmentation?\", <em>BMVC</em>, vol. 27, pp. 2013, 2013.","title":"What is a good evaluation measure for semantic segmentation?","context":[{"sec":"sec2a","text":" The Jaccard distance loss, which aims to maximize the overlap between the two maps [14], is defined as\n.","part":"1"}],"links":{"crossRefLink":"https://doi.org/10.5244/C.27.32","openUrlImgLoc":"/assets/img/btn.find-in-library.png"},"googleScholarLink":"https://scholar.google.com/scholar?as_q=What+is+a+good+evaluation+measure+for+semantic+segmentation%3F&as_occt=title&hl=en&as_sdt=0%2C31","refType":"biblio","id":"ref14"},{"order":"15","text":"Diederik P. Kingma and Jimmy Ba, \"Adam: A method for stochastic optimization\", <em>CoRR</em>, 2014.","title":"Adam: A method for stochastic optimization","context":[{"sec":"sec2a","text":"The minimization of the loss function is performed using the ADAM algorithm [15], a gradient descend variant where the learning rate is updated at each iteration using estimates of low-order moments.","part":"1"}],"googleScholarLink":"https://scholar.google.com/scholar?as_q=Adam%3A+A+method+for+stochastic+optimization&as_occt=title&hl=en&as_sdt=0%2C31","refType":"biblio","id":"ref15"}],"articleNumber":"8900441","getProgramTermsAccepted":false,"formulaStrippedArticleTitle":"Deep Learning Solutions for Tandem-X-Based Forest Classification","pubLink":"/xpl/conhome/8891871/proceeding","allowComments":false,"issueLink":"/xpl/tocresult.jsp?isnumber=null","isReadingRoomArticle":false,"isGetArticle":false,"isGetAddressInfoCaptured":false,"isMarketingOptIn":false,"publisher":"IEEE","xploreDocumentType":"Conference Publication","isPromo":false,"isNotDynamicOrStatic":false,"htmlAbstractLink":"/document/8900441/","isCustomDenial":false,"isSAE":false,"isDynamicHtml":true,"isFreeDocument":false,"displayDocTitle":"Deep Learning Solutions for Tandem-X-Based Forest Classification","isStandard":false,"isSMPTE":false,"isOUP":false,"isNow":false,"isProduct":false,"isMorganClaypool":false,"isJournal":false,"isBook":false,"isBookWithoutChapters":false,"isOpenAccess":false,"isEphemera":false,"isConference":true,"htmlLink":"/document/8900441/","isChapter":false,"isStaticHtml":false,"isEarlyAccess":false,"persistentLink":"https://ieeexplore.ieee.org/servlet/opac?punumber=8891871","articleId":"8900441","openAccessFlag":"F","ephemeraFlag":"false","title":"Deep Learning Solutions for Tandem-X-Based Forest Classification","contentTypeDisplay":"Conferences","html_flag":"false","ml_html_flag":"true","mlTime":"PT0.158217S","lastupdate":"2021-08-21","mediaPath":"/mediastore_new/IEEE/content/media/8891871/8897702/8900441","contentType":"conferences","definitions":"false","publicationNumber":"8891871"}`, http.StatusOK, nil),
			ID:     0,
			ExpectedReferences: &goieeeapi.GetReferencesResponse{
				UserInfo: goieeeapi.UserInfo{
					Institute:                     false,
					Member:                        false,
					Individual:                    false,
					Guest:                         false,
					SubscribedContent:             false,
					FileCabinetContent:            false,
					FileCabinetUser:               false,
					InstitutionalFileCabinetUser:  false,
					ShowPatentCitations:           true,
					ShowGet802Link:                false,
					ShowOpenURLLink:               b(false),
					Tracked:                       b(false),
					DelegatedAdmin:                b(false),
					Desktop:                       b(false),
					IsInstitutionDashboardEnabled: b(false),
					IsInstitutionProfileEnabled:   b(false),
					IsRoamingEnabled:              b(false),
					IsDelegatedAdmin:              b(false),
					IsMdl:                         b(false),
					IsCwg:                         b(false),
				},
				References: &[]goieeeapi.Reference{
					{
						Order: "1",
						Text:  str("A. Chakraborty, M. V. R. Seshasai, C. Sudhakar Reddy and V. K. Dadhwal, \"Persistent negative changes in seasonal greenness over different forest types of india using modis time series ndvi data (20012014)\", <em>Ecological Indicators</em>, vol. 85, pp. 887-903, 2018."),
						Title: str("Persistent negative changes in seasonal greenness over different forest types of india using modis time series ndvi data (20012014)"),
						Context: &[]goieeeapi.Context{
							{
								Sec:  "sec1",
								Text: " The Normalized Difference Vegetation Index (NDVI) is a notable example of a standard and simple vegetation indicator that is extracted through a straightforward combination of spectral bands [1].",
								Part: str("1"),
							},
						},
						Links: &goieeeapi.Links{
							CrossRefLink:  str("https://doi.org/10.1016/j.ecolind.2017.11.032"),
							OpenURLImgLoc: "/assets/img/btn.find-in-library.png",
						},
						GoogleScholarLink: str("https://scholar.google.com/scholar?as_q=Persistent+negative+changes+in+seasonal+greenness+over+different+forest+types+of+india+using+modis+time+series+ndvi+data+%2820012014%29&as_occt=title&hl=en&as_sdt=0%2C31"),
						RefType:           str("biblio"),
						ID:                str("ref1"),
					}, {
						Order: "2",
						Text:  str("V. J. Pasquarella, C. E. Holden and C. E. Woodcock, \"Improved mapping of forest type using spectral-temporal landsat features\", <em>Remote Sensing of Environment</em>, vol. 210, pp. 193-207, 2018."),
						Title: str("Improved mapping of forest type using spectral-temporal landsat features"),
						Context: &[]goieeeapi.Context{
							{
								Sec:  "sec1",
								Text: " More specific indicators can also be derived from multispectral images, like the Enhanced Vegetation Index (EVI), more suited to discriminate canopy [2].",
								Part: str("1"),
							},
						},
						Links: &goieeeapi.Links{
							CrossRefLink:  str("https://doi.org/10.1016/j.rse.2018.02.064"),
							OpenURLImgLoc: "/assets/img/btn.find-in-library.png",
						},
						GoogleScholarLink: str("https://scholar.google.com/scholar?as_q=Improved+mapping+of+forest+type+using+spectral-temporal+landsat+features&as_occt=title&hl=en&as_sdt=0%2C31"),
						RefType:           str("biblio"),
						ID:                str("ref2"),
					}, {
						Order: "3",
						Text:  str("J. Inglada et al., \"Assessment of an operational system for crop type map production using high temporal and spatial resolution satellite optical imagery\", <em>Remote Sensing</em>, vol. 7, no. 9, pp. 12356-12379, 2015."),
						Title: str("Assessment of an operational system for crop type map production using high temporal and spatial resolution satellite optical imagery"),
						Context: &[]goieeeapi.Context{
							{
								Sec:  "sec1",
								Text: " However, the use of optical data is severely undermined by their dependence on the weather conditions, which can be only partially mitigated through multitemporal processing and data fusion techniques [3], [4], [5].",
								Part: str("1"),
							},
						},
						GoogleScholarLink: str("https://scholar.google.com/scholar?as_q=Assessment+of+an+operational+system+for+crop+type+map+production+using+high+temporal+and+spatial+resolution+satellite+optical+imagery&as_occt=title&hl=en&as_sdt=0%2C31"),
						RefType:           str("biblio"),
						ID:                str("ref3"),
					}, {
						Order: "4",
						Text:  str("G. Scarpa, M. Gargiulo, A. Mazza and R. Gaetano, \"A CNN-Based Fusion Method for Feature Extraction from Sentinel Data\", <em>Remote Sensing</em>, vol. 10, no. 2, 2018."),
						Title: str("A CNN-Based Fusion Method for Feature Extraction from Sentinel Data"),
						Context: &[]goieeeapi.Context{
							{
								Sec:  "sec1",
								Text: " However, the use of optical data is severely undermined by their dependence on the weather conditions, which can be only partially mitigated through multitemporal processing and data fusion techniques [3], [4], [5].",
								Part: str("1"),
							},
						},
						GoogleScholarLink: str("https://scholar.google.com/scholar?as_q=A+CNN-Based+Fusion+Method+for+Feature+Extraction+from+Sentinel+Data&as_occt=title&hl=en&as_sdt=0%2C31"),
						RefType:           str("biblio"),
						ID:                str("ref4"),
					}, {
						Order: "5",
						Text:  str("A. Errico, C. V. Angelino, L. Cicala, D. P. Podobinski, G. Persechino, C. Ferrara, et al., \"SAR/multispectral image fusion for the detection of environmental hazards with a gis\", <em>Proceedings of SPIE - The International Society for Optical Engineering</em>, vol. 9245, 2014."),
						Title: str("SAR/multispectral image fusion for the detection of environmental hazards with a gis"),
						Context: &[]goieeeapi.Context{
							{
								Sec:  "sec1",
								Text: " However, the use of optical data is severely undermined by their dependence on the weather conditions, which can be only partially mitigated through multitemporal processing and data fusion techniques [3], [4], [5].",
								Part: str("1"),
							},
						},
						GoogleScholarLink: str("https://scholar.google.com/scholar?as_q=SAR%2Fmultispectral+image+fusion+for+the+detection+of+environmental+hazards+with+a+gis&as_occt=title&hl=en&as_sdt=0%2C31"),
						RefType:           str("biblio"),
						ID:                str("ref5"),
					}, {
						Order: "6",
						Text:  str("R. Gaetano, D. Amitrano, G. Masi, G. Poggi, G. Ruello, L. Verdoliva, et al., \"Exploration of multitemporal COSMO-skymed data via interactive tree-structured MRF segmentation\", <em>IEEE Journal of Selected Topics in Applied Earth Observations and Remote Sensing</em>, vol. 7, no. 7, pp. 2763-2775, 2014."),
						Title: str("Exploration of multitemporal COSMO-skymed data via interactive tree-structured MRF segmentation"),
						Context: &[]goieeeapi.Context{
							{
								Sec:  "sec1",
								Text: " On the contrary, SAR data are almost weather insensitive and carry precious information related to ground geometry and electromagnetic propagation [6].",
								Part: str("1"),
							},
						},
						Links: &goieeeapi.Links{
							DocumentLink:  str("/document/6881820"),
							PdfLink:       str("/stamp/stamp.jsp?tp=&arnumber=6881820"),
							Abstract:      str("We propose a new approach for remote sensing data exploration, based on a tight human-machine interaction. The analyst uses a number of powerful and user-friendly image classification/segmentation tools to obtain a satisfactory thematic map, based only on visual assessment and expertise. All processing tools are in the framework of the tree-structured MRF model, which allows for a flexible and spatially adaptive description of the data. We test the proposed approach for the exploration of multit..."),
							OpenURLImgLoc: "/assets/img/btn.find-in-library.png",
							PdfSize:       str("3176KB"),
						},
						GoogleScholarLink: str("https://scholar.google.com/scholar?as_q=Exploration+of+multitemporal+COSMO-skymed+data+via+interactive+tree-structured+MRF+segmentation&as_occt=title&hl=en&as_sdt=0%2C31"),
						RefType:           str("biblio"),
						ID:                str("ref6"),
					}, {
						Order: "7",
						Text:  str("R. Hagensieker and B. Waske, \"Evaluation of multi-frequency sar images for tropical land cover mapping\", <em>Remote Sensing</em>, vol. 10, no. 2, 2018."),
						Title: str("Evaluation of multi-frequency sar images for tropical land cover mapping"),
						Context: &[]goieeeapi.Context{
							{
								Sec:  "sec1",
								Text: " In [7] SAR images obtained in different bands are combined for land cover classification.",
								Part: str("1"),
							},
						},
						Links: &goieeeapi.Links{
							CrossRefLink:  str("https://doi.org/10.3390/rs10020257"),
							OpenURLImgLoc: "/assets/img/btn.find-in-library.png",
						},
						GoogleScholarLink: str("https://scholar.google.com/scholar?as_q=Evaluation+of+multi-frequency+sar+images+for+tropical+land+cover+mapping&as_occt=title&hl=en&as_sdt=0%2C31"),
						RefType:           str("biblio"),
						ID:                str("ref7"),
					}, {
						Order: "8",
						Text:  str("M. Martone, P. Rizzoli, C. Wecklich, C. Gonzlez, J.-L. Bueso-Bello, P. Valdo, et al., \"The global forest/non-forest map from tandem-x interferometric SAR data\", <em>Remote Sensing of Environment</em>, vol. 205, pp. 352-373, 2018."),
						Title: str("The global forest/non-forest map from tandem-x interferometric SAR data"),
						Context: &[]goieeeapi.Context{
							{
								Sec:  "sec1",
								Text: " In [8], [9], the TanDEM-X forest/non-forest map is generated from TanDEM-X bistatic interferometric images, by linking the presence of vegetation to the retrieved InSAR volume decorrelation.",
								Part: str("1"),
							}, {
								Sec:  "sec2a",
								Text: "In this work we exploit the same dataset used in [8], described in Section 3, including the ground-truth reference which is given in terms of density of forest in a squared area of 6 × 6 meters.",
								Part: str("1"),
							}, {
								Sec:  "sec3",
								Text: " Following the methodology used in [8] we have chosen the accuray indicator ACC, a widespread quality index for binary classification problems, defined as\n\nwhere TP, TN, FP, and FN count true positive, true negative, false positive, and false negative pixels, respectively.",
								Part: str("1"),
							}, {
								Sec:  "sec3",
								Text: " To this end we followed the same criterion used in [8], maximizing the Pearson coefficient ϕ with respect to the threshold pair (one for prediction, one for reference).",
								Part: str("1"),
							}, {
								Sec:  "sec3",
								Text: " The threshold pair that corresponds to the maximum value of ϕ is the optimal choice according to this criterion (see [8] for a deeper discussion).",
								Part: str("1"),
							}, {
								Sec:  "sec3",
								Text: " It has to be remarked that for a fair comparison with the proposed methods, the baseline solution of [8] was used without masking any class, contrarily to what is done in the original formulation.",
								Part: str("1"),
							}, {
								Sec:  "sec3",
								Text: " Specifically, in [8] city and water classes are excluded by means of available masks, because forests, cities and water classes all exhibit a low volume correlation, the core feature proposed to classify forests.",
								Part: str("1"),
							},
						},
						Links: &goieeeapi.Links{
							CrossRefLink:  str("https://doi.org/10.1016/j.rse.2017.12.002"),
							OpenURLImgLoc: "/assets/img/btn.find-in-library.png",
						},
						GoogleScholarLink: str("https://scholar.google.com/scholar?as_q=The+global+forest%2Fnon-forest+map+from+tandem-x+interferometric+SAR+data&as_occt=title&hl=en&as_sdt=0%2C31"),
						RefType:           str("biblio"),
						ID:                str("ref8"),
					}, {
						Order: "9",
						Text:  str("M. Martone, F. Sica, C. Gonzlez, J.-L. Bueso-Bello, P. Valdo and P. Rizzoli, \"High-resolution forest mapping from tandem-x interferometric data exploiting nonlocal filtering\", <em>Remote Sensing</em>, vol. 10, pp. 1477, 2018."),
						Title: str("High-resolution forest mapping from tandem-x interferometric data exploiting nonlocal filtering"),
						Context: &[]goieeeapi.Context{
							{
								Sec:  "sec1",
								Text: " In [8], [9], the TanDEM-X forest/non-forest map is generated from TanDEM-X bistatic interferometric images, by linking the presence of vegetation to the retrieved InSAR volume decorrelation.",
								Part: str("1"),
							},
						},
						GoogleScholarLink: str("https://scholar.google.com/scholar?as_q=High-resolution+forest+mapping+from+tandem-x+interferometric+data+exploiting+nonlocal+filtering&as_occt=title&hl=en&as_sdt=0%2C31"),
						RefType:           str("biblio"),
						ID:                str("ref9"),
					}, {
						Order: "10",
						Text:  str("K. He, X. Zhang, S. Ren and J. Sun, \"Deep residual learning for image recognition\", <em>CVPR</em>, June 2016."),
						Title: str("Deep residual learning for image recognition"),
						Context: &[]goieeeapi.Context{
							{
								Sec:  "sec1",
								Text: " Specifically, we define and train from scratch two DL architectures following the ResNet [10] and the DenseNet [11] models, respectively.",
								Part: str("1"),
							}, {
								Sec:  "sec2",
								Text: " In this work, we consider two state-of-the-art CNN models, ResNet [10] and DenseNet [11], which are particularly appealing as they can be reach a considerable depth avoiding vanishing gradient problems during training.",
								Part: str("1"),
							}, {
								Sec:  "sec2",
								Text: " For both models, we describe here only the main functional aspects of interest for the present work, referring to the original papers [10], [11] for a thorough description of the network architecture.",
								Part: str("1"),
							}, {
								Sec:  "sec2",
								Text: " Although functionally unnecessary, skip connections have proven to speed-up the training [10], [12].",
								Part: str("1"),
							},
						},
						Links: &goieeeapi.Links{
							DocumentLink:  str("/document/7780459"),
							PdfLink:       str("/stamp/stamp.jsp?tp=&arnumber=7780459"),
							Abstract:      str("Deeper neural networks are more difficult to train. We present a residual learning framework to ease the training of networks that are substantially deeper than those used previously. We explicitly reformulate the layers as learning residual functions with reference to the layer inputs, instead of learning unreferenced functions. We provide comprehensive empirical evidence showing that these residual networks are easier to optimize, and can gain accuracy from considerably increased depth. On the..."),
							OpenURLImgLoc: "/assets/img/btn.find-in-library.png",
							PdfSize:       str("282KB"),
						},
						GoogleScholarLink: str("https://scholar.google.com/scholar?as_q=Deep+residual+learning+for+image+recognition&as_occt=title&hl=en&as_sdt=0%2C31"),
						RefType:           str("biblio"),
						ID:                str("ref10"),
					}, {
						Order: "11",
						Text:  str("G. Huang, Z. Liu, L. Van Der Maaten and K. Q. Weinberger, \"Densely connected convolutional networks\", <em>CVPR</em>, 2017."),
						Title: str("Densely connected convolutional networks"),
						Context: &[]goieeeapi.Context{
							{
								Sec:  "sec1",
								Text: " Specifically, we define and train from scratch two DL architectures following the ResNet [10] and the DenseNet [11] models, respectively.",
								Part: str("1"),
							}, {
								Sec:  "sec2",
								Text: " In this work, we consider two state-of-the-art CNN models, ResNet [10] and DenseNet [11], which are particularly appealing as they can be reach a considerable depth avoiding vanishing gradient problems during training.",
								Part: str("1"),
							}, {
								Sec:  "sec2",
								Text: " For both models, we describe here only the main functional aspects of interest for the present work, referring to the original papers [10], [11] for a thorough description of the network architecture.",
								Part: str("1"),
							},
						},
						Links: &goieeeapi.Links{
							DocumentLink:  str("/document/8099726"),
							PdfLink:       str("/stamp/stamp.jsp?tp=&arnumber=8099726"),
							Abstract:      str("Recent work has shown that convolutional networks can be substantially deeper, more accurate, and efficient to train if they contain shorter connections between layers close to the input and those close to the output. In this paper, we embrace this observation and introduce the Dense Convolutional Network (DenseNet), which connects each layer to every other layer in a feed-forward fashion. Whereas traditional convolutional networks with L layers have L connections-one between each layer and its ..."),
							OpenURLImgLoc: "/assets/img/btn.find-in-library.png",
							PdfSize:       str("395KB"),
						},
						GoogleScholarLink: str("https://scholar.google.com/scholar?as_q=Densely+connected+convolutional+networks&as_occt=title&hl=en&as_sdt=0%2C31"),
						RefType:           str("biblio"),
						ID:                str("ref11"),
					}, {
						Order: "12",
						Text:  str("G. Scarpa, S. Vitale and D. Cozzolino, \"Target-adaptive cnn-based pansharpening\", <em>IEEE Transactions on Geoscience and Remote Sensing</em>, vol. 56, no. 9, pp. 5443-5457, Sep. 2018."),
						Title: str("Target-adaptive cnn-based pansharpening"),
						Context: &[]goieeeapi.Context{
							{
								Sec:  "sec2",
								Text: " Although functionally unnecessary, skip connections have proven to speed-up the training [10], [12].",
								Part: str("1"),
							},
						},
						Links: &goieeeapi.Links{
							DocumentLink:  str("/document/8334206"),
							PdfLink:       str("/stamp/stamp.jsp?tp=&arnumber=8334206"),
							Abstract:      str("We recently proposed a convolutional neural network (CNN) for remote sensing image pansharpening obtaining a significant performance gain over the state of the art. In this paper, we explore a number of architectural and training variations to this baseline, achieving further performance gains with a lightweight network that trains very fast. Leveraging on this latter property, we propose a target-adaptive usage modality that ensures a very good performance also in the presence of a mismatch wit..."),
							OpenURLImgLoc: "/assets/img/btn.find-in-library.png",
							PdfSize:       str("4862KB"),
						},
						GoogleScholarLink: str("https://scholar.google.com/scholar?as_q=Target-adaptive+cnn-based+pansharpening&as_occt=title&hl=en&as_sdt=0%2C31"),
						RefType:           str("biblio"),
						ID:                str("ref12"),
					}, {
						Order: "13",
						Text:  str("A. Krizhevsky, I. Sutskever and G. E Hinton, \"Imagenet classification with deep convolutional neural networks\", <em>Advances in neural information processing systems</em>, pp. 1097-1105, 2012."),
						Title: str("Imagenet classification with deep convolutional neural networks"),
						Context: &[]goieeeapi.Context{
							{
								Sec:  "sec2",
								Text: "Here we propose for both models a cascade architecture with six convolutional layers with 3 × 3 kernels interleaved by ReLU (Rectified Linear Unit) activation functions [13].",
								Part: str("1"),
							},
						},
						GoogleScholarLink: str("https://scholar.google.com/scholar?as_q=Imagenet+classification+with+deep+convolutional+neural+networks&as_occt=title&hl=en&as_sdt=0%2C31"),
						RefType:           str("biblio"),
						ID:                str("ref13"),
					}, {
						Order: "14",
						Text:  str("G. Csurka, D. Larlus, F. Perronnin and F. Meylan, \"What is a good evaluation measure for semantic segmentation?\", <em>BMVC</em>, vol. 27, pp. 2013, 2013."),
						Title: str("What is a good evaluation measure for semantic segmentation?"),
						Context: &[]goieeeapi.Context{
							{
								Sec:  "sec2a",
								Text: " The Jaccard distance loss, which aims to maximize the overlap between the two maps [14], is defined as\n.",
								Part: str("1"),
							},
						},
						Links: &goieeeapi.Links{
							CrossRefLink:  str("https://doi.org/10.5244/C.27.32"),
							OpenURLImgLoc: "/assets/img/btn.find-in-library.png",
						},
						GoogleScholarLink: str("https://scholar.google.com/scholar?as_q=What+is+a+good+evaluation+measure+for+semantic+segmentation%3F&as_occt=title&hl=en&as_sdt=0%2C31"),
						RefType:           str("biblio"),
						ID:                str("ref14"),
					}, {
						Order: "15",
						Text:  str("Diederik P. Kingma and Jimmy Ba, \"Adam: A method for stochastic optimization\", <em>CoRR</em>, 2014."),
						Title: str("Adam: A method for stochastic optimization"),
						Context: &[]goieeeapi.Context{
							{
								Sec:  "sec2a",
								Text: "The minimization of the loss function is performed using the ADAM algorithm [15], a gradient descend variant where the learning rate is updated at each iteration using estimates of low-order moments.",
								Part: str("1"),
							},
						},
						GoogleScholarLink: str("https://scholar.google.com/scholar?as_q=Adam%3A+A+method+for+stochastic+optimization&as_occt=title&hl=en&as_sdt=0%2C31"),
						RefType:           str("biblio"),
						ID:                str("ref15"),
					},
				},
				ArticleNumber:               str("8900441"),
				GetProgramTermsAccepted:     false,
				FormulaStrippedArticleTitle: str("Deep Learning Solutions for Tandem-X-Based Forest Classification"),
				PubLink:                     str("/xpl/conhome/8891871/proceeding"),
				AllowComments:               false,
				IssueLink:                   "/xpl/tocresult.jsp?isnumber=null",
				IsReadingRoomArticle:        false,
				IsGetArticle:                false,
				IsGetAddressInfoCaptured:    false,
				IsMarketingOptIn:            false,
				Publisher:                   str("IEEE"),
				XploreDocumentType:          str("Conference Publication"),
				IsPromo:                     false,
				IsNotDynamicOrStatic:        false,
				HTMLAbstractLink:            "/document/8900441/",
				IsCustomDenial:              false,
				IsSAE:                       false,
				IsDynamicHTML:               true,
				IsFreeDocument:              false,
				DisplayDocTitle:             str("Deep Learning Solutions for Tandem-X-Based Forest Classification"),
				IsStandard:                  false,
				IsSMPTE:                     false,
				IsOUP:                       false,
				IsNow:                       false,
				IsProduct:                   false,
				IsMorganClaypool:            false,
				IsJournal:                   false,
				IsBook:                      false,
				IsBookWithoutChapters:       false,
				IsOpenAccess:                false,
				IsEphemera:                  false,
				IsConference:                true,
				HTMLLink:                    str("/document/8900441/"),
				IsChapter:                   false,
				IsStaticHTML:                false,
				IsEarlyAccess:               false,
				PersistentLink:              str("https://ieeexplore.ieee.org/servlet/opac?punumber=8891871"),
				ArticleId:                   str("8900441"),
				OpenAccessFlag:              str("F"),
				EphemeraFlag:                str("false"),
				Title:                       str("Deep Learning Solutions for Tandem-X-Based Forest Classification"),
				ContentTypeDisplay:          str("Conferences"),
				HTMLFlag:                    str("false"),
				MlHTMLFlag:                  str("true"),
				MlTime:                      "PT0.158217S",
				LastUpdate:                  str("2021-08-21"),
				MediaPath:                   str("/mediastore_new/IEEE/content/media/8891871/8897702/8900441"),
				ContentType:                 str("conferences"),
				Definitions:                 str("false"),
				PublicationNumber:           str("8891871"),
			},
			ExpectedErr: nil,
		},
	}

	for testname, tt := range tests {
		t.Run(testname, func(t *testing.T) {
			refs, err := goieeeapi.GetReferences(tt.Client, tt.ID)

			if !reflect.DeepEqual(refs, tt.ExpectedReferences) {
				t.Errorf("Failed to get expected references: got \"%v\" instead of \"%v\".", refs, tt.ExpectedReferences)
			}
			checkErr(err, tt.ExpectedErr, t)
		})
	}
}
