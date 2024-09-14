// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/panoratech/go-sdk/internal/utils"
	"github.com/panoratech/go-sdk/models/components"
)

type ListEcommerceFulfillmentsRequest struct {
	// The connection token
	XConnectionToken string `header:"style=simple,explode=false,name=x-connection-token"`
	// Set to true to include data from the original software.
	RemoteData *bool `queryParam:"style=form,explode=true,name=remote_data"`
	// Set to get the number of records.
	Limit *float64 `default:"30" queryParam:"style=form,explode=true,name=limit"`
	// Set to get the number of records after this cursor.
	Cursor *string `queryParam:"style=form,explode=true,name=cursor"`
}

func (l ListEcommerceFulfillmentsRequest) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(l, "", false)
}

func (l *ListEcommerceFulfillmentsRequest) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &l, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *ListEcommerceFulfillmentsRequest) GetXConnectionToken() string {
	if o == nil {
		return ""
	}
	return o.XConnectionToken
}

func (o *ListEcommerceFulfillmentsRequest) GetRemoteData() *bool {
	if o == nil {
		return nil
	}
	return o.RemoteData
}

func (o *ListEcommerceFulfillmentsRequest) GetLimit() *float64 {
	if o == nil {
		return nil
	}
	return o.Limit
}

func (o *ListEcommerceFulfillmentsRequest) GetCursor() *string {
	if o == nil {
		return nil
	}
	return o.Cursor
}

type ListEcommerceFulfillmentsResponseBody struct {
	PrevCursor *string                                        `json:"prev_cursor"`
	NextCursor *string                                        `json:"next_cursor"`
	Data       []components.UnifiedEcommerceFulfillmentOutput `json:"data"`
}

func (o *ListEcommerceFulfillmentsResponseBody) GetPrevCursor() *string {
	if o == nil {
		return nil
	}
	return o.PrevCursor
}

func (o *ListEcommerceFulfillmentsResponseBody) GetNextCursor() *string {
	if o == nil {
		return nil
	}
	return o.NextCursor
}

func (o *ListEcommerceFulfillmentsResponseBody) GetData() []components.UnifiedEcommerceFulfillmentOutput {
	if o == nil {
		return []components.UnifiedEcommerceFulfillmentOutput{}
	}
	return o.Data
}

type ListEcommerceFulfillmentsResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	Object   *ListEcommerceFulfillmentsResponseBody

	Next func() (*ListEcommerceFulfillmentsResponse, error)
}

func (o *ListEcommerceFulfillmentsResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *ListEcommerceFulfillmentsResponse) GetObject() *ListEcommerceFulfillmentsResponseBody {
	if o == nil {
		return nil
	}
	return o.Object
}