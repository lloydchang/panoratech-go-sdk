// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

import (
	"github.com/panoratech/go-sdk/internal/utils"
	"time"
)

// UnifiedAccountingCashflowstatementOutputFieldMappings - The custom field mappings of the object between the remote 3rd party & Panora
type UnifiedAccountingCashflowstatementOutputFieldMappings struct {
}

// UnifiedAccountingCashflowstatementOutputRemoteData - The remote data of the cash flow statement in the context of the 3rd Party
type UnifiedAccountingCashflowstatementOutputRemoteData struct {
}

type UnifiedAccountingCashflowstatementOutput struct {
	// The name of the cash flow statement
	Name *string `json:"name,omitempty"`
	// The currency used in the cash flow statement
	Currency *string `json:"currency,omitempty"`
	// The UUID of the associated company
	CompanyID *string `json:"company_id,omitempty"`
	// The start date of the period covered by the cash flow statement
	StartPeriod *time.Time `json:"start_period,omitempty"`
	// The end date of the period covered by the cash flow statement
	EndPeriod *time.Time `json:"end_period,omitempty"`
	// The cash balance at the beginning of the period
	CashAtBeginningOfPeriod *float64 `json:"cash_at_beginning_of_period,omitempty"`
	// The cash balance at the end of the period
	CashAtEndOfPeriod *float64 `json:"cash_at_end_of_period,omitempty"`
	// The date when the cash flow statement was generated in the remote system
	RemoteGeneratedAt *time.Time `json:"remote_generated_at,omitempty"`
	// The report items associated with this cash flow statement
	LineItems []LineItem `json:"line_items,omitempty"`
	// The custom field mappings of the object between the remote 3rd party & Panora
	FieldMappings *UnifiedAccountingCashflowstatementOutputFieldMappings `json:"field_mappings,omitempty"`
	// The UUID of the cash flow statement record
	ID *string `json:"id,omitempty"`
	// The remote ID of the cash flow statement in the context of the 3rd Party
	RemoteID *string `json:"remote_id,omitempty"`
	// The remote data of the cash flow statement in the context of the 3rd Party
	RemoteData *UnifiedAccountingCashflowstatementOutputRemoteData `json:"remote_data,omitempty"`
	// The created date of the cash flow statement record
	CreatedAt *time.Time `json:"created_at,omitempty"`
	// The last modified date of the cash flow statement record
	ModifiedAt *time.Time `json:"modified_at,omitempty"`
}

func (u UnifiedAccountingCashflowstatementOutput) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(u, "", false)
}

func (u *UnifiedAccountingCashflowstatementOutput) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &u, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *UnifiedAccountingCashflowstatementOutput) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

func (o *UnifiedAccountingCashflowstatementOutput) GetCurrency() *string {
	if o == nil {
		return nil
	}
	return o.Currency
}

func (o *UnifiedAccountingCashflowstatementOutput) GetCompanyID() *string {
	if o == nil {
		return nil
	}
	return o.CompanyID
}

func (o *UnifiedAccountingCashflowstatementOutput) GetStartPeriod() *time.Time {
	if o == nil {
		return nil
	}
	return o.StartPeriod
}

func (o *UnifiedAccountingCashflowstatementOutput) GetEndPeriod() *time.Time {
	if o == nil {
		return nil
	}
	return o.EndPeriod
}

func (o *UnifiedAccountingCashflowstatementOutput) GetCashAtBeginningOfPeriod() *float64 {
	if o == nil {
		return nil
	}
	return o.CashAtBeginningOfPeriod
}

func (o *UnifiedAccountingCashflowstatementOutput) GetCashAtEndOfPeriod() *float64 {
	if o == nil {
		return nil
	}
	return o.CashAtEndOfPeriod
}

func (o *UnifiedAccountingCashflowstatementOutput) GetRemoteGeneratedAt() *time.Time {
	if o == nil {
		return nil
	}
	return o.RemoteGeneratedAt
}

func (o *UnifiedAccountingCashflowstatementOutput) GetLineItems() []LineItem {
	if o == nil {
		return nil
	}
	return o.LineItems
}

func (o *UnifiedAccountingCashflowstatementOutput) GetFieldMappings() *UnifiedAccountingCashflowstatementOutputFieldMappings {
	if o == nil {
		return nil
	}
	return o.FieldMappings
}

func (o *UnifiedAccountingCashflowstatementOutput) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *UnifiedAccountingCashflowstatementOutput) GetRemoteID() *string {
	if o == nil {
		return nil
	}
	return o.RemoteID
}

func (o *UnifiedAccountingCashflowstatementOutput) GetRemoteData() *UnifiedAccountingCashflowstatementOutputRemoteData {
	if o == nil {
		return nil
	}
	return o.RemoteData
}

func (o *UnifiedAccountingCashflowstatementOutput) GetCreatedAt() *time.Time {
	if o == nil {
		return nil
	}
	return o.CreatedAt
}

func (o *UnifiedAccountingCashflowstatementOutput) GetModifiedAt() *time.Time {
	if o == nil {
		return nil
	}
	return o.ModifiedAt
}
