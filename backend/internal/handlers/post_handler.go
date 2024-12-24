package handlers

import (
	"github.com/Onyekachukwu-Nweke/piko-blog/backend/internal/interfaces"
	"net/http"
)

type PostHandler struct {
	Service interfaces.PostStore
}

func NewPostHandler(service interfaces.PostStore) *PostHandler {
	return &PostHandler{Service: service}
}

func (handler *PostHandler) CreatePost(w http.ResponseWriter, r *http.Request) {

}
