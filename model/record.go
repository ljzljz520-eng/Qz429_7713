package model

import "time"

type Record struct {
	ID, Title, Body, Status, Owner string
	LegalNumber                    int
	CreatedAt, UpdatedAt           time.Time
}
type User struct {
	ID, Name, Role string
	Active         bool
}
type Event struct {
	ID, RecordID, Kind, Detail string
	At                         time.Time
}
type Audit struct {
	ID, RecordID, Actor, Action string
	At                          time.Time
}

func NewRecord(id, title, body, owner string, number int) Record {
	now := time.Now().UTC()
	return Record{ID: id, Title: title, Body: body, Owner: owner, LegalNumber: number, Status: "received", CreatedAt: now, UpdatedAt: now}
}
func (r Record) IsValid() bool {
	return r.ID != "" && r.Title != "" && r.Body != "" && r.LegalNumber > 0
}
func (r Record) IsClosed() bool           { return r.Status == "archived" || r.Status == "rejected" }
func (r *Record) SetStatus(status string) { r.Status = status; r.UpdatedAt = time.Now().UTC() }
func NewUser(id, name, role string) User  { return User{ID: id, Name: name, Role: role, Active: true} }
func NewEvent(id, recordID, kind, detail string) Event {
	return Event{ID: id, RecordID: recordID, Kind: kind, Detail: detail, At: time.Now().UTC()}
}
func NewAudit(id, recordID, actor, action string) Audit {
	return Audit{ID: id, RecordID: recordID, Actor: actor, Action: action, At: time.Now().UTC()}
}
