package mig

import (
	"io/ioutil"
	"testing"
)

func TestMarshalC0401(t *testing.T) {
	tc, _ := ioutil.ReadFile("testcases/C0401/01.xml")

	actual, _ := NewC0401Invoice(tc)

	expected := C0401Invoice{
		Main: &C0401InvoiceMain{
			InvoiceMain: InvoiceMain{
				InvoiceNumber: "AA00000000",
				InvoiceDate:   "20060102",
				InvoiceTime:   "15:04:05",
				Seller: &Seller{
					RoleDescription: RoleDescription{
						Identifier:     "54834795",
						Name:           "台灣智慧家庭股份有限公司",
						Address:        "Address",
						PersonInCharge: "PersonInCharge",
						EmailAddress:   "example@example.com",
					},
				},
				Buyer: &Buyer{
					RoleDescription: RoleDescription{
						Identifier: "0000000000",
						Name:       "Buyer Name",
					},
				},
				InvoiceType: "07",
				DonateMark:  "0",
			},
			CarrierType:  "EJ1507",
			CarrierId1:   "CarrierId1",
			CarrierId2:   "CarrierId2",
			PrintMark:    "N",
			RandomNumber: "1031",
		},
		Details: &C0401InvoiceDetail{
			ProductItem: []*InvoiceProductItem{
				{
					Description: "網紅小遙 回饋問卷的早鳥們 享53折優惠",
					ProductItem: ProductItem{
						Quantity:  "1",
						UnitPrice: "1650",
						Amount:    "1650",
					},
					RelateNumber:   "A735632420116",
					SequenceNumber: "1",
				},
			},
		},
		Amount: &C0401InvoiceAmount{
			InvoiceAmount: InvoiceAmount{
				SalesAmount: "1650",
				TaxType:     "1",
				TaxRate:     "0.05",
				TaxAmount:   "0",
				TotalAmount: "1650",
			},
			FreeTaxSalesAmount: "0",
			ZeroTaxSalesAmount: "0",
		},
	}

	if actual.Main.InvoiceNumber != expected.Main.InvoiceNumber {
		t.Errorf("invoice number not match, expected: %s got: %s", expected.Main.InvoiceNumber, actual.Main.InvoiceNumber)
	}

	if actual.Main.InvoiceDate != expected.Main.InvoiceDate {
		t.Errorf("invoice date not match, expected: %s got: %s", expected.Main.InvoiceDate, actual.Main.InvoiceDate)
	}

	if actual.Main.InvoiceTime != expected.Main.InvoiceTime {
		t.Errorf("invoice time not match, expected: %s got: %s", expected.Main.InvoiceTime, actual.Main.InvoiceTime)
	}

	if actual.Main.Seller.Identifier != expected.Main.Seller.Identifier {
		t.Errorf("invoice seller identifier not match, expected: %s got: %s", expected.Main.Seller.Identifier, actual.Main.Seller.Identifier)
	}
	if actual.Main.Seller.Name != expected.Main.Seller.Name {
		t.Errorf("invoice seller name not match, expected: %s got: %s", expected.Main.Seller.Name, actual.Main.Seller.Name)
	}
	if actual.Main.Seller.Address != expected.Main.Seller.Address {
		t.Errorf("invoice seller address not match, expected: %s got: %s", expected.Main.Seller.Address, actual.Main.Seller.Address)
	}
	if actual.Main.Seller.PersonInCharge != expected.Main.Seller.PersonInCharge {
		t.Errorf("invoice seller person in charge not match, expected: %s got: %s", expected.Main.Seller.PersonInCharge, actual.Main.Seller.PersonInCharge)
	}
	if actual.Main.Seller.EmailAddress != expected.Main.Seller.EmailAddress {
		t.Errorf("invoice email address not match, expected: %s got: %s", expected.Main.Seller.EmailAddress, actual.Main.Seller.EmailAddress)
	}

	if actual.Main.Buyer.Identifier != expected.Main.Buyer.Identifier {
		t.Errorf("invoice buyer identifier not match, expected: %s got: %s", expected.Main.Buyer.Identifier, actual.Main.Buyer.Identifier)
	}
	if actual.Main.Buyer.Name != expected.Main.Buyer.Name {
		t.Errorf("invoice buyer name not match, expected: %s got: %s", expected.Main.Buyer.Name, actual.Main.Buyer.Name)
	}

	if actual.Main.InvoiceType != expected.Main.InvoiceType {
		t.Errorf("invoice type not match, expected: %s got: %s", expected.Main.InvoiceType, actual.Main.InvoiceType)
	}
	if actual.Main.DonateMark != expected.Main.DonateMark {
		t.Errorf("invoice donate mark not match, expected: %s got: %s", expected.Main.DonateMark, actual.Main.DonateMark)
	}
	if actual.Main.CarrierType != expected.Main.CarrierType {
		t.Errorf("invoice carrier type not match, expected: %s got: %s", expected.Main.CarrierType, actual.Main.CarrierType)
	}
	if actual.Main.CarrierId1 != expected.Main.CarrierId1 {
		t.Errorf("invoice carrier id 1 not match, expected: %s got: %s", expected.Main.CarrierId1, actual.Main.CarrierId1)
	}
	if actual.Main.CarrierId2 != expected.Main.CarrierId2 {
		t.Errorf("invoice carrier id 2 not match, expected: %s got: %s", expected.Main.CarrierId2, actual.Main.CarrierId2)
	}
	if actual.Main.PrintMark != expected.Main.PrintMark {
		t.Errorf("invoice print mark not match, expected: %s got: %s", expected.Main.PrintMark, actual.Main.PrintMark)
	}
	if actual.Main.RandomNumber != expected.Main.RandomNumber {
		t.Errorf("invoice random number not match, expected: %s got: %s", expected.Main.RandomNumber, actual.Main.RandomNumber)
	}

	// Details

	if len(actual.Details.ProductItem) != len(expected.Details.ProductItem) {
		t.Errorf("length of product item not match, expected: %d got: %d", len(expected.Details.ProductItem), len(actual.Details.ProductItem))

	} else {
		for pi, it := range expected.Details.ProductItem {
			actualItem := actual.Details.ProductItem[pi]

			if actualItem.Description != it.Description {
				t.Errorf("invoice product item %d description not match, expected: %s got: %s", pi, it.Description, actualItem.Description)
			}

			if actualItem.Quantity != it.Quantity {
				t.Errorf("invoice product item %d quantity not match, expected: %s got: %s", pi, it.Quantity, actualItem.Quantity)
			}

			if actualItem.UnitPrice != it.UnitPrice {
				t.Errorf("invoice product item %d unit price not match, expected: %s got: %s", pi, it.UnitPrice, actualItem.UnitPrice)
			}

			if actualItem.Amount != it.Amount {
				t.Errorf("invoice product item %d amount not match, expected: %s got: %s", pi, it.Amount, actualItem.Amount)
			}

			if actualItem.SequenceNumber != it.SequenceNumber {
				t.Errorf("invoice product item %d sequence number not match, expected: %s got: %s", pi, it.SequenceNumber, actualItem.SequenceNumber)
			}

			if actualItem.RelateNumber != it.RelateNumber {
				t.Errorf("invoice product item %d relate number not match, expected: %s got: %s", pi, it.RelateNumber, actualItem.RelateNumber)
			}

		}
	}

	if actual.Amount.SalesAmount != expected.Amount.SalesAmount {
		t.Errorf("invoice sales amount not match, expected: %s got: %s", expected.Amount.SalesAmount, actual.Amount.SalesAmount)
	}
	if actual.Amount.FreeTaxSalesAmount != expected.Amount.FreeTaxSalesAmount {
		t.Errorf("invoice free tax sales amount not match, expected: %s got: %s", expected.Amount.FreeTaxSalesAmount, actual.Amount.FreeTaxSalesAmount)
	}
	if actual.Amount.ZeroTaxSalesAmount != expected.Amount.ZeroTaxSalesAmount {
		t.Errorf("invoice zero tax sales amount not match, expected: %s got: %s", expected.Amount.FreeTaxSalesAmount, actual.Amount.FreeTaxSalesAmount)
	}

	if actual.Amount.TaxType != expected.Amount.TaxType {
		t.Errorf("invoice tax type not match, expected: %s got: %s", expected.Amount.TaxType, actual.Amount.TaxType)
	}
	if actual.Amount.TaxRate != expected.Amount.TaxRate {
		t.Errorf("invoice tax rate not match, expected: %s got: %s", expected.Amount.TaxRate, actual.Amount.TaxRate)
	}
	if actual.Amount.TaxAmount != expected.Amount.TaxAmount {
		t.Errorf("invoice tax amount not match, expected: %s got: %s", expected.Amount.TaxAmount, actual.Amount.TaxAmount)
	}
	if actual.Amount.TotalAmount != expected.Amount.TotalAmount {
		t.Errorf("invoice total amount not match, expected: %s got: %s", expected.Amount.TotalAmount, actual.Amount.TotalAmount)
	}
}
