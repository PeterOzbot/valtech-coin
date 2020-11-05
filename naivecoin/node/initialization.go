package node

import (
	"fmt"
	"naivecoin/p2p"
)

//InitializeNode : Hooks nethods from p2p to handle comunication between peers.
func InitializeNode() {

	// Handles message when its received from peers.
	p2p.OnMessageReceived = func(requestMessage string, socketInfo *p2p.SocketInfo) {
		fmt.Println("OnMessageReceived request: ", requestMessage)
		responseMessage, err := HandleMessage(requestMessage)
		if err != nil {
			fmt.Println("send message : ", err)
			return
		}

		if len(responseMessage) > 0 {
			socketInfo.SendMessage(responseMessage)
		}
		fmt.Println("OnMessageReceived response: ", responseMessage)
	}

	//OnPeerConnected : Executes actions that are required when peer gets conencted to this node.
	p2p.OnPeerConnected = func(socketInfo *p2p.SocketInfo) {
		connectionMessage, err := OnPeerConnected()
		if err != nil {
			fmt.Println("send message : ", err)
			return
		}

		socketInfo.SendMessage(connectionMessage)
	}
}
