package api

import (
	"encoding/json"
	"net/http"

	"github.com/khanadnanxyz/konta/api/payload"
	"github.com/khanadnanxyz/konta/model"
	"github.com/khanadnanxyz/konta/repository"
	"github.com/khanadnanxyz/konta/util"
)

func AllQuestions(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello Go"))
}
func DetailQuestion(w http.ResponseWriter, r *http.Request) {}

func CreateQuestion(w http.ResponseWriter, r *http.Request) {
	body := payload.QuestionPayload{}
	if err := util.ParseBody(r, &body); err != nil {
		w.Write([]byte(err.Error()))
		return
	}
	if err := body.Validate(); err != nil {
		w.Write([]byte(err.Error()))
		return
	}

	//
	category, err := repository.GetCategoryById(body.CategoryId)
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}
	question := model.Question{}
	question.QText = body.QText
	question.CategoryID = category.ID
	var options []model.Option

	for _, opt := range body.Options {
		option := model.Option{}
		option.OptText = opt
		options = append(options, option)
	}

	question.Option = options
	question.UserID = 1
	q, err := repository.AddQuestion(&question)

	if err != nil {
		return
	}
	qj, err := json.Marshal(q)
	w.Write(qj)
}

func UpdateQuestion(w http.ResponseWriter, r *http.Request) {}
func DeleteQuestion(w http.ResponseWriter, r *http.Request) {}
