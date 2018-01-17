package types

import "net/http"

type Redirect struct {
	URI  string
	Code int
}

func (re *Redirect) Exec(w http.ResponseWriter, r *http.Request) {
	if re != nil {
		http.Redirect(w, r, re.URI, re.Code)
	}
}
