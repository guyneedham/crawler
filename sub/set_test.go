package crawler

import (
	"testing"
)

func TestSet(t *testing.T) {
	s := Generate()

	if checkMembership(s, "test") {
		t.Error("New set not empty")
	}

	AddElement(&s, "test")

	if checkMembership(s, "test") == false {
		t.Error("Element not in set")
	}
}
