package auth

import (
	"fmt"
	"go/advanced-demo/configs"
	"net/http"
)

type AuthHandler struct {
	Config *configs.AuthConfig
}

type AuthHandlerDeps struct {
	Config *configs.AuthConfig
}

func NewAuthHandler(router *http.ServeMux, deps AuthHandlerDeps) {
	handler := AuthHandler{
		Config: deps.Config,
	}
	router.HandleFunc("POST /auth/login", handler.Login())
	router.HandleFunc("POST /auth/register", handler.Register())
}

func (handler *AuthHandler) Login() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Login: %s", handler.Config.Secret)
	}
}

func (handler *AuthHandler) Register() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Register")
	}
}