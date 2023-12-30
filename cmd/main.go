package main

import (
	"netcat/Internal/App/server"
	"netcat/Internal/config"
	"fmt"
	"os"
)


func main() {
	l := os.Args[1:]
	if len(l) == 1 {
		port := l[0]
		s := server.NewServer(config.SetPort(port))
		s.Start()
	} else {
		fmt.Println("[USAGE]: ./TCPChat $port")
	}
}