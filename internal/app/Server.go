package app

import (
	"log"
	"net/http"
)

type Server struct {
	httpServer *http.Server
}

func (s *Server) Run(port string) error {
	s.httpServer = &http.Server{
		Addr: ":" + port,
		//Handler: handler,
	}
	log.Println("Server start")
	return s.httpServer.ListenAndServe()
}
