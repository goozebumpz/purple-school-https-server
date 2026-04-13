package main

import (
	"net/http"
	"purple-school/configs"
	"purple-school/internal/auth"
	"purple-school/pkg/db"
)

func main() {
	config := configs.LoadConfig()
	_ = db.CreateDbConnection(config)
	router := http.NewServeMux()

	auth.NewAuthHandler(router, auth.HandlerDeps{Config: config})

	server := http.Server{
		Addr:    ":8081",
		Handler: router,
	}

	server.ListenAndServe()
}
