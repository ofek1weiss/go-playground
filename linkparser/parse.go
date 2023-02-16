package linkparser

import (
	"io"
	"strings"

	"golang.org/x/net/html"
)

func Parse(reader io.Reader) ([]*Link, error) {
	node, err := html.Parse(reader)
	if err != nil {
		return nil, err
	}
	return parseNode(node), nil
}

func parseNode(node *html.Node) []*Link {
	if node.Type == html.ElementNode && node.Data == "a" {
		text := getAllTextFromNode(node)
		href := getAttributeFromNode(node, "href", "")
		link := &Link{
			Href: href,
			Text: text,
		}
		return []*Link{link}
	}
	links := make([]*Link, 0)
	for child := node.FirstChild; child != nil; child = child.NextSibling {
		childLinks := parseNode(child)
		if len(childLinks) != 0 {
			links = append(links, childLinks...)
		}
	}
	return links
}

func getAllTextFromNode(node *html.Node) string {
	if node.Type == html.TextNode {
		return strings.TrimSpace(node.Data)
	}
	if node.Type == html.CommentNode {
		return ""
	}
	textParts := make([]string, 0)
	for child := node.FirstChild; child != nil; child = child.NextSibling {
		childText := getAllTextFromNode(child)
		if childText != "" {
			textParts = append(textParts, childText)
		}
	}
	return strings.Join(textParts, " ")
}

func getAttributeFromNode(node *html.Node, name string, defaultValue string) string {
	for _, attribute := range node.Attr {
		if attribute.Key == name {
			return attribute.Val
		}
	}
	return defaultValue
}
