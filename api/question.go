package api

import (
	"encoding/json"
	"github.com/khanadnanxyz/konta/model"
	"github.com/khanadnanxyz/konta/repository"
	"net/http"
)

func AllQuestion(w http.ResponseWriter, r *http.Request) {}
func DetailQuestion(w http.ResponseWriter, r *http.Request) {}
func CreateQuestion(w http.ResponseWriter, r *http.Request) {
	body := model.Question2{}
	if err := parseBody(r, &body); err != nil {
		w.Write([]byte(err.Error()))
		return
	}

	if err := body.Validate(); err != nil {
		w.Write([]byte(err.Error()))
		return
	}

	repository.CreateQuestion2(&body)

}
func UpdateQuestion(w http.ResponseWriter, r *http.Request) {}
func DeleteQuestion(w http.ResponseWriter, r *http.Request) {}

// parseBody parses request body to given data struct
func parseBody(r *http.Request, v interface{}) error {
	b := json.NewDecoder(r.Body)
	b.DisallowUnknownFields()
	return b.Decode(v)
}