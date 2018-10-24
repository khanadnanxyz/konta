package model

import "time"

// User represents a user's account model
type User struct {
	ID        uint64    `gorm:"AUTO_INCREMENT;primary_key"`
	Role      Role      `gorm:"foreignkey:RoleId"`
	RoleID    int64     `gorm:"not null"`
	FullName  string    `gorm:"type:varchar(36);not null"`
	username  string    `gorm:"type:varchar(36);not null"`
	password  string    `gorm:"type:varchar(36);not null"`
	email     string    `gorm:"type:varchar(36);not null"`
	phone     string    `gorm:"type:varchar(36);"`
	CreatedAt time.Time `gorm:"default:'now()';not null"`
	UpdatedAt time.Time `gorm:"default:'now()';not null"`
}

// TableName set a custom table name for model
func (*User) TableName() string {
	return "users"
}
