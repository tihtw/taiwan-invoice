package mig

import "fmt"

// ProductItem 有六個地方會出現:
// 1. A0101Invoice.Details.ProductItem (表 5-5)
// 2. B0101Allowance.Details.ProductItem (表 7-5)
// 3. F0401Invoice.Details.ProductItem (表 15-5)
// 4. G0401Allowance.Details.ProductItem (表 18-5)
// 5. InputInvoice.Details.ProductItem (表 23-4)
// 6. InputAllowance.Details.ProductItem (表 24-5)

type ProductItem struct {
	Text      string `xml:",chardata"`
	Quantity  string `xml:"Quantity"`
	Unit      string `xml:"Unit,omitempty"`
	UnitPrice string `xml:"UnitPrice"`

	Amount string `xml:"Amount"`
}

type InvoiceProductItem struct {
	ProductItem

	Description string `xml:"Description"`

	TaxType TaxTypeEnum `xml:"TaxType"`

	SequenceNumber string `xml:"SequenceNumber"`
	Remark         string `xml:"Remark,omitempty"`
	RelateNumber   string `xml:"RelateNumber,omitempty"`
}

type AllowanceProductItem struct {
	ProductItem

	OriginalInvoiceDate     string `xml:"OriginalInvoiceDate"`
	OriginalInvoiceNumber   string `xml:"OriginalInvoiceNumber"`
	OriginalSequenceNumber  string `xml:"OriginalSequenceNumber,omitempty"`
	OriginalDescription     string `xml:"OriginalDescription"`
	AllowanceSequenceNumber string `xml:"AllowanceSequenceNumber"`

	TaxType TaxTypeEnum `xml:"TaxType"`
}

type F0401ProductItem struct {
	InvoiceProductItem
}

func NewF0401ProductItem(description string) *F0401ProductItem {
	return &F0401ProductItem{
		InvoiceProductItem{
			Description: description,
			TaxType:     TaxTypeTaxable,
		},
	}
}

func (item *ProductItem) Validate() error {

	if item.Quantity == "" {
		return fmt.Errorf("數量 (Quantity) 為必填")
	}
	// TODO: check Quantity in type of decimal(20,7)

	if len(item.Unit) > 6 {
		return fmt.Errorf("單位 (Unit) 長度不得大於6個字元")
	}

	if item.UnitPrice == "" {
		return fmt.Errorf("單價 (UnitPrice) 為必填")
	}
	// TODO: check UnitPrice in type of decimal(20,7)

	if item.Amount == "" {
		return fmt.Errorf("金額 (Amount) 為必填")
	}
	// TODO: check Amount in type of decimal(20,7)

	// TODO: check amount = quantity * unit price * (1 + tax rate)

	return nil
}

func (item *InvoiceProductItem) Validate() error {
	err := item.ProductItem.Validate()
	if err != nil {
		return err
	}

	if item.Description == "" {
		return fmt.Errorf("品名 (Description) 為必填")
	}

	if item.Quantity == "" {
		return fmt.Errorf("數量 (Quantity) 為必填")
	}
	// TODO: check Quantity in type of decimal(20,7)

	if item.TaxType == "" {
		return fmt.Errorf("課稅別 (TaxType) 為必填")
	}
	err = item.TaxType.Validate()
	if err != nil {
		return err
	}

	if item.SequenceNumber == "" {
		return fmt.Errorf("明細排列序號 (SequenceNumber) 為必填")
	}
	if len(item.SequenceNumber) > 4 {
		return fmt.Errorf("明細排列序號 (SequenceNumber) 長度不得大於4個字元")
	}
	if len(item.SequenceNumber) < 1 {
		return fmt.Errorf("明細排列序號 (SequenceNumber) 長度不得小於1個字元")
	}
	// 規範中並沒要求一定要是數字

	if len(item.Remark) > 120 {
		return fmt.Errorf("單一欄位備註 (Remark) 長度不得大於120個字元")
	}

	if len(item.RelateNumber) > 50 {
		return fmt.Errorf("相關號碼 (RelateNumber) 長度不得大於50個字元")
	}
	return nil
}

func (item *AllowanceProductItem) Validate() error {
	err := item.ProductItem.Validate()
	if err != nil {
		return err
	}

	if item.OriginalInvoiceDate == "" {
		return fmt.Errorf("原發票日期 (OriginalInvoiceDate) 為必填")
	}

	if item.OriginalInvoiceNumber == "" {
		return fmt.Errorf("原發票號碼 (OriginalInvoiceNumber) 為必填")
	}

	if len(item.OriginalSequenceNumber) > 4 {
		return fmt.Errorf("原明細排列序號 (OriginalSequenceNumber) 長度不得大於4個字元")
	}

	if item.OriginalDescription == "" {
		return fmt.Errorf("原品名 (OriginalDescription) 為必填")
	}
	if len(item.OriginalDescription) > 500 {
		return fmt.Errorf("原品名 (OriginalDescription) 長度不得大於500個字元")
	}

	if item.AllowanceSequenceNumber == "" {
		return fmt.Errorf("折讓證明單號 (AllowanceSequenceNumber) 為必填")
	}

	if item.TaxType == "" {
		return fmt.Errorf("課稅別 (TaxType) 為必填")
	}
	err = item.TaxType.Validate()
	if err != nil {
		return err
	}
	return nil
}

func (item *F0401ProductItem) Validate() error {
	err := item.InvoiceProductItem.Validate()
	if err != nil {
		return err
	}
	if len(item.Description) > 500 {
		return fmt.Errorf("品名 (Description) 長度不得大於500個字元")
	}

	return nil
}
