package mig

import "fmt"

type A0101InvoiceDetail struct {
	Text        string                `xml:",chardata"`
	ProductItem []*InvoiceProductItem `xml:"ProductItem"`
}

// Deprecated in Mig 4.0
type C0401InvoiceDetail struct {
	Text        string                `xml:",chardata"`
	ProductItem []*InvoiceProductItem `xml:"ProductItem"`
}

type F0401InvoiceDetail struct {
	Text        string              `xml:",chardata"`
	ProductItem []*F0401ProductItem `xml:"ProductItem"`
}

func (block *A0101InvoiceDetail) Validate() error {
	if len(block.ProductItem) == 0 {
		return nil
	}
	if len := len(block.ProductItem); len > 9999 {
		return fmt.Errorf("發票明細項目數量不得超過9999個，目前為%d", len)
	}

	for _, item := range block.ProductItem {
		if err := item.Validate(); err != nil {
			return err
		}
	}
	return nil
}

func (block *F0401InvoiceDetail) Validate() error {
	if len(block.ProductItem) == 0 {
		return nil
	}
	if len := len(block.ProductItem); len > 9999 {
		return fmt.Errorf("發票明細項目數量不得超過9999個，目前為%d", len)
	}

	for _, item := range block.ProductItem {
		if err := item.Validate(); err != nil {
			return err
		}
	}
	return nil
}
