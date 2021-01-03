package blockchain

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"strconv"
	"strings"
)

//CalculateHash : Calculates block hash from its properties.
func (block *Block) CalculateHash() (string, error) {
	if block == nil {
		return "", nil
	}

	// convert timestamp to string
	var timestamp = strconv.FormatInt(block.Timestamp.Unix(), 10)

	// convert block transactions to string
	transactionsBytes, err := json.Marshal(block.Transactions)
	if err != nil {
		return "", err
	}
	transactions := string(transactionsBytes)

	// create string value from block properties
	blockStringRepresentation := strings.Join([]string{strconv.Itoa(block.Index), block.PreviousHash, timestamp, block.Message, transactions, strconv.Itoa(block.Difficulty), strconv.Itoa(block.Nonce)}, "")

	// create hasher and write the bytes to it
	hasher := sha256.New()
	hasher.Write([]byte(blockStringRepresentation))

	// return string
	return hex.EncodeToString(hasher.Sum(nil)), nil
}
