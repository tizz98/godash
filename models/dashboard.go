package models

type Dashboard struct {
	Model

	StockTickers []string `sql:"stock_tickers,array"`
	WidgetOrder  []string `sql:"widget_order,array"`

	Background *string `sql:"background"`
	Foreground *string `sql:"foreground"`

	TemperatureUnit *string `sql:"temperature_unit"`
	TimeUnit        *string `sql:"time_unit"`

	Location *string `sql:"location"`
}
