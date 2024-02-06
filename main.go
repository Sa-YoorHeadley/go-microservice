package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	// "github.com/go-chi/chi/v5/middleware"
)


func main() {
	router := chi.NewRouter()
	// logger := middleware.Logger

	router.Get("/hello", basicHandler)
	// router.Use(logger)

	server := &http.Server{
		Addr: ":3000",
		Handler: router,
	}

	err := server.ListenAndServe()

	if err != nil {
		fmt.Println("Failed to listen to server", err)
	}
}

func basicHandler(writer http.ResponseWriter, req *http.Request){
	writer.Write([]byte("Hello, world!"))
}