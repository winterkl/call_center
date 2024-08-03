package http_server

import (
	"net/http"
	"time"
)

type Server struct {
	httpServer *http.Server
}

func New(host, port string, handler http.Handler) *Server {
	return &Server{
		httpServer: &http.Server{
			Addr:           host + ":" + port,
			Handler:        handler,
			MaxHeaderBytes: 1 << 20,
			ReadTimeout:    10 * time.Second,
			WriteTimeout:   10 * time.Second,
		},
	}
}

// Start - Запуск сервера
func (s *Server) Start() error {
	return s.httpServer.ListenAndServe()
}
