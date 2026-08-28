package main

import (
	"log"
	"math/rand/v2"
	"net/http"
	"strconv"
)

func randomHandler(w http.ResponseWriter, r *http.Request) {
	randomNum := rand.IntN(6)

	w.Write([]byte(strconv.Itoa(randomNum)))
}

func main() {
	server := http.Server{
      Addr: ":8081",
      Handler: nil,
    }

    http.HandleFunc("/random", randomHandler)

    err := server.ListenAndServe()
    if err != nil {
      log.Fatal("Error")
    }
}