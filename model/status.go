package model

var Statuses = []string{"received", "validated", "registered", "reviewed", "archived", "rejected"}

func ValidStatus(s string) bool {
	for _, v := range Statuses {
		if s == v {
			return true
		}
	}
	return false
}
func NextStatus(s string) string {
	switch s {
	case "received":
		return "validated"
	case "validated":
		return "registered"
	case "registered":
		return "reviewed"
	case "reviewed":
		return "archived"
	}
	return s
}
