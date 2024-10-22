package server

type User struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type Post struct {
	Title    string   `json:"title"`
	Content  string   `json:"content"`
	Category string   `json:"category,omitempty"`
	Tags     []string `json:"tags,omitempty"`
}

type Response struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

type TokenData struct {
	Token string `json:"token"`
}
