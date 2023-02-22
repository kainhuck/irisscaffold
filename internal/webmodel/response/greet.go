package response

type GreetResp struct {
	Name string `json:"name"`
}

type LoginResp struct {
	Token string `json:"token"`
}

type JwtDemoResp struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
