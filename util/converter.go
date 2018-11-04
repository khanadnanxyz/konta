package util

import (
	"encoding/json"
	"net/http"
)

// parseBody parses request body to given data struct
func ParseBody(r *http.Request, v interface{}) error {
	b := json.NewDecoder(r.Body)
	b.DisallowUnknownFields()
	return b.Decode(v)
}