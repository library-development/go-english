package english

import (
	"unicode"
)

// validatePascalCase returns an error if the string is not in PascalCase.
func validatePascalCase(s string) error {
	if s == "" {
		return ErrEmptyString
	}
	if !unicode.IsUpper(rune(s[0])) {
		return ErrExpectedPascalCase
	}
	for _, r := range s {
		if !unicode.IsUpper(r) && !unicode.IsLower(r) && !unicode.IsDigit(r) {
			return ErrExpectedPascalCase
		}
	}
	return nil
}
