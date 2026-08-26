package model

import "testing"

func TestStatus(t *testing.T) {
	if NextStatus("received") != "validated" {
		t.Fatal()
	}
}
