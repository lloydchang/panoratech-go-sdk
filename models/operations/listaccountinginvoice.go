// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/panoratech/go-sdk/models/components"
)

type ListAccountingInvoiceRequest struct {
	// The connection token
	XConnectionToken string `header:"style=simple,explode=false,name=x-connection-token"`
	// Set to true to include data from the original software.
	RemoteData *bool `queryParam:"style=form,explode=true,name=remote_data"`
	// Set to get the number of records.
	Limit *float64 `queryParam:"style=form,explode=true,name=limit"`
	// Set to get the number of records after this cursor.
	Cursor *string `queryParam:"style=form,explode=true,name=cursor"`
}

func (o *ListAccountingInvoiceRequest) GetXConnectionToken() string {
	if o == nil {
		return ""
	}
	return o.XConnectionToken
}

func (o *ListAccountingInvoiceRequest) GetRemoteData() *bool {
	if o == nil {
		return nil
	}
	return o.RemoteData
}

func (o *ListAccountingInvoiceRequest) GetLimit() *float64 {
	if o == nil {
		return nil
	}
	return o.Limit
}

func (o *ListAccountingInvoiceRequest) GetCursor() *string {
	if o == nil {
		return nil
	}
	return o.Cursor
}

type ListAccountingInvoiceResponseBody struct {
	PrevCursor *string                                     `json:"prev_cursor"`
	NextCursor *string                                     `json:"next_cursor"`
	Data       []components.UnifiedAccountingInvoiceOutput `json:"data"`
}

func (o *ListAccountingInvoiceResponseBody) GetPrevCursor() *string {
	if o == nil {
		return nil
	}
	return o.PrevCursor
}

func (o *ListAccountingInvoiceResponseBody) GetNextCursor() *string {
	if o == nil {
		return nil
	}
	return o.NextCursor
}

func (o *ListAccountingInvoiceResponseBody) GetData() []components.UnifiedAccountingInvoiceOutput {
	if o == nil {
		return []components.UnifiedAccountingInvoiceOutput{}
	}
	return o.Data
}

type ListAccountingInvoiceResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	Object   *ListAccountingInvoiceResponseBody

	Next func() (*ListAccountingInvoiceResponse, error)
}

func (o *ListAccountingInvoiceResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *ListAccountingInvoiceResponse) GetObject() *ListAccountingInvoiceResponseBody {
	if o == nil {
		return nil
	}
	return o.Object
}
