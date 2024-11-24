package main

import (
	"findip/internals/server"
)

func main() {
	serverNode := server.NewServer("127.0.0.1", 8000)
	serverNode.Run()
}
