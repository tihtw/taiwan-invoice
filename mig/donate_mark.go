package mig

import "fmt"

// DonateMarkEnum 捐贈註記 (表 4-11)
type DonateMarkEnum string

const (
	// 非捐贈發票
	DonateMarkNo DonateMarkEnum = "0"
	// 捐贈發票
	DonateMarkYes DonateMarkEnum = "1"
)

// Validate 檢查捐贈註記是否符合規範
func (t DonateMarkEnum) Validate() error {
	switch t {
	case DonateMarkNo, DonateMarkYes:
		return nil
	}
	return fmt.Errorf("捐贈註記 (DonateMark) 欄位格式錯誤")
}
