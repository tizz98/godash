package api

import (
	"encoding/json"
	"fmt"
	"github.com/tizz98/godash/app"
	"net/http"
)

type Api struct {
	App *app.App
}

func NewApi(app *app.App) *Api {
	return &Api{App: app}
}

func (a *Api) Context() *app.Context {
	return a.App.Context
}

func (a *Api) ErrorMessage(w http.ResponseWriter, msg string, status int) {
	a.Error(w, fmt.Errorf(msg), status)
}

func (a *Api) Error(w http.ResponseWriter, err error, status int) {
	resp := GenericResponse{
		Ok:      false,
		Message: err.Error(),
	}
	w.WriteHeader(status)
	a.WriteJSON(w, resp)
}

func (a *Api) WriteJSON(w http.ResponseWriter, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}

type GenericResponse struct {
	Ok      bool   `json:"ok"`
	Message string `json:"message"`
}
