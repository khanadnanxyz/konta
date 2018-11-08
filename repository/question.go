package repository

import (
	"github.com/khanadnanxyz/konta/conn"
	"github.com/khanadnanxyz/konta/model"
)

func GetAllQuestion() {}

func AddQuestion(question *model.Question) (*model.Question, error) {
	if err := db.Db.Create(&question).Error; err != nil {
		return nil, err
	}
	return question, nil
}
