package main

import (
	"fmt"
	"linkparser/sitemap"
)

func main() {
	flags := GetFlags()
	sm := sitemap.Crawl(flags.Domain, flags.Depth)
	xml, err := sm.ToXML()
	if err != nil {
		panic(err)
	}
	fmt.Print(xml)
}
