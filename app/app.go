package app

import (
	"github.com/gorilla/mux"
)

type app struct {
	Router *mux.Router
}

func New() *app {
	app:= &app {
		Router: mux.NewRouter(),
	}
	app.initRoutes()
	return app
}

func (a *app) initRoutes() {
	a.Router.HandleFunc("/", a.IndexHandler()).Methods("GET")
	a.Router.HandleFunc("/health", a.HealthHandler()).Methods("GET")
	a.Router.HandleFunc("/readiness", a.ReadinessHandler()).Methods("GET")
}