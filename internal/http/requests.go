package http

import (
	"fmt"

	"github.com/go-playground/validator/v10"
)

type BookRequest struct {
	Title string `json:"title" validate:"required"`
	Author string `json:"author" validate:"required"`
	CategoryId int `json:"category_id" validate:"required"`
}


func (b BookRequest) Validate() error {
	validate := validator.New(validator.WithRequiredStructEnabled())

	err := validate.Struct(b)

	if err != nil {
		return fmt.Errorf("validate error: %w", err)
	}
	
	return nil
}