package app

import (
	"github.com/go-pg/pg"
	"github.com/sirupsen/logrus"
	"github.com/tizz98/godash/db"
	"golang.org/x/sync/semaphore"
)

type App struct {
	Context *Context
	Db      *pg.DB
}

type Context struct {
	App    *App
	Logger *logrus.Logger

	StockInfoSem   *semaphore.Weighted
	StockSearchSem *semaphore.Weighted
}

func NewApp() *App {
	app := &App{Db: db.Connect()}
	app.Context = &Context{
		App:            app,
		Logger:         logrus.New(),
		StockInfoSem:   semaphore.NewWeighted(4),
		StockSearchSem: semaphore.NewWeighted(4),
	}
	return app
}
