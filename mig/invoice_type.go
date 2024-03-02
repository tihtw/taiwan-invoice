package mig

import "fmt"

// InvoiceTypeEnum 電子發票類別 (表 4-4)
type InvoiceTypeEnum string

const (
	// 一般稅額計算之電子發票
	InvoiceTypeGeneral InvoiceTypeEnum = "07"
	// 特種稅額計算之電子發票
	InvoiceTypeSpecial InvoiceTypeEnum = "08"
)

// Validate 檢查發票類別是否符合規範
func (t InvoiceTypeEnum) Validate() error {
	switch t {
	case InvoiceTypeGeneral, InvoiceTypeSpecial:
		return nil
	}
	return fmt.Errorf("發票類別 (InvoiceType) 欄位格式錯誤")
}
