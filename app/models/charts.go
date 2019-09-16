package models

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type Chart struct {
	gorm.Model
	Title      string `gorm:"column:title"`
	Slug       string `gorm:"column:slug" json:"slug"`
	IsFavorite string `gorm:"column:isFavorite" json:"isFavorite"`
}

func (c *Chart) TableName() string {
	return "charts"
}

func InsertChart(db *gorm.DB, c *Chart) (err error) {
	if err = db.Save(c).Error; err != nil {
		return err
	}
	return nil
}

func GetAll(db *gorm.DB, c *Chart) (err error) {
	if err = db.Order("title desc").Error; err != nil {
		return err
	}
	return nil
}

func GetChart(db *gorm.DB, slug string, b *Chart) (err error) {
	if err := db.Where("slug = ?", ids).First(&c).Error; err != nil {
		return err
	}
	return nil
}
