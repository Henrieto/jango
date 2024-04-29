package jango

import (
	"net/http"

	"github.com/rs/cors"
)

var (
	// cors credentials
	AllowedOrigins = []string{}

	AllowedHeaders = []string{}

	AllowedMethods = []string{}

	AllowCredentials = true
)

type OptionFunc func(*Server)

func UseCors(server *Server) {
	options := cors.Options{
		AllowedOrigins:   AllowedOrigins,
		AllowedHeaders:   AllowedHeaders,
		AllowedMethods:   AllowedMethods,
		AllowCredentials: AllowCredentials,
	}
	server.Handler = cors.New(options).Handler(server.Handler)
}

func UseHealthCheck(handlers ...http.HandlerFunc) OptionFunc {
	return func(server *Server) {
		if len(handlers) == 0 {
			server.router.Mux.HandleFunc("/server/health", HealthCheck)
		} else {
			server.router.Mux.HandleFunc("/server/health", handlers[0])
		}
	}
}
