package registerAuth

type InputRegister struct {
	Username string `json:"username" validate:"required,username"`
	Password string `json:"password" validate:"required,gte=8"`
}
