package models

// type CncStatusEnum string

// const (
// 	Working CncStatusEnum = "Working"
// 	Stopped CncStatusEnum = "Stopped"
// 	Broken  CncStatusEnum = "Broken"
// )

// func (st *CncStatusEnum) Scan(value interface{}) error {
// 	*st = CncStatusEnum(value.([]byte))
// 	return nil
// }

// func (st CncStatusEnum) Value() (driver.Value, error) {
// 	return st, nil
// }

type CncChecker struct {
	ID       uint64 `gorm:"column:id;type:bigint;primaryKey" json:"id"`
	Title    string `gorm:"column:title;type:varchar(80);" json:"title"`
	StatusID uint64
	Status   CncStatus `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	// Status CncStatus `gorm:"references:id"`

	Credentials
	TimingAt
}
