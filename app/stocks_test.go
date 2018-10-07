package app

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestContext_GetStockInfo(t *testing.T) {
	app := NewApp()
	stock, err := app.Context.GetStockInfo("UA")
	assert.NoError(t, err)
	assert.Equal(t, "Under Armour Inc", stock.Name)
}

func TestContext_SearchStocks(t *testing.T) {
	app := NewApp()
	stocks, err := app.Context.SearchStocks("under")
	assert.NoError(t, err)
	assert.NotNil(t, stocks[0].Symbol)
}
