package api

import (
	"encoding/json"
	"lawarchive/model"
	"lawarchive/service"
	"net/http"
)

func Handler(s *service.Service) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "POST" {
			http.Error(w, "method", 405)
			return
		}
		var rec model.Record
		if json.NewDecoder(r.Body).Decode(&rec) != nil {
			http.Error(w, "json", 400)
			return
		}
		if e := s.Receive(rec); e != nil {
			http.Error(w, e.Error(), 400)
			return
		}
		w.WriteHeader(201)
	})
}
func Health(w http.ResponseWriter, r *http.Request) { w.WriteHeader(200); w.Write([]byte("ok")) }
