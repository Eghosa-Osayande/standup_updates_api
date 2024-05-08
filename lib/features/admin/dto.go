package admin

type AdminLoginInputDto struct {
	Password string `json:"password" validate:"required"`
}

type AdminLoginOutputDto struct {
	Token string `json:"token"`
}