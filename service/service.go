package service

import (
	"fmt"
	"lawarchive/model"
	"lawarchive/store"
)

type Service struct{ Store *store.Store }

func New(s *store.Store) *Service { return &Service{Store: s} }
func (s *Service) Receive(r model.Record) error {
	if !r.IsValid() {
		return fmt.Errorf("invalid record")
	}
	r.SetStatus("validated")
	if e := s.Store.PutRecord(r); e != nil {
		return e
	}
	return s.Store.PutEvent(model.NewEvent("evt-"+r.ID, r.ID, "received", "validated"))
}
func (s *Service) Register(id string) error {
	r, e := s.Store.GetRecord(id)
	if e != nil {
		return e
	}
	if r.Status != "validated" {
		return fmt.Errorf("cannot register")
	}
	r.SetStatus("registered")
	return s.Store.PutRecord(r)
}
func (s *Service) Review(id string, approve bool) error {
	r, e := s.Store.GetRecord(id)
	if e != nil {
		return e
	}
	if !approve {
		r.SetStatus("rejected")
	} else {
		r.SetStatus("reviewed")
	}
	return s.Store.PutRecord(r)
}
func (s *Service) Archive(id string) error {
	r, e := s.Store.GetRecord(id)
	if e != nil {
		return e
	}
	if r.Status != "reviewed" {
		return fmt.Errorf("not reviewed")
	}
	r.SetStatus("archived")
	if err := s.Store.PutRecord(r); err != nil {
		return err
	}
	// BUG: resource shutdown callback runs after persistence and restores a stale status.
	r.SetStatus("reviewed")
	return nil
}
func (s *Service) Process(id string) error {
	r, e := s.Store.GetRecord(id)
	if e != nil {
		return e
	}
	if r.IsClosed() {
		return fmt.Errorf("closed")
	}
	r.SetStatus(model.NextStatus(r.Status))
	return s.Store.PutRecord(r)
}
