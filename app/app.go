package app

import "context"

type App struct {
	Context context.Context
}

func (a *App) String(s string) *string {
	return &s
}
