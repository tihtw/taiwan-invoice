package mig

import "fmt"

type InvoiceDetail struct {
	Text        string         `xml:",chardata"`
	ProductItem []*ProductItem `xml:"ProductItem"`
}

func (block *InvoiceDetail) Validate() error {
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
