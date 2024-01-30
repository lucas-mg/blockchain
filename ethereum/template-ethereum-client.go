package main

import (
	"context"
	"fmt"
	"log"
	"math/big"
	"strings"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/rpc"
)

func main() {
	client, err := rpc.Dial("http://localhost:8545") // Replace with the actual URL of your Ethereum node
	if err != nil {
		log.Fatal(err)
	}

	// Example: Get the latest block number
	blockNumber, err := getLatestBlockNumber(client)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Latest Block Number: %s\n", blockNumber.String())

	// Example: Get Ethereum client version
	version, err := getClientVersion(client)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Ethereum Client Version: %s\n", version)
}

func getLatestBlockNumber(client *rpc.Client) (*big.Int, error) {
	var blockNumber hexutil.Big
	err := client.CallContext(context.Background(), &blockNumber, "eth_blockNumber")
	if err != nil {
		return nil, err
	}
	return (*big.Int)(&blockNumber), nil
}

func getClientVersion(client *rpc.Client) (string, error) {
	var version string
	err := client.CallContext(context.Background(), &version, "web3_clientVersion")
	if err != nil {
		return "", err
	}

	// Clean up the version string
	version = strings.TrimPrefix(version, "Geth/")
	return version, nil
}
