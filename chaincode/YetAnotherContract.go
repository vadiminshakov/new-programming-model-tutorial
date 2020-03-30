package main

import (
	"fmt"

	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

type YetAnotherContract struct {
	contractapi.Contract
}

func (sc *YetAnotherContract) SayHi(ctx contractapi.TransactionContextInterface) string {
	return "Hi there"
}

func After(ctx contractapi.TransactionContextInterface, beforeValue interface{}) error {
	fmt.Println(ctx.GetStub().GetTxID())
	fmt.Println("beforeValue", beforeValue)
	return nil
}
