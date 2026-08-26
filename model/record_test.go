package model

import "testing"

func TestRecordValid(t *testing.T) {
	if !NewRecord("a", "b", "c", "d", 1).IsValid() {
		t.Fatal()
	}
}
