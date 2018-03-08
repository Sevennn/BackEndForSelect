package main

import (
	"os"
	"go-select/routes"
	flag "github.com/spf13/pflag"
)

const (
	PORT string = "3000"
)

func main() {
	port := os.Getenv("PORT")
	if len(port) == 0 {
		port = PORT
	}

	pPort := flag.StringP("port", "p", PORT, "PORT for httpd listening")
	flag.Parse()
	if len(*pPort) != 0 {
		port = *pPort
	}

	server := routes.NewServer()
	server.Run(":" + port)
}