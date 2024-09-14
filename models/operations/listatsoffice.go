// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/panoratech/go-sdk/internal/utils"
	"github.com/panoratech/go-sdk/models/components"
)

type ListAtsOfficeRequest struct {
	// The connection token
	XConnectionToken string `header:"style=simple,explode=false,name=x-connection-token"`
	// Set to true to include data from the original software.
	RemoteData *bool `queryParam:"style=form,explode=true,name=remote_data"`
	// Set to get the number of records.
	Limit *float64 `default:"50" queryParam:"style=form,explode=true,name=limit"`
	// Set to get the number of records after this cursor.
	Cursor *string `queryParam:"style=form,explode=true,name=cursor"`
}

func (l ListAtsOfficeRequest) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(l, "", false)
}

func (l *ListAtsOfficeRequest) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &l, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *ListAtsOfficeRequest) GetXConnectionToken() string {
	if o == nil {
		return ""
	}
	return o.XConnectionToken
}

func (o *ListAtsOfficeRequest) GetRemoteData() *bool {
	if o == nil {
		return nil
	}
	return o.RemoteData
}

func (o *ListAtsOfficeRequest) GetLimit() *float64 {
	if o == nil {
		return nil
	}
	return o.Limit
}

func (o *ListAtsOfficeRequest) GetCursor() *string {
	if o == nil {
		return nil
	}
	return o.Cursor
}

type ListAtsOfficeResponseBody struct {
	PrevCursor *string                             `json:"prev_cursor"`
	NextCursor *string                             `json:"next_cursor"`
	Data       []components.UnifiedAtsOfficeOutput `json:"data"`
}

func (o *ListAtsOfficeResponseBody) GetPrevCursor() *string {
	if o == nil {
		return nil
	}
	return o.PrevCursor
}

func (o *ListAtsOfficeResponseBody) GetNextCursor() *string {
	if o == nil {
		return nil
	}
	return o.NextCursor
}

func (o *ListAtsOfficeResponseBody) GetData() []components.UnifiedAtsOfficeOutput {
	if o == nil {
		return []components.UnifiedAtsOfficeOutput{}
	}
	return o.Data
}

type ListAtsOfficeResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	Object   *ListAtsOfficeResponseBody

	Next func() (*ListAtsOfficeResponse, error)
}

func (o *ListAtsOfficeResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *ListAtsOfficeResponse) GetObject() *ListAtsOfficeResponseBody {
	if o == nil {
		return nil
	}
	return o.Object
}
