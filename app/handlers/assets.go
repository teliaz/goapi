package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"

	"github.com/teliaz/goapi/app/models"
)

// GetAssets - Get All Assets
func GetAssets(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	assets := []models.Asset{}
	db.Find(&assets)
	respondJSON(w, http.StatusOK, assets)
}

// GetAsset - Get Asset by Id
func GetAsset(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.ParseUint(vars["id"], 10, 32)
	asset := getAssetOr404(db, uint(id), w, r)
	if asset == nil {
		return
	}
	respondJSON(w, http.StatusOK, asset)
}

// UpdateAssetTitle Set Asset's Title
func UpdateAssetTitle(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id, _ := strconv.ParseUint(vars["id"], 10, 32)
	asset := getAssetOr404(db, uint(id), w, r)
	if asset == nil {
		return
	}

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&asset); err != nil {
		respondError(w, http.StatusBadRequest, err.Error())
		return
	}
	defer r.Body.Close()

	asset.SetAssetTitle(vars["title"])

	if err := db.Save(&asset).Error; err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondJSON(w, http.StatusOK, asset)
}

// UpdateAssetIsFavorite Add/Remove from favorites
func UpdateAssetIsFavorite(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id, _ := strconv.ParseUint(vars["id"], 10, 32)
	asset := getAssetOr404(db, uint(id), w, r)
	if asset == nil {
		return
	}

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&asset); err != nil {
		respondError(w, http.StatusBadRequest, err.Error())
		return
	}
	defer r.Body.Close()

	isFavorite, _ := strconv.ParseBool(vars["isFavorite"])

	asset.SetAssetIsFavorite(isFavorite)

	if err := db.Save(&asset).Error; err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondJSON(w, http.StatusOK, asset)
}

// DeleteAsset Remove Asset by Id
func DeleteAsset(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.ParseUint(vars["id"], 10, 32)

	if err := db.Where("Id = ?", id).Remove().Error; err != nil {
		respondError(w, http.StatusNotFound, err.Error())
		return
	}
	respondJSON(w, http.StatusOK, id)
}

// Gets an Asset instance if exists, or respond the 404 error otherwise
func getAssetOr404(db *gorm.DB, id uint, w http.ResponseWriter, r *http.Request) *models.Asset {
	asset := models.Asset{}
	if err := db.First(&asset, models.Asset{Id: id}).Error; err != nil {
		respondError(w, http.StatusNotFound, err.Error())
		return nil
	}
	return &asset
}
