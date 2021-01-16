package p2p

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

//GetPeers : Returns all connected peers.
func GetPeers(c *gin.Context) {
	peerLock.Lock()
	defer peerLock.Unlock()

	var friendlyPeers = make([]*PeerData, 0, len(Peers))
	for _, peer := range Peers {
		friendlyPeers = append(friendlyPeers, &PeerData{
			Address: peer.Connection.RemoteAddr().String(),
			ID:      peer.ID,
		})
	}
	c.JSON(http.StatusOK, friendlyPeers)
}

//AddPeer : Adds peer to the list of all peers.
func AddPeer(c *gin.Context) {
	peerLock.Lock()
	defer peerLock.Unlock()

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
			Peers = append(Peers, peer)

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
	peerLock.Lock()
	defer peerLock.Unlock()

	// get caller id
	callerID := c.Request.Header.Get(IdentifierHeader)

	// initialize and check if peer does is already connected
	peer, alreadyConnected := InitializeServerPeer(c, callerID)

	//  add if peer does not already exists
	if !alreadyConnected {
		if peer != nil {

			// add to list
			Peers = append(Peers, peer)

			// peer connected
			OnPeerConnected(peer)

			// log success
			fmt.Println("Server added.")
		} else {
			fmt.Println("Server initialization failed.")
		}
	} else {
		fmt.Println("Server is already added.")
	}
}

//RemovePeer : Removes peer from current collection.
func RemovePeer(callerID string) {
	peerLock.Lock()
	defer peerLock.Unlock()

	// find peer to remove
	var peerIndex = 0
	for index, peer := range Peers {
		if peer.ID == callerID {
			peerIndex = index
		}
	}

	// remove
	copy(Peers[peerIndex:], Peers[:peerIndex])
	Peers = Peers[:len(Peers)-1]
}

//NotifyPeers : Notifies all peers with new message.
func NotifyPeers(message string) {
	for _, peer := range Peers {
		peer.SendMessage(message)
	}
}
