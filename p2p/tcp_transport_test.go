package p2p

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTcpTransport(t *testing.T) {
	opts := TCPTransportOpts{
		HandshakeFunc: NOPHandshakeFunc,
		Decoder:       GOBDecoder{},
		ListenAddr:    ":4000",
	}

	listenAddr := ":4000"
	tr := NewTcpTransport(opts)

	assert.Equal(t, listenAddr, tr.TCPTransportOpts.ListenAddr)

	assert.Nil(t, tr.ListenAndAccept())

	select {}
}
