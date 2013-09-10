package sitemap

import (
	"bytes"
	"testing"
	"time"
)

func TestSiteMap(t *testing.T) {
	tests := map[string]struct {
		urlset URLSet
		xml    []byte
	}{
		"basic": {
			urlset: URLSet{
				URLs: []URL{
					{
						Loc:        "http://www.example.com/",
						LastMod:    time.Date(2013, time.September, 9, 11, 22, 33, 0, time.UTC),
						ChangeFreq: Daily,
						Priority:   0.7,
					},
				},
			},
			xml: []byte(`
<?xml version="1.0" encoding="UTF-8"?>
<urlset xmlns="http://www.sitemaps.org/schemas/sitemap/0.9">
  <url>
    <loc>http://www.example.com/</loc>
    <lastmod>2013-09-09T11:22:33Z</lastmod>
    <changefreq>daily</changefreq>
    <priority>0.7</priority>
  </url>
</urlset>`),
		},
	}

	for label, test := range tests {
		xml, err := Marshal(&test.urlset)
		if err != nil {
			t.Errorf("%s: Marshal: %s", label, err)
			continue
		}

		// Trim whitespace from the expected XML so we can compare it to the
		// actual XML output.
		test.xml = bytes.Replace(bytes.Replace(test.xml, []byte("  "), nil, -1), []byte("\n"), nil, -1)

		if !bytes.Equal(test.xml, xml) {
			t.Errorf("%s: want XML %q, got %q", label, test.xml, xml)
		}
	}
}
