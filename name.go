package english

import (
	"encoding/json"
	"io"
	"net/http"
)

type Name []string

func (d *Name) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		json.NewEncoder(w).Encode(d)
		return
	}
	if r.Method == http.MethodPost {
		b, err := io.ReadAll(r.Body)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}
		var in struct {
			Method string
		}
		err = json.Unmarshal(b, &in)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}
		switch in.Method {
		case "String":
			var in struct {
				Args struct {
				}
			}
			err = json.Unmarshal(b, &in)
			if err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
			}
			// out1 := d.String()
			var out struct {
				Out1 string
			}
			// out.Out1 = out1
			json.NewEncoder(w).Encode(out)
		default:
			http.Error(w, "unknown method", http.StatusBadRequest)
		}
		return
	}
}
