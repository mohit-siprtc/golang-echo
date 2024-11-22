package request

import (
	"github.com/go-playground/validator/v10"
)

var validate = validator.New()

type Request struct {
	ID     int     `json:"id" bson:"_id"`
	Name   string  `json:"name" validate:"required,min=1,max=100"`
	Gender string  `json:"gender" validate:"required"`
	Age    float64 `json:"age" validate:"required,gt=0"`
}

func (r *Request) Validate() error {
	return validate.Struct(r)
}
