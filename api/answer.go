package api

import (
	"encoding/json"
	"github.com/khanadnanxyz/konta/api/payload"
	"github.com/khanadnanxyz/konta/model"
	"github.com/khanadnanxyz/konta/repository"
	"github.com/khanadnanxyz/konta/util"
	"net/http"
)

func AllAnswers(w http.ResponseWriter, r *http.Request) {}
func DetailAnswer(w http.ResponseWriter, r *http.Request) {}

func CreateAnswer(w http.ResponseWriter, r *http.Request) {
	body := payload.AnswerPayload{}

	if err := util.ParseBody(r, &body); err != nil {
		w.Write([]byte(err.Error()))
		return
	}

	if err := body.Validate(); err != nil {
		w.Write([]byte(err.Error()))
		return
	}

	answer := model.Answer{}
	answer.UserID = body.UserID
	answer.QuestionID = body.QuestionID
	answer.OptionID = body.OptionID

	ans, err := repository.AddAnswer(&answer)

	if err != nil {
		return
	}
	ansj, err := json.Marshal(ans)
	w.Write(ansj)
}

func UpdateAnswer(w http.ResponseWriter, r *http.Request) {}
func DeleteAnswer(w http.ResponseWriter, r *http.Request) {}
