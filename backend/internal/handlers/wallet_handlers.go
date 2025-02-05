package handlers

import (
	"encoding/json"
	"net/http"
)

func GetWalletAddresses(w http.ResponseWriter, r *http.Request) {
	// Fetch and return wallet addresses
	addresses := []string{"0xAddress1", "0xAddress2"}
	json.NewEncoder(w).Encode(addresses)
}

func CreateTransaction(w http.ResponseWriter, r *http.Request) {
	// Handle transaction creation
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Transaction created"))
}
