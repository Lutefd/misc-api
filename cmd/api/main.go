package main

import (
	"challenge-api/internal/server"
	"fmt"
)

func main() {

	server := server.NewServer()

	err := server.ListenAndServe()
	if err != nil {
		fmt.Println(err)
		panic("cannot start server")
	}
}
