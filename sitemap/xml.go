package sitemap

import (
	"encoding/xml"
)

type root struct {
	XMLName xml.Name  `xml:"urlset"`
	Urls    []*urlloc `xml:"url"`
	Xmlns   string    `xml:"xmlns,attr"`
}

type urlloc struct {
	XMLName xml.Name `xml:"url"`
	Loc     string   `xml:"loc"`
}

func (s *SiteMap) ToXML() (string, error) {
	urllocs := make([]*urlloc, len(s.urls))
	for i, u := range s.urls {
		urllocs[i] = &urlloc{Loc: u}
	}
	root := &root{Urls: urllocs, Xmlns: "http://www.sitemaps.org/schemas/sitemap/0.9"}
	parsed, err := xml.MarshalIndent(root, " ", " ")
	if err != nil {
		return "", err
	}
	return xml.Header + string(parsed), nil
}
