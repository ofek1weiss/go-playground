package linkparser

import "fmt"

type Link struct {
	Href string
	Text string
}

func (l *Link) String() string {
	return fmt.Sprintf("Link(Href=%#v, Text=%#v)", l.Href, l.Text)
}
