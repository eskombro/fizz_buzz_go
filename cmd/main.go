package main

import (
	"log"

	"fizz_buzz_go/fizzbuzz"
	"fizz_buzz_go/server"
)

func main() {
	err := fizzbuzz.ConnectDB()
	if err != nil {
		log.Fatal("Counldn't connect to DB")
	}
	server.StartServer()
}
