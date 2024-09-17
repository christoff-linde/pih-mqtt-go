package db

import "fmt"

func InitDb(serviceName string) {
	fmt.Println("Intitalising db...")
	fmt.Printf("...Done initialising db for service %s", serviceName)
}
