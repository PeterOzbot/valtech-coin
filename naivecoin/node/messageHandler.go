package node

import (
	"encoding/json"
	"errors"
	"naivecoin/blockchain"
)

//OnPeerConnected : Prepares message to send when new peer gets connected.
func OnPeerConnected() (string, error) {
	message := &MessageData{
		Type: GetBlockchainType,
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
	case GetBlockchainType:
		return handleGetBlockchain(requestMessage)
	case BlockchainType:
		return handleBlockchain(requestMessage)
	case GetLatestBlockType:
		return handleGetLatestBlock(requestMessage)
	case LatestBlockType:
		return handleLatestBlock(requestMessage)
	}

	return "", errors.New("HandleMessage: request message type not supported")
}

func handleGetBlockchain(requestMessage *MessageData) (string, error) {
	blockchain := blockchain.GetBlockchain()
	serializedData, err := serialize(blockchain)
	if err != nil {
		return "", err
	}

	responseMessage := &MessageData{
		Data: serializedData,
		Type: BlockchainType,
	}

	return serialize(responseMessage)
}

func handleBlockchain(requestMessage *MessageData) (string, error) {
	newBlockChain, err := deserializeBlockchain(requestMessage.Data)
	if err != nil {
		return "", err
	}
	SelectChain(newBlockChain)
	return "", nil
}

func handleLatestBlock(requestMessage *MessageData) (string, error) {
	newBlock, err := deserializeBlock(requestMessage.Data)
	if err != nil {
		return "", err
	}
	AddBlock(newBlock)
	return "", nil
}

func handleGetLatestBlock(requestMessage *MessageData) (string, error) {
	latestBlock := blockchain.GetLatestBlock()
	serializedData, err := serialize(latestBlock)
	if err != nil {
		return "", err
	}

	responseMessage := &MessageData{
		Data: serializedData,
		Type: LatestBlockType,
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
