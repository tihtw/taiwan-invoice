package mig

import (
	"fmt"
	"regexp"
	"strings"
)

// Mig 4.0 的圖 5-1, 5-3, 15-1 和 15-3 少放一個 ZeroTaxRateReason, GroupMark

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
	GroupMark   string `xml:"GroupMark,omitempty"`
	DonateMark  string `xml:"DonateMark"`

	ZeroTaxRateReason string `xml:"ZeroTaxRateReason,omitempty"`
	Reserved1         string `xml:"Reserved1,omitempty"`
	Reserved2         string `xml:"Reserved2,omitempty"`
}

type A0401InvoiceMain struct {
	InvoiceMain
}

type F0401InvoiceMain struct {
	InvoiceMain

	CarrierType       string `xml:"CarrierType,omitempty"`
	CarrierId1        string `xml:"CarrierId1,omitempty"`
	CarrierId2        string `xml:"CarrierId2,omitempty"`
	PrintMark         string `xml:"PrintMark"`
	NPOBAN            string `xml:"NPOBAN,omitempty"`
	RandomNumber      string `xml:"RandomNumber,omitempty"`
	BondedAreaConfirm string `xml:"BondedAreaConfirm,omitempty"`
}

func (block *InvoiceMain) Validate() error {
	if block.InvoiceNumber == "" {
		return fmt.Errorf("發票號碼 (InvoiceNumber) 為必填")
	}
	// TODO: validate InvoiceNumber in type of InvoiceNumberType

	if block.InvoiceDate == "" {
		return fmt.Errorf("發票開立日期 (InvoiceDate) 為必填")
	}
	// TODO: validate InvoiceDate in type of DateType

	if block.InvoiceTime == "" {
		return fmt.Errorf("發票開立時間 (InvoiceTime) 為必填")
	}
	// TODO: validate InvoiceTime in type of TimeType

	if block.Seller == nil {
		return fmt.Errorf("賣方資訊 (Seller) 為必填")
	}
	if err := block.Seller.Validate(); err != nil {
		return fmt.Errorf("賣方資訊 (Seller) 不符規範: %w", err)
	}

	if block.Buyer == nil {
		return fmt.Errorf("買方資訊 (Buyer) 為必填")
	}
	if err := block.Buyer.Validate(); err != nil {
		return fmt.Errorf("買方資訊 (Buyer) 不符規範: %w", err)
	}

	// TODO: validate BuyerRemark in type of BuyerRemarkEnum

	if len(block.MainRemark) > 200 {
		return fmt.Errorf("總備註 (MainRemark) 長度不得大於200個字元")
	}

	if block.CustomerClearanceMark != "" {
		// TODO: validate CustomerClearanceMark in CustomerClearanceMarkEnum
	}

	if len(block.Category) > 2 {
		return fmt.Errorf("沖帳別 (Category) 長度不得大於2個字元")
	}

	if len(block.RelateNumber) > 20 {
		return fmt.Errorf("相關號碼 (RelateNumber) 長度不得大於20個字元")
	}

	if block.InvoiceType == "" {
		return fmt.Errorf("發票類別 (InvoiceType) 為必填")
	}
	// TODO: validate InvoiceType in InvoiceTypeEnum

	if len(block.GroupMark) > 1 {
		return fmt.Errorf("彙開註記 (GroupMark) 長度不得大於1個字元")
	}

	if block.DonateMark == "" {
		return fmt.Errorf("捐贈註記 (DonateMark) 為必填")
	}

	if block.ZeroTaxRateReason == "" {
		return fmt.Errorf("零稅率原因 (ZeroTaxRateReason) 為必填")
	}

	if len(block.Reserved1) > 20 {
		return fmt.Errorf("保留欄位 (Reserved1) 長度不得大於20個字元")
	}

	if len(block.Reserved2) > 100 {
		return fmt.Errorf("保留欄位 (Reserved2) 長度不得大於100個字元")
	}

	return nil
}

func (block *A0401InvoiceMain) Validate() error {
	err := block.InvoiceMain.Validate()
	if err != nil {
		return err
	}
	return nil
}

func (block *F0401InvoiceMain) Validate() error {
	err := block.InvoiceMain.Validate()
	if err != nil {
		return err
	}

	// TODO: validate CarrierType in CarrierTypeEnum

	if len(block.CarrierId1) > 400 {
		return fmt.Errorf("載具顯碼 id (CarrierId1) 長度不得大於400個字元")
	}
	if strings.HasPrefix(block.CarrierId1, " ") ||
		strings.HasPrefix(block.CarrierId1, "　") ||
		strings.HasSuffix(block.CarrierId1, " ") ||
		strings.HasSuffix(block.CarrierId1, "　") {
		return fmt.Errorf("載具顯碼 id (CarrierId1) 前後不允許空白及全形空白")
	}

	if len(block.CarrierId2) > 400 {
		return fmt.Errorf("載具隱碼 id (CarrierId2) 長度不得大於400個字元")
	}
	if strings.HasPrefix(block.CarrierId2, " ") ||
		strings.HasPrefix(block.CarrierId2, "　") ||
		strings.HasSuffix(block.CarrierId2, " ") ||
		strings.HasSuffix(block.CarrierId2, "　") {
		return fmt.Errorf("載具隱碼 id (CarrierId2) 前後不允許空白及全形空白")
	}

	if block.PrintMark == "" {
		return fmt.Errorf("電子發票證明聯已列印註記 (PrintMark) 為必填")
	}
	if block.PrintMark != "Y" && block.PrintMark != "N" {
		return fmt.Errorf("電子發票證明聯已列印註記 (PrintMark) 應為 Y 或 N")
	}
	if block.PrintMark == "Y" {
		if block.CarrierType != "" ||
			block.CarrierId1 != "" ||
			block.CarrierId2 != "" ||
			block.DonateMark != "0" {
			return fmt.Errorf(
				"電子發票證明聯已列印註記 (PrintMark) 為 Y 時，載具類別號碼 (CarrierType)、載具顯碼 id (CarrierId1)、載具隱碼 id (CarrierId2) 必須為空白，捐贈註記 (DonateMark) 欄位為 0")

		}
	}

	if len(block.NPOBAN) > 10 {
		return fmt.Errorf("發票捐贈對象 (NPOBAN) 長度不得大於10個字元")
	}

	if len(block.RandomNumber) != 4 {
		return fmt.Errorf("發票防偽隨機碼 (RandomNumber) 長度必須為4個字元")
	}
	pattern := "[0-9]{4}"
	if match, _ := regexp.MatchString(pattern, block.RandomNumber); !match {
		return fmt.Errorf("發票防偽隨機碼 (RandomNumber) 應為4位數字")
	}

	// TODO: validate BondedAreaConfirm in BondedAreaConfirmEnum

	return nil
}
