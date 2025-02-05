package main

import (
	"hd-wallet/internal/handlers"
	"log"
	"net/http"
)

func main() {
	// Register handlers
	http.HandleFunc("/auth/register/begin", handlers.BeginRegistration)
	http.HandleFunc("/auth/register/finish", handlers.FinishRegistration)
	http.HandleFunc("/auth/login/begin", handlers.BeginLogin)
	http.HandleFunc("/auth/login/finish", handlers.FinishLogin)
	http.HandleFunc("/wallet/addresses", handlers.GetWalletAddresses)
	http.HandleFunc("/wallet/transactions", handlers.CreateTransaction)

	// Start the server
	log.Println("Server started on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}