package app

import (
	"net/http"

	"github.com/jinzhu/gorm"
	"github.com/teliaz/goapi/app/handlers"
	"github.com/teliaz/goapi/app/middlewares"
)

func (a *App) setRouters() {

	// Health Check
	a.Get("/", middlewares.SetMiddlewareJSON(handlers.Ping, a.DB))

	// Login Route
	a.Post("/login", middlewares.SetMiddlewareJSON(handlers.Login, a.DB))

	//Users routes
	a.Post("/users", middlewares.SetMiddlewareAuthentication(handlers.CreateUser, a.DB))
	a.Get("/users", middlewares.SetMiddlewareAuthentication(handlers.GetUsers, a.DB))
	a.Get("/users/{id}", middlewares.SetMiddlewareAuthentication(handlers.GetUser, a.DB))
	a.Put("/users/{id}", middlewares.SetMiddlewareAuthentication(handlers.UpdateUser, a.DB))
	a.Delete("/users/{id}", middlewares.SetMiddlewareAuthentication(handlers.DeleteUser, a.DB))

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
