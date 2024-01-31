package main

import (
	"context"
	"encoding/hex"
	"fmt"
	"log"
	"math/big"
	"net/http"
	"strings"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/rpc"
)

func main() {
	http.HandleFunc("/balance", getBalanceHandler)
	http.HandleFunc("/sendTransaction", sendTransactionHandler)

	fmt.Println("EVM API server listening on :8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func getBalanceHandler(w http.ResponseWriter, r *http.Request) {
	address := r.URL.Query().Get("address")
	if address == "" {
		http.Error(w, "Address parameter is required", http.StatusBadRequest)
		return
	}

	client, err := rpc.Dial("http://localhost:8545")
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to connect to Ethereum node: %v", err), http.StatusInternalServerError)
		return
	}
	defer client.Close()

	ctx := context.Background()
	var balanceHex string
	err = client.CallContext(ctx, &balanceHex, "eth_getBalance", address, "latest")
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to get balance: %v", err), http.StatusInternalServerError)
		return
	}

	balance, success := new(big.Int).SetString(strings.TrimPrefix(balanceHex, "0x"), 16)
	if !success {
		http.Error(w, "Failed to convert balance to integer", http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, "Balance for address %s: %s wei", address, balance.String())
}

func sendTransactionHandler(w http.ResponseWriter, r *http.Request) {
	// Implement the logic to send a transaction
	// This is just a placeholder; you would need to handle the creation and signing of transactions
	http.Error(w, "Not implemented yet", http.StatusNotImplemented)
}
