package createCart

type InputCreateCart struct {
	User_Id    string `json:"user_id" validate:"required"`
	Order_Id   string `json:"order_id" validate:"required"`
	Product_Id string `json:"product_id" validate:"required"`
	Quantity   int    `json:"quantity" validate:"required"`
	Note       string `json:"note"`
}
