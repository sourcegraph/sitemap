// Package sitemap generates sitemap.xml files based on the sitemaps.org
// protocol.
package sitemap

import (
	"encoding/xml"
	"time"
)

const xmlns = "http://www.sitemaps.org/schemas/sitemap/0.9"

// URLSet represents a set of URLs in a sitemap.
//
// Refer to http://www.sitemaps.org/protocol.html#xmlTagDefinitions for more
// information.
type URLSet struct {
	XMLName xml.Name `xml:"urlset"`
	XMLNS   string   `xml:"xmlns,attr"`
	URLs    []URL    `xml:"url"`
}

// URL presents a URL and associated sitemap information.
//
// Refer to http://www.sitemaps.org/protocol.html#xmlTagDefinitions for more
// information.
type URL struct {
	Loc        string     `xml:"loc"`
	LastMod    time.Time  `xml:"lastmod"`
	ChangeFreq ChangeFreq `xml:"changefreq"`
	Priority   float64    `xml:"priority"`
}

// ChangeFreq indicates how frequently the page is likely to change. This value
// provides general information to search engines and may not correlate exactly
// to how often they crawl the page.
//
// Refer to http://www.sitemaps.org/protocol.html#xmlTagDefinitions for more
// information.
type ChangeFreq string

const (
	Always  ChangeFreq = "always"
	Hourly  ChangeFreq = "hourly"
	Daily   ChangeFreq = "daily"
	Weekly  ChangeFreq = "weekly"
	Monthly ChangeFreq = "monthly"
	Yearly  ChangeFreq = "yearly"
	Never   ChangeFreq = "never"
)

const preamble = `<?xml version="1.0" encoding="UTF-8"?>`

// Marshal serializes the sitemap URLSet to XML, with the <urlset> xmlns added
// and the XML preamble prepended.
func Marshal(urlset *URLSet) (sitemapXML []byte, err error) {
	urlset.XMLNS = xmlns
	sitemapXML = []byte(preamble)
	var urlsetXML []byte
	urlsetXML, err = xml.Marshal(urlset)
	if err == nil {
		sitemapXML = append(sitemapXML, urlsetXML...)
	}
	return
}
