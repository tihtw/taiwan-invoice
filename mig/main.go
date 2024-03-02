package mig

import (
	"encoding/xml"
	"log"
)

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
	// log.Println("xmlname", f.XMLName.Space, f.XMLName.Local)

	return &f, nil
}
