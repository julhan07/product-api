package handler

import (
	"api-product/command"
	"net/http"
)

func (h handler) GetSourceProduct(w http.ResponseWriter, r *http.Request) {
	pagination := command.GetPagination(r)
	res := h.service.GetSourceProduct(w, &pagination)
	h.response.PrintJson(w, res)
}

func (h handler) GetDestinationProduct(w http.ResponseWriter, r *http.Request) {

	pagination := command.GetPagination(r)
	res := h.service.GetDestinationProduct(w, &pagination)
	h.response.PrintJson(w, res)
}
