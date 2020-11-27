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
	blockStringRepresentation := strings.Join([]string{strconv.Itoa(block.Index), block.PreviousHash, timestamp, block.Data, strconv.Itoa(block.Difficulty), strconv.Itoa(block.Nonce)}, "")

	// create hasher and write the bytes to it
	hasher := sha256.New()
	hasher.Write([]byte(blockStringRepresentation))

	// return string
	return hex.EncodeToString(hasher.Sum(nil))
}

//MineBlock : Tries to find nonce to get hash that matches difficulty.
func (block *Block) MineBlock() {
	for {
		// get hash from current block values
		block.Hash = block.CalculateHash()

		// checks if hash is valid for current difficulty
		if block.HashMatchesDifficulty() {
			return
		}

		// increase nonce if the produced hash is not valid
		block.Nonce++
	}
}
