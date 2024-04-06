package main

import (
	httpHandler "project/handler"

	"github.com/labstack/echo/v4"
)

func main() {
	handler := httpHandler.InitNews()
	echoServer := echo.New()

	// Register the handler
	echoServer.GET("/articles", handler.FetchNews)

	// Start the server
	echoServer.Start(":9090")
}