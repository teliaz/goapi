package handlers

import (
	"net/http"
	"strconv"

	// "time"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	"github.com/teliaz/goapi/app/models"
	"github.com/teliaz/goapi/app/responses"
)

// GetAssets will bring asset of the user
func GetAssets(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	asset := models.Asset{}

	vars := mux.Vars(r)
	id, err := strconv.ParseUint(vars["id"], 10, 32)

	// Select only User Assets
	// auth.ExtractTokenID()

	assets, err := asset.GetAsset(db, id)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(w, http.StatusOK, assets)
}
