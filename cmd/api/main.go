package main

import (
	"fmt"
	"rixlog/internal/controllers"
	"rixlog/internal/webserver"
)

var WebServer = webserver.New(controllers.RouteMap{
	"/":         controllers.Home(),
	"/articles": controllers.Articles(),
	"/auth":     controllers.Auth(),
})

// TODO: Welcome message + Server info message
func main() {
	fmt.Println("Starting webserver http://localhost:3333")
	if err := WebServer.ListenAndServe(); err != nil {
		panic(fmt.Sprintf("Cannot start server: %v", err))
	}
}
