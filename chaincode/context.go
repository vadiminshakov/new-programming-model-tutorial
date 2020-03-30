package main

import (
	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

// CustomTransactionContextInterface interface to define interaction with custom transaction context
type CustomTransactionContextInterface interface {
	contractapi.TransactionContextInterface
	GetData() []byte
	SetData([]byte)
}

// CustomTransactionContext adds methods of storing and retrieving additional data for use
// with before and after transaction hooks
type CustomTransactionContext struct {
	contractapi.TransactionContext
	data []byte
}

// GetData return set data
func (ctc *CustomTransactionContext) GetData() []byte {
	return ctc.data
}

// SetData provide a value for data
func (ctc *CustomTransactionContext) SetData(data []byte) {
	ctc.data = data
}
