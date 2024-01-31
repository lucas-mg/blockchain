package main

import (
	"bytes"
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"log"

	"github.com/btcsuite/btcd/btcec"
	"github.com/btcsuite/btcutil"
)

func main() {
	// Step 1: Generate a new private key
	privateKey, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	if err != nil {
		log.Fatal(err)
	}

	// Step 2: Derive the public key from the private key
	publicKey := append(privateKey.PublicKey.X.Bytes(), privateKey.PublicKey.Y.Bytes()...)

	// Step 3: Create a Bitcoin address from the public key
	address, err := btcutil.NewAddressPubKey(publicKey, &btcutil.MainNetParams)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Generated Bitcoin address: %s\n", address.EncodeAddress())

	// Step 4: Create a simple transaction
	senderAddress := address
	recipientAddress := "recipient_address_here"
	amount := 0.001

	transaction := createTransaction(privateKey, senderAddress, recipientAddress, amount)
	fmt.Printf("Created transaction: %s\n", transaction)
}

func createTransaction(privateKey *ecdsa.PrivateKey, sender btcutil.Address, recipient string, amount float64) string {
	// Construct a simple transaction (not suitable for real-world scenarios)
	tx := fmt.Sprintf("Transaction: %s sent %f BTC to %s", sender.EncodeAddress(), amount, recipient)

	// Sign the transaction with the private key
	signature, err := signTransaction(privateKey, tx)
	if err != nil {
		log.Fatal(err)
	}

	// Include the signature in the transaction
	signedTransaction := fmt.Sprintf("%s\nSignature: %s", tx, signature)
	return signedTransaction
}

func signTransaction(privateKey *ecdsa.PrivateKey, data string) (string, error) {
	// Hash the transaction data
	hashedData := sha256.Sum256([]byte(data))

	// Sign the hashed data with the private key
	r, s, err := ecdsa.Sign(rand.Reader, privateKey, hashedData[:])
	if err != nil {
		return "", err
	}

	// Encode the signature
	signature := append(r.Bytes(), s.Bytes()...)

	// Return the hex-encoded signature
	return hex.EncodeToString(signature), nil
}
