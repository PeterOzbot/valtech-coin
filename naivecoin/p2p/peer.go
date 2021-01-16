package p2p

import (
	"fmt"

	"github.com/gorilla/websocket"
)

//OnMessageReceived : Handles message when its received from peers.
var OnMessageReceived func(requestMessage string, socketInfo *SocketInfo)

//OnPeerConnected : Executes actions that are required when peer gets conencted to this node.
var OnPeerConnected func(socketInfo *SocketInfo)

//InitializePeer : Common initialization logic for both server/caller peer,
func InitializePeer(callerID string, conn *websocket.Conn) (*SocketInfo, bool) {

	// check if already connected
	if doesPeerExists(callerID) {
		conn.Close()
		return nil, true
	}

	// create new info
	socketInfo := &SocketInfo{
		Connection: conn,
		ID:         callerID,
	}

	// hook send message
	socketInfo.SendMessage = func(message string) {
		err := conn.WriteMessage(websocket.TextMessage, []byte(message))
		if err != nil {
			fmt.Println("send message : ", err)
		}
	}

	// hook on message received
	socketInfo.OnMessageReceived = func(requestMessage string) {
		OnMessageReceived(requestMessage, socketInfo)
	}

	go func() {

		// wait for incoming message
		for {

			// reading messages
			_, msg, err := conn.ReadMessage()
			if err != nil {
				fmt.Println("Failed to read message or peer disconnected: ", err)
				break
			}

			// notify message received
			if socketInfo.OnMessageReceived != nil {
				socketInfo.OnMessageReceived(string(msg))
			}
		}

		// remove peer
		RemovePeer(callerID)

		// close connection
		err := conn.Close()
		if err != nil {
			fmt.Println("Failed closing connection: ", err)
		}
	}()

	return socketInfo, false
}

//doesPeerExists : return if peer is already connected
func doesPeerExists(callerID string) bool {
	for _, peer := range Peers {
		if peer.ID == callerID {
			return true
		}
	}

	return false
}
