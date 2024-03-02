package mig

import (
	"encoding/xml"
	"log"
)

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

type C0401Invoice struct {
	XMLName xml.Name `xml:"Invoice"`
	Text    string   `xml:",chardata"`
	Xmlns   string   `xml:"xmlns,attr"`

	Main    *C0401InvoiceMain   `xml:"Main"`
	Details *C0401InvoiceDetail `xml:"Details"`
	Amount  *C0401InvoiceAmount `xml:"Amount"`
}

func NewC0401Invoice(b []byte) (*C0401Invoice, error) {
	f := C0401Invoice{}

	if err := xml.Unmarshal(b, &f); err != nil {
		log.Fatal(err)
	}
	log.Println("xmlname", f.XMLName.Space, f.XMLName.Local)

	return &f, nil
}
