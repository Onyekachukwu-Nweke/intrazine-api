package http

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)


type Handler struct {
	Router *mux.Router
	PostService PostService
	Server *http.Server
}

func NewHandler(postService PostService) *Handler  {
	h := &Handler{
		PostService: postService,
	}
	h.Router = mux.NewRouter()

	h.Server = &http.Server{
		Addr: "0.0.0.0:8080",
		Handler: h.Router,
	}

	return h
}

func (h *Handler) mapRoutes() {
	h.Router.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Pong, API is up")
	})
}