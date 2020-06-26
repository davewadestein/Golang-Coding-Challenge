package bank

import "errors"

type Customer struct {
	name string
	// reserved for future expansion
}

func (c *Customer) ChangeName(newname string) error {
	if newname == "" {
		return errors.New("Invalid customer name")
	}
	c.name = newname
	return nil
}
