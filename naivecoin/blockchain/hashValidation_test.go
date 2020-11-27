package blockchain

import (
	"testing"
)

//Test_HashMatchesDifficulty : Tests if hash pasing fails.
func Test_HashMatchesDifficulty_HashCantBeParsed(t *testing.T) {
	// create example block
	var block *Block = &Block{
		Hash: "blablalbalblablablalbalbl",
	}

	// validate block
	result := block.HashMatchesDifficulty()

	// result should be false
	if result {
		t.Errorf("hashin parsing fail should return false")
	}
}

//Test_HashMatchesDifficulty : Checks if valid hash is correctly validated.
func Test_HashMatchesDifficulty_Valid(t *testing.T) {
	// create example block
	var block *Block = &Block{
		Hash:       "08222a4e5e618f5929dbc0e3168baa53d77f49c31875bdc13d8cd4f9c738df76",
		Difficulty: 4,
	}

	// validate block
	validHash := block.HashMatchesDifficulty()

	// validate result
	if !validHash {
		t.Errorf("validating hash: '%s' with difficulty: '%d' should succeed.", block.Hash, block.Difficulty)
	}
}

//Test_HashMatchesDifficulty : Checks if invalid hash is correctly validated.
func Test_HashMatchesDifficulty_Invalid(t *testing.T) {
	// create example block
	var block *Block = &Block{
		Hash:       "08222a4e5e618f5929dbc0e3168baa53d77f49c31875bdc13d8cd4f9c738df76",
		Difficulty: 6,
	}

	// validate block
	validHash := block.HashMatchesDifficulty()

	// validate result
	if validHash {
		t.Errorf("validating hash: '%s' with difficulty: '%d' should fail.", block.Hash, block.Difficulty)
	}
}

//Test_HashMatchesDifficulty : Checks if valid hash is correctly validated when difficulty is so large that multiple bytes are checked.
func Test_HashMatchesDifficulty_ValidMultipleBytesCheck(t *testing.T) {
	// create example block
	var block *Block = &Block{
		Hash:       "00082a4e5e618f5929dbc0e3168baa53d77f49c31875bdc13d8cd4f9c738df76",
		Difficulty: 10,
	}

	// validate block
	validHash := block.HashMatchesDifficulty()

	// validate result
	if !validHash {
		t.Errorf("validating hash: '%s' with difficulty: '%d' should succeed.", block.Hash, block.Difficulty)
	}
}

//Test_HashMatchesDifficulty : Checks if invalid hash is correctly validated when difficulty is so large that multiple bytes are checked.
func Test_HashMatchesDifficulty_InvalidMultipleBytesCheck(t *testing.T) {
	// create example block
	var block *Block = &Block{
		Hash:       "00882a4e5e618f5929dbc0e3168baa53d77f49c31875bdc13d8cd4f9c738df76",
		Difficulty: 10,
	}

	// validate block
	validHash := block.HashMatchesDifficulty()

	// validate result
	if validHash {
		t.Errorf("validating hash: '%s' with difficulty: '%d' should fail.", block.Hash, block.Difficulty)
	}
}

//Test_HashMatchesDifficulty : Checks if valid hash is correctly validated when leading zeros match exactly.
func Test_HashMatchesDifficulty_ValidExactly(t *testing.T) {
	// create example block
	var block *Block = &Block{
		Hash:       "00882a4e5e618f5929dbc0e3168baa53d77f49c31875bdc13d8cd4f9c738df76",
		Difficulty: 8,
	}

	// validate block
	validHash := block.HashMatchesDifficulty()

	// validate result
	if !validHash {
		t.Errorf("validating hash: '%s' with difficulty: '%d' should succeed.", block.Hash, block.Difficulty)
	}
}
