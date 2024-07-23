package controllers

import (
	v "github.com/go-playground/validator/v10"
)

type Controller struct {
	Validator *v.Validate
}

func NewController(validator *v.Validate) *Controller {
	return &Controller{
		Validator: validator,
	}
}
