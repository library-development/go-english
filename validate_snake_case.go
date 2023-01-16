package english

import "unicode"

// validateSnakeCase returns an error if the string is not valid snake_case.
func validateSnakeCase(s string) error {
	if s == "" {
		return ErrEmptyString
	}
	for _, r := range s {
		if !unicode.IsLower(r) && !unicode.IsDigit(r) && r != '_' {
			return ErrExpectedSnakeCase
		}
	}
	return nil
}
