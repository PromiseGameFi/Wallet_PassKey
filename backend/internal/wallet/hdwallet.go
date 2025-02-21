package wallet

import (
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

}

func generateMasterSeed() ([]byte, error) {

}

func deriveMasterKey(seed []byte) (*accounts.Account, error) {

}
