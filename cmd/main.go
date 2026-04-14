package main

import (
	"fmt"
	"net/http"
	"purple-school/configs"
	"purple-school/internal/auth"
	"purple-school/internal/link"
	"purple-school/pkg/db"
)

func main() {
	config := configs.LoadConfig()
	database := db.NewDb(config)
	router := http.NewServeMux()

	auth.NewAuthHandler(router, auth.HandlerDeps{
		Config: config,
	})
	link.NewLinksHandler(router, link.HandlerDeps{
		Config:     config,
		Repository: link.NewRepository(database),
	})

	server := http.Server{
		Addr:    ":8081",
		Handler: router,
	}

	fmt.Println("Server start work on port 8081")
	server.ListenAndServe()
}
