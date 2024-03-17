package server

import (
	"net/http"

	service "github.com/LiaNestazi/moviesApi/pkg/service"
)

type Handler struct{
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes(){
	http.HandleFunc("/signin", h.signIn)
	http.HandleFunc("/signup", h.signUp)
	http.HandleFunc("/api", h.userIdentity)
	http.HandleFunc("/api/movies", h.moviesAction)
	http.HandleFunc("/api/actors", h.actorsActions)
}