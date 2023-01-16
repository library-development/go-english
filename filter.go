package english

import "strings"

func filter(word, allowed string) string {
	var b strings.Builder
	for _, r := range word {
		if strings.ContainsRune(allowed, r) {
			b.WriteRune(r)
		}
	}
	return b.String()
}
