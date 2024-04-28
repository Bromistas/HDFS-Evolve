package main

import (
	"HDFS-Evolve/p2p"
	"bytes"
	"log"
	"time"
)

func makeServer(listenAddr string, nodes ...string) *FileServer {

	tcpTransPortOpts := p2p.TCPTransportOpts{
		HandshakeFunc: p2p.NOPHandshakeFunc,
		Decoder:       p2p.GOBDecoder{},
		ListenAddr:    listenAddr,
	}

	tcpTransport := p2p.NewTcpTransport(tcpTransPortOpts)

	fsOpts := FileServerOpts{
		ListenAddr:        listenAddr,
		StorageRoot:       listenAddr + "_network",
		PathTransformFunc: CASPathTransformFunc,
		Transport:         tcpTransport,
		BootstrapNodes:    nodes,
	}

	fs := NewFileServer(fsOpts)
	tcpTransport.OnPeer = fs.OnPeer

	return fs
}

func main() {
	s1 := makeServer(":3000", "")
	s2 := makeServer(":4000", ":3000")

	go func() {
		log.Fatal(s1.Start())
	}()
	time.Sleep(time.Second * 1)

	go s2.Start()
	time.Sleep(time.Second * 1)

	data := bytes.NewReader([]byte("Hello, World!"))
	s2.StoreData("myprivatedata", data)

}
