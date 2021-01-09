package node

import (
	"encoding/json"
	"naivecoin/blockchain"
)

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
