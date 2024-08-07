// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/panoratech/go-sdk/internal/utils"
	"github.com/panoratech/go-sdk/models/components"
)

type ListAtsOfferRequest struct {
	// The connection token
	XConnectionToken string `header:"style=simple,explode=false,name=x-connection-token"`
	// Set to true to include data from the original software.
	RemoteData *bool `queryParam:"style=form,explode=true,name=remote_data"`
	// Set to get the number of records.
	Limit *float64 `default:"30" queryParam:"style=form,explode=true,name=limit"`
	// Set to get the number of records after this cursor.
	Cursor *string `queryParam:"style=form,explode=true,name=cursor"`
}

func (l ListAtsOfferRequest) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(l, "", false)
}

func (l *ListAtsOfferRequest) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &l, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *ListAtsOfferRequest) GetXConnectionToken() string {
	if o == nil {
		return ""
	}
	return o.XConnectionToken
}

func (o *ListAtsOfferRequest) GetRemoteData() *bool {
	if o == nil {
		return nil
	}
	return o.RemoteData
}

func (o *ListAtsOfferRequest) GetLimit() *float64 {
	if o == nil {
		return nil
	}
	return o.Limit
}

func (o *ListAtsOfferRequest) GetCursor() *string {
	if o == nil {
		return nil
	}
	return o.Cursor
}

type ListAtsOfferResponseBody struct {
	PrevCursor *string                            `json:"prev_cursor"`
	NextCursor *string                            `json:"next_cursor"`
	Data       []components.UnifiedAtsOfferOutput `json:"data"`
}

func (o *ListAtsOfferResponseBody) GetPrevCursor() *string {
	if o == nil {
		return nil
	}
	return o.PrevCursor
}

func (o *ListAtsOfferResponseBody) GetNextCursor() *string {
	if o == nil {
		return nil
	}
	return o.NextCursor
}

func (o *ListAtsOfferResponseBody) GetData() []components.UnifiedAtsOfferOutput {
	if o == nil {
		return []components.UnifiedAtsOfferOutput{}
	}
	return o.Data
}

type ListAtsOfferResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	Object   *ListAtsOfferResponseBody

	Next func() (*ListAtsOfferResponse, error)
}

func (o *ListAtsOfferResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *ListAtsOfferResponse) GetObject() *ListAtsOfferResponseBody {
	if o == nil {
		return nil
	}
	return o.Object
}
