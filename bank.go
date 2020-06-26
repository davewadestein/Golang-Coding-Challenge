package bank

import (
	"errors"
	"fmt"
)

// Config options for bank package
var bankConfig struct {
	minOpeningBalance float64
	// reserved for future expansion
}

// init() will run when the package is imported
func init() {
	bankConfig.minOpeningBalance = 25.0
}

type bankAccount struct {
	Customer
	balance float64
}

func (b bankAccount) String() string {
	return fmt.Sprintf("Name: %s\nBalance: £%.2f\n", b.name, b.balance)
}

var invalidAmountError = errors.New("Invalid amount!")

// Helper function to create a BankAccount object...basically
// a constructor that returns an init'ed BankAccount.
func CreateAccount(name string, balance float64) (*bankAccount, error) {
	if balance < bankConfig.minOpeningBalance {
		return nil, invalidAmountError
	}
	return &bankAccount{Customer{name}, balance}, nil
}

// Change account owner's name
func (b *bankAccount) ChangeName(newname string) error {
	// We could possibly do something different here, as we might
	// have more restricted rules for changing a bank customer's name.
	return b.Customer.ChangeName(newname)
}

// Deposit amount, error if amount <= 0
func (b *bankAccount) Deposit(amount float64) (float64, error) {
	if amount > 0.0 {
		b.balance += amount
		return b.balance, nil
	}
	return 0.0, invalidAmountError
}

var overdrawnError = errors.New("You can't withdraw that much, you'd be overdrawn!")

// Withdraw amount
// Error if overdrawn
// Error if amount <= 0
func (b *bankAccount) Withdraw(amount float64) (float64, error) {
	if amount > 0.0 {
		if amount <= b.balance {
			b.balance -= amount
			return b.balance, nil
		}
		return 0.0, overdrawnError
	}
	return 0.0, invalidAmountError
}

// Any error cases here?
func (b *bankAccount) CheckBalance() float64 {
	return b.balance
}

func MakeJointAccount(one, two *bankAccount) (*bankAccount, error) {
	// What are the error cases here?
	new := &bankAccount{Customer{one.name + " + " + two.name}, one.balance + two.balance - 5.95}
	one.balance, two.balance = 0.0, 0.0
	return new, nil
}
