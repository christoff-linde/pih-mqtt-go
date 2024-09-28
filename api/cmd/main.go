package main

import (
	"fmt"

	"github.com/christoff-linde/pih-core-go/api/internal/server"
)

func main() {
	server := server.NewServer()

	err := server.ListenAndServe()
	if err != nil {
		panic(fmt.Sprintf("cannot start server: %s", err))
	}
}
