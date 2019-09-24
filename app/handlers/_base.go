package handlers

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres" 

	"github.com/teliaz/goapi/app/models"
)

type App struct {
	DB     *gorm.DB
	Router *mux.Router
}

func (app *App) Initialize(DbUser, DbPassword, DbPort, DbHost, DbName string) {

	var err error
	
	DBURL := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable password=%s", DbHost, DbPort, DbUser, DbName, DbPassword)
	App.DB, err = gorm.Open(app., DBURL)
	if err != nil {
		fmt.Printf("Cannot connect to %s database", Dbdriver)
		log.Fatal("This is the error:", err)
	} else {
		fmt.Printf("We are connected to the %s database", Dbdriver)
	}
	

	app.DB.Debug().AutoMigrate(&models.User{}, &models.Post{}) //database migration
	app.Router = mux.NewRouter()

	app.initializeRoutes()
}

func (app *App) StartApplication(addr string) {
	log.Fatal(http.ListenAndServe(addr, app.Router))
}