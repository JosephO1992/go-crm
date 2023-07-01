// package http

// import (
// 	"net/http"

// 	"github.com/go-chi/chi/v5"
// )

// type UserService interface{}

// type Handler struct {
// 	Router  *chi.Router
// 	Service UserService
// 	Server  *http.Server
// }

// func NewHandler(service UserService) *Handler {
// 	h := Handler{
// 		Service: service,
// 	}

// 	h.Router = chi.NewRouter()

// 	h.Server = &http.Server{
// 		Addr: "0.0.0.0",
// 	}

// }
