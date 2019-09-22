package models

import (
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

// Assets
type Asset struct {
	gorm.Model
	Id         uint       `gorm:"primary_key"`
	Title      string     `gorm:"column:title"`
	IsFavorite bool       `gorm:"column:isFavorite" json:"isFavorite"`
	CreatedAt  *time.Time `gorm:"default:null" json:"createdAt"`
}

func (c *Asset) TableName() string {
	return "Assets"
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
	Id         uint   `gorm:"primary_key"`
	AssetId    uint   `json:"assetId"`
	XAxisTitle string `json:"xAxisTitle"`
	YAxisTitle string `json:"yAxisTitle"`
}

// DBMigrate will create and migrate the tables, and then make the some relationships if necessary
func DBMigrate(db *gorm.DB) *gorm.DB {
	db.AutoMigrate(&Asset{}, &Chart{})
	db.Model(&Chart{}).AddForeignKey("AssetId", "Assets(id)", "CASCADE", "CASCADE")
	return db
}
