package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

type Participant struct {
	ID                      uint64    `gorm:"primary_key;auto_increment" json:"id"`
	HoursSpendOnSocialDaily uint8     `json:"hoursSpendDailyOnSocialMedia"`
	Age                     uint8     `gorm:"size:1" json:"age"`
	Gender                  string    `json:"gender"`
	CreatedAt               time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
}

func (p *Participant) TableName() string {
	return "participants"
}

func (p *Participant) SaveParticipant(db *gorm.DB) (*Participant, error) {
	err := db.Debug().Model(&Participant{}).Create(&p).Error
	if err != nil {
		return &Participant{}, err
	}
	if p.ID != 0 {
		err = db.Debug().Model(&Participant{}).Where("id = ?", p.ID).Take(&p.CreatedAt).Error
		if err != nil {
			return &Participant{}, err
		}
	}
	return p, nil
}
