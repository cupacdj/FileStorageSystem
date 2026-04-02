package p2p

// Peer is an interface that represents the remote node.
type Peer interface {
	Close() error
}

// Transport is anything that handles the communication
// between nodes in the network. This can be of the form
// of TCP, UDP, websocket, or any other protocol that can be used to send messages between nodes.
type Transport interface {
	ListenAndAccept() error
	Consume() <-chan RPC
}
