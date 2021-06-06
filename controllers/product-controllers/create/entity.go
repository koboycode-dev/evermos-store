package createProduct

type InputCreateProduct struct {
	Id_Product string `json:"id_product" validate:"required"`
	Name       string `json:"name" validate:"required" unique:"name"`
	Price      int    `json:"price"`
	Quantity   int    `json:"quantity"`
}
