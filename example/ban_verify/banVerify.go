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

	apiKey := os.Getenv("PICHU_EINVOICE_APIKEY")
	if apiKey == "" {
		fmt.Println("Please use export PICHU_EINVOICE_APIKEY={your_api_key} for setting api key")
		return
	}

	host := os.Getenv("INVOICE_CONNECTION_HOST")

	c := invoice.NewClient(appId)
	c.ApiKey = apiKey

	if host != "" {
		c.ConnectionHost = host
	}
	banUnitTpStatus, err := c.CheckBusinessAdministrationNumberExist("54834795")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("status:", banUnitTpStatus)
}
