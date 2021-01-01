package transactions

import (
	"crypto"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/hex"
	"encoding/pem"
)

//IsValid : Checks if transaction is valid one.
func (transaction *Transaction) IsValid(unspentTransactionOutputs []*UnspentTransactionOutput) (bool, error) {
	if transaction == nil {
		return false, nil
	}

	// get the proper transaction ID
	var calculatedTransactionID = transaction.CalculateID()
	// check if transaction has proper ID
	if transaction.ID != calculatedTransactionID {
		return false, nil
	}

	// calculate sum of transaction outputs amount
	var outputsAmountSum float64 = 0
	for _, transactionOutput := range transaction.Outputs {
		outputsAmountSum += transactionOutput.Amount
	}

	// calculate sum of unspent outputs amount
	var unspentOutputsAmountSum float64 = 0
	for _, transactionInput := range transaction.Inputs {

		// try to find unspent output
		var unspentTransactionOutput = FindUnspentTransactionOutput(unspentTransactionOutputs, transactionInput)

		// if found add to the total amount and validate signature to the address
		if unspentTransactionOutput != nil {

			// add to total amount
			unspentOutputsAmountSum += unspentTransactionOutput.Amount

			// validate signature to the transaction ID with unspent address
			signatureValid, err := validateSignature(transaction.ID, transactionInput.Signature, unspentTransactionOutput.Address)
			if !signatureValid || err != nil {
				return false, err
			}

		} else {
			// if not found transaction is not valid
			return false, nil
		}
	}

	// validate amounts
	if outputsAmountSum != unspentOutputsAmountSum {
		return false, nil
	}

	return true, nil
}

// validates signature to the transaction ID with unspent address
func validateSignature(transactionID string, transactionInputSignature string, unspentOutputAddress string) (bool, error) {

	// unspentOutputAddress is actually public key
	// first try to decode it to PEM block
	pemBlock, _ := pem.Decode([]byte(unspentOutputAddress))
	if pemBlock == nil {
		return false, nil
	}
	// try to get the public key out of the PEM block
	pub, err := x509.ParsePKIXPublicKey(pemBlock.Bytes)
	if err != nil {
		return false, err
	}

	// get the string value out of signature which is hex encoded
	decodedTransactionInputSignature, err := hex.DecodeString(transactionInputSignature)
	if err != nil {
		return false, err
	}

	// hash the unsigned transactionID so we can use the value in signature verification
	hashedID := sha256.Sum256([]byte(transactionID))

	// verify signed decoded transactionID to the hashed unsigned transactionID
	var verificationError = rsa.VerifyPKCS1v15(pub.(*rsa.PublicKey), crypto.SHA256, hashedID[:], []byte(decodedTransactionInputSignature))

	// verification failed
	if verificationError != nil {
		return false, verificationError
	}

	// verification was success if there is no error
	return true, nil
}
