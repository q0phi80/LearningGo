package main

import (
	"crypto/rand" // The rand.Reader from this package will be used to seed the generation of the rsa.PrivateKey
	"crypto/rsa"  // Required to generate the private key
	"crypto/sha256"
	"fmt"
	"os"
)

func main() {
	privateKey, err := rsa.GenerateKey(rand.Reader, 1024)
	if err != nil {
		fmt.Printf("[!] Error generating RSA private key: %v", err)
	}
	publicKey := privateKey.PublicKey
	text := []byte("My Secret Text")

	// Encrypt the data using publicKey
	ciphertext, err := rsa.EncryptOAEP(sha256.New(), rand.Reader, &publicKey, text, nil)
	if err != nil {
		fmt.Printf("[!] Error encrypting data: %v", err)
		os.Exit(1)
	}
	fmt.Printf("Encrypted ciphertext: %x\n", ciphertext)
	// Use privateKey to decrypt the ciphertext
	decrypted, err := rsa.DecryptOAEP(sha256.New(), rand.Reader, privateKey, ciphertext, nil)
	if err != nil {
		fmt.Printf("[!] Error decrypting data: %v", err)
		os.Exit(1)
	}
	fmt.Println("Decrypted text: ", string(decrypted))
}
