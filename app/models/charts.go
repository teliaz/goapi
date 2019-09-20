package models

import (
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

// Assets
type Asset struct {
	gorm.Model
	Id         uint       `gorm:"column:id" json:"id"`
	Title      string     `gorm:"column:title"`
	IsFavorite bool       `gorm:"column:isFavorite" json:"isFavorite"`
	CreatedAt  *time.Time `gorm:"default:null" json:"createdAt"`
}

func (c *Asset) TableName() string {
	return "assets"
}

func (a *Asset) isFavorite(isFavorite bool) {
	a.IsFavorite = isFavorite
}

func (a *Asset) SetAssetTitle(title string) {
	a.Title = title
}

// Chart Section
type Chart struct {
	gorm.Model
	AssetId    uint   `json:"assetId"`
	XAxisTitle string `json:"xTitle"`
	YAxisTitle string `json:"yTitle"`
	// yAxisTitle string `gorm:"type:ENUM('0', '1', '2', '3');default:'0'" json:"priority"`
}

// DBMigrate will create and migrate the tables, and then make the some relationships if necessary
func DBMigrate(db *gorm.DB) *gorm.DB {
	db.AutoMigrate(&Asset{}, &Chart{})
	db.Model(&Chart{}).AddForeignKey("AssetId", "assets(id)", "CASCADE", "CASCADE")
	return db
}
