package main

import (
	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

func main() {
	simpleContract := new(SimpleContract)

	simpleContract.TransactionContextHandler = new(CustomTransactionContext)
	simpleContract.BeforeTransaction = GetWorldState
	simpleContract.UnknownTransaction = UnknownTransactionHandler

	yetAnotherContract := new(YetAnotherContract)
	yetAnotherContract.AfterTransaction = After

	cc, err := contractapi.NewChaincode(simpleContract, yetAnotherContract)

	if err != nil {
		panic(err.Error())
	}

	if err := cc.Start(); err != nil {
		panic(err.Error())
	}
}
