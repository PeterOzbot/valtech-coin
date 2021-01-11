package node

import (
	"fmt"
	"io/ioutil"
	"naivecoin/p2p"
	"naivecoin/wallet"
)

//InitializeNode : Hooks nethods from p2p to handle comunication between peers.
func InitializeNode(nodeName string) bool {

	// Handles message when its received from peers.
	p2p.OnMessageReceived = func(requestMessage string, socketInfo *p2p.SocketInfo) {
		fmt.Println("OnMessageReceived request: ", requestMessage)
		responseMessage, err := HandleMessage(requestMessage)
		if err != nil {
			fmt.Println("send message : ", err)
			return
		}

		// send messages to peer if there are any
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

	// generate new wallet and save it
	var err error = nil
	Wallet, err = wallet.GenerateAddress()
	if err != nil {
		fmt.Println("Generating wallet failed : ", err)
		return false
	}
	serializedWallet, err := serialize(Wallet)
	if err != nil {
		fmt.Println("Serializing wallet failed : ", err)
		return false
	}
	fileName := fmt.Sprintf("%s.json", nodeName)
	err = ioutil.WriteFile(fileName, []byte(serializedWallet), 0664)
	if err != nil {
		fmt.Println("Serializing wallet failed : ", err)
		return false
	}

	return true
}
