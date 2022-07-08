package bountie

import (
	"net/http"

	"github.com/go-playground/validator/v10"
)

type Bountie interface {
	DecimalToBinary(w http.ResponseWriter, r *http.Request)
	LongestPalindromicString(w http.ResponseWriter, r *http.Request)
}

type bountie struct {
	validator *validator.Validate
}

func NewBountieClient(validator *validator.Validate) Bountie {
	return &bountie{
		validator: validator,
	}
}
