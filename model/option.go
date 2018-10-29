package model

import "time"

// Option represents a user's role model
type Option struct {
	ID         uint64   `gorm:"AUTO_INCREMENT;primary_key"`
	Question   Question `gorm:"foreignkey:QuestionID"`
	QuestionID uint64
	OptText    string    `gorm:"type:varchar(255);not null"`
	CreatedAt  time.Time `gorm:"default:'now()';not null"`
	UpdatedAt  time.Time `gorm:"default:'now()'"`
}

// TableName set a custom table name for Option model
func (*Option) TableName() string {
	return "poll.options"
}
