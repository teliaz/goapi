package handlers

import (
	"net/http"

	// "time"

	"github.com/jinzhu/gorm"
	"github.com/teliaz/goapi/app/auth"
	"github.com/teliaz/goapi/app/models"
	"github.com/teliaz/goapi/app/responses"
)

// GetAssets will bring asset of the user
func GetAssets(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	assets := []models.Asset{}

	// vars := mux.Vars(r)
	// id, err := strconv.ParseUint(vars["id"], 10, 32)

	// Select only User Assets
	uid, err := auth.ExtractTokenID(r)

	assets, err := asset.GetAssets(db, id)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(w, http.StatusOK, assets)
}
