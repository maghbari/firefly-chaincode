// SPDX-License-Identifier: Apache-2.0

package main

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/golang/protobuf/ptypes/timestamp"
	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

// SmartContract provides functions for managing an Asset
type SmartContract struct {
	contractapi.Contract
}

// PutValue - Adds a key value pair to the world state
func (sc *SmartContract) PutValue(ctx contractapi.TransactionContextInterface, key string, value string) error {
	return ctx.GetStub().PutState(key, []byte(value))
}

// GetValue - Gets the value for a key from the world state
func (sc *SmartContract) GetValue(ctx contractapi.TransactionContextInterface, key string) (string, error) {
	bytes, err := ctx.GetStub().GetState(key)

	if err != nil {
		return "", nil
	}

	return string(bytes), nil
}

func main() {
	SmartContract := new(SmartContract)

	cc, err := contractapi.NewChaincode(SmartContract)

	if err != nil {
		panic(err.Error())
	}

	if err := cc.Start(); err != nil {
		panic(err.Error())
	}
}

type BatchPinEvent struct {
	Signer     string               `json:"signer"`
	Timestamp  *timestamp.Timestamp `json:"timestamp"`
	Action     string               `json:"action"`
	Uuids      string               `json:"uuids"`
	BatchHash  string               `json:"batchHash"`
	PayloadRef string               `json:"payloadRef"`
	Contexts   []string             `json:"contexts"`
}

func (s *SmartContract) PinBatch(ctx contractapi.TransactionContextInterface, uuids, batchHash, payloadRef string, contexts []string) error {
	cid := ctx.GetClientIdentity()
	id, err := cid.GetID()
	if err != nil {
		return fmt.Errorf("Failed to obtain client identity's ID: %s", err)
	}
	idString, err := base64.StdEncoding.DecodeString(id)
	if err != nil {
		return fmt.Errorf("Failed to decode client identity ID: %s", err)
	}
	mspId, err := cid.GetMSPID()
	if err != nil {
		return fmt.Errorf("Failed to obtain client identity's MSP ID: %s", err)
	}
	timestamp, err := ctx.GetStub().GetTxTimestamp()
	if err != nil {
		return fmt.Errorf("Failed to get transaction timestamp: %s", err)
	}
	event := BatchPinEvent{
		Signer:     fmt.Sprintf("%s::%s", mspId, idString),
		Timestamp:  timestamp,
		Action:     "",
		Uuids:      uuids,
		BatchHash:  batchHash,
		PayloadRef: payloadRef,
		Contexts:   contexts,
	}
	bytes, err := json.Marshal(event)
	if err != nil {
		return fmt.Errorf("Failed to marshal event: %s", err)
	}
	ctx.GetStub().SetEvent("BatchPin", bytes)
	return nil
}

func (s *SmartContract) NetworkAction(ctx contractapi.TransactionContextInterface, action, payload string) error {
	cid := ctx.GetClientIdentity()
	id, err := cid.GetID()
	if err != nil {
		return fmt.Errorf("Failed to obtain client identity's ID: %s", err)
	}
	idString, err := base64.StdEncoding.DecodeString(id)
	if err != nil {
		return fmt.Errorf("Failed to decode client identity ID: %s", err)
	}
	mspId, err := cid.GetMSPID()
	if err != nil {
		return fmt.Errorf("Failed to obtain client identity's MSP ID: %s", err)
	}
	timestamp, err := ctx.GetStub().GetTxTimestamp()
	if err != nil {
		return fmt.Errorf("Failed to get transaction timestamp: %s", err)
	}
	event := BatchPinEvent{
		Signer:     fmt.Sprintf("%s::%s", mspId, idString),
		Timestamp:  timestamp,
		Action:     action,
		Uuids:      "",
		BatchHash:  "",
		PayloadRef: payload,
		Contexts:   []string{},
	}
	bytes, err := json.Marshal(event)
	if err != nil {
		return fmt.Errorf("Failed to marshal event: %s", err)
	}
	ctx.GetStub().SetEvent("BatchPin", bytes)
	return nil
}

func (s *SmartContract) NetworkVersion() int {
	return 2
}
