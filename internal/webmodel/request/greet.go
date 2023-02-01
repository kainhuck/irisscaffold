package request

type GreetReq struct {
	Name string `url:"name"`
}

type LoginReq struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}
