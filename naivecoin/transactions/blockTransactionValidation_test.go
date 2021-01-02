package transactions

import "testing"

// Test_ValidateBlockTransactions_EmptyTransaction : Tests validation when there are no transactions.
func Test_ValidateBlockTransactions_EmptyTransaction(t *testing.T) {

	// create inputs transaction
	var blockIndex = 12
	var transactions = []*Transaction{}
	var unspentTransactionOutputs = []*UnspentTransactionOutput{}

	// create coinbase transaction
	result, _ := ValidateBlockTransactions(transactions, unspentTransactionOutputs, blockIndex)

	// validate expected
	if result {
		t.Errorf("block transactions are not valid so the result should be false")
	}
}

// Test_ValidateBlockTransactions_Coinbase : Tests validation when fist transaction is not valid coinbase transaction.
func Test_ValidateBlockTransactions_Coinbase(t *testing.T) {

	// create inputs transaction
	var blockIndex = 12
	var transactions = []*Transaction{
		// coinbase transaction
		{
			ID: "invalid-coinbase",
			Inputs: []*TransactionInput{
				{
					OutputID:    "",
					Signature:   "",
					OutputIndex: blockIndex,
				},
			},
			Outputs: []*TransactionOutput{
				{
					Address: "coinbase-address",
					Amount:  CoinbaseAmount,
				},
			},
		},
		{
			ID: "3e5d88c061d2b79dd2ac79daf877232203089307d4576b2c1b3851b4920eb952",
			Inputs: []*TransactionInput{
				{
					OutputID:    "1",
					Signature:   "567af38a37a36b25e45a63f477c2b66a8e221a27831dc87624c6ebbe92ff16c5936eb0100490cbbc0b1658a2db4a0190b55c19b6756a9907ec9bddb8dfe0c141a7208fc1be073350e17a0d65aa9511d19e6713b28e37f3c732373d77aeac8e8a3e998721a64e235a7e84e0f16d61a6e556d329988f03f546e9906f7731f1aa78955666a65fa3739ef4198d7af2babe00c0fc268078c3992d1f1d6bed6be34ed3d475bb18437dc2aac31dbd90f891d6a0c9dbeefab6d40dd7b69c1b426eaa482841a637445988518fea20969bfa99312b16a95ba2d155e44d898ca8b8f189941ced763aa111826a45b669ff0f904419e475fce41829f9f2f26b11e9a9fb4f38a10bd12bf5a629c97dda67a61431bd3839a8a28e55646bf864286bc805002164a562b4ccc874dce4b9b9f08b33df5e5063af91d58fa4edd6d5f85d6d8a28c99534881ffaebac09e5990642fa4b14d349c1c4e23d3bd4d600f2e521b803c57c0b3fb820f81d8ba915cea300dc722f4ee1a5d2a339d5a85633151e17cb129ed6b750e69eb9e2f4aa43cefa94adf99675a2b01e0e837a80538e774839f4f27fc30034ae0a2d179da3eb34c1d46ba863f67c2efe4ff2405d89ad4f98acc57e411a3e85e3ba82dbe0e3e3c9a09dd99cfede261271a7cd442db4a34cbdd7fe11f1e3a8564e6e340a0c175e2ee5e950c2503a05caedabcb8c563c1157ed99eb0f2f8b7844",
					OutputIndex: 10,
				},
			},
			Outputs: []*TransactionOutput{
				{
					Address: testPublicKey,
					Amount:  100,
				},
			},
		},
		{
			ID: "7ef0ab206de97f0906adbaccb68bdd7039b86893cbeede8ef9311858b8187fdb",
			Inputs: []*TransactionInput{
				{
					OutputID:    "2",
					Signature:   "45b364938dcde0267e019ad19518abd5174659e45341b624174cc6c941a98fd40cb9b230a2dce924f17982403b88d20afd8d7a11e046c6dfa8d2cd2b713232016d95a4c30990fb02f2b6b611311a02c490bb0f0a7e941a26ddc0b41ebb2356c2a250ab1ae34190463a1e63f896eb7a2f20edaffbd5fd715a819b3ba9c36ce3fe4006fc476add623e874cdb880ca9e2962ec369e6b097930652948c4a175231716e24cefec3b90908139dfe1ae29ca469d00bfaa127838c73e135ad5a66a34242d2518fd66a35857d0d1f897b7035862642c0d07c45b9094039dc278572c06045c09acd568dc0adda60e022b591f76061ede28010cbba7f2758e1a1dbc1e374a8266421ad9fb79e2d4532f1466b687ded5c02aeed4020ea23b4c184181453ea111b3b6db6c8e381f1467e56aecc02475463d713fb1300c5c38379763b26c6b87cb0f27b7d3603e83416dae8f2cd06e2c48090c3b08b5cd6525c669f5a730eec9062c6cffb916db2ce90d41b1734b16eb7be54be19910e9f2669e254c880346aec5756ee8e0520e076838414fafb8348ee350258fd18910f4cca3d8630aa621642fc2b437c6d74a151383beb95aacfe3c7fdf31e372c1d9330abb9ba0be27af1ed745735bd8c09bab1fbc3e7f4f1baf070a260bdbe439b119ae09d87a09989f0cdfdc4f99a109b62a2db862d5ded19daf20d28aafb098efdeefedd935053bd0796",
					OutputIndex: 20,
				},
			},
			Outputs: []*TransactionOutput{
				{
					Address: testPublicKey,
					Amount:  200,
				},
			},
		},
	}
	var unspentTransactionOutputs = []*UnspentTransactionOutput{
		{
			OutputID:    "1",
			OutputIndex: 10,
			Address:     testPublicKey,
			Amount:      100,
		},
		{
			OutputID:    "2",
			OutputIndex: 20,
			Address:     testPublicKey,
			Amount:      200,
		},
	}

	// create coinbase transaction
	result, _ := ValidateBlockTransactions(transactions, unspentTransactionOutputs, blockIndex)

	// validate expected
	if result {
		t.Errorf("block transactions are not valid so the result should be false")
	}
}

// Test_ValidateBlockTransactions_Transactions : Tests validation when transactions are not valid.
func Test_ValidateBlockTransactions_Transactions(t *testing.T) {

	// create inputs transaction
	var blockIndex = 12
	var transactions = []*Transaction{
		// coinbase transaction
		{
			ID: "ebafa7518cac709e160f201a888bdf3c969c36993eefbf852cc30c9eb1a553b8",
			Inputs: []*TransactionInput{
				{
					OutputID:    "",
					Signature:   "",
					OutputIndex: blockIndex,
				},
			},
			Outputs: []*TransactionOutput{
				{
					Address: "coinbase-address",
					Amount:  CoinbaseAmount,
				},
			},
		},
		{
			ID: "3e5d88c061d2b79dd2ac79daf877232203089307d4576b2c1b3851b4920eb952",
			Inputs: []*TransactionInput{
				{
					OutputID:    "1",
					Signature:   "invalid",
					OutputIndex: 10,
				},
			},
			Outputs: []*TransactionOutput{
				{
					Address: testPublicKey,
					Amount:  100,
				},
			},
		},
		{
			ID: "7ef0ab206de97f0906adbaccb68bdd7039b86893cbeede8ef9311858b8187fdb",
			Inputs: []*TransactionInput{
				{
					OutputID:    "2",
					Signature:   "invalid",
					OutputIndex: 20,
				},
			},
			Outputs: []*TransactionOutput{
				{
					Address: testPublicKey,
					Amount:  200,
				},
			},
		},
	}
	var unspentTransactionOutputs = []*UnspentTransactionOutput{
		{
			OutputID:    "1",
			OutputIndex: 10,
			Address:     testPublicKey,
			Amount:      100,
		},
		{
			OutputID:    "2",
			OutputIndex: 20,
			Address:     testPublicKey,
			Amount:      200,
		},
	}

	// create coinbase transaction
	result, _ := ValidateBlockTransactions(transactions, unspentTransactionOutputs, blockIndex)

	// validate expected
	if result {
		t.Errorf("block transactions are not valid so the result should be false")
	}
}

// Test_ValidateBlockTransactions_IsValid : Tests validation when transactions are valid.
func Test_ValidateBlockTransactions_IsValid(t *testing.T) {

	// create inputs transaction
	var blockIndex = 12
	var transactions = []*Transaction{
		// coinbase transaction
		{
			ID: "ebafa7518cac709e160f201a888bdf3c969c36993eefbf852cc30c9eb1a553b8",
			Inputs: []*TransactionInput{
				{
					OutputID:    "",
					Signature:   "",
					OutputIndex: blockIndex,
				},
			},
			Outputs: []*TransactionOutput{
				{
					Address: "coinbase-address",
					Amount:  CoinbaseAmount,
				},
			},
		},
		{
			ID: "3e5d88c061d2b79dd2ac79daf877232203089307d4576b2c1b3851b4920eb952",
			Inputs: []*TransactionInput{
				{
					OutputID:    "1",
					Signature:   "567af38a37a36b25e45a63f477c2b66a8e221a27831dc87624c6ebbe92ff16c5936eb0100490cbbc0b1658a2db4a0190b55c19b6756a9907ec9bddb8dfe0c141a7208fc1be073350e17a0d65aa9511d19e6713b28e37f3c732373d77aeac8e8a3e998721a64e235a7e84e0f16d61a6e556d329988f03f546e9906f7731f1aa78955666a65fa3739ef4198d7af2babe00c0fc268078c3992d1f1d6bed6be34ed3d475bb18437dc2aac31dbd90f891d6a0c9dbeefab6d40dd7b69c1b426eaa482841a637445988518fea20969bfa99312b16a95ba2d155e44d898ca8b8f189941ced763aa111826a45b669ff0f904419e475fce41829f9f2f26b11e9a9fb4f38a10bd12bf5a629c97dda67a61431bd3839a8a28e55646bf864286bc805002164a562b4ccc874dce4b9b9f08b33df5e5063af91d58fa4edd6d5f85d6d8a28c99534881ffaebac09e5990642fa4b14d349c1c4e23d3bd4d600f2e521b803c57c0b3fb820f81d8ba915cea300dc722f4ee1a5d2a339d5a85633151e17cb129ed6b750e69eb9e2f4aa43cefa94adf99675a2b01e0e837a80538e774839f4f27fc30034ae0a2d179da3eb34c1d46ba863f67c2efe4ff2405d89ad4f98acc57e411a3e85e3ba82dbe0e3e3c9a09dd99cfede261271a7cd442db4a34cbdd7fe11f1e3a8564e6e340a0c175e2ee5e950c2503a05caedabcb8c563c1157ed99eb0f2f8b7844",
					OutputIndex: 10,
				},
			},
			Outputs: []*TransactionOutput{
				{
					Address: testPublicKey,
					Amount:  100,
				},
			},
		},
		{
			ID: "7ef0ab206de97f0906adbaccb68bdd7039b86893cbeede8ef9311858b8187fdb",
			Inputs: []*TransactionInput{
				{
					OutputID:    "2",
					Signature:   "45b364938dcde0267e019ad19518abd5174659e45341b624174cc6c941a98fd40cb9b230a2dce924f17982403b88d20afd8d7a11e046c6dfa8d2cd2b713232016d95a4c30990fb02f2b6b611311a02c490bb0f0a7e941a26ddc0b41ebb2356c2a250ab1ae34190463a1e63f896eb7a2f20edaffbd5fd715a819b3ba9c36ce3fe4006fc476add623e874cdb880ca9e2962ec369e6b097930652948c4a175231716e24cefec3b90908139dfe1ae29ca469d00bfaa127838c73e135ad5a66a34242d2518fd66a35857d0d1f897b7035862642c0d07c45b9094039dc278572c06045c09acd568dc0adda60e022b591f76061ede28010cbba7f2758e1a1dbc1e374a8266421ad9fb79e2d4532f1466b687ded5c02aeed4020ea23b4c184181453ea111b3b6db6c8e381f1467e56aecc02475463d713fb1300c5c38379763b26c6b87cb0f27b7d3603e83416dae8f2cd06e2c48090c3b08b5cd6525c669f5a730eec9062c6cffb916db2ce90d41b1734b16eb7be54be19910e9f2669e254c880346aec5756ee8e0520e076838414fafb8348ee350258fd18910f4cca3d8630aa621642fc2b437c6d74a151383beb95aacfe3c7fdf31e372c1d9330abb9ba0be27af1ed745735bd8c09bab1fbc3e7f4f1baf070a260bdbe439b119ae09d87a09989f0cdfdc4f99a109b62a2db862d5ded19daf20d28aafb098efdeefedd935053bd0796",
					OutputIndex: 20,
				},
			},
			Outputs: []*TransactionOutput{
				{
					Address: testPublicKey,
					Amount:  200,
				},
			},
		},
	}
	var unspentTransactionOutputs = []*UnspentTransactionOutput{
		{
			OutputID:    "1",
			OutputIndex: 10,
			Address:     testPublicKey,
			Amount:      100,
		},
		{
			OutputID:    "2",
			OutputIndex: 20,
			Address:     testPublicKey,
			Amount:      200,
		},
	}

	// create coinbase transaction
	result, _ := ValidateBlockTransactions(transactions, unspentTransactionOutputs, blockIndex)

	// validate expected
	if !result {
		t.Errorf("block transactions are valid so the result should be true")
	}
}
