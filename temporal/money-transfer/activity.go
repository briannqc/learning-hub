package app

import (
	"context"
	"fmt"
	"log"
)

func Withdraw(ctx context.Context, data PaymentDetails) (string, error) {
	log.Printf("Withdrawing $%d from account %s.\n", data.Amount, data.SourceAccount)
	referenceID := fmt.Sprintf("%s-withdraw", data.ReferenceID)
	bank := BankingService{"bank-api.example.com"}
	confirmation, err := bank.Withdraw(data.SourceAccount, data.Amount, referenceID)
	return confirmation, err
}

func Deposit(ctx context.Context, data PaymentDetails) (string, error) {
	log.Printf("Depositing $%d to account %s.\n", data.Amount, data.TargetAccount)
	referenceID := fmt.Sprintf("%s-deposit", data.ReferenceID)
	bank := BankingService{"bank-api.example.com"}
	confirmation, err := bank.Deposit(data.TargetAccount, data.Amount, referenceID)
	return confirmation, err
}

func Refund(ctx context.Context, data PaymentDetails) (string, error) {
	log.Printf("Refunding $%d to account %s.\n", data.Amount, data.SourceAccount)
	referenceID := fmt.Sprintf("%s-refund", data.ReferenceID)
	bank := BankingService{"bank-api.example.com"}
	confirmation, err := bank.Deposit(data.SourceAccount, data.Amount, referenceID)
	return confirmation, err
}
