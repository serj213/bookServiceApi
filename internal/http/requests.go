package http

import (
	"fmt"

	"github.com/go-playground/validator/v10"
)

type CreateBookReq struct {
	Title string `json:"title" validate:"required"`
	Author string `json:"author" validate:"required"`
	CategoryId int `json:"category_id" validate:"required"`
}

type BookRequest struct {
	Id int `json:"id" validate:"required"`
	Title string `json:"title,omitempty" validate:"required"`
	Author string `json:"author,omitempty"`
	CategoryId int `json:"category_id,omitempty"`
}


func (b CreateBookReq) Validate() error {
	validate := validator.New(validator.WithRequiredStructEnabled())

	err := validate.Struct(b)

	if err != nil {
		return fmt.Errorf("validate error: %w", err)
	}
	
	return nil
}


func (b BookRequest) Validate() error {
	validate := validator.New(validator.WithRequiredStructEnabled())

	err := validate.Struct(b)
	if err != nil {
		return fmt.Errorf("validate error: %w", err)
	}
	
	return nil
}