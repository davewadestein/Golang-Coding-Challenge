package bank

import "testing"

var initialBalance = 100.0
var increment = 10.0

func TestGood(t *testing.T) {
	var bal float64
	a, err := CreateAccount("Foo Bar", initialBalance)

	if err != nil {
		t.Log("CreateAccount failed with " + err.Error())
		t.Fail()
	}
	t.Run("deposit", func(t *testing.T) {
		if bal, err = a.Deposit(increment); bal != initialBalance+increment {
			t.Log("Deposit failed with " + err.Error())
			t.Fail()
		}
	})
	t.Run("withdraw", func(t *testing.T) {
		if newbal, err := a.Withdraw(increment); newbal != bal-increment {
			t.Log("Withdraw failed with " + err.Error())
			t.Fail()
		}
	})
	t.Run("check", func(t *testing.T) {
		if newbal := a.CheckBalance(); newbal != a.balance {
			t.Log("CheckBalance failed")
			t.Fail()
		}
	})
}
