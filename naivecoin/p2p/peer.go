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

	// reading messages from other peers
	go func() {
		for {
			_, msg, err := conn.ReadMessage()
			if err != nil {
				fmt.Println("Failed to read message: ", err)
				break
			}
			if socketInfo.OnMessageReceived != nil {
				socketInfo.OnMessageReceived(string(msg))
			}
		}

		conn.Close()
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
