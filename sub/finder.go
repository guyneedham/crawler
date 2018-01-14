package crawler

import (
	"golang.org/x/net/html"
	"io"
	"regexp"
)

func Find(httpBody io.Reader, domain string) ([]string, []string) {
	links := Generate()
	resources := Generate()
	page := html.NewTokenizer(httpBody)
	for {
		htmlType := page.Next()
		if htmlType == html.ErrorToken {
			return links, resources
		}
		token := page.Token()
		if htmlType == html.StartTagToken {
			for _, attr := range token.Attr {
				if attr.Key == "href" {
					cleaned := clean(attr.Val)
					d := extractDomain(cleaned)
					if d == domain && cleaned != "" {
						AddElement(&links, cleaned)
					}
				}
				if attr.Key == "src" {
					AddElement(&resources, attr.Val)
				}
			}
		}
	}
}

var re = regexp.MustCompile("[#?]")

func clean(link string) string {
	s := re.Split(link, -1)
	return s[0]
}
