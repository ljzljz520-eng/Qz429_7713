package model

func SearchCatalog(prefix string) []string {
	out := []string{}
	for _, v := range LegalNoticeCatalog {
		if len(prefix) <= len(v) && v[:len(prefix)] == prefix {
			out = append(out, v)
		}
	}
	return out
}
