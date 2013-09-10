package sitemap

import (
	"encoding/xml"
	"time"
)

// Index represents a collection of sitemaps in a sitemap index.
//
// Refer to http://www.sitemaps.org/protocol.html#index for more information.
type Index struct {
	XMLName  xml.Name  `xml:"sitemapindex"`
	XMLNS    string    `xml:"xmlns,attr"`
	Sitemaps []Sitemap `xml:"sitemap"`
}

// Sitemap represents information about an individual sitemap in a sitemap
// index.
//
// Refer to http://www.sitemaps.org/protocol.html#index for more information.
type Sitemap struct {
	Loc     string    `xml:"loc"`
	LastMod time.Time `xml:"lastmod"`
}

// MarshalIndex serializes the sitemap index to XML, with the <sitemapindex>
// xmlns added and the XML preamble prepended.
func MarshalIndex(index *Index) (indexXML []byte, err error) {
	index.XMLNS = xmlns
	indexXML = []byte(preamble)
	var smiXML []byte
	smiXML, err = xml.Marshal(index)
	if err == nil {
		indexXML = append(indexXML, smiXML...)
	}
	return
}
