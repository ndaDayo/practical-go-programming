package main

import (
	"errors"
	"fmt"
	"net/http"
)

type transientError struct {
	err error
}

func (t transientError) Error() string {
	return fmt.Sprintf("transient error: %v", t.err)
}

func GetTransactionAmountHandler(w http.ResponseWriter, r *http.Request) (float32, error) {
	transactionID := r.URL.Query().Get("transaction")
	amount, err := getTransactionAmount(transactionID)
	if err != nil {
		if errors.As(err, &transientError{}) {
			http.Error(w, err.Error(), http.StatusServiceUnavailable)
			return 0, err
		}
		http.Error(w, err.Error(), http.StatusBadRequest)
		return 0, err
	}

	return amount, nil
}

func getTransactionAmount(transactionID string) (float32, error) {
	if len(transactionID) != 5 {
		return 0, fmt.Errorf("id is invalid: %s", transactionID)
	}

	amount, err := getTransactionAmountFromDB(transactionID)
	if err != nil {
		return 0, fmt.Errorf("failed to get transaction %s: %w", transactionID, err)
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
