package handlers

import (
	"net/http"
	"strconv"

	"gwiapi/app/auth"
	"gwiapi/app/models"
	"gwiapi/app/responses"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
)

// GetAssets will bring asset of the user
func GetAssets(db *gorm.DB, w http.ResponseWriter, r *http.Request) {

	// Get Pagination variables
	page, found := mux.Vars(r)["page"]
	if !found {
		page = "1"
	}
	pg, err := strconv.ParseUint(page, 10, 32)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
	}

	limit, found := mux.Vars(r)["limit"]
	if !found {
		limit = "10"
	}
	lmt, err := strconv.ParseUint(limit, 10, 32)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
	}

	asset := models.Asset{}

	// Select only User Assets from Token
	uid, err := auth.ExtractTokenID(r)

	assets, err := asset.GetAssets(db, uid, uint32(pg), uint32(lmt))
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(w, http.StatusOK, assets)
}

func GetAsset(db *gorm.DB, w http.ResponseWriter, r *http.Request) {

	id, found := mux.Vars(r)["id"]
	if !found {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	assetId, err := strconv.ParseUint(id, 10, 32)

	// Select only User Asset from Token
	uid, err := auth.ExtractTokenID(r)

	a := &models.Asset{}
	a, err = a.GetAsset(db, uint32(assetId), uid)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}

	if &a.Audience != nil {

	}

	responses.JSON(w, http.StatusOK, a)
}
