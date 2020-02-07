package server

import (
	"log"
	"net/http"

	"github.com/labstack/echo/v4"

	fb "fizz_buzz_go/fizzbuzz"
)

func StartServer() {
	e := echo.New()
	e.POST("/", func(c echo.Context) error {
		log.Println("Request:  POST request recieved")
		resp := fb.Fizzbuzz(
			3,
			5,
			15,
			"str1",
			"str2",
		)
		log.Printf("Response: %s", resp)
		return c.String(http.StatusOK, resp)
	})
	e.Logger.Fatal(e.Start(":1323"))
}
