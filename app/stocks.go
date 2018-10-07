package app

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

type Stock struct {
	Status        string  `json:",omitempty"`
	Message       string  `json:",omitempty"`
	Name          string  `json:",omitempty"`
	Symbol        string  `json:",omitempty"`
	Exchange      string  `json:",omitempty"`
	LastPrice     float64 `json:",omitempty"`
	ChangePercent float64 `json:",omitempty"`
}

func (ctx *Context) GetStockInfo(symbol string) (*Stock, error) {
	if err := ctx.StockInfoSem.Acquire(context.Background(), 1); err != nil {
		return nil, err
	}
	defer ctx.StockInfoSem.Release(1)

	resp, err := http.Get("http://dev.markitondemand.com/Api/v2/Quote/json?symbol=" + symbol)
	if err != nil {
		return nil, err
	}

	var stock Stock
	if err := json.NewDecoder(resp.Body).Decode(&stock); err != nil {
		return nil, err
	}

	if stock.Status != "SUCCESS" {
		return nil, fmt.Errorf("error retreiving stock info: %s", stock.Message)
	}

	return &stock, nil
}

func (ctx *Context) SearchStocks(query string) ([]*Stock, error) {
	if err := ctx.StockSearchSem.Acquire(context.Background(), 1); err != nil {
		return nil, err
	}
	defer ctx.StockSearchSem.Release(1)

	resp, err := http.Get("http://dev.markitondemand.com/MODApis/Api/v2/Lookup/json?input=" + query)
	if err != nil {
		return nil, err
	}

	stocks := make([]*Stock, 0)
	if err := json.NewDecoder(resp.Body).Decode(&stocks); err != nil {
		return nil, err
	}

	return stocks, nil
}
