package main

import (
	"os"
	
	"github.com/linyuy/editor/service"
	flag "github.com/spf13/pflag"
	
	// "log"
)

const (
	// PORT : default port
	PORT string = "8080"
)

func main() {
	var port string

	pPort := flag.StringP("port", "p", PORT, "PORT for httpd listening")
	flag.Parse()
	if len(*pPort) != 0 {
		port = *pPort
	}

	if envPort := os.Getenv("PORT"); len(envPort) != 0 {
		// log.Fatal("$PORT must be set")
		port = envPort
	}

	server := service.NewServer()
	server.Run(":" + port)
}
