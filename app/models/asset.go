package models

import (
	"errors"
	"html"
	"strings"
	"time"

	"github.com/jinzhu/gorm"
)

type Asset struct {
	ID             uint64 `gorm:"primary_key;auto_increment" json:"id"`
	UserID         uint64 `json:"uid"`
	IsFavorite     bool   `gorm:"not null" json:"isFavorite"`
	Title          string `gorm:"size:255;not null" json:"title"`
	TitleGenerated string `gorm:"-" json:"title_generated"`

	Chart    ChartDetails    `json:"chart"`
	Insight  InsightDetails  `json:"insight"`
	Audience AudienceDetails `json:"audience"`

	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
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
	assets := []Asset{}
	err := db.Debug().Model(&Asset{}).Limit(100).Order("created_at DESC").Where("user_id = ?", uid).Find(&assets).Error
	if err != nil {
		return &[]Asset{}, err
	}
	return &assets, err
}

func (a *Asset) SaveAsset(db *gorm.DB, uid uint64) (*Asset, error) {
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
	err := db.Debug().Model(&Asset{}).Where("id = ? and user_id = ?", id, uid).Take(&a).Error
	if err != nil {
		return &Asset{}, err
	}
	if a.ID != 0 {
		err = db.Debug().Model(&Asset{}).Where("id = ? and user_id = ?", a.Title).Error
		if err != nil {
			return &Asset{}, err
		}
	}
	return a, nil
}

func (a *Asset) DeleteAsset(db *gorm.DB, id uint64, uid uint32) (int64, error) {

	db = db.Debug().Model(&Asset{}).Where("id = ? and user_id = ?", id, uid).Take(&Asset{}).Delete(&Asset{})

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
	ID            uint   `gorm:"primary_key"`
	AssetId       uint   `json:"assetId"`
	GroupedMetric string `json:"groupedMetric"` // Could be age,
}

func (c *Chart) TableName() string {
	return "charts"
}

// Audience Section
type Audience struct {
	ID      uint64 `gorm:"primary_key"`
	AssetId uint64 `json:"assetId"`

	AgeFrom     uint8  `json:"ageFrom"`
	AgeTo       uint8  `json:"ageTo"`
	CountryCode string `json:"countryCode"`
	Gender      string `gorm:"size:1" json:"gender"`
}

func (a *Audience) TableName() string {
	return "audiences"
}

func (a *Audience) CreateAssetAudience(db *gorm.DB, uid uint64) (*Asset, *Audience, error) {
	asset := &Asset{}
	asset.UserID = uid
	asset, err := asset.SaveAsset(db, uid)
	if err != nil {
		return nil, nil, err
	}

	a.AssetId = asset.ID
	err = db.Debug().Create(&a).Error
	if err != nil {
		return asset, a, err
	}
	return asset, a, nil
}

// Insight Section
type Insight struct {
	ID      uint64 `gorm:"primary_key"`
	AssetId uint64 `json:"assetId"`

	Gender       string `gorm:"size:1" json:"insightDescription"`
	BirthCountry string `gorm:"size:2" json:"birthCountry"`
}

func (i *Insight) TableName() string {
	return "insights"
}

func (i *Insight) CreateAssetInsight(db *gorm.DB, uid uint64) (*Asset, *Insight, error) {
	asset := &Asset{}
	asset.UserID = uid
	asset, err := asset.SaveAsset(db, uid)
	if err != nil {
		return nil, nil, err
	}

	i.AssetId = asset.ID
	err = db.Debug().Create(&i).Error
	if err != nil {
		return asset, i, err
	}

	return asset, i, nil
}

// Generated Structs Section
type ChartDetails struct {
	Title  string
	XTitle string
	YTitle string
	Data   map[string]string
}

type InsightDetails struct {
	Title string
}

type AudienceDetails struct {
	Title string
}
