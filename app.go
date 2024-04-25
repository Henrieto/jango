package jango

import "net/http"

type App struct {
	Label       string
	Routes      []Route
	Middlewares []func(http.Handler) http.Handler
}

func (app *App) Ready(router *Router) {
	for _, route := range app.Routes {
		route.Handle(router)
	}
	router.middlewares = append(router.middlewares, app.Middlewares...)
}
