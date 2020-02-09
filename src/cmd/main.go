package main

import (
	"log"

	"fizz_buzz_go/db"
	"fizz_buzz_go/server"
)

func main() {
	err := db.ConnectDB()
	if err != nil {
		log.Fatal("Counldn't connect to DB")
	}
	server.StartServer()
}
