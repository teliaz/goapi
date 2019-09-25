package handlers

import (
	"net/http"

	// "time"

	"github.com/jinzhu/gorm"
	"gwiapi/app/auth"
	"gwiapi/app/models"
	"gwiapi/app/responses"
)

// GetAssets will bring asset of the user
func GetAssets(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	asset := models.Asset{}

	// Select only User Assets
	uid, err := auth.ExtractTokenID(r)

	assets, err := asset.GetAssets(db, uid)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(w, http.StatusOK, assets)
}
