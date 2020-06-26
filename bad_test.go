package bank

import (
	"fmt"
	"testing"
)

func TestBad(t *testing.T) {
	a, err := CreateAccount("Foo Bar", initialBalance)

	if err != nil {
		t.Log("CreateAccount failed with " + err.Error())
		t.Fail()
	}
	t.Run("deposit", func(t *testing.T) {
		if _, err = a.Deposit(0.0); err == nil {
			t.Log("Deposit(0.0) did not fail")
			t.Fail()
		}
	})
	t.Run("withdraw", func(t *testing.T) {
		if _, err := a.Withdraw(0.0); err == nil {
			t.Log("Withdraw(0.0) did not fail ")
			t.Fail()
		}
	})
	t.Run("withdraw", func(t *testing.T) {
		if _, err := a.Withdraw(a.balance + 1); err != nil {
			t.Log(fmt.Sprintf("Withdraw(%.2f) did not fail ", a.balance+1))
			t.Fail()
		}
	})
}
