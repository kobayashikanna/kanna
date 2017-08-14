package main

import (
	"github.com/songgao/water"
	"net"
	log "github.com/sirupsen/logrus"
	"fmt"
)

type udpPacket struct {
	addr *net.UDPAddr
	data []byte
}

type Server struct {
	iface *water.Interface

	sentChan chan *udpPacket
	receivedChan chan *udpPacket
	ifaceFromChan chan []byte
	ifaceSendChan chan []byte
}

func New(cfg *ServerConfig) (server *Server) {
	server = new(Server)

	server.receivedChan = make(chan *udpPacket, 2048)
	server.sentChan = make(chan *udpPacket, 2048)
	server.ifaceFromChan = make(chan []byte, 2048)
	server.ifaceSendChan = make(chan []byte, 2048)

	return
}

func (s *Server) readFromConn() {

}

func (s *Server) listenAndServe(addr string, port uint32) {
	listenPort := fmt.Sprintf("%s:%d", addr, port)
	log.Debug("Server listen at ", listenPort)


	udpAddr, err := net.ResolveUDPAddr("udp", listenPort)
	if err != nil {
		log.Error("Invalid port: ", port)
		return
	}
	udpConn, err := net.ListenUDP("udp", udpAddr)
	if err != nil {
		log.Error("Failed to listen udp port %s: %s", port, err.Error())
		return
	}

	// write
	go func() {
		packet := <- s.sentChan
		log.Info("send packet")
		udpConn.WriteTo(packet.data, packet.addr)
	}()

	// read
	for {
		packet := new(udpPacket)
		buffer := make([]byte, 2048)
		n, addr, err := udpConn.ReadFromUDP(buffer)

		if err != nil {
			log.Error(err)
		}
		packet.addr = addr
		packet.data = buffer[:n]

		log.Info("recv packet")
		s.receivedChan <- packet
	}

}