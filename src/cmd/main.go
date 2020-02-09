package main

import (
	"log"

	"github.com/eskombro/fizz_buzz_go/src/db"
	"github.com/eskombro/fizz_buzz_go/src/server"
)

func main() {
	err := db.ConnectDB()
	if err != nil {
		log.Fatal("Counldn't connect to DB")
	}
	server.StartServer()
}
