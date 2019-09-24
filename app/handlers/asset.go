package handlers

import (
	"time"
	"github.com/teliaz/goapi/app/auth"
	"github.com/teliaz/goapi/app/responses"
	"github.com/jinzhu/gorm"
)

func GetAssets(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	post := models.Asset{}
	posts, err := post.FindAllPosts(db)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(w, http.StatusOK, posts)
}
