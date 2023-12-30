package main

import (
	"netcat/Internal/App/server"
	"netcat/Internal/config"
	"fmt"
	"os"
)


func main() {
	if len(os.Args) < 2 {
		var s server.Server
		if len(os.Args) == 1 {
			s = *server.NewServer()
		} else {
			port := os.Args[1]
			s = *server.NewServer(config.SetPort(port))
		}	
		s.Start()
	} else {
		fmt.Println("[USAGE]: ./TCPChat $port")
	}
}