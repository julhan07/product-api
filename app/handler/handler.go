package handler

import (
	"api-product/command"
	"api-product/service"
	"net/http"
)

type handler struct {
	service  service.IService
	response command.Response
}

type IHandler interface {
	GetSourceProduct(w http.ResponseWriter, r *http.Request)
	GetDestinationProduct(w http.ResponseWriter, r *http.Request)
	ErrorPage(w http.ResponseWriter, r *http.Request)
}

func NewHandler(service service.IService, response command.Response) IHandler {
	return handler{service, response}
}

func (h handler) ErrorPage(w http.ResponseWriter, r *http.Request) {
	resp := command.Response{
		StatusCode: 400,
		Message:    "invalid url",
	}
	h.response.PrintJson(w, resp)
}
