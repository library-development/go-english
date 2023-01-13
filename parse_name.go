package english

import (
	"strings"
)

// ParseName parses a string into a Name.
func ParseName(s string) Name {
	words := strings.Split(s, " ")
	name := Name{}
	for _, word := range words {
		if word == "" {
			continue
		}
		name = append(name, Word(word))
	}
	return name
}
