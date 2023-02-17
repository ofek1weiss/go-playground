package sitemap

import (
	"linkparser/linkparser"
	"log"
	"net/http"
	"net/url"
	"sync"

	mapset "github.com/deckarep/golang-set/v2"
)

func Crawl(domain string, depth int) *SiteMap {
	visitedUrls := mapset.NewSet[string]()
	targetUris := mapset.NewSet("/")
	for i := 0; i < depth && len(targetUris.ToSlice()) != 0; i++ {
		var wg sync.WaitGroup
		for _, targetUri := range targetUris.ToSlice() {
			targetUris.Remove(targetUri)
			targetUrl, _ := url.JoinPath("https://"+domain, targetUri)
			if visitedUrls.Contains(targetUrl) {
				continue
			}
			wg.Add(1)
			go func() {
				defer wg.Done()
				visitedUrls.Add(targetUrl)
				for _, uri := range worker(domain, targetUrl) {
					targetUris.Add(uri)
				}
			}()
		}
		wg.Wait()
	}
	return &SiteMap{urls: visitedUrls.ToSlice()}
}

func worker(domain string, targetUrl string) []string {
	log.Printf("Crawling url %s\n", targetUrl)
	uris, err := crawlForUris(domain, targetUrl)
	if err != nil {
		log.Printf("Error while crawling %s\n", targetUrl)
	}
	log.Printf("Found %d uris in url %s\n", len(uris), targetUrl)
	return uris
}

func crawlForUris(domain string, url string) ([]string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	links, err := linkparser.Parse(resp.Body)
	var uris []string
	for _, link := range links {
		if uri, ok := getUriFromLink(domain, link); ok {
			uris = append(uris, uri)
		}
	}
	return uris, err
}

func getUriFromLink(domain string, link *linkparser.Link) (string, bool) {
	u, err := url.Parse(link.Href)
	if err != nil {
		return link.Href, true
	}
	if u.Host != domain && u.Host != "" {
		return "", false
	}
	return u.Path, true
}
