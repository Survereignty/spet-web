package api

import (
	"SPET/internal/spet-api/db"
	"SPET/internal/spet-api/settings"
	"SPET/pkg/lig"
	"net/http"

	"github.com/gorilla/mux"
)

// Объект с конфигурацией API сервера
type Api struct {
	config *settings.Config
	router *mux.Router
	store  *db.Store
}

// Создание экземпляра API
func New(config *settings.Config) *Api {
	return &Api{
		config: config,
		router: mux.NewRouter(),
	}
}
func (api *Api) Routers() {
	api.router.HandleFunc("/ping", api.Ping()).Methods(http.MethodGet)
}

// Запуск API
func (api *Api) Start() error {
	if err := api.Store(); err != nil {
		return err
	}

	api.Routers()

	lig.Info("Server listen port: " + api.config.Address)

	return http.ListenAndServe(":"+api.config.Address, api.router)
}

func (api *Api) Store() error {
	store := db.New(api.config)

	if api.config.Database.Name == "postgres" {
		if err := store.OpenPostgres(); err != nil {
			return err
		}
	}

	api.store = store

	return nil
}
