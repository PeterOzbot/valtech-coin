package transactions

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/hex"
	"encoding/pem"
)

//Sign : Creates string content from transaction inputs, id and unspent data. The whole thing is signed with private key.
func (transaction *Transaction) Sign(privateKey string) (string, error) {
	if transaction == nil {
		return "", nil
	}

	// try to decode private key
	pemBlock, _ := pem.Decode([]byte(privateKey))
	if pemBlock == nil {
		panic("Cant decode private key.")
	}
	// try to parse private key
	parsedPrivateKey, parseError := x509.ParsePKCS1PrivateKey(pemBlock.Bytes)
	if parseError != nil {
		return "", parseError
	}

	// hash message to be signed
	hashedTransactionID := sha256.Sum256([]byte(transaction.ID))

	// sign
	signedTransactionID, signingError := rsa.SignPKCS1v15(rand.Reader, parsedPrivateKey,
		crypto.SHA256, hashedTransactionID[:])
	if signingError != nil {
		return "", signingError
	}

	// encode to hex
	encodedSignedTransactionID := hex.EncodeToString(signedTransactionID)

	// convert to string and return
	return string(encodedSignedTransactionID), nil
}

//CanSign : Checks if unspent transaction can be located with transaction input and validates that the public key matches address on unspent transaction.
func CanSign(unspentTransactions []*UnspentTransactionOutput, transactionInput *TransactionInput, publicKey string) bool {
	if transactionInput == nil {
		return false
	}

	// find the unspent transaction that matched transaction input
	var matchingUnspentTransaction = FindUnspentTransactionOutput(unspentTransactions, transactionInput)

	// check if unspent transaction was found
	if matchingUnspentTransaction == nil {
		return false
	}

	// match key
	return matchingUnspentTransaction.Address == publicKey
}
