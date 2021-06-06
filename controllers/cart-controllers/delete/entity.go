package deleteCart

type InputDeleteCart struct {
	Id string `validate:"required,uuid"`
}
