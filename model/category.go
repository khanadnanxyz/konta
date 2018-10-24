package model

import "time"

type Category struct {
	ID				uint64    	`gorm:"AUTO_INCREMENT;primary_key"`
	CategoryName 	string		`gorm:"type:varchar(200);not null"`
	CreatedAt     	time.Time	`gorm:"default:'now()';not null"`
	UpdatedAt     	time.Time	`gorm:"default:'now()'"`
}


// Set QuestionCategory's table name to be `question_categories`
func (*Category) TableName() string {
	return "test.categories"
}