package wallet

import (
	"crypto/rand"

	"github.com/ethereum/go-ethereum/accounts"
)

type HDWallet struct {
	MasterSeed []byte
	MasterKey  *accounts.Account
}

func NewHDWallet() (*HDWallet, error) {
	seed, err := generateMasterSeed()
	if err != nil {
		return nil, err
	}

	masterKey, err := deriveMasterKey(seed)
	if err != nil {
		return nil, err
	}

	return &HDWallet{
		MasterSeed: seed,
		MasterKey:  masterKey,
	}, nil
}

func generateMasterSeed() ([]byte, error) {
	seed := make([]byte, 64)
	_, err := rand.Read(seed)
	if err != nil {
		return nil, err
	}
	return seed, nil
}

func deriveMasterKey(seed []byte) (*accounts.Account, error) {
	masterKey, err := hd.NewMaster(seed, []byte("m/44'/60'/0'/0"))
	if err != nil {
		return nil, err
	}
	return masterKey, nil
}
