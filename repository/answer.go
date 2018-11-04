package repository

import (
	"github.com/khanadnanxyz/konta/conn"
	"github.com/khanadnanxyz/konta/model"
)

func AddAnswer(answer *model.Answer) (*model.Answer, error)  {
	if err := db.Db.Create(&answer).Error; err != nil {
		return nil, err
	}
	return answer, nil
}