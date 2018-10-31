package repository

import (
	"github.com/khanadnanxyz/konta/conn"
	"github.com/khanadnanxyz/konta/model"
)

func CreateQuestion2(question *model.Question2) (*model.Question2, error)  {
	if err := db.Db.Create(&question).Error; err != nil {
		return nil, err
	}
	return question, nil
}
