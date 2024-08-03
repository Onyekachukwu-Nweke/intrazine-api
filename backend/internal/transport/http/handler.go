package http

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gorilla/mux"
)


type Handler struct {
	Router *mux.Router
	PostService PostService
	UserService UserService
	Server *http.Server
}

type Response struct {
	Message string
}

func NewHandler(postService PostService, userService UserService) *Handler  {
	h := &Handler{
		PostService: postService,
		UserService: userService,
	}
	h.Router = mux.NewRouter()
	h.mapRoutes()

	h.Server = &http.Server{
		Addr: "0.0.0.0:8080",
		Handler: h.Router,
	}

	return h
}

func (h *Handler) mapRoutes(){
	h.Router.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Pong, API is up")
	})

	// User Service Routes
	h.Router.HandleFunc("/api/v1/users/signup", (h.Signup)).Methods("POST")

	// Post Service Routes
	h.Router.HandleFunc("/api/v1/posts", (h.CreatePost)).Methods("POST")
	h.Router.HandleFunc("/api/v1/posts/{id}", (h.GetPostByID)).Methods("GET")
	h.Router.HandleFunc("/api/v1/posts", (h.GetAllPosts)).Methods("GET")
	h.Router.HandleFunc("/api/v1/posts/{id}", (h.UpdatePost)).Methods("PUT")
	h.Router.HandleFunc("/api/v1/posts/{id}", (h.DeletePost)).Methods("DELETE")

	// Comment Service Routes
}

func (h *Handler) Serve() error {
	go func ()  {
		if err := h.Server.ListenAndServe(); err != nil {
			log.Println(err.Error())
		}
	}()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel ()
	h.Server.Shutdown(ctx)

	log.Println("shutdown ")

	return nil
}