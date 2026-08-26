package model

func NormalizeStatus(s string) string {
	if !ValidStatus(s) {
		return "received"
	}
	return s
}
