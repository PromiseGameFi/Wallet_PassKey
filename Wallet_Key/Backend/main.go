package main

import (
	"crypto/ed25519"
	"crypto/sha256"
	"fmt"
	"sync"
)

type HDWallet struct {
	mu           sync.RWMutex
	masterKey    ed25519.PrivateKey
	derivedKeys  map[uint32]ed25519.PublicKey
	passkeyHash  []byte
}

func NewHDWallet() *HDWallet {
	return &HDWallet{
		derivedKeys: make(map[uint32]ed25519.PublicKey),
	}
}

func (w *HDWallet) CreatePasskey(passkey []byte) ([]byte, error) {
	w.mu.Lock()
	defer w.mu.Unlock()

	// Generate master seed from passkey
	masterSeed := sha256.Sum256(passkey)
	w.masterKey = ed25519.NewKeyFromSeed(masterSeed[:])
	
	// Derive first address
	firstAddress := w.deriveAddress(0)

	// Store passkey hash
	w.passkeyHash = sha256.Sum256(passkey)

	return firstAddress, nil
}

func (w *HDWallet) deriveAddress(index uint32) []byte {
	derivationKey := append(w.masterKey, byte(index))
	childKey := ed25519.NewKeyFromSeed(derivationKey)
	publicKey := childKey.Public().(ed25519.PublicKey)

	w.derivedKeys[index] = publicKey
	return publicKey
}

func main() {
	wallet := NewHDWallet()
	
	// Example usage
	firstAddress, err := wallet.CreatePasskey([]byte("mypasskey"))
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("First Address:", string(firstAddress))
}