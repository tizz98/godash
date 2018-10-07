package api

import (
	"net/http"
)

func (a *Api) SearchStocks(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query().Get("q")
	if query == "" {
		a.ErrorMessage(w, "no query provided", http.StatusBadRequest)
		return
	}

	stocks, err := a.Context().SearchStocks(query)
	if err != nil {
		a.Error(w, err, http.StatusInternalServerError)
		return
	}

	a.WriteJSON(w, stocks)
}
