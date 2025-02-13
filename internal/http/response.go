package http

import (
	"encoding/json"
	"net/http"

	"github.com/serj213/bookServiceApi/internal/domain"
)

const (
	StatusFailed = "failed"
)

type ResponseErr struct {
	Status string `json:"status"`
	Msg string `json:"msg"`
}

type BookResponse struct {
	Id int `json:"id"`
	Title string `json:"title"`
	Author string `json:"author"`
	CategoryId int `json:"category_id"`
}

type ResponseOkBody struct {
	Status string `json:"status"`
	Data domain.Book `json:"data"`
}


func ErrResponse(msg string, w http.ResponseWriter, r *http.Request, status int) {

	resp := ResponseErr{
		Status: StatusFailed,
		Msg: msg,
	}
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(&resp)
}

func ResponseOk(data any, w http.ResponseWriter) {
	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(data)
}