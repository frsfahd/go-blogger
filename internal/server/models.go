package server

type User struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type Response struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

type TokenData struct {
	Token string `json:"token"`
}
