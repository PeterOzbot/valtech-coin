package blockchain

import (
	"encoding/hex"
	"fmt"
)

//HashMatchesDifficulty : Check if block has valid hash for block difficulty
func (block *Block) HashMatchesDifficulty() bool {

	// convert hash to byte array
	byteHash, err := hex.DecodeString(block.Hash)
	if err != nil {
		fmt.Println("converting hash to bytes failed: ", err)
		return false
	}

	// temp difficulty value, used in validation process
	leadingZerosNeeded := block.Difficulty

	// go through hash bytes and covert each into binary then check for zeroes
	for _, singleHashByte := range byteHash {
		// convert byte into binary
		binary := fmt.Sprintf("%08b", singleHashByte)

		// go through 'bits' and check if number of leading zeros matched difficulty
		for _, bit := range binary {

			// if it came to the point when difficulty is zero then there were enough zeros
			if leadingZerosNeeded == 0 {
				return true
			}

			// if bit is zero decrease temp difficulty
			if string(bit) == "0" {
				leadingZerosNeeded--
			} else {
				return false
			}
		}
	}

	return true
}
