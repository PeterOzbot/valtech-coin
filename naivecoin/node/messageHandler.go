package node

import (
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

//OnNewBlock : Prepares message to send when new block is added.
func OnNewBlock(newBlock *blockchain.Block) (string, error) {
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

// the chain is requested
func handleQueryBlockchain() (string, error) {
	blockchain := ChainState.GetBlockchain(CurrentBlockchain)
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

// whole chain is received
func handleResponseBlockchain(requestMessage *MessageData) (string, error) {
	newBlockChain, err := deserializeBlockchain(requestMessage.Data)
	if err != nil {
		return "", err
	}
	err = SelectChain(newBlockChain)
	return "", err
}

// new block is received
func handleResponseLatestBlock(requestMessage *MessageData) (string, error) {
	newBlock, err := deserializeBlock(requestMessage.Data)
	if err != nil {
		return "", err
	}

	// try to add new block and determine if sender's whole chain may be required
	queryWholeChain, err := ReceivedBlock(newBlock)
	if queryWholeChain && err == nil {
		message := &MessageData{
			Type: QueryBlockchainType,
		}

		return serialize(message)
	}
	return "", err
}

// latest block was requested
func handleQueryLatestBlock() (string, error) {
	currentBlockchain := ChainState.GetBlockchain(CurrentBlockchain)
	latestBlock := currentBlockchain[len(currentBlockchain)-1]
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
