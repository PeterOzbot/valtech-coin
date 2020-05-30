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

	// convert timestamp to string
	var timestamp = strconv.FormatInt(block.Timestamp.Unix(), 10)

	// create string value from block properties
	blockStringRepresentation := strings.Join([]string{strconv.Itoa(block.Index), block.PreviousHash, timestamp, block.Data}, "")

	// create hasher and write the bytes to it
	hasher := sha256.New()
	hasher.Write([]byte(blockStringRepresentation))

	// return string
	return hex.EncodeToString(hasher.Sum(nil))
}
