package db

import "time"

var models = []interface{}{
	&User{},
}

type User struct {
	ID        uint64    `gorm:"primarykey"`
	CreatedAt time.Time `gorm:"index"`
	UpdatedAt time.Time `gorm:"index"`
	/*TelegramChatID   int64     `gorm:"index"`
	TelegramNickname string    `gorm:"type:varchar(24);default:Anonymous"`
	UTMSource        string    `gorm:"type:varchar(24);index"`
	UTMCampaign      string    `gorm:"type:varchar(24);index"`
	Yclid            string    `gorm:"type:varchar(32)"` // yandex click ID
	LastReminderAt   time.Time `gorm:"index;default:CURRENT_TIMESTAMP()"`
	IsUserBlockBot   bool      `gorm:"index;default:false"`*/
}
