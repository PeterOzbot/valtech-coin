package p2p

import "github.com/beevik/guid"

//Peers : Current peers connected to this one.
var Peers []*SocketInfo

//Identifier : current peer identifier to identify the peer among all peer
var Identifier *PeerIdentifier = &PeerIdentifier{
	ID: guid.New().String(),
}
