package models

import (
	"errors"
	"html"
	"strings"
	"time"

	"github.com/jinzhu/gorm"
)

type Asset struct {
	gorm.Model
	ID         uint64    `gorm:"primary_key;auto_increment" json:"id"`
	Title      string    `gorm:"size:255;not null;unique" json:"title"`
	IsFavorite bool      `gorm:"not null" json:"isFavorite"`
	CreatedAt  time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt  time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}

func (a *Asset) Prepare() {
	a.ID = 0
	a.Title = html.EscapeString(strings.TrimSpace(a.Title))
	a.CreatedAt = time.Now()
	a.UpdatedAt = time.Now()
}

func (a *Asset) SaveAsset(db *gorm.DB) (*Asset, error) {
	var err error
	err = db.Debug().Model(&Asset{}).Create(&a).Error
	if err != nil {
		return &Asset{}, err
	}
	if a.ID != 0 {
		err = db.Debug().Model(&User{}).Where("id = ?", a.Title).Error
		if err != nil {
			return &Asset{}, err
		}
	}
	return a, nil
}

func (a *Asset) GetAsset(db *gorm.DB, id uint64) (*Asset, error) {
	var err error
	err = db.Debug().Model(&Asset{}).Where("id = ?", id).Take(&a).Error
	if err != nil {
		return &Asset{}, err
	}
	if a.ID != 0 {
		err = db.Debug().Model(&Asset{}).Where("id = ?", a.Title).Error
		if err != nil {
			return &Asset{}, err
		}
	}
	return a, nil
}

func (a *Asset) DeleteAsset(db *gorm.DB, id uint64, uid uint32) (int64, error) {

	db = db.Debug().Model(&Asset{}).Where("id = ? and userID = ?", id, uid).Take(&Asset{}).Delete(&Asset{})

	if db.Error != nil {
		if gorm.IsRecordNotFoundError(db.Error) {
			return 0, errors.New("Asset not found")
		}
		return 0, db.Error
	}
	return db.RowsAffected, nil
}
