package mig

import "fmt"

// 在 Mig 4.0 裡面的 Invoice/Amount 有兩種定義，一個是 A0401 開立發泡
// 另一個是 F0401 平台存證開立發票訊息，相同欄位名稱的驗證規則不一定相同
// 舉例來說，A0401 的 SalesAmount 的 fractionDigits 是 0
// 但是 F0401 的 SalesAmount 的 fractionDigits 是 7
// 相同部分的驗證會在 InvoiceAmount 物件被驗證，如果規則有不同時則會被拆分驗證

type InvoiceAmount struct {
	Text                   string `xml:",chardata"`
	SalesAmount            string `xml:"SalesAmount"`
	TaxType                string `xml:"TaxType"`
	TaxRate                string `xml:"TaxRate"`
	TaxAmount              string `xml:"TaxAmount"`
	TotalAmount            string `xml:"TotalAmount"`
	DiscountAmount         string `xml:"DiscountAmount,omitempty"`
	OriginalCurrencyAmount string `xml:"OriginalCurrencyAmount,omitempty"`
	ExchangeRate           string `xml:"ExchangeRate,omitempty"`
	Currency               string `xml:"Currency,omitempty"`
}

type A0401InvoiceAmount struct {
	InvoiceAmount
}

// Deprecated in Mig 4.0
type C0401InvoiceAmount struct {
	InvoiceAmount

	FreeTaxSalesAmount string `xml:"FreeTaxSalesAmount"`
	ZeroTaxSalesAmount string `xml:"ZeroTaxSalesAmount"`
}

type F0401InvoiceAmount struct {
	InvoiceAmount

	FreeTaxSalesAmount string `xml:"FreeTaxSalesAmount"`
	ZeroTaxSalesAmount string `xml:"ZeroTaxSalesAmount"`
}

func (block *InvoiceAmount) Validate() error {
	if block.SalesAmount == "" {
		return fmt.Errorf("銷售額 (SalesAmount) 為必填")
	}

	if block.TaxType == "" {
		return fmt.Errorf("課稅別 (TaxType) 為必填")
	}
	// TODO: validate TaxType in TaxTypeEnum

	if block.TaxRate == "" {
		return fmt.Errorf("稅率 (TaxRate) 為必填")
	}
	// TODO: validate TaxRate in TaxRateEnum

	if block.TaxAmount == "" {
		return fmt.Errorf("營業稅額 (TaxAmount) 為必填")
	}
	// TODO: validate TaxAmount in type of decimal(20,0)

	if block.TotalAmount == "" {
		return fmt.Errorf("總金額 (TotalAmount) 為必填")
	}

	if block.OriginalCurrencyAmount != "" {
		// TODO: validate OriginalCurrencyAmount in type of decimal(20,7)
	}

	if block.ExchangeRate != "" {
		// TODO: validate ExchangeRate in type of decimal(13,5)
	}

	if block.Currency != "" {
		// TODO: validate Currency in CurrencyCodeEnum
	}

	return nil
}

func (block *A0401InvoiceAmount) Validate() error {
	err := block.InvoiceAmount.Validate()
	if err != nil {
		return err
	}
	// TODO validate SalesAmount in type of decimal(20,0)
	// TODO validate TaxAmount in type of decimal(20,0)
	// TODO validate TotalAmount in type of decimal(20,0)
	// TODO validate DiscountAmount in type of decimal(20,0)

	return nil
}

func (block *F0401InvoiceAmount) Validate() error {
	err := block.InvoiceAmount.Validate()
	if err != nil {
		return err
	}

	// TODO validate SalesAmount in type of decimal(20,7)
	if block.FreeTaxSalesAmount == "" {
		return fmt.Errorf("免稅銷售額 (FreeTaxSalesAmount) 為必填")
	}
	// TODO validate FreeTaxSalesAmount in type of decimal(20,7)

	if block.ZeroTaxSalesAmount == "" {
		return fmt.Errorf("零稅率銷售額 (ZeroTaxSalesAmount) 為必填")
	}
	// TODO validate ZeroTaxSalesAmount in type of decimal(20,7)

	// TODO validate TaxAmount in type of decimal(20,0)
	// TODO validate TotalAmount in type of decimal(20,7)
	// TODO validate DiscountAmount in type of decimal(20,7)
	// TODO validate OriginalCurrencyAmount in type of decimal(20,7)
	// TODO validate ExchangeRate in type of decimal(13,5)
	// TODO validate Currency in CurrencyCodeEnum

	return nil
}
