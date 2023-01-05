package app

import (
	"api-product/app/handler"
	"api-product/middleware"

	"github.com/gorilla/mux"
)

func router(h handler.IHandler) *mux.Router {
	r := mux.NewRouter()
	r.Use(middleware.LoggingMiddleware)
	r.HandleFunc("/source-product", h.GetSourceProduct).Methods("GET")
	r.HandleFunc("/destination-product", h.GetDestinationProduct).Methods("GET")
	r.HandleFunc("/", h.ErrorPage)
	return r
}
