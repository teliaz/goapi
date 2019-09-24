package handlers

import (
	"net/http"
	"time"

	"github.com/jinzhu/gorm"
	"github.com/teliaz/goapi/app/responses"
)

// Ping will respond with Time of Call
func Ping(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	responses.JSON(w, http.StatusOK, time.Now())
}
