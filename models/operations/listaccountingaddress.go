// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/panoratech/go-sdk/internal/utils"
	"github.com/panoratech/go-sdk/models/components"
)

type ListAccountingAddressRequest struct {
	// The connection token
	XConnectionToken string `header:"style=simple,explode=false,name=x-connection-token"`
	// Set to true to include data from the original software.
	RemoteData *bool `queryParam:"style=form,explode=true,name=remote_data"`
	// Set to get the number of records.
	Limit *float64 `default:"30" queryParam:"style=form,explode=true,name=limit"`
	// Set to get the number of records after this cursor.
	Cursor *string `queryParam:"style=form,explode=true,name=cursor"`
}

func (l ListAccountingAddressRequest) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(l, "", false)
}

func (l *ListAccountingAddressRequest) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &l, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *ListAccountingAddressRequest) GetXConnectionToken() string {
	if o == nil {
		return ""
	}
	return o.XConnectionToken
}

func (o *ListAccountingAddressRequest) GetRemoteData() *bool {
	if o == nil {
		return nil
	}
	return o.RemoteData
}

func (o *ListAccountingAddressRequest) GetLimit() *float64 {
	if o == nil {
		return nil
	}
	return o.Limit
}

func (o *ListAccountingAddressRequest) GetCursor() *string {
	if o == nil {
		return nil
	}
	return o.Cursor
}

type ListAccountingAddressResponseBody struct {
	PrevCursor *string                                     `json:"prev_cursor"`
	NextCursor *string                                     `json:"next_cursor"`
	Data       []components.UnifiedAccountingAddressOutput `json:"data"`
}

func (o *ListAccountingAddressResponseBody) GetPrevCursor() *string {
	if o == nil {
		return nil
	}
	return o.PrevCursor
}

func (o *ListAccountingAddressResponseBody) GetNextCursor() *string {
	if o == nil {
		return nil
	}
	return o.NextCursor
}

func (o *ListAccountingAddressResponseBody) GetData() []components.UnifiedAccountingAddressOutput {
	if o == nil {
		return []components.UnifiedAccountingAddressOutput{}
	}
	return o.Data
}

type ListAccountingAddressResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	Object   *ListAccountingAddressResponseBody

	Next func() (*ListAccountingAddressResponse, error)
}

func (o *ListAccountingAddressResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *ListAccountingAddressResponse) GetObject() *ListAccountingAddressResponseBody {
	if o == nil {
		return nil
	}
	return o.Object
}
