package migrate

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

// Migrate Models re-write to avoid circular dependencies
type Chart struct {
	gorm.Model
	Id         uint `gorm:"primary_key"`
	Title      int  `gorm:"column:title" json:"title"`
	IsFavorite bool `gorm:"column:isFavorite" json:"name"`
}

func DBMigrate(db *gorm.DB) *gorm.DB {

	db.AutoMigrate(&Chart{})
	tableChartsExists := db.HasTable(&Chart{})
	fmt.Println("Table Chart is ", tableChartsExists)
	if !tableChartsExists {
		db.CreateTable(&Chart{})
	}

	return db
}
