package p2p

import "net"

type Peer interface {
	RemoteAddr() net.Addr
	Close() error
}

type Transport interface {
	ListenAndAccept() error
	Consume() <-chan RPC
	Close() error
	Dial(addr string) error
}
