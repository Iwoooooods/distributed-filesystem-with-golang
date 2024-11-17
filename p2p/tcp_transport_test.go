package p2p

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTCPTransport(t *testing.T) {
	transport := NewTCPTransport(":3000")
	assert.Equal(t, transport.listenAddress, ":3000")

	assert.NoError(t, transport.ListenAndAccept())
}
