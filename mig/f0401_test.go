package mig

import (
	"fmt"
	"time"
)

func ExampleNewF0401Invoice() {
	// 這是一個範例
	// 這個範例是用來展示如何使用 NewF0401Invoice 來建立一個 F0401 發票物件
	// 這個範例會建立一個空的 F0401 發票物件

	// 首先，我們需要建立賣方資訊
	seller := NewSeller()
	seller.Identifier = "12345678"
	seller.Name = "網路書店"
	seller.Address = "台北市中正區和平西路一段 1 號"
	seller.PersonInCharge = "王小明"
	seller.EmailAddress = "wang@example.com"

	// 接著，我們需要建立買方資訊
	buyer := NewBuyer()
	buyer.Identifier = "87654321"
	buyer.Name = "網路購物者"
	buyer.Address = "台北市信義區信義路五段 7 號"
	buyer.PersonInCharge = "陳小美"
	buyer.EmailAddress = "mei@example.com"

	// 最後，我們需要建立發票明細
	details := []*F0401ProductItem{}
	item := NewF0401ProductItem("網紅小遙")
	item.RelateNumber = "A1234567890"
	item.Quantity = "1"
	item.Unit = "個"
	item.UnitPrice = "105"
	item.Amount = "105"
	details = append(details, item)

	item = NewF0401ProductItem("30m USB 3.0 延長線")
	item.RelateNumber = "A1234567891"
	item.Quantity = "2"
	item.Unit = "條"
	item.UnitPrice = "210"
	item.Amount = "420"
	details = append(details, item)

	// 現在，我們可以使用 NewF0401Invoice 來建立一個 F0401 發票物件
	invoice, err := NewF0401Invoice(seller, buyer, details)
	if err != nil {
		fmt.Println(err)
		return
	}
	invoice.Main.InvoiceNumber = "QQ18927486"
	invoice.SetDateAndTime(time.Date(2024, 3, 2, 11, 39, 40, 0, time.Local))
	invoice.Details.FillSequenceNumber()
	invoice.FillAmount()

	// 最後，我們可以檢查這個發票是否符合規範
	if err := invoice.Validate(); err != nil {
		fmt.Println(err)
		return
	}

	// Output:
	f, err := invoice.Bytes()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(f))

	// 這個範例會輸出一個符合 F0401 發票規範的 XML 字串
	// Output:
	// <Invoice xmlns="urn:GEINV:eInvoiceMessage:F0401:4.0">
	// 	<Main>
	// 		<InvoiceNumber>A1234567890</InvoiceNumber>
	// 		<InvoiceDate>20210101</InvoiceDate>
	// 		<InvoiceTime>00:00:00</InvoiceTime>
	// 		<Seller>
	// 			<Identifier>12345678</Identifier>
	// 			<Name>網路書店</Name>
	// 			<Address>台北市中正區和平西路一段 1 號</Address>
	// 			<PersonInCharge>王小明</PersonInCharge>
	// 			<EmailAddress>
	// 				<
	// 			</EmailAddress>
	// 		</Seller>
	// 		<Buyer>
	// 			<Identifier>87654321</Identifier>
	// 			<Name>網路購物者</Name>
	// 			<Address>台北市信義區信義路五段 7 號</Address>
	// 			<PersonInCharge>陳小美</PersonInCharge>
	// 			<EmailAddress>
	// 				<
	// 			</EmailAddress>
	// 		</Buyer>
	// 		<InvoiceType>07</InvoiceType>
	// 		<DonateMark>N</DonateMark>
	// 		<PrintMark>Y</PrintMark>
	// 	</Main>
	// 	<Details>
	// 		<SequenceNumber>1</SequenceNumber>
	// 		<ProductItem>
	// 			<Description>網紅小遙</Description>
	// 			<Quantity>1</Quantity>
	// 			<Unit>個</Unit>
	// 			<UnitPrice>100</UnitPrice>
	// 			<Amount>105</Amount>
	// 			<RelateNumber>A1234567890</RelateNumber>
	// 		</ProductItem>
	// 		<SequenceNumber>2</SequenceNumber>
	// 		<ProductItem>
	// 			<Description>30m USB 3.0 延長線</Description>
	// 			<Quantity>1</Quantity>
	// 			<Unit>條</Unit>
	// 			<UnitPrice>1550</UnitPrice>
	// 			<Amount>1628</Amount>
	// 			<RelateNumber>A1234567891</RelateNumber>
	// 		</ProductItem>
	// 	</Details>
	// 	<Amount>
	// 		<SalesAmount>1733</SalesAmount>
	// 		<TaxType>1</TaxType>
	// 		<TaxRate>0.05</TaxRate>
	// 		<TaxAmount>83</TaxAmount>
	// 		<TotalAmount>1816</TotalAmount>
	// 		<FreeTaxSalesAmount>0</FreeTaxSalesAmount>
	// 		<ZeroTaxSalesAmount>0</ZeroTaxSalesAmount>
	// 	</Amount>
	// </Invoice>

}
