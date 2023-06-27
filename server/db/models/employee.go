package models

type Employee struct {
	ID uint64 `gorm:"column:id;type:bigint;primaryKey" json:"id"`

	Credentials
	TimingAt
}
