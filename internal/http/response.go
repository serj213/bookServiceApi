package http

import (
	"encoding/json"
	"net/http"

	"github.com/serj213/bookServiceApi/internal/domain"
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

	var response ResponseErr

	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(&response)
}

func ResponseOk(data any, w http.ResponseWriter) {
	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(data)
}