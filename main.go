package main

import (
	"HDFS-Evolve/p2p"
	"log"
)

func main() {
	opts := p2p.TCPTransportOpts{
		HandshakeFunc: p2p.NOPHandshakeFunc,
		Decoder:       p2p.DefaultDecoder{},
		ListenAddr:    ":3000",
	}

	tr := p2p.NewTcpTransport(opts)

	go func() {
		for {
			msg := <-tr.Consume()
			log.Printf("Received message: %+v\n", msg)
		}
	}()

	if err := tr.ListenAndAccept(); err != nil {
		log.Fatal(err)
	}

	select {}
}
