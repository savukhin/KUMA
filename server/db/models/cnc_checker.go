package models

import "database/sql/driver"

type CncStatusEnum int64

const (
	Working CncStatusEnum = iota
	Stopped
	Broken
)

func (st CncStatusEnum) String() string {
	return [...]string{"Working", "Stopped", "Broken"}[st]
}

type CncStatus struct {
	Status CncStatusEnum
}

func (st *CncStatusEnum) Scan(value interface{}) error {
	*st = CncStatusEnum(value.(int64))
	return nil
}

func (st CncStatusEnum) Value() (driver.Value, error) {
	return st.String(), nil
}

type carType string

const (
	SEDAN     carType = "SEDAN"
	HATCHBACK carType = "HATCHBACK"
	MINIVAN   carType = "MINIVAN"
)

func (self *carType) Scan(value interface{}) error {
	*self = carType(value.([]byte))
	return nil
}

func (self carType) Value() (driver.Value, error) {
	return string(self), nil
}

type CncChecker struct {
	ID uint64 `gorm:"column:id;type:bigint;primaryKey" json:"id"`

	// Status CncStatusEnum `sql:"type:ENUM('SEDAN', 'HATCHBACK', 'MINIVAN')" gorm:"column:car_type"`
	Status  CncStatusEnum
	CarType carType `sql:"type:ENUM('SEDAN', 'HATCHBACK', 'MINIVAN')" gorm:"column:car_type"`
	// Car3Type carType `gorm:"type:car2_type"`
	// Car3Type carType `gorm:"type:enum('SEDAN', 'HATCHBACK', 'MINIVAN')"`

	Credentials
	TimingAt
}
