package command

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Response struct {
	StatusCode int         `json:"status_code"`
	Count      int64       `json:"count,omitempty"`
	Message    string      `json:"message"`
	Pagination *Pagination `json:"pagination,omitempty"`
	Data       interface{} `json:"data,omitempty"`
}

func NewReponse() Response {
	return Response{}
}

func (r *Response) PrintJson(w http.ResponseWriter, res interface{}) {
	resp, err := json.Marshal(res)
	if err != nil {
		fmt.Println("err", err.Error())
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(resp)
}

func (r *Response) Success(pagination *Pagination, data interface{}, count int64) Response {
	r.StatusCode = http.StatusOK
	r.Message = http.StatusText(r.StatusCode)
	r.Data = data
	r.Count = count
	r.Pagination = pagination
	return *r
}

func (r *Response) Created(data interface{}) Response {
	r.StatusCode = http.StatusCreated
	r.Message = http.StatusText(r.StatusCode)
	r.Data = data
	return *r
}

func (r *Response) BadRequest(msg string) Response {
	r.StatusCode = http.StatusBadRequest
	r.Message = msg
	return *r
}
