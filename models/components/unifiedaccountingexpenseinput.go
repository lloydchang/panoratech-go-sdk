// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

import (
	"github.com/panoratech/go-sdk/internal/utils"
	"time"
)

// UnifiedAccountingExpenseInputFieldMappings - The custom field mappings of the object between the remote 3rd party & Panora
type UnifiedAccountingExpenseInputFieldMappings struct {
}

type UnifiedAccountingExpenseInput struct {
	// The date of the expense transaction
	TransactionDate *time.Time `json:"transaction_date,omitempty"`
	// The total amount of the expense
	TotalAmount *float64 `json:"total_amount,omitempty"`
	// The sub-total amount of the expense (before tax)
	SubTotal *float64 `json:"sub_total,omitempty"`
	// The total tax amount of the expense
	TotalTaxAmount *float64 `json:"total_tax_amount,omitempty"`
	// The currency of the expense
	Currency *string `json:"currency,omitempty"`
	// The exchange rate applied to the expense
	ExchangeRate *string `json:"exchange_rate,omitempty"`
	// A memo or description for the expense
	Memo *string `json:"memo,omitempty"`
	// The UUID of the associated account
	AccountID *string `json:"account_id,omitempty"`
	// The UUID of the associated contact
	ContactID *string `json:"contact_id,omitempty"`
	// The UUID of the associated company info
	CompanyInfoID *string `json:"company_info_id,omitempty"`
	// The UUIDs of the tracking categories associated with the expense
	TrackingCategories []string `json:"tracking_categories,omitempty"`
	// The line items associated with this expense
	LineItems []LineItem `json:"line_items,omitempty"`
	// The custom field mappings of the object between the remote 3rd party & Panora
	FieldMappings *UnifiedAccountingExpenseInputFieldMappings `json:"field_mappings,omitempty"`
}

func (u UnifiedAccountingExpenseInput) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(u, "", false)
}

func (u *UnifiedAccountingExpenseInput) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &u, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *UnifiedAccountingExpenseInput) GetTransactionDate() *time.Time {
	if o == nil {
		return nil
	}
	return o.TransactionDate
}

func (o *UnifiedAccountingExpenseInput) GetTotalAmount() *float64 {
	if o == nil {
		return nil
	}
	return o.TotalAmount
}

func (o *UnifiedAccountingExpenseInput) GetSubTotal() *float64 {
	if o == nil {
		return nil
	}
	return o.SubTotal
}

func (o *UnifiedAccountingExpenseInput) GetTotalTaxAmount() *float64 {
	if o == nil {
		return nil
	}
	return o.TotalTaxAmount
}

func (o *UnifiedAccountingExpenseInput) GetCurrency() *string {
	if o == nil {
		return nil
	}
	return o.Currency
}

func (o *UnifiedAccountingExpenseInput) GetExchangeRate() *string {
	if o == nil {
		return nil
	}
	return o.ExchangeRate
}

func (o *UnifiedAccountingExpenseInput) GetMemo() *string {
	if o == nil {
		return nil
	}
	return o.Memo
}

func (o *UnifiedAccountingExpenseInput) GetAccountID() *string {
	if o == nil {
		return nil
	}
	return o.AccountID
}

func (o *UnifiedAccountingExpenseInput) GetContactID() *string {
	if o == nil {
		return nil
	}
	return o.ContactID
}

func (o *UnifiedAccountingExpenseInput) GetCompanyInfoID() *string {
	if o == nil {
		return nil
	}
	return o.CompanyInfoID
}

func (o *UnifiedAccountingExpenseInput) GetTrackingCategories() []string {
	if o == nil {
		return nil
	}
	return o.TrackingCategories
}

func (o *UnifiedAccountingExpenseInput) GetLineItems() []LineItem {
	if o == nil {
		return nil
	}
	return o.LineItems
}

func (o *UnifiedAccountingExpenseInput) GetFieldMappings() *UnifiedAccountingExpenseInputFieldMappings {
	if o == nil {
		return nil
	}
	return o.FieldMappings
}
