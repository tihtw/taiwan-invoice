package mig

import (
	"encoding/xml"
	"fmt"
	"math/big"
	"time"
)

type F0401Invoice struct {
	XMLName xml.Name `xml:"Invoice"`
	Text    string   `xml:",chardata"`
	Xmlns   string   `xml:"xmlns,attr"`

	Main    *F0401InvoiceMain   `xml:"Main"`
	Details *F0401InvoiceDetail `xml:"Details"`
	Amount  *F0401InvoiceAmount `xml:"Amount"`
}

// NewF0401Invoice 會回傳一個新的F0401發票，輸入參數有賣方資訊 (seller) 與買方資訊 (buyer)以及發票明細
func NewF0401Invoice(seller *Seller, buyer *Buyer, details []*F0401ProductItem) (*F0401Invoice, error) {

	ret := &F0401Invoice{
		Xmlns: "urn:GEINV:eInvoiceMessage:F0401:4.0",
	}

	ret.Main = &F0401InvoiceMain{
		InvoiceMain: InvoiceMain{
			InvoiceDate: time.Now().Format("20060102"),
			InvoiceTime: time.Now().Format("15:04:05"),
			Seller:      seller,
			Buyer:       buyer,
			InvoiceType: InvoiceTypeGeneral,
			DonateMark:  DonateMarkNo,
		},
		PrintMark: "Y",
	}

	ret.Details = &F0401InvoiceDetail{
		ProductItem: details,
	}

	amount := &F0401InvoiceAmount{
		InvoiceAmount: InvoiceAmount{
			TaxRate: "0.05",
			TaxType: TaxTypeTaxable,
		},
		FreeTaxSalesAmount: "0",
		ZeroTaxSalesAmount: "0",
	}

	ret.Amount = amount
	return ret, nil
}

func (invoice *F0401Invoice) SetDateAndTime(t time.Time) {
	invoice.Main.InvoiceDate = t.Format("20060102")
	invoice.Main.InvoiceTime = t.Format("15:04:05")
}

// IsB2C 會回傳發票是否為B2C發票，判斷根據 (表 4-7 BAN 資料元規格)
func (invoice *F0401Invoice) IsB2C() bool {
	return invoice.Main.Buyer.Identifier == "0000000000"
}

// FillAmount 會根據發票明細填入 SalesAmount, TaxAmount, TotalAmount
// SalesAmount 為明細的金額加總
// 當此發票為 B2B 發票時，TaxAmount 為 SalesAmount * TaxRate (四捨五入至整數)
// 當此發票為 B2C 發票時，TaxAmount 為 0
// TotalAmount 為 SalesAmount + TaxAmount
func (invoice *F0401Invoice) FillAmount() error {
	salesAmount := new(big.Float)
	taxAmount := new(big.Float)
	for _, item := range invoice.Details.ProductItem {
		amount := new(big.Float)
		amount, ok := amount.SetString(item.Amount)
		if !ok {
			return fmt.Errorf("parse amount failed")
		}
		salesAmount.Add(salesAmount, amount)
	}
	invoice.Amount.SalesAmount = salesAmount.Text('f', 0)

	if invoice.IsB2C() {
		invoice.Amount.TaxAmount = "0"
	} else {
		taxRate := new(big.Float)
		taxRate, ok := taxRate.SetString(invoice.Amount.TaxRate)
		if !ok {
			return fmt.Errorf("parse tax rate failed")
		}
		taxAmount := new(big.Float).Mul(salesAmount, taxRate)

		// 四捨五入至整數
		taxAmount = taxAmount.Add(taxAmount, big.NewFloat(0.5))
		bint, _ := taxAmount.Int(nil)
		invoice.Amount.TaxAmount = bint.String()
	}
	taxAmount.SetString(invoice.Amount.TaxAmount)
	totalAmount := new(big.Float)
	totalAmount.Add(salesAmount, taxAmount)
	invoice.Amount.TotalAmount = totalAmount.Text('f', 0)
	return nil
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
	return xml.MarshalIndent(f, "", "  ")
}
