package mig

import "fmt"

type RoleDescription struct {
	Identifier      string `xml:"Identifier"`
	Name            string `xml:"Name"`
	Address         string `xml:"Address,omitempty"`
	PersonInCharge  string `xml:"PersonInCharge,omitempty"`
	TelephoneNumber string `xml:"TelephoneNumber,omitempty"`
	FacsimileNumber string `xml:"FacsimileNumber,omitempty"`
	EmailAddress    string `xml:"EmailAddress,omitempty"`
	CustomerNumber  string `xml:"CustomerNumber,omitempty"`
	RoleRemark      string `xml:"RoleRemark,omitempty"`
}

// Validate 檢查賣方資料是否符合規範
func (item *RoleDescription) Validate() error {
	if item.Identifier == "" {
		return fmt.Errorf("識別碼 (Identifier) 為必填")
	}
	// TODO: validate Identifier in type of BAN

	if item.Name == "" {
		return fmt.Errorf("名稱 (Name) 為必填")
	}
	if len(item.Name) < 1 {
		return fmt.Errorf("名稱 (Name) 長度不得小於1個字元")
	}
	if len(item.Name) > 60 {
		return fmt.Errorf("名稱 (Name) 長度不得大於60個字元")
	}

	if len(item.Address) > 100 {
		return fmt.Errorf("地址 (Address) 長度不得大於100")
	}

	if len(item.PersonInCharge) > 30 {
		return fmt.Errorf("負責人姓名 (PersonInCharge) 長度不得大於30個字元")
	}

	if len(item.TelephoneNumber) > 26 {
		return fmt.Errorf("電話號碼 (TelephoneNumber) 長度不得大於26個字元")
	}

	if len(item.FacsimileNumber) > 26 {
		return fmt.Errorf("傳真號碼 (FacsimileNumber) 長度不得大於26個字元")
	}

	if len(item.EmailAddress) > 400 {
		return fmt.Errorf("電子郵件地址 (EmailAddress) 長度不得大於400個字元")
	}

	if len(item.CustomerNumber) > 20 {
		return fmt.Errorf("客戶編號 (CustomerNumber) 長度不得大於20個字元")
	}

	if len(item.RoleRemark) > 40 {
		return fmt.Errorf("營業人角色註記 (RoleRemark) 長度不得大於40個字元")
	}

	return nil
}
