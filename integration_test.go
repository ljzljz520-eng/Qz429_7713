package lawarchive

import (
	"lawarchive/model"
	"lawarchive/service"
	"lawarchive/store"
	"os"
	"testing"
)

func testSvc(t *testing.T) *service.Service {
	p := "test.db"
	os.Remove(p)
	s, e := store.Open(p)
	if e != nil {
		t.Fatal(e)
	}
	t.Cleanup(func() { s.Close(); os.Remove(p) })
	return service.New(s)
}
func TestRecordFlow46(t *testing.T) {
	svc := testSvc(t)
	r := model.NewRecord("r46", "公告46", "资料", "u", 46)
	if svc.Receive(r) != nil {
		t.Fatal()
	}
	if e := svc.Register("r46"); e != nil {
		t.Fatal(e)
	}
	svc.Review("r46", true)
	if e := svc.Archive("r46"); e != nil {
		t.Fatal(e)
	}
	got, e := svc.Store.GetRecord("r46")
	if e != nil || got.Status != "archived" {
		t.Fatalf("status=%s", got.Status)
	}
}
func TestWorkflowOne(t *testing.T) {
	svc := testSvc(t)
	if e := svc.Receive(model.NewRecord("one", "公告", "内容", "u", 1)); e != nil {
		t.Fatal(e)
	}
}
func TestWorkflowTwo(t *testing.T) {
	svc := testSvc(t)
	r := model.NewRecord("two", "公告", "内容", "u", 2)
	svc.Receive(r)
	svc.Register(r.ID)
	svc.Review(r.ID, true)
	if e := svc.Archive(r.ID); e != nil {
		t.Fatal(e)
	}
}
func TestWorkflowThree(t *testing.T) {
	svc := testSvc(t)
	r := model.NewRecord("three", "公告", "内容", "u", 3)
	svc.Receive(r)
	if e := svc.Process(r.ID); e != nil {
		t.Fatal(e)
	}
}
func TestPersistenceSurvivesReopen(t *testing.T) {
	p := "reopen.db"
	os.Remove(p)
	s, _ := store.Open(p)
	s.PutRecord(model.NewRecord("persist", "x", "y", "u", 1))
	s.Close()
	s, e := store.Open(p)
	if e != nil {
		t.Fatal(e)
	}
	defer s.Close()
	if _, e = s.GetRecord("persist"); e != nil {
		t.Fatal(e)
	}
	os.Remove(p)
}
