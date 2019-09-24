package handlers

import (
	"net/http"
	"time"

	"github.com/teliaz/goapi/app/responses"
)

// Ping will respond with Time of Call
func Ping(w http.ResponseWriter, r *http.Request) {
	responses.JSON(w, http.StatusOK, time.Now())
}
