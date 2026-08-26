package store

func (s *Store) HasRecord(id string) bool { _, e := s.GetRecord(id); return e == nil }
