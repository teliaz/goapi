package middlewares

import (
	"errors"
	"net/http"

	"gwiapi/app/auth"
	"gwiapi/app/responses"

	"github.com/jinzhu/gorm"
)

// RequestHandlerFunction Overloading HandlerFunc with db
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
			w.Header().Set("Content-Type", "application/json")
			responses.ERROR(w, http.StatusUnauthorized, errors.New("Unauthorized"))
			return
		}
		next(db, w, r)
	}
}
