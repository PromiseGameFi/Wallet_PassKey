package wallet

import (
	"crypto/ed25519"
	"crypto/sha256"
	"encoding/base64"
	"errors"
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

func (w *HDWallet) CreatePasskey(passkey []byte) (string, error) {
	w.mu.Lock()
	defer w.mu.Unlock()

	// Generate master seed
	masterSeed := sha256.Sum256(passkey)
	w.masterKey = ed25519.NewKeyFromSeed(masterSeed[:])
	
	// Derive first address
	firstAddress := w.deriveAddress(0)

	// Store passkey hash
	w.passkeyHash = sha256.Sum256(passkey)[:]

	// Convert to base64 for readability
	return base64.StdEncoding.EncodeToString(firstAddress), nil
}

func (w *HDWallet) DeriveAddress(index uint32) (string, error) {
	w.mu.Lock()
	defer w.mu.Unlock()

	if w.masterKey == nil {
		return "", errors.New("wallet not initialized")
	}

	address := w.deriveAddress(index)
	return base64.StdEncoding.EncodeToString(address), nil
}

func (w *HDWallet) deriveAddress(index uint32) []byte {
	derivationKey := append(w.masterKey, byte(index))
	childKey := ed25519.NewKeyFromSeed(derivationKey)
	publicKey := childKey.Public().(ed25519.PublicKey)

	w.derivedKeys[index] = publicKey
	return publicKey
}

func (w *HDWallet) RemoveAddress(index uint32) error {
	w.mu.Lock()
	defer w.mu.Unlock()

	delete(w.derivedKeys, index)
	return nil
}

func (w *HDWallet) ListAddresses() []string {
	w.mu.RLock()
	defer w.mu.RUnlock()

	addresses := make([]string, 0, len(w.derivedKeys))
	for _, pubKey := range w.derivedKeys {
		addresses = append(addresses, base64.StdEncoding.EncodeToString(pubKey))
	}
	return addresses
}