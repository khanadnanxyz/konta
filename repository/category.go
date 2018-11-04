package repository

import (
	"github.com/khanadnanxyz/konta/conn"
	"github.com/khanadnanxyz/konta/model"
)

func CreateCategory(category *model.Category) (*model.Category, error) {
	if err := db.Db.Create(category).Error; err != nil {
		return nil, err
	}
	return category, nil
}

func GetCategoryById(id uint64) (*model.Category, error) {
	category := model.Category{}
	if err := db.Db.First(&category, id).Error; err != nil {
		return nil, err
	}
	return &category, nil
}