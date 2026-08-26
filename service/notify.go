package service

import "lawarchive/model"

func (s *Service) Notify(id string) error {
	r, e := s.Store.GetRecord(id)
	if e != nil {
		return e
	}
	return s.Store.PutEvent(model.NewEvent("notify-"+id, r.ID, "notify", r.Status))
}
