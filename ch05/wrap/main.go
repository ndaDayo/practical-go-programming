package main

import (
	"fmt"
)

type transientError struct {
	err error
}

func (t transientError) Error() string {
	return fmt.Sprintf("transient error: %v", t.err)
}

func getTransactionAmount(transactionID string) (float32, error) {
	if len(transactionID) != 5 {
		return 0, fmt.Errorf("id is invalid: %s", transactionID)
	}

	amount, err := getTransactionAmountFromDB(transactionID)
	if err != nil {
		return 0, transientError{err: err}
	}

	return amount, nil
}

func getTransactionAmountFromDB(transactionID string) (float32, error) {
	var amount float32
	err := db.QueryRow("SELECT amount FROM transactions WHERE id = $1", transactionID).Scan(&amount)
	if err != nil {
		return 0, transientError{err: err}
	}

	return amount, nil
}
