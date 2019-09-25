package app

import (
	"fmt"
	"log"
	"net/http"

	"gwiapi/app/models"
	"gwiapi/config"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

// App Structure
type App struct {
	Router *mux.Router
	DB     *gorm.DB
}

// Initialize initializes the app with predefined configuration
func (a *App) Initialize(config *config.Config) {
	dbURI := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable password=%s",
		config.DB.Host,
		config.DB.Port,
		config.DB.Username,
		config.DB.Name,
		config.DB.Password,
		// config.DB.Charset
	)

	db, err := gorm.Open(config.DB.Dialect, dbURI)
	if err != nil {
		log.Fatal("Could not connect database", err)
	}

	a.DB = models.DBMigrate(db)
	a.Router = mux.NewRouter()
	a.setRouters()
}

// Run the app on Mux router
func (a *App) Run(host string) {
	fmt.Printf("Serving on http://127.0.0.1%s", host)
	// In case this port is used
	log.Fatal(http.ListenAndServe(host, a.Router))
}
