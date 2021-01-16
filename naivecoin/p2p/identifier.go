package p2p

//PeerIdentifier : Holds basic data identifying peer.
type PeerIdentifier struct {
	ID string `json:"id"`
}

//IdentifierHeader : header name carrying the identifier
var IdentifierHeader string = "Caller-Identifier"
