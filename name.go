package english

import "strings"

type Name []string

func (n Name) String() string {
	s := strings.Builder{}
	for i, word := range n {
		if i > 0 {
			s.WriteString(" ")
		}
		s.WriteString(word)
	}
	return s.String()
}

func (n Name) PascalCase() string {
	s := strings.Builder{}
	for _, w := range n {
		s.WriteString(strings.Title(w))
	}
	return s.String()
}

func (n Name) CamelCase() string {
	s := strings.Builder{}
	for i, w := range n {
		if i == 0 {
			s.WriteString(strings.ToLower(w))
		} else {
			s.WriteString(strings.Title(w))
		}
	}
	return s.String()
}

func (n Name) SnakeCase() string {
	s := strings.Builder{}
	for i, w := range n {
		if i > 0 {
			s.WriteString("_")
		}
		s.WriteString(strings.ToLower(w))
	}
	return s.String()
}

func (n Name) KebabCase() string {
	s := strings.Builder{}
	for i, w := range n {
		if i > 0 {
			s.WriteString("-")
		}
		s.WriteString(strings.ToLower(w))
	}
	return s.String()
}

func (n Name) ScreamingSnakeCase() string {
	s := strings.Builder{}
	for i, w := range n {
		if i > 0 {
			s.WriteString("_")
		}
		s.WriteString(strings.ToUpper(w))
	}
	return s.String()
}
