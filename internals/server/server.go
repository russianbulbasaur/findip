package server

import (
	"findip/internals/dns_message"
	"fmt"
	"log"
	"net"
)

type Server struct {
	address string
	port    int
}

func NewServer(address string, port int) *Server {
	return &Server{
		address,
		port,
	}
}

func (server *Server) Run() {
	udpAddr, err := net.ResolveUDPAddr("udp", fmt.Sprintf("%s:%d", server.address, server.port))
	serverSocket, err := net.ListenUDP("udp", udpAddr)
	if err != nil {
		log.Fatalln(err)
	}
	var buffer []byte = make([]byte, 1024)
	for {
		_, client, err := serverSocket.ReadFromUDP(buffer)
		if err != nil {
			log.Fatalln(err)
		}
		message := dns_message.NewDNSMessage()
		response := message.Serialize()
		_, err = serverSocket.WriteToUDP(response, client)
	}
}