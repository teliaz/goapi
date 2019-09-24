package models

import (
	"github.com/jinzhu/gorm"
)

func DBMigrate(db *gorm.DB) *gorm.DB {

	db.DropTable(&Asset{}, &User{}, &Data{})
	db.AutoMigrate(&Asset{})
	db.AutoMigrate(&User{})
	db.AutoMigrate(&Data{})
	// db.Model(&Chart{}).AddForeignKey("AssetId", "Assets(id)", "CASCADE", "CASCADE")
	return db
}
