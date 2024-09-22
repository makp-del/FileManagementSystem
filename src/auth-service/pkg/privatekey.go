package pkg

import (
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"os"
)

// PrivateKey stores the parsed RSA private key
var PrivateKey *rsa.PrivateKey

// LoadPrivateKey loads the RSA private key from the given file path
func LoadPrivateKey(privateKeyPath string) error {
	// Read the private key file
	keyData, err := os.ReadFile(privateKeyPath + "/private_key.pem")
	if err != nil {
		return fmt.Errorf("failed to read private key: %w", err)
	}

	// Decode the PEM block
	block, _ := pem.Decode(keyData)
	if block == nil || block.Type != "RSA PRIVATE KEY" {
		return fmt.Errorf("failed to decode PEM block containing the private key")
	}

	// Parse the RSA private key
	privateKey, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return fmt.Errorf("failed to parse private key: %w", err)
	}

	// Assign to the global PrivateKey variable
	PrivateKey = privateKey
	return nil
}