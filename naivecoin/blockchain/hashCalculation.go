package blockchain

import (
	"crypto/sha256"
	"encoding/hex"
	"strconv"
	"strings"
)

//CalculateHash : Calculates block hash from its properties.
func (block *Block) CalculateHash() string {
	if block == nil {
		return ""
	}

	// create string value from block properties
	blockStringRepresentation := strings.Join([]string{strconv.Itoa(block.index), block.previousHash, strconv.FormatInt(block.timestamp.Unix(), 10), block.data}, "")

	// create hasher and write the bytes to it
	hasher := sha256.New()
	hasher.Write([]byte(blockStringRepresentation))

	// return string
	return hex.EncodeToString(hasher.Sum(nil))
}
