package models

import "database/sql/driver"

type CncStatusEnum string

const (
	Working CncStatusEnum = "Working"
	Stopped               = "Stopped"
	Broken                = "Broken"
)

func (st *CncStatusEnum) Scan(value interface{}) error {
	*st = CncStatusEnum(value.([]byte))
	return nil
}

func (st CncStatusEnum) Value() (driver.Value, error) {
	return st, nil
}

type CncChecker struct {
	ID uint64 `gorm:"column:id;type:bigint;primaryKey" json:"id"`

	Status CncStatusEnum `gorm:"cnc_status"`

	Credentials
	TimingAt
}
