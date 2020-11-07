package p2p

import (
	"github.com/beevik/guid"
)

//PeerIdentifier : Holds basic data identifying peer.
type PeerIdentifier struct {
	ID string `json:"id"`
}

//Identifier : current peer identifier to identify the peer among all peer
var Identifier *PeerIdentifier = &PeerIdentifier{
	ID: guid.New().String(),
}

//IdentifierHeader : header name carrying the identifier
var IdentifierHeader string = "Caller-Identifier"
