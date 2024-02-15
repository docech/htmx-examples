package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("test")

	router := http.NewServeMux()

	router.HandleFunc("GET /test", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)	
		w.Write([]byte("OK"))
	})
	
	server := http.Server{
		Addr: ":8080",
		Handler: router,
	}

	fmt.Println("Starting server at", server.Addr)
	server.ListenAndServe()
}
