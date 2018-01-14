package crawler

import (
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

	if resources[0] != "/js/cookiechoices.js" {
		t.Error("The first resources should be js")
	}

	if resources[1] != "cat.png" {
		t.Error("The second resource should be cat.png")
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
