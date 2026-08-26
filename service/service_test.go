package service

import (
	"lawarchive/model"
	"lawarchive/store"
	"os"
	"testing"
)

func TestReceiveRejects(t *testing.T) {
	p := "v.db"
	defer os.Remove(p)
	st, _ := store.Open(p)
	defer st.Close()
	if New(st).Receive(model.Record{}) == nil {
		t.Fatal()
	}
}
