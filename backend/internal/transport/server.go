package transport

import (
	"context"
	"errors"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gorilla/mux"
)

type Server struct {
	Router *mux.Router
	Server *http.Server
}

func NewServer(handlerMappings func(router *mux.Router)) *Server {
	router := mux.NewRouter()

	// Map routes to handlers
	handlerMappings(router)

	server := &http.Server{
		Addr:    "0.0.0.0:8080",
		Handler: router,
	}

	return &Server{
		Router: router,
		Server: server,
	}
}

func (s *Server) Serve() error {
	go func() {
		if err := s.Server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatalf("Could not listen on %s: %v\n", s.Server.Addr, err)
		}
	}()

	log.Println("Server is running on", s.Server.Addr)

	// Graceful shutdown
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	if err := s.Server.Shutdown(ctx); err != nil {
		log.Fatalf("Server Shutdown Failed:%+v", err)
	}

	log.Println("Server exited properly")
	return nil
}
