package main

import (
	"filestoragesystem/p2p"
	"fmt"
	"log"
)

func OnPeer(p2p.Peer) error {
	fmt.Println("doing some logic with the peer otside of TCPTransport")
	return nil
}
func main() {

	tcpOpts := p2p.TCPTransportOpts{
		ListenAddr:    ":3000",
		HandshakeFunc: p2p.NOPHandshakeFunc,
		Decoder:       p2p.DefaultDecoder{},
		OnPeer:        OnPeer,
	}

	tr := p2p.NewTCPTransport(tcpOpts)

	go func() {
		for {
			msg := <-tr.Consume()
			fmt.Printf("%+v\n", msg)
		}
	}()

	if err := tr.ListenAndAccept(); err != nil {
		log.Fatal(err)
	}

	select {}

}
