package server

import (
	"context"
	"github.com/gorilla/mux"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"time"
)

type Server struct {
	Client *http.Server
}

func NewServer(port string, router *mux.Router) *Server {
	return &Server{
		Client: &http.Server{
			Addr:         net.JoinHostPort("0.0.0.0", port),
			WriteTimeout: time.Second * 15,
			ReadTimeout:  time.Second * 15,
			IdleTimeout:  time.Second * 60,
			Handler:      router,
		},
	}
}

func (s *Server) Serve() error {
	go func() {

		if err := s.Client.ListenAndServe(); err != nil {
			log.Println(err)
		}
	}()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c

	// Create a deadline to wait for
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	s.Client.Shutdown(ctx)

	return nil
}
