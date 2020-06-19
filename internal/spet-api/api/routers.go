package api

import (
	"encoding/json"
	"net/http"
)

// Функция проверки сервера на запуск
func (api *Api) Ping() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		api.Respond(w, r, http.StatusOK, nil)
	}
}

func (api *Api) Error(w http.ResponseWriter, r *http.Request, code int, err error) {
	api.Respond(w, r, code, map[string]string{"Error": err.Error()})
}

func (api *Api) Respond(w http.ResponseWriter, r *http.Request, code int, data interface{}) {
	w.WriteHeader(code)
	if data != nil {
		err := json.NewEncoder(w).Encode(data)
		if err != nil {
			api.Error(w, r, http.StatusInternalServerError, err)
		}
	}
}
