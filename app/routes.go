package app

import (
	"net/http"

	"gwiapi/app/handlers"
	"gwiapi/app/middlewares"

	"github.com/jinzhu/gorm"
)

func (a *App) setRouters() {

	// Health Check
	a.Get("/", middlewares.SetMiddlewareJSON(handlers.Ping, a.DB))

	// Auth Routes
	a.Post("/auth/signup", middlewares.SetMiddlewareJSON(handlers.CreateUser, a.DB))
	a.Post("/auth/login", middlewares.SetMiddlewareJSON(handlers.Login, a.DB))

	// Users routes
	a.Get("/users", middlewares.SetMiddlewareAuthentication(handlers.GetUsers, a.DB))
	a.Get("/users/{id}", middlewares.SetMiddlewareAuthentication(handlers.GetUser, a.DB))
	a.Put("/users/{id}", middlewares.SetMiddlewareAuthentication(handlers.UpdateUser, a.DB))
	a.Delete("/users/{id}", middlewares.SetMiddlewareAuthentication(handlers.DeleteUser, a.DB))

	// Assets Routes
	a.Get("/assets", middlewares.SetMiddlewareAuthentication(handlers.GetAssets, a.DB))
	a.Get("/assets/{id:[0-9]+}", middlewares.SetMiddlewareAuthentication(handlers.GetAsset, a.DB))
	a.Patch("/assets/{id:[0-9]+}", middlewares.SetMiddlewareAuthentication(handlers.UpdateAsset, a.DB))
	a.Post("/assets/charts", middlewares.SetMiddlewareAuthentication(handlers.CreateAssetChart, a.DB))
	a.Post("/assets/insights", middlewares.SetMiddlewareAuthentication(handlers.CreateAssetInsight, a.DB))
	a.Post("/assets/audiences", middlewares.SetMiddlewareAuthentication(handlers.CreateAssetAudience, a.DB))

	// Participants
	a.Get("/participants", middlewares.SetMiddlewareAuthentication(handlers.GetParticipants, a.DB))
	a.Post("/participants", middlewares.SetMiddlewareAuthentication(handlers.AddParticipant, a.DB))
	a.Get("/countries", middlewares.SetMiddlewareJSON(handlers.GetCountries, a.DB))

}

// RequestHandlerFunction HandlerRequest extension
type RequestHandlerFunction func(db *gorm.DB, w http.ResponseWriter, r *http.Request)

func (a *App) handleRequest(handler RequestHandlerFunction) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		handler(a.DB, w, r)
	}
}
func (a *App) handleMiddleware(handler RequestHandlerFunction) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		handler(a.DB, w, r)
	}
}

// Get wrap HandleFunc for GET method
func (a *App) Get(path string, f func(w http.ResponseWriter, r *http.Request)) {
	a.Router.HandleFunc(path, f).Methods("GET")
}

// Get wrap HandleFunc for POST method
func (a *App) Post(path string, f func(w http.ResponseWriter, r *http.Request)) {
	a.Router.HandleFunc(path, f).Methods("POST")
}

// Get wrap HandleFunc for Patch method
func (a *App) Patch(path string, f func(w http.ResponseWriter, r *http.Request)) {
	a.Router.HandleFunc(path, f).Methods("POST")
}

// Get wrap HandleFunc for PUT method
func (a *App) Put(path string, f func(w http.ResponseWriter, r *http.Request)) {
	a.Router.HandleFunc(path, f).Methods("PUT")
}

// Get wrap HandleFunc for DELETE method
func (a *App) Delete(path string, f func(w http.ResponseWriter, r *http.Request)) {
	a.Router.HandleFunc(path, f).Methods("DELETE")
}
