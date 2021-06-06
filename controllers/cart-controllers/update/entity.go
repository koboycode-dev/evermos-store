package updateCart

import (
	"time"
)

type InputUpdateCart struct {
	Id         string `validate:"required,uuid"`
	Order_Id   string `json:"order_id" validate:"required"`
	Product_Id string `json:"product_id" validate:"required"`
	Quantity   int    `json:"quantity" validate:"required"`
	Note       string `json:"note"`
	UpdatedAt  time.Time
}
