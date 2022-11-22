package english

import "strings"

type Name []Word

func (n Name) String() string {
	s := strings.Builder{}
	for i, word := range n {
		if i > 0 {
			s.WriteString(" ")
		}
		s.WriteString(word.String)
	}
	return s.String()
}

func (n Name) PascalCase() string {
	s := strings.Builder{}
	for _, w := range n {
		s.WriteString(w.Title())
	}
	return s.String()
}

func (n Name) CamelCase() string {
	s := strings.Builder{}
	for i, w := range n {
		if i == 0 {
			s.WriteString(w.Lower())
		} else {
			s.WriteString(w.Title())
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
		s.WriteString(w.Lower())
	}
	return s.String()
}

func (n Name) KebabCase() string {
	s := strings.Builder{}
	for i, w := range n {
		if i > 0 {
			s.WriteString("-")
		}
		s.WriteString(w.Lower())
	}
	return s.String()
}

func (n Name) ScreamingSnakeCase() string {
	s := strings.Builder{}
	for i, w := range n {
		if i > 0 {
			s.WriteString("_")
		}
		s.WriteString(w.Upper())
	}
	return s.String()
}
