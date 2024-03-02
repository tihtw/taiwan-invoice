package mig

import "fmt"

func NewSeller() *RoleDescription {
	return &RoleDescription{}
}

// Validate 檢查賣方資料是否符合規範
func (seller *RoleDescription) Validate() error {
	if seller.Identifier == "" {
		return fmt.Errorf("賣方識別碼(統一編號)為必填")
	}
	if len(seller.Identifier) > 10 {
		// 通常統一編號長度為8個字元，但是規範最大值為10。
		return fmt.Errorf("賣方識別碼(統一編號)長度不得大於10個字元")
	}

	if seller.Name == "" {
		return fmt.Errorf("賣方營業人名稱為必填")
	}
	if len(seller.Name) < 1 {
		return fmt.Errorf("賣方營業人名稱長度不得小於1個字元")
	}
	if len(seller.Name) > 60 {
		return fmt.Errorf("賣方營業人名稱長度不得大於60個字元")
	}

	if seller.Address == "" {
		return fmt.Errorf("賣方地址欄位為必填")
	}
	if len(seller.Address) > 100 {
		return fmt.Errorf("賣方營業地址長度不得大於100")
	}

	if len(seller.PersonInCharge) > 30 {
		return fmt.Errorf("賣方負責人姓名長度不得大於30個字元")
	}

	if len(seller.TelephoneNumber) > 26 {
		return fmt.Errorf("賣方電話號碼長度不得大於26個字元")
	}

	if len(seller.FacsimileNumber) > 26 {
		return fmt.Errorf("賣方傳真號碼長度不得大於26個字元")
	}

	if len(seller.EmailAddress) > 400 {
		return fmt.Errorf("賣方電子郵件地址長度不得大於400個字元")
	}

	if len(seller.CustomerNumber) > 20 {
		return fmt.Errorf("賣方客戶編號長度不得大於20個字元")
	}

	if len(seller.RoleRemark) > 40 {
		return fmt.Errorf("賣方營業人角色註記長度不得大於40個字元")
	}

	return nil
}
