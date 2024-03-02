package mig

type Buyer struct {
	RoleDescription
}

func NewBuyer() *Buyer {
	return &Buyer{
		RoleDescription: RoleDescription{
			Identifier: "0000000000",
		},
	}
}

func (buyer *Buyer) Validate() error {
	err := buyer.RoleDescription.Validate()
	if err != nil {
		return err
	}

	return nil
}
