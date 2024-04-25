package jango

import (
	"fmt"
	"net/http"
	"strings"
)

type Route struct {
	Path    string
	Handler http.HandlerFunc
	method  string
	name    string
}

func Path(path string, handler http.HandlerFunc, method string) *Route {
	return &Route{
		Path:    path,
		Handler: handler,
		method:  method,
	}
}

func (rte *Route) Name(name string) *Route {
	rte.name = name
	return rte
}

func (rte *Route) Handle(rter *Router) {
	_path := rter.prefix + rte.Path
	rter.RoutePaths[rte.name] = _path
	if rte.method != "" {
		_path = fmt.Sprintf("%v %v", rte.method, rte.Path)
	}
	rter.Mux.HandleFunc(_path, rte.Handler)
}

type Router struct {
	Mux         *http.ServeMux
	prefix      string
	version     string
	middlewares []func(http.Handler) http.Handler
	RoutePaths  map[string]string
}

func NewRouter() *Router {
	return &Router{
		Mux:         http.NewServeMux(),
		middlewares: []func(http.Handler) http.Handler{},
		RoutePaths:  map[string]string{},
	}
}

func (rter *Router) VersionString(version_string string) *Router {
	version_string = strings.Replace(version_string, "/", "", -1)
	rter.version = version_string
	return rter
}

func (rter *Router) Version(version int) *Router {
	switch rter.version {
	case "":
		rter.version = fmt.Sprintf("/api/%v", version)
	default:
		rter.version = fmt.Sprintf("%v/%v", rter.version, version)
	}
	return rter
}

func (rter *Router) Prefix(prefix string) *Router {
	rter.prefix = rter.version + prefix
	return rter
}

func (rter *Router) SubPrefix(sub_prefix string) *Router {
	rter.prefix = rter.prefix + sub_prefix
	return rter
}

func (rter *Router) Middlewares(middlewares ...func(http.Handler) http.Handler) {
	rter.middlewares = append(rter.middlewares, middlewares...)
}
