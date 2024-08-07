// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/panoratech/go-sdk/internal/utils"
	"github.com/panoratech/go-sdk/models/components"
)

type ListEcommerceProductsRequest struct {
	// The connection token
	XConnectionToken string `header:"style=simple,explode=false,name=x-connection-token"`
	// Set to true to include data from the original software.
	RemoteData *bool `queryParam:"style=form,explode=true,name=remote_data"`
	// Set to get the number of records.
	Limit *float64 `default:"30" queryParam:"style=form,explode=true,name=limit"`
	// Set to get the number of records after this cursor.
	Cursor *string `queryParam:"style=form,explode=true,name=cursor"`
}

func (l ListEcommerceProductsRequest) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(l, "", false)
}

func (l *ListEcommerceProductsRequest) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &l, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *ListEcommerceProductsRequest) GetXConnectionToken() string {
	if o == nil {
		return ""
	}
	return o.XConnectionToken
}

func (o *ListEcommerceProductsRequest) GetRemoteData() *bool {
	if o == nil {
		return nil
	}
	return o.RemoteData
}

func (o *ListEcommerceProductsRequest) GetLimit() *float64 {
	if o == nil {
		return nil
	}
	return o.Limit
}

func (o *ListEcommerceProductsRequest) GetCursor() *string {
	if o == nil {
		return nil
	}
	return o.Cursor
}

type ListEcommerceProductsResponseBody struct {
	PrevCursor *string                                    `json:"prev_cursor"`
	NextCursor *string                                    `json:"next_cursor"`
	Data       []components.UnifiedEcommerceProductOutput `json:"data"`
}

func (o *ListEcommerceProductsResponseBody) GetPrevCursor() *string {
	if o == nil {
		return nil
	}
	return o.PrevCursor
}

func (o *ListEcommerceProductsResponseBody) GetNextCursor() *string {
	if o == nil {
		return nil
	}
	return o.NextCursor
}

func (o *ListEcommerceProductsResponseBody) GetData() []components.UnifiedEcommerceProductOutput {
	if o == nil {
		return []components.UnifiedEcommerceProductOutput{}
	}
	return o.Data
}

type ListEcommerceProductsResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	Object   *ListEcommerceProductsResponseBody

	Next func() (*ListEcommerceProductsResponse, error)
}

func (o *ListEcommerceProductsResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *ListEcommerceProductsResponse) GetObject() *ListEcommerceProductsResponseBody {
	if o == nil {
		return nil
	}
	return o.Object
}
