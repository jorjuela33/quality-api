package server

import (
	"net/http"
	"time"

	"github.com/codegangsta/negroni"
	"github.com/gorilla/context"
	"gopkg.in/tylerb/graceful.v1"
)

// JSON body limit is set to 5MB
const BodyLimitBytes uint32 = 1048576 * 5

type Server struct {
	negroni        *negroni.Negroni
	router         *Router
	gracefulServer *graceful.Server
	timeout        time.Duration
}

// Options for running the server
type Options struct {
	Timeout         time.Duration
	ShutdownHandler func()
}

func NewServer() *Server {
	negroni := negroni.Classic()
	server := &Server{negroni, nil, nil, 0}
	return server
}

func (server *Server) Serve(responseWriter http.ResponseWriter, request *http.Request) *Server {
	server.negroni.ServeHTTP(responseWriter, request)
	return server
}

func (server *Server) Stop() {
	server.gracefulServer.Stop(server.timeout)
}

func (server *Server) Run(address string, options Options) *Server {
	server.gracefulServer = &graceful.Server{Timeout: options.Timeout, Server: &http.Server{Addr: address, Handler: server.negroni}, ShutdownInitiated: options.ShutdownHandler}
	server.gracefulServer.ListenAndServe()
	return server
}

func (server *Server) UserRouter(router *Router) *Server {
	server.negroni.UseHandler(context.ClearHandler(router))
	return server
}
