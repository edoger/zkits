package rest

import (
	"context"
	"net/http"
	"sync"
)

type Server struct {
	mutex  sync.Mutex
	driver *http.Server
	err    error
}

func NewServer() *Server {
	return &Server{}
}

func (s *Server) Run() error {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	if s.driver != nil {
		return nil
	}
	s.driver = &http.Server{
		Addr:    "127.0.0.1:9898",
		Handler: new(multiplexer),
	}
	go s.run()
	return nil
}

func (s *Server) run() {
	err := s.driver.ListenAndServe()
	if err != nil && err != http.ErrServerClosed {
		s.err = err
	}
}

func (s *Server) Close() error {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	if driver := s.driver; driver != nil {
		s.driver = nil
        return driver.Shutdown(context.TODO())
	}
	return nil
}
