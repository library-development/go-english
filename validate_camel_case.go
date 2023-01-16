package english

import (
	"unicode"
)

// ValidateCamelCase returns an error if the string is not in camelCase.
func ValidateCamelCase(s string) error {
	if s == "" {
		return ErrEmptyString
	}
	if !unicode.IsLower(rune(s[0])) {
		return ErrExpectedCamelCase
	}
	for _, r := range s {
		if !unicode.IsUpper(r) && !unicode.IsLower(r) && !unicode.IsDigit(r) {
			return ErrExpectedCamelCase
		}
	}
	return nil
}
