package main


import (
	"log"
	"net/http"

	"GoDataOpsAPI/internal/handler"
	"GoDataOpsAPI/internal/store"
)


func main() {
	store := store.NewInMemoryStore()
	handler := handler.NewItemHandler(store)

	http.HandleFunc("/items", handler.ItemsHandler)
	http.HandleFunc("/items/", handler.ItemHandler)

	log.Println("Starting server on :8080...")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("Server failed to start: ", err)
	}
}
