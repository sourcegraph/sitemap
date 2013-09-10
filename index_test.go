package sitemap

import (
	"bytes"
	"testing"
	"time"
)

func TestMarshalIndex(t *testing.T) {
	tests := map[string]struct {
		index Index
		xml   []byte
	}{
		"basic": {
			index: Index{
				Sitemaps: []Sitemap{
					{
						Loc:     "http://www.example.com/sitemap.xml.gz",
						LastMod: time.Date(2013, time.September, 9, 11, 22, 33, 0, time.UTC),
					},
				},
			},
			xml: []byte(`
<?xml version="1.0" encoding="UTF-8"?>
<sitemapindex xmlns="http://www.sitemaps.org/schemas/sitemap/0.9">
  <sitemap>
    <loc>http://www.example.com/sitemap.xml.gz</loc>
    <lastmod>2013-09-09T11:22:33Z</lastmod>
  </sitemap>
</sitemapindex>`),
		},
	}

	for label, test := range tests {
		xml, err := MarshalIndex(&test.index)
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
