package handlers

import (
	"net/http"

	// "time"

	"gwiapi/app/auth"
	"gwiapi/app/models"
	"gwiapi/app/responses"

	"github.com/jinzhu/gorm"
)

// GetAssets will bring asset of the user
func GetAssets(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	asset := models.Asset{}

	// Select only User Assets from Token
	uid, err := auth.ExtractTokenID(r)

	assets, err := asset.GetAssets(db, uid)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(w, http.StatusOK, assets)
}

func GetParticipants(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	participant := models.Participant{}

	participants, err := participant.GetAllParticipants(db)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(w, http.StatusOK, participants)
}
