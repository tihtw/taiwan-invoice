package mig

import "fmt"

type Seller struct {
	RoleDescription
}

func NewSeller() *Seller {
	return &Seller{}
}

func (seller *Seller) Validate() error {
	err := seller.RoleDescription.Validate()
	if err != nil {
		return err
	}

	if seller.Address == "" {
		return fmt.Errorf("賣方地址欄位 (Address) 為必填")
	}
	return nil
}
