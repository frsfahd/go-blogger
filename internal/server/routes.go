package server

import (
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"log"
	"net/http"

	"github.com/frsfahd/go-blogger/internal/sqlc"
	"golang.org/x/crypto/bcrypt"
)

func (s *Server) RegisterRoutes() http.Handler {

	mux := http.NewServeMux()
	mux.HandleFunc("/health", s.healthHandler)

	mux.HandleFunc("/hello", Chain(s.HelloWorldHandler, Auth(), Logging()))
	mux.HandleFunc("POST /register", Chain(s.RegisterHandler, Logging()))
	mux.HandleFunc("POST /login", Chain(s.LoginHandler, Logging()))

	mux.HandleFunc("POST /posts", Chain(s.AddPostHandler, Auth(), Logging()))
	mux.HandleFunc("GET /posts", Chain(s.ListPostsHandler, Logging()))

	return mux
}

func (s *Server) HelloWorldHandler(w http.ResponseWriter, r *http.Request) {
	resp := make(map[string]string)
	resp["message"] = "Hello World"

	jsonResp, err := json.Marshal(resp)
	if err != nil {
		log.Fatalf("error handling JSON marshal. Err: %v", err)
	}

	_, _ = w.Write(jsonResp)
}

func (s *Server) LoginHandler(w http.ResponseWriter, r *http.Request) {
	var user User
	json.NewDecoder(r.Body).Decode(&user)

	var res Response

	existingUser, err := s.db.Query().GetUser(context.Background(), user.Email)
	if err != nil && errors.Is(err, sql.ErrNoRows) {
		res = Response{
			Message: "incorrect email",
		}
		json.NewEncoder(w).Encode(res)
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(existingUser.Password), []byte(user.Password))
	if err != nil && errors.Is(err, bcrypt.ErrMismatchedHashAndPassword) {
		res = Response{
			Message: "incorect password",
		}
	} else {
		token := signToken(existingUser)
		res = Response{
			Message: "logged in",
			Data: TokenData{
				Token: token,
			},
		}
	}

	json.NewEncoder(w).Encode(res)

}

func (s *Server) RegisterHandler(w http.ResponseWriter, r *http.Request) {
	var user User
	json.NewDecoder(r.Body).Decode(&user)

	var res Response

	bytes, _ := bcrypt.GenerateFromPassword([]byte(user.Password), 5)

	newUser, err := s.db.Query().AddUser(context.Background(), sqlc.AddUserParams{Email: user.Email, Password: string(bytes), Role: "admin"})

	if err != nil {
		res = Response{
			Message: err.Error(),
		}
	} else {
		res = Response{
			Message: "user added",
			Data:    newUser,
		}
	}

	json.NewEncoder(w).Encode(res)
}

func (s *Server) AddPostHandler(w http.ResponseWriter, r *http.Request) {
	var post Post
	json.NewDecoder(r.Body).Decode(&post)

	var res Response

	//validate input
	var category = sql.NullString{
		Valid: false,
	}
	if post.Category != "" {
		category = sql.NullString{
			Valid:  true,
			String: post.Category,
		}
	}

	newPost, err := s.db.Query().AddPost(context.Background(), sqlc.AddPostParams{Title: post.Title, Content: post.Content, Category: category, Tags: post.Tags})

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		res = Response{
			Message: err.Error(),
		}
		json.NewEncoder(w).Encode(res)
		return
	}

	res = Response{
		Message: "new post added",
		Data:    newPost,
	}

	json.NewEncoder(w).Encode(res)
}

func (s *Server) ListPostsHandler(w http.ResponseWriter, r *http.Request) {
	var res Response

	keyword := r.URL.Query().Get("search")

	var listPost []sqlc.Post
	var err error

	if keyword != "" {
		listPost, err = s.db.Query().FilterPosts(context.Background(), sql.NullString{Valid: true, String: keyword})
	} else {
		listPost, err = s.db.Query().ListPosts(context.Background())
	}

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		res = Response{
			Message: err.Error(),
		}
		json.NewEncoder(w).Encode(res)
		return
	}

	//check if slice is empty
	if listPost == nil {
		listPost = []sqlc.Post{}
	}

	res = Response{
		Message: "success",
		Data:    listPost,
	}
	json.NewEncoder(w).Encode(res)
}

func (s *Server) healthHandler(w http.ResponseWriter, r *http.Request) {
	jsonResp, err := json.Marshal(s.db.Health())

	if err != nil {
		log.Fatalf("error handling JSON marshal. Err: %v", err)
	}

	_, _ = w.Write(jsonResp)
}
