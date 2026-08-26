package config

import "testing"

func TestLoad(t *testing.T) {
	if Load().Port == "" {
		t.Fatal()
	}
}
