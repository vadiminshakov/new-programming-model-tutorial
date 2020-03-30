package main

import "fmt"

// UnknownTransactionHandler returns a shim error
// with details of a bad transaction request
func UnknownTransactionHandler(ctx CustomTransactionContextInterface) error {
	fcn, args := ctx.GetStub().GetFunctionAndParameters()
	return fmt.Errorf("Invalid function %s passed with args %v", fcn, args)
}
