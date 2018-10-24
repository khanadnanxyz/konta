package model

import "time"

// Role represents a user's role model
type Role struct {
	ID        uint64    `gorm:"AUTO_INCREMENT;primary_key"`
	RoleName  string    `gorm:"type:varchar(255);not null"`
	CreatedAt time.Time `gorm:"default:'now()';not null"`
	UpdatedAt time.Time `gorm:"default:'now()'"`
}

// TableName set a custom table name for model
func (*Role) TableName() string {
	return "test.roles"
}
