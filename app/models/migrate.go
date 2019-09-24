package models 

import (

	"github.com/jinzhu/gorm"

)

func DBMigrate(db *gorm.DB) *gorm.DB {
	db.AutoMigrate(&Asset{})
	db.AutoMigrate(&User{})
	db.AutoMigrate(&Data{})
	// db.Model(&Chart{}).AddForeignKey("AssetId", "Assets(id)", "CASCADE", "CASCADE")
	return db
}