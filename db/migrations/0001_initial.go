package migrations

import (
	"fmt"
	"github.com/go-pg/migrations"
)

func init() {
	migrations.Register(func(db migrations.DB) error {
		fmt.Println("creating dashboards table...")
		_, err := db.Exec(`
			CREATE TABLE dashboards (
				id char(32) PRIMARY KEY,
				stock_tickers text[],
				widget_order text[],
				background text,
				foreground text,
				temperature_unit text,
				time_unit text,
				location text
			)
		`)
		return err
	}, func(db migrations.DB) error {
		fmt.Println("dropping table dashboards...")
		_, err := db.Exec(`DROP TABLE dashboards`)
		return err
	})
}
