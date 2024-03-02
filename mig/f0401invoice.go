package mig

import (
	"encoding/xml"
	"fmt"
)

type F0401Invoice struct {
	XMLName xml.Name `xml:"Invoice"`
	Text    string   `xml:",chardata"`
	Xmlns   string   `xml:"xmlns,attr"`

	Main    *InvoiceMain   `xml:"Main"`
	Details *InvoiceDetail `xml:"Details"`
	Amount  *InvoiceAmount `xml:"Amount"`
}

// NewF0401Invoice 會回傳一個新的F0401發票，輸入參數有賣方資訊 (seller) 與買方資訊 (buyer)以及發票明細
func NewF0401Invoice(seller *RoleDescription, buyer *RoleDescription, details []*ProductItem) (*F0401Invoice, error) {

	ret := &F0401Invoice{
		Xmlns: "urn:GEINV:eInvoiceMessage:F0401:4.0",
	}

	ret.Main = &InvoiceMain{
		Seller: seller,
		Buyer:  buyer,
	}

	ret.Details = &InvoiceDetail{
		ProductItem: details,
	}

	amount := &InvoiceAmount{}
	for _, item := range details {
		amount.SalesAmount += item.Amount
	}
	amount.TotalAmount = amount.SalesAmount
	ret.Amount = amount
	return ret, nil
}

func (invoice *F0401Invoice) Validate() error {
	if invoice.Main == nil {
		return fmt.Errorf("發票主要資訊為必填")
	}
	if err := invoice.Main.Validate(); err != nil {
		return err
	}
	if invoice.Details == nil {
		return fmt.Errorf("發票明細為必填")
	}
	if err := invoice.Details.Validate(); err != nil {
		return err
	}
	if invoice.Amount == nil {
		return fmt.Errorf("發票金額為必填")
	}
	if err := invoice.Amount.Validate(); err != nil {
		return err
	}
	return nil
}

func (f *F0401Invoice) Bytes() ([]byte, error) {
	return xml.Marshal(f)
}
