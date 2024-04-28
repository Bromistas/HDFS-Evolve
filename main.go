package main

import (
	"HDFS-Evolve/p2p"
	"fmt"
	"log"
	"time"
)

func main() {

	listenAddr := ":3000"

	tcpTransPortOpts := p2p.TCPTransportOpts{
		HandshakeFunc: p2p.NOPHandshakeFunc,
		Decoder:       p2p.GOBDecoder{},
		ListenAddr:    listenAddr,
	}

	tcpTransport := p2p.NewTcpTransport(tcpTransPortOpts)

	fsOpts := FileServerOpts{
		ListenAddr:        listenAddr,
		StorageRoot:       "3000_network",
		PathTransformFunc: CASPathTransformFunc,
		Transport:         tcpTransport,
	}

	s := NewFileServer(fsOpts)

	go func() {
		fmt.Printf("temp")
		time.Sleep(time.Second * 3)
		s.Stop()
	}()

	if err := s.Start(); err != nil {
		log.Fatal(err)
	}
}
