package node

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"naivecoin/p2p"
	"naivecoin/wallet"
	"os"
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

	// initialize current node wallet
	err := initializeWallet(nodeName)
	if err != nil {
		fmt.Println("Generating wallet failed : ", err)
		return false
	}

	return true
}

func initializeWallet(nodeName string) error {
	// generate file name
	fileName := fmt.Sprintf("%s.json", nodeName)

	// check if file exists and read it
	file, err := os.OpenFile(fileName, os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		return nil
	}

	// close on exit
	defer file.Close()

	// read file
	fileData, err := ioutil.ReadAll(file)
	if err != nil {
		return nil
	}

	// if no data then create
	if len(fileData) == 0 {
		// generate new wallet and save it
		var err error = nil
		Wallet, err = wallet.GenerateAddress()
		if err != nil {
			return err
		}
		// serialize
		serializedWallet, err := serialize(Wallet)
		if err != nil {
			return err
		}
		// save to file
		file.WriteString(serializedWallet)

	} else {
		// get wallet out of file
		err = json.Unmarshal(fileData, &Wallet)
		if err != nil {
			return err
		}
	}

	return nil
}
