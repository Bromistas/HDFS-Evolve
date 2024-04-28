package p2p

type HandshakeFunc func(peer Peer) error

func NOPHandshakeFunc(peer Peer) error { return nil }
