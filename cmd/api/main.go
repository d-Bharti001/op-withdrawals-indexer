package main

import (
	"context"
	"log"

	"github.com/ethereum/go-ethereum/ethclient"
)

func main() {
	cl, err := ethclient.Dial("http://localhost:8545")
	if err != nil {
		log.Fatalln(err)
	}

	chainid, err := cl.ChainID(context.Background())
	if err != nil {
		log.Fatalf("Error in fetching chain id: %v", err)
	}

	log.Println("Chain ID:", chainid)
}
