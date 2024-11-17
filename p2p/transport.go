package p2p

// Peer is a peer in the network that can send and receive messages
type Peer interface{}

// Transport is a transport layer for the network that handles the communication between peers
type Transport interface {
	ListenAndAccept() error
}
