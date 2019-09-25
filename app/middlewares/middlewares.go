package middlewares

import (
	"errors"
	"net/http"

	"github.com/jinzhu/gorm"
	"gwiapi/app/auth"
	"gwiapi/app/responses"
)

type RequestHandlerFunction func(db *gorm.DB, w http.ResponseWriter, r *http.Request)

// SetMiddlewareJSON Adds Json Content Type
func SetMiddlewareJSON(next RequestHandlerFunction, db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		next(db, w, r)
	}
}

// SetMiddlewareAuthentication Validates Tokens
func SetMiddlewareAuthentication(next RequestHandlerFunction, db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		err := auth.TokenValid(r)
		if err != nil {
			responses.ERROR(w, http.StatusUnauthorized, errors.New("Unauthorized"))
			return
		}
		next(db, w, r)
	}
}
