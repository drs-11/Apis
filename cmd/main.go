package main

import (
	"MightyLiteDB/internal/handlers"
	"MightyLiteDB/pkg/storage"
)

func main() {
	storage.DataMap = make(map[string]*storage.Data)
	handlers.HandleRequests()
}
