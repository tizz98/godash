package app

import (
	"github.com/go-pg/pg"
	"github.com/tizz98/godash/db"
)

type App struct {
	Context *Context
	Db      *pg.DB
}

type Context struct {
	App *App
}

func NewApp() *App {
	app := &App{Db: db.Connect()}
	app.Context = &Context{App: app}
	return app
}
