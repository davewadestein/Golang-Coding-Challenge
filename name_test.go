package bank

import "testing"

// Create a bankAccount, change the owner's name and then
// ensure it has been changed.

func TestName(t *testing.T) {
	a, err := CreateAccount("Foo Bar", initialBalance)
	if err != nil {
		t.Log("CreateAccount failed with " + err.Error())
		t.Fail()
	}
	t.Run("name change", func(t *testing.T) {
		a.ChangeName("New Name")
		if a.Customer.name != "New Name" {
			t.Log("ChangeName failed with " + a.Customer.name)
			t.Fail()
		}
	})
}
