package main

// 手機條碼驗證

import (
	"fmt"
	"github.com/tihtw/taiwan-invoice"
	"os"
)

func main() {

	appId := os.Getenv("PICHU_EINVOICE_APPID")
	if appId == "" {
		fmt.Println("Please use export PICHU_EINVOICE_APPID={your_app_id} for setting appId")
		return
	}

	host := os.Getenv("INVOICE_CONNECTION_HOST")

	c := invoice.NewClient(appId)

	if host != "" {
		c.ConnectionHost = host
	}
	isExist, err := c.CheckMobilePhoneExist("/H33M4PR")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("is exist:", isExist)
}
