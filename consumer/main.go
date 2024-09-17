package main

import (
	"fmt"

	"github.com/christoff-linde/pih-core-go/db"
)

func main() {
	fmt.Println("Hello from consumer")

	db.InitDb("consumer")
}
