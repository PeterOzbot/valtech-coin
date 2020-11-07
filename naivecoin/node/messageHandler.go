package node

import (
	"encoding/json"
	"errors"
	"naivecoin/blockchain"
)

//OnPeerConnected : Prepares message to send when new peer gets connected.
func OnPeerConnected() (string, error) {
	message := &MessageData{
		Type: QueryLatestBlockType,
	}

	return serialize(message)
}

//OnMinedBlock : Prepares message to send when new block is mined.
func OnMinedBlock(newBlock *blockchain.Block) (string, error) {
	// serialize new block
	serializedData, err := serialize(newBlock)
	if err != nil {
		return "", err
	}

	// create message
	message := &MessageData{
		Type: ResponseLatestBlockType,
		Data: serializedData,
	}

	return serialize(message)
}

//HandleMessage : parses message and determines appropriate actions.
func HandleMessage(message string) (string, error) {
	requestMessage, err := deserialize(message)
	if err != nil {
		return "", err
	}

	switch requestMessage.Type {
	case QueryBlockchainType:
		return handleQueryBlockchain()
	case ResponseBlockchainType:
		return handleResponseBlockchain(requestMessage)
	case QueryLatestBlockType:
		return handleQueryLatestBlock()
	case ResponseLatestBlockType:
		return handleResponseLatestBlock(requestMessage)
	}

	return "", errors.New("HandleMessage: request message type not supported")
}

func handleQueryBlockchain() (string, error) {
	blockchain := blockchain.GetBlockchain()
	serializedData, err := serialize(blockchain)
	if err != nil {
		return "", err
	}

	responseMessage := &MessageData{
		Data: serializedData,
		Type: ResponseBlockchainType,
	}

	return serialize(responseMessage)
}

func handleResponseBlockchain(requestMessage *MessageData) (string, error) {
	newBlockChain, err := deserializeBlockchain(requestMessage.Data)
	if err != nil {
		return "", err
	}
	SelectChain(newBlockChain)
	return "", nil
}

func handleResponseLatestBlock(requestMessage *MessageData) (string, error) {
	newBlock, err := deserializeBlock(requestMessage.Data)
	if err != nil {
		return "", err
	}
	blockAdded := ReceivedBlock(newBlock)
	if blockAdded {
		return handleQueryLatestBlock()
	}
	return "", nil
}

func handleQueryLatestBlock() (string, error) {
	latestBlock := blockchain.GetLatestBlock()
	serializedData, err := serialize(latestBlock)
	if err != nil {
		return "", err
	}

	responseMessage := &MessageData{
		Data: serializedData,
		Type: ResponseLatestBlockType,
	}

	return serialize(responseMessage)
}

func serialize(data interface{}) (string, error) {
	b, err := json.Marshal(data)
	if err != nil {
		return "", err
	}
	return string(b), nil
}

func deserialize(data string) (*MessageData, error) {
	messageData := MessageData{}
	err := json.Unmarshal([]byte(data), &messageData)
	if err != nil {
		return nil, err
	}
	return &messageData, nil
}

func deserializeBlock(data string) (*blockchain.Block, error) {
	messageData := &blockchain.Block{}
	err := json.Unmarshal([]byte(data), messageData)
	if err != nil {
		return nil, err
	}
	return messageData, nil
}

func deserializeBlockchain(data string) ([]*blockchain.Block, error) {
	blockchain := make([]*blockchain.Block, 0)
	err := json.Unmarshal([]byte(data), &blockchain)
	if err != nil {
		return nil, err
	}
	return blockchain, nil
}
