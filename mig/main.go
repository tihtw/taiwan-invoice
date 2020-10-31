package mig

import (
	"encoding/xml"
	"log"
)

type RoleDescription struct {
	Identifier     string `xml:"Identifier"`
	Name           string `xml:"Name"`
	Address        string `xml:"Address"`
	PersonInCharge string `xml:"PersonInCharge"`
	EmailAddress   string `xml:"EmailAddress"`
}

type MigMain struct {
	InvoiceNumber string          `xml:"InvoiceNumber"`
	InvoiceDate   string          `xml:"InvoiceDate"`
	InvoiceTime   string          `xml:"InvoiceTime"`
	Seller        RoleDescription `xml:"Seller"`
	Buyer         RoleDescription `xml:"Buyer"`

	InvoiceType  string `xml:"InvoiceType"`
	DonateMark   string `xml:"DonateMark"`
	CarrierType  string `xml:"CarrierType"`
	CarrierId1   string `xml:"CarrierId1"`
	CarrierId2   string `xml:"CarrierId2"`
	PrintMark    string `xml:"PrintMark"`
	RandomNumber string `xml:"RandomNumber"`
}

type ProductItem struct {
	Text           string `xml:",chardata"`
	Description    string `xml:"Description"`
	Quantity       string `xml:"Quantity"`
	UnitPrice      string `xml:"UnitPrice"`
	Amount         string `xml:"Amount"`
	SequenceNumber string `xml:"SequenceNumber"`
	RelateNumber   string `xml:"RelateNumber"`
}

type MigDetail struct {
	Text        string        `xml:",chardata"`
	ProductItem []ProductItem `xml:"ProductItem"`
}
type MigAmount struct {
	Text               string `xml:",chardata"`
	SalesAmount        string `xml:"SalesAmount"`
	FreeTaxSalesAmount string `xml:"FreeTaxSalesAmount"`
	ZeroTaxSalesAmount string `xml:"ZeroTaxSalesAmount"`
	TaxType            string `xml:"TaxType"`
	TaxRate            string `xml:"TaxRate"`
	TaxAmount          string `xml:"TaxAmount"`
	TotalAmount        string `xml:"TotalAmount"`
}

type MigFile struct {
	XMLName xml.Name `xml:"Invoice"`
	Text    string   `xml:",chardata"`
	Xmlns   string   `xml:"xmlns,attr"`

	Main    MigMain   `xml:"Main"`
	Details MigDetail `xml:"Details"`
	Amount  MigAmount `xml:"Amount"`
}

func NewMigFile(b []byte) (*MigFile, error) {
	f := MigFile{}

	if err := xml.Unmarshal(b, &f); err != nil {
		log.Fatal(err)
	}
	log.Println("xmlname", f.XMLName.Space, f.XMLName.Local)

	return &f, nil

}
