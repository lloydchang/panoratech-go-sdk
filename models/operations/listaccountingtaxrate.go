// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/panoratech/go-sdk/internal/utils"
	"github.com/panoratech/go-sdk/models/components"
)

type ListAccountingTaxRateRequest struct {
	// The connection token
	XConnectionToken string `header:"style=simple,explode=false,name=x-connection-token"`
	// Set to true to include data from the original software.
	RemoteData *bool `queryParam:"style=form,explode=true,name=remote_data"`
	// Set to get the number of records.
	Limit *float64 `default:"30" queryParam:"style=form,explode=true,name=limit"`
	// Set to get the number of records after this cursor.
	Cursor *string `queryParam:"style=form,explode=true,name=cursor"`
}

func (l ListAccountingTaxRateRequest) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(l, "", false)
}

func (l *ListAccountingTaxRateRequest) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &l, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *ListAccountingTaxRateRequest) GetXConnectionToken() string {
	if o == nil {
		return ""
	}
	return o.XConnectionToken
}

func (o *ListAccountingTaxRateRequest) GetRemoteData() *bool {
	if o == nil {
		return nil
	}
	return o.RemoteData
}

func (o *ListAccountingTaxRateRequest) GetLimit() *float64 {
	if o == nil {
		return nil
	}
	return o.Limit
}

func (o *ListAccountingTaxRateRequest) GetCursor() *string {
	if o == nil {
		return nil
	}
	return o.Cursor
}

type ListAccountingTaxRateResponseBody struct {
	PrevCursor *string                                     `json:"prev_cursor"`
	NextCursor *string                                     `json:"next_cursor"`
	Data       []components.UnifiedAccountingTaxrateOutput `json:"data"`
}

func (o *ListAccountingTaxRateResponseBody) GetPrevCursor() *string {
	if o == nil {
		return nil
	}
	return o.PrevCursor
}

func (o *ListAccountingTaxRateResponseBody) GetNextCursor() *string {
	if o == nil {
		return nil
	}
	return o.NextCursor
}

func (o *ListAccountingTaxRateResponseBody) GetData() []components.UnifiedAccountingTaxrateOutput {
	if o == nil {
		return []components.UnifiedAccountingTaxrateOutput{}
	}
	return o.Data
}

type ListAccountingTaxRateResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	Object   *ListAccountingTaxRateResponseBody

	Next func() (*ListAccountingTaxRateResponse, error)
}

func (o *ListAccountingTaxRateResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *ListAccountingTaxRateResponse) GetObject() *ListAccountingTaxRateResponseBody {
	if o == nil {
		return nil
	}
	return o.Object
}
