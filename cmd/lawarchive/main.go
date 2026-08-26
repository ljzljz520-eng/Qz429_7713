package main

import (
	"lawarchive/api"
	"lawarchive/config"
	"lawarchive/service"
	"lawarchive/store"
	"log"
	"net/http"
)

func main() {
	c := config.Load()
	s, e := store.Open(c.Path)
	if e != nil {
		log.Fatal(e)
	}
	defer s.Close()
	svc := service.New(s)
	mux := http.NewServeMux()
	mux.Handle("/records", api.Handler(svc))
	mux.HandleFunc("/health", api.Health)
	log.Fatal(http.ListenAndServe(c.Address(), mux))
}
