package main

import (
	"Apis/internal/handlers"
	"Apis/pkg/storage"
)

func main() {
	db := storage.InitDB()
	handler := handlers.Handler{DB: db}
	handler.HandleRequests()
}
