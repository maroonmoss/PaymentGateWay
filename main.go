package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type PrompayService struct{}

type TransactionReq struct {
	AccNo string `json:"accNo"`
	Id    string `json:"id"`
}

type TransactionRes struct {
	Name     string `json:"name"`
	AccNo    string `json:"accNo"`
	Verified bool   `json:"verified"`
}

func (p *PrompayService) VerifyTransaction(accNo, id string) (TransactionRes, error) {
	return TransactionRes{
		Name:     "test",
		AccNo:    accNo,
		Verified: true,
	}, nil
}

func (p *PrompayService) ConfirmTransaction(transaction TransactionReq) (string, error) {
	return "OK", nil
}

func verifyTransactionHandler(w http.ResponseWriter, r *http.Request) {
	var request TransactionReq
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	prompay := &PrompayService{}
	transaction, err := prompay.VerifyTransaction(request.AccNo, request.Id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(transaction)
}

func confirmTransactionHandler(w http.ResponseWriter, r *http.Request) {
	var transaction TransactionReq
	if err := json.NewDecoder(r.Body).Decode(&transaction); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	prompay := &PrompayService{}
	result, err := prompay.ConfirmTransaction(transaction)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, `{"result":"%s"}`, result)
}

func main() {
	http.HandleFunc("/verify-transaction", verifyTransactionHandler)
	http.HandleFunc("/confirm-transaction", confirmTransactionHandler)

	fmt.Println("Server is running at http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
