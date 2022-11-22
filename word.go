package english

import "strings"

type Word struct {
	Acronym bool   `json:"acronym"`
	String  string `json:"string"`
}

func (w *Word) Title() string {
	return strings.Title(w.String)
}

func (w *Word) Lower() string {
	return strings.ToLower(w.String)
}

func (w *Word) Upper() string {
	return strings.ToUpper(w.String)
}
