package mig

import "fmt"

// Mig 4.0 的圖 5-1 和 15-1 少放一個 ZeroTaxRateReason

type InvoiceMain struct {
	InvoiceNumber string           `xml:"InvoiceNumber"`
	InvoiceDate   string           `xml:"InvoiceDate"`
	InvoiceTime   string           `xml:"InvoiceTime"`
	Seller        *RoleDescription `xml:"Seller"`
	Buyer         *RoleDescription `xml:"Buyer"`

	BuyerRemark           string `xml:"BuyerRemark,omitempty"`
	MainRemark            string `xml:"MainRemark,omitempty"`
	CustomerClearanceMark string `xml:"CustomerClearanceMark,omitempty"`
	Category              string `xml:"Category,omitempty"`
	RelateNumber          string `xml:"RelateNumber,omitempty"`

	InvoiceType string `xml:"InvoiceType"`
	DonateMark  string `xml:"DonateMark"`

	CarrierType       string `xml:"CarrierType,omitempty"`
	CarrierId1        string `xml:"CarrierId1,omitempty"`
	CarrierId2        string `xml:"CarrierId2,omitempty"`
	PrintMark         string `xml:"PrintMark"`
	NPOBAN            string `xml:"NPOBAN,omitempty"`
	RandomNumber      string `xml:"RandomNumber,omitempty"`
	BondedAreaConfirm string `xml:"BondedAreaConfirm,omitempty"`

	ZeroTaxRateReason string `xml:"ZeroTaxRateReason,omitempty"`
	Reserved1         string `xml:"Reserved1,omitempty"`
	Reserved2         string `xml:"Reserved2,omitempty"`
}

type A0401InvoiceMain struct {
	InvoiceMain
}

func (block *InvoiceMain) Validate() error {
	if block.InvoiceNumber == "" {
		return fmt.Errorf("發票號碼為必填")
	}
	// TODO: validate InvoiceNumber in type of InvoiceNumberType

	if block.InvoiceDate == "" {
		return fmt.Errorf("發票日期為必填")
	}
	// TODO: validate InvoiceDate in type of DateType

	if block.InvoiceTime == "" {
		return fmt.Errorf("發票時間為必填")
	}
	// TODO: validate InvoiceTime in type of TimeType

	if block.Seller == nil {
		return fmt.Errorf("賣方為必填")
	}
	if err := block.Seller.Validate(); err != nil {
		return fmt.Errorf("賣方資料不符規範: %w", err)
	}

	if block.Buyer == nil {
		return fmt.Errorf("買方為必填")
	}
	if err := block.Buyer.Validate(); err != nil {
		return fmt.Errorf("買方資料不符規範: %w", err)
	}

	// TODO: validate BuyerRemark in type of BuyerRemarkEnum

	if len(block.MainRemark) > 200 {
		return fmt.Errorf("發票主要註記長度不得大於200個字元")
	}
	return nil
}
