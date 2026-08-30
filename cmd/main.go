package main

import (
	"go/advanced-demo/configs"
	"go/advanced-demo/internal/auth"
	"log"
	"net/http"
)


func main() {
  conf := configs.LoadConfig()
  router := http.NewServeMux()
	server := http.Server{
      Addr: ":8081",
      Handler: router,
    }

  auth.NewAuthHandler(router, auth.AuthHandlerDeps{
    Config: &conf.Auth,
  })

  err := server.ListenAndServe()
  if err != nil {
    log.Fatal(err.Error())
  }
}