package wallet

import (
	"bytes"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
)

//GenerateAddress : Generates new private and public key and returns them both in the Address.
func GenerateAddress() (*Address, error) {

	// generate private key
	privatekey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		return nil, err
	}

	// get private key PEM format
	privateKeyBytes := x509.MarshalPKCS1PrivateKey(privatekey)
	privateKeyString, err := encodePEM(privateKeyBytes, "RSA PRIVATE KEY")
	if err != nil {
		return nil, err
	}

	// get public key PEM format
	publicKeyBytes, err := x509.MarshalPKIXPublicKey(&privatekey.PublicKey)
	publicKeyString, err := encodePEM(publicKeyBytes, "PUBLIC KEY")
	if err != nil {
		return nil, err
	}

	// return address
	return &Address{
		PrivateKey: privateKeyString,
		PublicKey:  publicKeyString,
	}, nil
}

// converts key as PEM into string
func encodePEM(keyBytes []byte, pemType string) (string, error) {
	// create the PEM block
	privateKeyBlock := &pem.Block{
		Type:  pemType,
		Bytes: keyBytes,
	}

	// encode it into buffer so we can convert it into string
	var keyPem bytes.Buffer
	err := pem.Encode(&keyPem, privateKeyBlock)
	if err != nil {
		return "", err
	}

	// return as string
	return keyPem.String(), nil
}
