/*
SPDX-License-Identifier: Apache-2.0
*/

package main

import (
	"log"

	"github.com/hyperledger/fabric-contract-api-go/contractapi"
	"github.com/hyperledger/maghbari/firefly-chaincode/contracts/firefly"
)

func main() {
	fireflyChaincode, err := contractapi.NewChaincode(&firefly.SmartContract{})
	if err != nil {
		log.Panicf("Error creating firefly chaincode: %v", err)
	}

	if err := fireflyChaincode.Start(); err != nil {
		log.Panicf("Error starting afirefly chaincode: %v", err)
	}
}
