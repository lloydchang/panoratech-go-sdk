// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

import (
	"github.com/panoratech/go-sdk/internal/utils"
	"time"
)

// UnifiedAccountingPaymentOutputFieldMappings - The custom field mappings of the object between the remote 3rd party & Panora
type UnifiedAccountingPaymentOutputFieldMappings struct {
}

// UnifiedAccountingPaymentOutputRemoteData - The remote data of the payment in the context of the 3rd Party
type UnifiedAccountingPaymentOutputRemoteData struct {
}

type UnifiedAccountingPaymentOutput struct {
	// The UUID of the associated invoice
	InvoiceID *string `json:"invoice_id,omitempty"`
	// The date of the transaction
	TransactionDate *time.Time `json:"transaction_date,omitempty"`
	// The UUID of the associated contact
	ContactID *string `json:"contact_id,omitempty"`
	// The UUID of the associated account
	AccountID *string `json:"account_id,omitempty"`
	// The currency of the payment
	Currency *string `json:"currency,omitempty"`
	// The exchange rate applied to the payment
	ExchangeRate *string `json:"exchange_rate,omitempty"`
	// The total amount of the payment in cents
	TotalAmount *float64 `json:"total_amount,omitempty"`
	// The type of payment
	Type *string `json:"type,omitempty"`
	// The UUID of the associated company info
	CompanyInfoID *string `json:"company_info_id,omitempty"`
	// The UUID of the associated accounting period
	AccountingPeriodID *string `json:"accounting_period_id,omitempty"`
	// The UUIDs of the tracking categories associated with the payment
	TrackingCategories []string `json:"tracking_categories,omitempty"`
	// The line items associated with this payment
	LineItems []LineItem `json:"line_items,omitempty"`
	// The custom field mappings of the object between the remote 3rd party & Panora
	FieldMappings *UnifiedAccountingPaymentOutputFieldMappings `json:"field_mappings,omitempty"`
	// The UUID of the payment record
	ID *string `json:"id,omitempty"`
	// The remote ID of the payment in the context of the 3rd Party
	RemoteID *string `json:"remote_id,omitempty"`
	// The date when the payment was last updated in the remote system
	RemoteUpdatedAt *time.Time `json:"remote_updated_at,omitempty"`
	// The remote data of the payment in the context of the 3rd Party
	RemoteData *UnifiedAccountingPaymentOutputRemoteData `json:"remote_data,omitempty"`
	// The created date of the payment record
	CreatedAt *time.Time `json:"created_at,omitempty"`
	// The last modified date of the payment record
	ModifiedAt *time.Time `json:"modified_at,omitempty"`
}

func (u UnifiedAccountingPaymentOutput) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(u, "", false)
}

func (u *UnifiedAccountingPaymentOutput) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &u, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *UnifiedAccountingPaymentOutput) GetInvoiceID() *string {
	if o == nil {
		return nil
	}
	return o.InvoiceID
}

func (o *UnifiedAccountingPaymentOutput) GetTransactionDate() *time.Time {
	if o == nil {
		return nil
	}
	return o.TransactionDate
}

func (o *UnifiedAccountingPaymentOutput) GetContactID() *string {
	if o == nil {
		return nil
	}
	return o.ContactID
}

func (o *UnifiedAccountingPaymentOutput) GetAccountID() *string {
	if o == nil {
		return nil
	}
	return o.AccountID
}

func (o *UnifiedAccountingPaymentOutput) GetCurrency() *string {
	if o == nil {
		return nil
	}
	return o.Currency
}

func (o *UnifiedAccountingPaymentOutput) GetExchangeRate() *string {
	if o == nil {
		return nil
	}
	return o.ExchangeRate
}

func (o *UnifiedAccountingPaymentOutput) GetTotalAmount() *float64 {
	if o == nil {
		return nil
	}
	return o.TotalAmount
}

func (o *UnifiedAccountingPaymentOutput) GetType() *string {
	if o == nil {
		return nil
	}
	return o.Type
}

func (o *UnifiedAccountingPaymentOutput) GetCompanyInfoID() *string {
	if o == nil {
		return nil
	}
	return o.CompanyInfoID
}

func (o *UnifiedAccountingPaymentOutput) GetAccountingPeriodID() *string {
	if o == nil {
		return nil
	}
	return o.AccountingPeriodID
}

func (o *UnifiedAccountingPaymentOutput) GetTrackingCategories() []string {
	if o == nil {
		return nil
	}
	return o.TrackingCategories
}

func (o *UnifiedAccountingPaymentOutput) GetLineItems() []LineItem {
	if o == nil {
		return nil
	}
	return o.LineItems
}

func (o *UnifiedAccountingPaymentOutput) GetFieldMappings() *UnifiedAccountingPaymentOutputFieldMappings {
	if o == nil {
		return nil
	}
	return o.FieldMappings
}

func (o *UnifiedAccountingPaymentOutput) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *UnifiedAccountingPaymentOutput) GetRemoteID() *string {
	if o == nil {
		return nil
	}
	return o.RemoteID
}

func (o *UnifiedAccountingPaymentOutput) GetRemoteUpdatedAt() *time.Time {
	if o == nil {
		return nil
	}
	return o.RemoteUpdatedAt
}

func (o *UnifiedAccountingPaymentOutput) GetRemoteData() *UnifiedAccountingPaymentOutputRemoteData {
	if o == nil {
		return nil
	}
	return o.RemoteData
}

func (o *UnifiedAccountingPaymentOutput) GetCreatedAt() *time.Time {
	if o == nil {
		return nil
	}
	return o.CreatedAt
}

func (o *UnifiedAccountingPaymentOutput) GetModifiedAt() *time.Time {
	if o == nil {
		return nil
	}
	return o.ModifiedAt
}
