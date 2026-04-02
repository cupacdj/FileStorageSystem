package p2p

// HandshackeFunc is
type HandshakeFunc func(Peer) error

func NOPHandshakeFunc(Peer) error {
	return nil
}
