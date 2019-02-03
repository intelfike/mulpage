package types

import "net/http"

type Redirect struct {
	URI  string
	Code int
}

func NewRedirect(URI string, Code int) *Redirect {
	return &Redirect{URI: URI, Code: Code}
}

func (re *Redirect) Exec(w http.ResponseWriter, r *http.Request) {
	if re != nil {
		http.Redirect(w, r, re.URI, re.Code)
	}
}
