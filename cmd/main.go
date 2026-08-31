package main

import (
	"go/advanced-demo/configs"
	"go/advanced-demo/internal/auth"
	"go/advanced-demo/internal/verify"
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
  
  verify.NewVerifyHandler(router, verify.VerifyHandlerDeps{
    Config: &conf.SMTP,
  })

  err := server.ListenAndServe()
  if err != nil {
    log.Fatal(err.Error())
  }
}