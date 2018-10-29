package model

import "time"

// User represents a user's account model
type User struct {
	ID        uint64    `gorm:"AUTO_INCREMENT;primary_key"`
	Role      Role      `gorm:"foreignkey:RoleId"`
	RoleID    uint64     `gorm:"not null"`
	FullName  string    `gorm:"type:varchar(255);not null"`
	Username  string    `gorm:"type:varchar(100);not null"`
	Password  string    `gorm:"type:varchar(100);not null"`
	Email     string    `gorm:"type:varchar(100);not null"`
	Phone     string    `gorm:"type:varchar(100);"`
	CreatedAt time.Time `gorm:"default:'now()';not null"`
	UpdatedAt time.Time `gorm:"default:'now()'"`
}

// TableName set a custom table name for model
func (*User) TableName() string {
	return "account.users"
}
