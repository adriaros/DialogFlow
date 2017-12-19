package main

import (
	"log"
	"github.com/labstack/echo"
)

func init() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
}

func main() {

	e := echo.New()

	e.GET("/", Root)
	e.POST("/post", Post)

	e.Logger.Fatal(e.Start(":8080"))
}
