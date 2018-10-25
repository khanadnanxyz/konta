package model

import "time"

// Answer represents a user's role model
type Answer struct {
	ID        uint64    `gorm:"AUTO_INCREMENT;primary_key"`
	User      User      `gorm:"foreignkey:UserID"`
	UserID    int64     `gorm:"not null"`
	Option    Option    `gorm:"foreignkey:OptionID"`
	OptionID  int64     `gorm:"not null"`
	RoleName  string    `gorm:"type:varchar(255);not null"`
	CreatedAt time.Time `gorm:"default:'now()';not null"`
	UpdatedAt time.Time `gorm:"default:'now()'"`
}

// TableName set a custom table name for Answer model
func (*Answer) TableName() string {
	return "account.answers"
}
