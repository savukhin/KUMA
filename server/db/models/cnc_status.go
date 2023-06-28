package models

type CncStatus struct {
	ID         uint64 `gorm:"column:id;type:bigint;primaryKey" json:"id"`
	StatusName string `gorm:"column:status_name;type:varchar(70);unique" json:"status_name"`
}
