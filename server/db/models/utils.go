package models

import (
	"gorm.io/gorm"
)

type IDType int

type Credentials struct {
	Username     string `gorm:"column:username;type:varchar(20);not null;unique" json:"username"`
	PasswordHash string `gorm:"column:password_hash;type:varchar(72);not null;unique" json:"password_hash"`
}

type TimingAt struct {
	CreatedAt uint64         `gorm:"column:created_at;type:timestamp with time zone;not null;now" json:"created_at"`
	UpdatedAt uint64         `gorm:"column:updated_at;type:timestamp with time zone;not null;now" json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at;type:timestamp with time zone" json:"deleted_at"`
}

var MigrateModels = []interface{}{
	CncStatus{},
	CncChecker{},
	Employee{},
}
