package model

import "time"

// Question represents a question model
type Question struct {
	ID         uint64    `gorm:"AUTO_INCREMENT;primary_key"`
	User       User      `gorm:"foreignkey:UserID"`
	UserID     int64     `gorm:"not null"`
	Category   Category  `gorm:"foreignkey:CategoryID"`
	CategoryID int64     `gorm:"not null"`
	QText      string    `gorm:"type:varchar(255);not null"`
	CreatedAt  time.Time `gorm:"default:'now()';not null"`
	UpdatedAt  time.Time `gorm:"default:'now()'"`
}

// TableName set a custom table name for Question model
func (*Question) TableName() string {
	return "account.question"
}
