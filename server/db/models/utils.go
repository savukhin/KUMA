package models

import (
	"time"

	"gorm.io/gorm"
)

type Credentials struct {
	Username     string `gorm:"column:username;type:varchar(20);not null;unique" json:"username"`
	PasswordHash string `gorm:"column:password_hash;type:varchar(72);not null;unique" json:"password_hash"`
}

type TimingAt struct {
	CreatedAt time.Time      `gorm:"column:created_at;type:timestamp with time zone;not null;default:now()" json:"created_at"`
	UpdatedAt time.Time      `gorm:"column:updated_at;type:timestamp with time zone;not null;default:now()" json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at;type:timestamp with time zone" json:"deleted_at"`
}

var MigrateModels = []interface{}{
	CncChecker{},
	Employee{},
}
