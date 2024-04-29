package jango

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

var (
	// server timeouts
	WRITETIMEOUT = time.Second * 30

	READTIMEOUT = time.Second * 30

	IDLETIMEOUT = time.Second * 30

	// server tls credentials
	SERVER_CERTIFICATE = ""

	SERVER_CERTIFICATE_KEY = ""

	//other server credentials
	SERVER_ADDRESS = ""
)

type Server struct {
	router  *Router
	use_tls bool
	*http.Server
}

func NewServer(router *Router, use_tls bool) *Server {
	return &Server{
		router,
		use_tls,
		&http.Server{
			ReadTimeout:  READTIMEOUT,
			WriteTimeout: WRITETIMEOUT,
			IdleTimeout:  IDLETIMEOUT,
		},
	}
}

func (server *Server) initialize_middlewares() {
	for _, middlewares := range server.router.middlewares {
		server.Handler = middlewares(server.Handler)
	}
}

func (server *Server) Use(middlewares ...func(http.Handler) http.Handler) {
	for _, middleware := range middlewares {
		server.Handler = middleware(server.Handler)
	}
}

func (server *Server) Start() (err error) {

	if SERVER_ADDRESS == "" {
		SERVER_ADDRESS = ":8080"
	}
	server.Addr = SERVER_ADDRESS

	start_server_message := "server started : " + SERVER_ADDRESS

	log.Println(start_server_message)

	if !server.use_tls {
		err = server.ListenAndServe()
	} else {
		if SERVER_CERTIFICATE == "" || SERVER_CERTIFICATE_KEY == "" {
			panic(" server certificate and key is needed ")
		}
		err = server.ListenAndServeTLS(SERVER_CERTIFICATE, SERVER_CERTIFICATE_KEY)
	}
	return
}

func (server *Server) applyOptions(options ...OptionFunc) {
	for _, optionFunc := range options {
		optionFunc(server)
	}
}

func (server *Server) Listen(options ...OptionFunc) error {

	server.Handler = server.router.Mux

	server.initialize_middlewares()

	server.applyOptions(options...)

	var err error

	go func() {

		err = server.Start()

	}()

	if err != nil {

		return err

	}

	ch := make(chan os.Signal, 4)

	signal.Notify(ch, os.Interrupt)

	signal.Notify(ch, syscall.SIGTERM)

	<-ch

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*30)

	defer cancel()

	shutdown_server_message := "server is shutting down"

	log.Println(shutdown_server_message)

	server.Shutdown(ctx)

	return nil
}
