package updateProduct

import (
	"time"
)

type InputUpdateProduct struct {
	Id        string `json:"id" validate:"required,uuid"`
	Name      string `json:"name" validate:"required"`
	Price     int    `json:"price"`
	Quantity  int    `json:"quantity"`
	UpdatedAt time.Time `json:"updated_at"`
}
