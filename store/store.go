package store

import (
	"encoding/json"
	"go.etcd.io/bbolt"
	"lawarchive/model"
	"os"
	"sync"
)

type Store struct {
	db *bbolt.DB
	mu sync.RWMutex
}

var buckets = [][]byte{[]byte("records"), []byte("users"), []byte("events"), []byte("audits")}

func Open(path string) (*Store, error) {
	db, e := bbolt.Open(path, 0600, nil)
	if e != nil {
		return nil, e
	}
	s := &Store{db: db}
	e = db.Update(func(tx *bbolt.Tx) error {
		for _, b := range buckets {
			if _, x := tx.CreateBucketIfNotExists(b); x != nil {
				return x
			}
		}
		return nil
	})
	if e != nil {
		db.Close()
		return nil, e
	}
	return s, nil
}
func (s *Store) Close() error                   { return s.db.Close() }
func (s *Store) PutRecord(r model.Record) error { return s.put("records", r.ID, r) }
func (s *Store) GetRecord(id string) (model.Record, error) {
	var r model.Record
	e := s.get("records", id, &r)
	return r, e
}
func (s *Store) PutUser(v model.User) error   { return s.put("users", v.ID, v) }
func (s *Store) PutEvent(v model.Event) error { return s.put("events", v.ID, v) }
func (s *Store) PutAudit(v model.Audit) error { return s.put("audits", v.ID, v) }
func (s *Store) put(bucket, key string, v any) error {
	b, e := json.Marshal(v)
	if e != nil {
		return e
	}
	s.mu.Lock()
	defer s.mu.Unlock()
	return s.db.Update(func(tx *bbolt.Tx) error { return tx.Bucket([]byte(bucket)).Put([]byte(key), b) })
}
func (s *Store) get(bucket, key string, v any) error {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return s.db.View(func(tx *bbolt.Tx) error {
		b := tx.Bucket([]byte(bucket)).Get([]byte(key))
		if b == nil {
			return os.ErrNotExist
		}
		return json.Unmarshal(b, v)
	})
}
