package main

import (
	"fmt"
	"net/http"
)


func main() {
	server := &http.Server{
		Addr: ":3000",
		Handler: http.HandlerFunc(basicHandler),
	}

	err := server.ListenAndServe()

	if err != nil {
		fmt.Println("Failed to listen to server", err)
	}
}

func basicHandler(writer http.ResponseWriter, req *http.Request){
	writer.Write([]byte("Hello, world!"))
}