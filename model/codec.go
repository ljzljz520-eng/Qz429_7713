package model

import "encoding/json"

func Encode(v any) ([]byte, error) { return json.Marshal(v) }
func DecodeRecord(b []byte) (Record, error) {
	var v Record
	err := json.Unmarshal(b, &v)
	return v, err
}
func DecodeUser(b []byte) (User, error)   { var v User; err := json.Unmarshal(b, &v); return v, err }
func DecodeEvent(b []byte) (Event, error) { var v Event; err := json.Unmarshal(b, &v); return v, err }
func DecodeAudit(b []byte) (Audit, error) { var v Audit; err := json.Unmarshal(b, &v); return v, err }
