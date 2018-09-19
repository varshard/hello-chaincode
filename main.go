package main

import (
	"fmt"

	"github.com/hyperledger/fabric/core/chaincode/shim"
)

func main() {
	if err := shim.Start(new(SimpleAsset)); err != nil {
		fmt.Printf("Error starting chaincode: %s", err)
	}
}
