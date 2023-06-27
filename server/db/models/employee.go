package models

type Employee struct {
	ID             uint64 `gorm:"column:id;type:bigint;primaryKey" json:"id"`
	Name           string `gorm:"column:name;type:varchar(72);primaryKey" json:"name"`
	TelegramUserID string `gorm:"column:telegram_user_id;type:varchar(60);primaryKey" json:"telegram_user_id"`

	Credentials
	TimingAt
}
