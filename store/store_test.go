package store

import (
	"lawarchive/model"
	"os"
	"testing"
)

func TestStoreRecord(t *testing.T) {
	p := "s.db"
	defer os.Remove(p)
	s, _ := Open(p)
	defer s.Close()
	s.PutRecord(model.NewRecord("x", "t", "b", "o", 1))
	if !s.HasRecord("x") {
		t.Fatal()
	}
}
