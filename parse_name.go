package english

import (
	"strings"
)

func ParseName(s string) Name {
	words := strings.Split(s, " ")
	name := make(Name, len(words))
	for i, word := range words {
		name[i] = word
	}
	return name
}
