package crawler

import (
	"fmt"
	"testing"
)

func TestPullPage(t *testing.T) {
	//p := pullPage("http://www.unofficialgoogledatascience.com/")
	//fmt.Println(p)
}

func TestExtractDomaint(t *testing.T) {
	d1 := extractDomain("http://www.unofficialgoogledatascience.com/")
	d2 := extractDomain("http://www.unofficialgoogledatascience.com/2017/03/attributing-deep-networks-prediction-to.html")
	if d1 != d2 {
		fmt.Println(d1)
		fmt.Println(d2)
		t.Error("domains not the same")
	}
}
