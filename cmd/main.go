package main

import (
	"net/http"
	"purple-school/configs"
	"purple-school/internal/auth"
	"purple-school/internal/hello"
	"purple-school/internal/random"
)

func main() {
	_ = configs.LoadConfig()
	router := http.NewServeMux()

	hello.NewHelloHandler(router)
	random.NewRandomHandler(router)
	auth.NewAuthHandler(router)

	server := http.Server{
		Addr:    ":8081",
		Handler: router,
	}

	server.ListenAndServe()
}
