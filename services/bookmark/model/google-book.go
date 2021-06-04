package model

type GoogleBook struct {
	Kind       string  `json:"kind"`
	Totalitems int     `json:"totalItems"`
	Items      []Items `json:"items"`
}
type Industryidentifiers struct {
	Type       string `json:"type"`
	Identifier string `json:"identifier"`
}
type Readingmodes struct {
	Text  bool `json:"text"`
	Image bool `json:"image"`
}
type Imagelinks struct {
	Smallthumbnail string `json:"smallThumbnail"`
	Thumbnail      string `json:"thumbnail"`
}
//type Saleinfo struct {
//	Country     string `json:"country"`
//	Saleability string `json:"saleability"`
//	Isebook     bool   `json:"isEbook"`
//}
type Epub struct {
	Isavailable bool `json:"isAvailable"`
}
type Pdf struct {
	Isavailable bool `json:"isAvailable"`
}
type Accessinfo struct {
	Country                string `json:"country"`
	Viewability            string `json:"viewability"`
	Embeddable             bool   `json:"-"`
	Publicdomain           bool   `json:"-"`
	Texttospeechpermission string `json:"textToSpeechPermission"`
	Epub                   Epub   `json:"epub"`
	Pdf                    Pdf    `json:"-"`
	Webreaderlink          string `json:"webReaderLink"`
	Accessviewstatus       string `json:"-"`
	Quotesharingallowed    bool   `json:"-"`
}
type Searchinfo struct {
	Textsnippet string `json:"-"`
}
type Panelizationsummary struct {
	Containsepubbubbles  bool `json:"containsEpubBubbles"`
	Containsimagebubbles bool `json:"containsImageBubbles"`
}

type Listprice struct {
	Amount       interface{} `json:"amount"`
	Currencycode string  `json:"currencyCode"`
}
type Retailprice struct {
	Amount       interface{} `json:"amount"`
	Currencycode string  `json:"currencyCode"`
}
//type Listprice struct {
//	Amountinmicros int    `json:"amountInMicros"`
//	Currencycode   string `json:"currencyCode"`
//}
//type Retailprice struct {
//	Amountinmicros int    `json:"amountInMicros"`
//	Currencycode   string `json:"currencyCode"`
//}
type Offers struct {
	Finskyoffertype int         `json:"finskyOfferType"`
	Listprice       Listprice   `json:"listPrice"`
	Retailprice     Retailprice `json:"retailPrice"`
}
type Saleinfo struct {
	Country     string      `json:"country"`
	Saleability string      `json:"saleability"`
	Isebook     bool        `json:"isEbook"`
	Listprice   Listprice   `json:"listPrice"`
	Retailprice Retailprice `json:"retailPrice"`
	Buylink     string      `json:"buyLink"`
	Offers      []Offers    `json:"offers"`
}
type Volumeinfo struct {
	Title               string                `json:"title"`
	Subtitle            string                `json:"subtitle"`
	Authors             []string              `json:"authors"`
	Publisher           string                `json:"publisher"`
	Publisheddate       string                `json:"publishedDate"`
	Description         string                `json:"description"`
	Industryidentifiers []Industryidentifiers `json:"industryIdentifiers"`
	Readingmodes        Readingmodes          `json:"-"`
	Pagecount           int                   `json:"pageCount"`
	Printtype           string                `json:"printType"`
	Categories          []string              `json:"categories"`
	Maturityrating      string                `json:"maturityRating"`
	Allowanonlogging    bool                  `json:"allowAnonLogging"`
	Contentversion      string                `json:"contentVersion"`
	Panelizationsummary Panelizationsummary   `json:"-"`
	Imagelinks          Imagelinks            `json:"-"`
	Language            string                `json:"language"`
	Previewlink         string                `json:"previewLink"`
	Infolink            string                `json:"infoLink"`
	Canonicalvolumelink string                `json:"canonicalVolumeLink"`
}
type Items struct {
	Kind       string     `json:"kind"`
	ID         string     `json:"id"`
	Etag       string     `json:"etag"`
	Selflink   string     `json:"selfLink"`
	Volumeinfo Volumeinfo `json:"volumeInfo,omitempty"`
	Saleinfo   Saleinfo   `json:"saleInfo,omitempty"`
	Accessinfo Accessinfo `json:"accessInfo"`
	Searchinfo Searchinfo `json:"searchInfo,omitempty"`
}