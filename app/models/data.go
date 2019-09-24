package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

type Data struct {
	ID        uint64    `gorm:"primary_key;auto_increment" json:"id"`
	HoursSpendOnSocialDaily uint `gorm: json:"hoursSpendDailyOnSocialMedia"`
	Age 				uint `gorm: json:"hoursSpendDailyOnSocialMedia"`
	Gender string `gorm: json:"hoursSpendDailyOnSocialMedia"`
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
}

func (d *Data) SaveData(db *gorm.DB) (*Data, error) {
	err := db.Debug().Model(&Data{}).Create(&d).Error
	if err != nil {
		return &Data{}, err
	}
	if d.ID != 0 {
		err = db.Debug().Model(&Data{}).Where("id = ?", d.ID).Take(&d.CreatedAt).Error
		if err != nil {
			return &Data{}, err
		}
	}
	return d, nil
}