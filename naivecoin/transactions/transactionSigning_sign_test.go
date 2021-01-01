package transactions

import (
	"crypto"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/hex"
	"encoding/pem"
	"log"
	"testing"
)

// Test_Sign_NilInput : Tests if nil input is handled without failing.
func Test_Sign_NilInput(t *testing.T) {

	// create nil transaction
	var transaction *Transaction

	// sign
	result, _ := transaction.Sign("")

	// result should be empty string
	if len(result) != 0 {
		t.Errorf("result of nil transaction should be empty string.")
	}
}

// Test_Sign_Correct : Tests if singing is correct.
func Test_Sign_Correct(t *testing.T) {

	// create transaction
	var transaction *Transaction = &Transaction{
		ID: "7d89bc87859d24f03736ab81c618527b7e128d0e7f98754dec76bb231b5b533c",
	}

	// sign
	result, _ := transaction.Sign(privateKey)

	// check if encrypted result can be decrypted
	pemBlock, _ := pem.Decode([]byte(publicKey))
	pub, err := x509.ParsePKIXPublicKey(pemBlock.Bytes)
	if err != nil {
		log.Fatal(err)
	}
	hashedID := sha256.Sum256([]byte(transaction.ID))
	decodedResult, err := hex.DecodeString(result)
	if err != nil {
		log.Fatal(err)
	}
	var verificationError = rsa.VerifyPKCS1v15(pub.(*rsa.PublicKey), crypto.SHA256, hashedID[:], []byte(decodedResult))
	if verificationError != nil {
		log.Fatal(verificationError)
		t.Errorf("Signing is not correct. Signature verification failed.")
	}
}
