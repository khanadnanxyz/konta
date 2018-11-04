package model

import "time"

// Answer represents a user's role model
type Answer struct {
	ID        uint64    `gorm:"AUTO_INCREMENT;primary_key"`
	User      User      `gorm:"foreignkey:UserID"`
	UserID    uint64     `gorm:"not null"`
	Question   Question `gorm:"foreignkey:QuestionID"`
	QuestionID uint64		`gorm:"not null"`
	Option    Option    `gorm:"foreignkey:OptionID"`
	OptionID  uint64     	`gorm:"not null"`
	CreatedAt time.Time `gorm:"default:'now()';not null"`
	UpdatedAt time.Time `gorm:"default:'now()'"`
}

// TableName set a custom table name for Answer model
func (*Answer) TableName() string {
	return "poll.answers"
}
