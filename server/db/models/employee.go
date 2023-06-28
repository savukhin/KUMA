package models

type Employee struct {
	ID               uint64 `gorm:"column:id;type:bigint;primaryKey" json:"id"`
	Name             string `gorm:"column:name;type:varchar(72)" json:"name"`
	TelegramUserName string `gorm:"column:telegram_user_name;type:varchar(60)" json:"telegram_user_name"`
	CheckedIN        bool   `gorm:"column:checked_in;type:boolean" json:"checked_in"`

	Credentials
	TimingAt
}
