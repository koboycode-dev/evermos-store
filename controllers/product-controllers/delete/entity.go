package deleteProduct

type InputDeleteProduct struct {
	Id string `validate:"required,uuid"`
}
