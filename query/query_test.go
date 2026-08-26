package query

import (
	"lawarchive/model"
	"testing"
)

func TestMatch(t *testing.T) {
	if !Match(model.NewRecord("x", "Title", "Body", "o", 1), "title") {
		t.Fatal()
	}
}
