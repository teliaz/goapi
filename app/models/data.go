package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

type Data struct {
	gorm.Model
	ID                      uint64    `gorm:"primary_key;auto_increment" json:"id"`
	HoursSpendOnSocialDaily uint      `json:"hoursSpendDailyOnSocialMedia"`
	Age                     uint      `json:"age"`
	Gender                  string    `json:"gender"`
	CreatedAt               time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
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
