package mig

import "fmt"

type ProductItem struct {
	Text           string `xml:",chardata"`
	Description    string `xml:"Description"`
	Quantity       string `xml:"Quantity"`
	Unit           string `xml:"Unit,omitempty"`
	UnitPrice      string `xml:"UnitPrice"`
	TaxType        string `xml:"TaxType"`
	Amount         string `xml:"Amount"`
	SequenceNumber string `xml:"SequenceNumber"`
	Remark         string `xml:"Remark,omitempty"`
	RelateNumber   string `xml:"RelateNumber,omitempty"`
}

func (item *ProductItem) Validate() error {
	if item.Description == "" {
		return fmt.Errorf("品名 (Description) 為必填")
	}
	if len(item.Description) < 1 {
		return fmt.Errorf("品名 (Description) 長度不得小於1個字元")
	}
	if len(item.Description) > 500 {
		return fmt.Errorf("品名 (Description) 長度不得大於500個字元")
	}

	if item.Quantity == "" {
		return fmt.Errorf("數量 (Quantity) 為必填")
	}
	// TODO: check Quantity is a number and totalDigits <= 20, fractionDigits <= 7

	if len(item.Unit) > 6 {
		return fmt.Errorf("單位 (Unit) 長度不得大於6個字元")
	}

	if item.UnitPrice == "" {
		return fmt.Errorf("單價 (UnitPrice) 為必填")
	}
	// TODO: check UnitPrice is a number and totalDigits <= 20, fractionDigits <= 7

	if item.TaxType == "" {
		return fmt.Errorf("課稅別 (TaxType) 為必填")
	}

	if item.Amount == "" {
		return fmt.Errorf("金額 (Amount) 為必填")
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
