package main

import (
	"fmt"
	"passkey-wallet/wallet"
)

func main() {
	// Create wallet
	hdWallet := wallet.NewHDWallet()

	
	// Create passkey and first address
	firstAddress, err := hdWallet.CreatePasskey([]byte("myStrongPasskey"))
	if err != nil {
		fmt.Println("Wallet creation error:", err)
		return
	}
	fmt.Println("First Address:", firstAddress)

	// Derive additional addresses
	for i := 1; i < 3; i++ {
		addr, err := hdWallet.DeriveAddress(uint32(i))
		if err != nil {
			fmt.Println("Address derivation error:", err)
			continue
		}
		fmt.Printf("Derived Address %d: %s\n", i, addr)
	}

	// List all addresses
	addresses := hdWallet.ListAddresses()
	fmt.Println("All Addresses:", addresses)
}