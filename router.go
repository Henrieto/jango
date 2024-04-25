package jango

import "net/http"

type Router struct {
	Mux         *http.ServeMux
	prefix      string
	version     string
	middlewares []func(http.Handler) http.Handler
	RoutePaths  map[string]string
}
