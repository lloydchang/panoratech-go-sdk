// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

type TaxItem struct {
	// The name of the tax
	Name *string `json:"name,omitempty"`
	// The amount of the tax
	Amount *float64 `json:"amount,omitempty"`
	// Indicates if this is an employer tax
	EmployerTax *bool `json:"employer_tax,omitempty"`
}

func (o *TaxItem) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

func (o *TaxItem) GetAmount() *float64 {
	if o == nil {
		return nil
	}
	return o.Amount
}

func (o *TaxItem) GetEmployerTax() *bool {
	if o == nil {
		return nil
	}
	return o.EmployerTax
}