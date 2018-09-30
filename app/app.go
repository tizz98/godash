package app

import (
	"context"
	"github.com/go-pg/pg"
	"github.com/tizz98/godash/db"
)

type App struct {
	Context context.Context
	Db      *pg.DB
}

func New() *App {
	return &App{
		Context: context.Background(),
		Db:      db.Connect(),
	}
}

func (a *App) String(s string) *string {
	return &s
}
