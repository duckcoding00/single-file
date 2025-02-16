package api

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Application struct {
	router *mux.Router
	config AppConfig
}

type AppConfig struct {
	addr string
}

func NewApp(config AppConfig) *Application {
	return &Application{
		router: mux.NewRouter(),
		config: config,
	}
}

func (a *Application) RegisterRoute() {
	apiRouter := a.router.PathPrefix("/api/v1").Subrouter()

	apiRouter.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(map[string]string{
			"message": "ok",
		})
	}).Methods("GET")
}

func (a *Application) Run() {
	log.Println("server running on ", a.config.addr)
	if err := http.ListenAndServe(a.config.addr, a.router); err != nil {
		log.Fatalf("failed to running server :%s", err)
	}
}
