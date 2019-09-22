package handlers

import (
	"net/http"

	"github.com/jinzhu/gorm"
	"github.com/teliaz/goapi/app/utils"
)

func SignIn(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	token, err := utils.GenerateToken()
	if err != nil {
		respondError(w, http.StatusBadRequest, err.Error())
		return
	}
	respondJSON(w, http.StatusOK, token)
}
