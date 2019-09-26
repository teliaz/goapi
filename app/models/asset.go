package models

import (
	"errors"
	"html"
	"strings"
	"time"

	"github.com/jinzhu/gorm"
)

type Asset struct {
	ID         uint64    `gorm:"primary_key;auto_increment" json:"id"`
	Title      string    `gorm:"size:255;not null;unique" json:"title"`
	IsFavorite bool      `gorm:"not null" json:"isFavorite"`
	CreatedAt  time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt  time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}

func (a *Asset) TableName() string {
	return "assets"
}

func (a *Asset) Prepare() {
	a.ID = 0
	a.Title = html.EscapeString(strings.TrimSpace(a.Title))
	a.CreatedAt = time.Now()
	a.UpdatedAt = time.Now()
}

func (a *Asset) GetAssets(db *gorm.DB, uid uint64) (*[]Asset, error) {
	asset := Asset{}
	// Add Performance improvements
	assets, err := asset.GetAssets(db, uid)
	if err != nil {
		return nil, err
	}
	return assets, nil
}

func (a *Asset) SaveAsset(db *gorm.DB) (*Asset, error) {
	err := db.Debug().Model(&Asset{}).Create(&a).Error
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

func (a *Asset) GetAsset(db *gorm.DB, id uint64, uid uint64) (*Asset, error) {
	err := db.Debug().Model(&Asset{}).Where("id = ? and userId = ?", id, uid).Take(&a).Error
	if err != nil {
		return &Asset{}, err
	}
	if a.ID != 0 {
		err = db.Debug().Model(&Asset{}).Where("id = ? and userId = ?", a.Title).Error
		if err != nil {
			return &Asset{}, err
		}
	}
	return a, nil
}

func (a *Asset) DeleteAsset(db *gorm.DB, id uint64, uid uint32) (int64, error) {

	db = db.Debug().Model(&Asset{}).Where("id = ? and UserID = ?", id, uid).Take(&Asset{}).Delete(&Asset{})

	if db.Error != nil {
		if gorm.IsRecordNotFoundError(db.Error) {
			return 0, errors.New("Asset not found")
		}
		return 0, db.Error
	}
	return db.RowsAffected, nil
}

// Chart Section
type Chart struct {
	ID         uint   `gorm:"primary_key"`
	AssetId    uint   `json:"assetId"`
	XAxisTitle string `json:"xAxisTitle"`
	YAxisTitle string `json:"yAxisTitle"`
}

func (c *Chart) TableName() string {
	return "charts"
}

// Insight Section
type Insight struct {
	ID                 uint64 `gorm:"primary_key"`
	AssetId            uint64 `json:"assetId"`
	InsightDescription string `json:"insightDescription"`
}

func (i *Insight) TableName() string {
	return "insights"
}

// Audience Section
type Audience struct {
	ID                           uint64 `gorm:"primary_key"`
	AssetId                      uint64 `json:"assetId"`
	Gender                       string `json:"insightDescription"`
	BirthCountry                 string `json:"birthCountry"`
	AgeGroup                     string `json:"ageGroup"`
	HoursSpentDailyOnSocialMedia uint   `json:"hoursSpentDailyOnSocialMedia"`
}

func (a *Audience) TableName() string {
	return "audiences"
}
