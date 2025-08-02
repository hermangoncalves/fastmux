package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func main() {
	r := New()

	r.HandleFunc("GET", "/ping", func(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode(map[string]string{
			"message": "pong",
		})
	})

	r.NotFound(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode(map[string]string{
			"message": "endpoint not found",
		})
	}))

	fmt.Println("Server is running on: http:localhost:8080")
	http.ListenAndServe(":8080", r)
}
