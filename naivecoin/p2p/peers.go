package p2p

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

var peers []*SocketInfo

//GetPeers : Returns all connected peers.
func GetPeers(c *gin.Context) {
	var friendlyPeers []*PeerData
	for _, peer := range peers {
		friendlyPeers = append(friendlyPeers, &PeerData{
			Address: peer.Connection.RemoteAddr().String(),
		})
	}
	c.JSON(http.StatusOK, friendlyPeers)
}

//AddPeer : Adds peer to the list of all peers.
func AddPeer(c *gin.Context) {

	// deserialize peer from body
	var peerData *PeerData = &PeerData{}
	if err := c.BindJSON(peerData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// validate
	if peerData == nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Deserialization failed."})
		return
	}

	// initialize peer and check if peer exists id
	peer, alreadyConnected := TryInitializeCallerPeer(peerData)

	//  add if peer does not already exists
	if !alreadyConnected {

		if peer != nil {

			// add to list
			peers = append(peers, peer)

			// peer connected
			OnPeerConnected(peer)

			// return success
			c.JSON(http.StatusOK, "Peer added.")

		} else {
			c.JSON(http.StatusInternalServerError, "Peer initialization failed.")
		}
	} else {
		c.JSON(http.StatusOK, "Peer is already added.")
	}
}

//AddServerPeer : Adds to peer list but handles the 'Server' side of connection, That is when other node dials for connection.
func AddServerPeer(c *gin.Context) {
	// get caller id
	callerID := c.Request.Header.Get(IdentifierHeader)

	// initialize and check if peer does is already connected
	peer, alreadyConnected := InitializeServerPeer(c, callerID)

	//  add if peer does not already exists
	if !alreadyConnected {
		if peer != nil {

			// add to list
			peers = append(peers, peer)

			// log success
			fmt.Println("Server added.")
		} else {
			fmt.Println("Server initialization failed.")
		}
	} else {
		fmt.Println("Server is already added.")
	}
}

//doesPeerExists : return if peer is already connected
func doesPeerExists(id string) bool {
	for _, peer := range peers {
		if peer.ID == id {
			return true
		}
	}

	return false
}

//NotifyPeers : Notifies all peers with new message.
func NotifyPeers(message string, c *gin.Context) {
	for _, peer := range peers {
		peer.SendMessage(message)
	}
}
