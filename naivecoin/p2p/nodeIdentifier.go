package p2p

import (
	"github.com/beevik/guid"
)

//NodeIdentifier : Holds basic data describing node.
type NodeIdentifier struct {
	ID string `json:"id"`
}

//Identifier : current node identifier to identify the node among all nodes
var Identifier *NodeIdentifier = &NodeIdentifier{
	ID: guid.New().String(),
}

//IdentifierHeader : header name carrying the identifier
var IdentifierHeader string = "Caller-Identifier"
