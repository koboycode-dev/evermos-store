package loginAuth

type InputLogin struct {
	Username string `json:"username" validate:"required,username"`
	Password string `json:"password" validate:"required"`
}
