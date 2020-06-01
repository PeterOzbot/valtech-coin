package blockchain

import (
	"testing"
	"time"
)

//Test_isGenesisBlockValid : Tests if the genesis block validation correctly validates valid block.
func Test_isGenesisBlockValid(t *testing.T) {
	// create testing block
	block := GenesisBlock()

	// validate block
	result := isGenesisBlockValid(block)

	// result should be positive
	if !result {
		t.Errorf("block validation failed even if the blocks are the same")
	}
}

//Test_isGenesisBlockValid_Index : Tests if the genesis block validation validates Index.
func Test_isGenesisBlockValid_Index(t *testing.T) {
	// create testing block and brake it
	block := GenesisBlock()
	block.Index = 234

	// validate block
	result := isGenesisBlockValid(block)

	// result should be false as Index is incorrect
	if result {
		t.Errorf("block validation passed even if Index is not the same")
	}
}

//Test_isGenesisBlockValid_Hash : Tests if the genesis block validation validates Hash.
func Test_isGenesisBlockValid_Hash(t *testing.T) {
	// create testing block and brake it
	block := GenesisBlock()
	block.Hash = "asdasd"

	// validate block
	result := isGenesisBlockValid(block)

	// result should be false as Hash is incorrect
	if result {
		t.Errorf("block validation passed even if Hash is not the same")
	}
}

//Test_isGenesisBlockValid_Data : Tests if the genesis block validation validates Data.
func Test_isGenesisBlockValid_Data(t *testing.T) {
	// create testing block and brake it
	block := GenesisBlock()
	block.Data = "asdasd"

	// validate block
	result := isGenesisBlockValid(block)

	// result should be false as Data is incorrect
	if result {
		t.Errorf("block validation passed even if Data is not the same")
	}
}

//Test_isGenesisBlockValid_PreviousHash : Tests if the genesis block validation validates PreviousHash.
func Test_isGenesisBlockValid_PreviousHash(t *testing.T) {
	// create testing block and brake it
	block := GenesisBlock()
	block.PreviousHash = "asdasd"

	// validate block
	result := isGenesisBlockValid(block)

	// result should be false as PreviousHash is incorrect
	if result {
		t.Errorf("block validation passed even if PreviousHash is not the same")
	}
}

//Test_isGenesisBlockValid_Timestamp : Tests if the genesis block validation validates Timestamp.
func Test_isGenesisBlockValid_Timestamp(t *testing.T) {
	// create testing block and brake it
	block := GenesisBlock()
	block.Timestamp = time.Unix(74747474, 324234)

	// validate block
	result := isGenesisBlockValid(block)

	// result should be false as Timestamp is incorrect
	if result {
		t.Errorf("block validation passed even if Timestamp is not the same")
	}
}
