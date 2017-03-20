package server

import "time"

type Server struct {
	router  *Router
	timeout time.Duration
}

func NewServer() *Server {

}
