package blockchain

//MineBlock : Tries to find nonce to get hash that matches difficulty.
func (block *Block) MineBlock() error {
	for {
		// get hash from current block values
		newHash, err := block.CalculateHash()
		if err != nil {
			return err
		}

		// set new hash
		block.Hash = newHash

		// checks if hash is valid for current difficulty
		if block.HashMatchesDifficulty() {
			return nil
		}

		// increase nonce if the produced hash is not valid
		block.Nonce++
	}
}
