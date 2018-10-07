package api

import (
	"github.com/stretchr/testify/assert"
	"github.com/tizz98/godash/app"
	"io/ioutil"
	"net/http/httptest"
	"testing"
)

func TestApi_SearchStocks(t *testing.T) {
	app := app.NewApp()
	api := NewApi(app)

	req := httptest.NewRequest("GET", "http://example.com/stocks?q=under", nil)
	w := httptest.NewRecorder()
	api.SearchStocks(w, req)

	resp := w.Result()
	body, _ := ioutil.ReadAll(resp.Body)

	assert.Equal(t, resp.StatusCode, 200)
	assert.Equal(t, resp.Header.Get("Content-Type"), "application/json")
	assert.Contains(t, string(body), "Under Armour Inc")
}
