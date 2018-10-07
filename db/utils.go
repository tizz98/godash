package db

import (
	"fmt"

	"github.com/go-pg/pg"
	"github.com/spf13/viper"
)

func Connect() *pg.DB {
	return pg.Connect(&pg.Options{
		Addr:     fmt.Sprintf("%s:%s", viper.GetString("Database.Host"), viper.GetString("Database.Port")),
		User:     viper.GetString("Database.User"),
		Password: viper.GetString("Database.Password"),
		Database: viper.GetString("Database.Name"),
	})
}
