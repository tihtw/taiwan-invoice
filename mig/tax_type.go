package mig

import "fmt"

type TaxTypeEnum string

const (
	// 應稅
	TaxTypeTaxable TaxTypeEnum = "1"
	// 零稅率
	TaxTypeZeroRated TaxTypeEnum = "2"
	// 免稅
	TaxTypeFreeTax TaxTypeEnum = "3"
	// 應稅(特種稅率)
	TaxTypeTaxableSpecial TaxTypeEnum = "4"
	// 混合應稅與免稅或零稅率 (限訊息 F0401 使用)
	TaxTypeMixed TaxTypeEnum = "9"
)

func (t TaxTypeEnum) Validate() error {
	switch t {
	case TaxTypeTaxable, TaxTypeZeroRated, TaxTypeFreeTax, TaxTypeTaxableSpecial, TaxTypeMixed:
		return nil
	}
	return fmt.Errorf("課稅別 (TaxType) 欄位格式錯誤")
}
