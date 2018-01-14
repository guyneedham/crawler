package crawler

import (
        "github.com/deckarep/golang-set"
	"golang.org/x/net/html"
	"io"
	"regexp"
)

func Find(httpBody io.Reader, domain string) ([]string, []string) {
        links := mapset.NewSet()
        resources := mapset.NewSet()
	page := html.NewTokenizer(httpBody)
	for {
		htmlType := page.Next()
		if htmlType == html.ErrorToken {
			return toSlice(links), toSlice(resources)
		}
		token := page.Token()
		if htmlType == html.StartTagToken {
			for _, attr := range token.Attr {
				if attr.Key == "href" {
					cleaned := clean(attr.Val)
					d := extractDomain(cleaned)
					if d == domain && cleaned != "" {
						links.Add(cleaned)
					}
				}
				if attr.Key == "src" {
					resources.Add(attr.Val)
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

func toSlice(s mapset.Set) []string {
   var sl = []string{}
   it := s.Iterator()
   for i := range it.C {
       s, ok := i.(string)
       if !ok {
          panic(ok)
       }
       sl = append(sl, s)
   }
   return sl
}
