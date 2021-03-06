package crawler

import (
        "github.com/deckarep/golang-set"
        "fmt"
	"strings"
	"testing"
)

func TestFind(t *testing.T) {
	reader := strings.NewReader(`<p>
	<a href="http://domain.com/1"</a>
        <link href="http://domain.com/2"</link>
	<a href="http://differentdomain.com/"</link>
	<script defer='' src='/js/cookiechoices.js'></script>
	<img alt='Image' sizes='(max-width: 800px) 12vw, 128px' src='cat.png'</img>
	</p>`)

	links, resources := Find(reader, "domain.com")

	if len(resources) != 2 {
		t.Error("Wrong number of resources.")
	}
        
        for _, r := range resources {
            if r != "/js/cookiechoices.js" && r != "cat.png" {
                t.Error("unexpected resource: "+r)
            }
        }

	if len(links) != 2 {
		t.Error("Wrong number of links.")
	}
}

func TestClean(t *testing.T) {
	link1 := "http://domain.com/2#comments"
	link2 := "http://domain.com/2?query"
	if clean(link1) != clean(link2) {
		t.Error("Cleaned links not equal")
	}
	link3 := "http://domain.com/2"
	if clean(link1) != clean(link3) {
		t.Error("Already clean link damaged")
	}
}

func TestToSlice(t *testing.T) {
        set := mapset.NewSet()
        set.Add("first")
        set.Add("second")
        slice := toSlice(set)
        if len(slice) != 2 {
            t.Error("Slice length is not 2")
        }
        fmt.Println(slice)
        for _, s := range slice {
            if s != "first" && s != "second" {
                t.Error("unexpected value in slice: "+s)
            }
        }
}
