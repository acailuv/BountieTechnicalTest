package models

type DecimalToBinary struct {
	Decimal int `json:"decimal" validate:"numeric,required"`
}

type LongestPalindromicString struct {
	Query string `json:"query" validate:"required,min=1"`
}
