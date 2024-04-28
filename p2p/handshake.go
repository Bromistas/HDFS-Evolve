package p2p

type HandshakeFunc func(any) error

func NOPHandshakeFunc(any) error { return nil }
