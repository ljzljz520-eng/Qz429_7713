package query

import (
	"lawarchive/model"
	"lawarchive/store"
	"strings"
)

func Find(s *store.Store, id string) (model.Record, error) { return s.GetRecord(id) }
func Match(r model.Record, term string) bool {
	return strings.Contains(strings.ToLower(r.Title), strings.ToLower(term)) || strings.Contains(r.Body, term)
}
func Display(r model.Record) string { return r.ID + " | " + r.Title + " | " + r.Status }
