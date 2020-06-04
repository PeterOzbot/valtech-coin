package p2p

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

var peers []*Peer

//GetPeers : Returns all connected peers.
func GetPeers(c *gin.Context) {
	c.JSON(200, peers)
}

//AddPeers : Adds peer to the list of all peers.
func AddPeers(c *gin.Context) {

	// deserialize peer from body
	var peer *Peer = &Peer{}
	if err := c.BindJSON(peer); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// validate
	if peer == nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Deserialization failed."})
		return
	}

	// add to list
	peers = append(peers, peer)
}
