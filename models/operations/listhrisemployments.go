// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/panoratech/go-sdk/internal/utils"
	"github.com/panoratech/go-sdk/models/components"
)

type ListHrisEmploymentsRequest struct {
	// The connection token
	XConnectionToken string `header:"style=simple,explode=false,name=x-connection-token"`
	// Set to true to include data from the original software.
	RemoteData *bool `queryParam:"style=form,explode=true,name=remote_data"`
	// Set to get the number of records.
	Limit *float64 `default:"30" queryParam:"style=form,explode=true,name=limit"`
	// Set to get the number of records after this cursor.
	Cursor *string `queryParam:"style=form,explode=true,name=cursor"`
}

func (l ListHrisEmploymentsRequest) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(l, "", false)
}

func (l *ListHrisEmploymentsRequest) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &l, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *ListHrisEmploymentsRequest) GetXConnectionToken() string {
	if o == nil {
		return ""
	}
	return o.XConnectionToken
}

func (o *ListHrisEmploymentsRequest) GetRemoteData() *bool {
	if o == nil {
		return nil
	}
	return o.RemoteData
}

func (o *ListHrisEmploymentsRequest) GetLimit() *float64 {
	if o == nil {
		return nil
	}
	return o.Limit
}

func (o *ListHrisEmploymentsRequest) GetCursor() *string {
	if o == nil {
		return nil
	}
	return o.Cursor
}

type ListHrisEmploymentsResponseBody struct {
	PrevCursor *string                                  `json:"prev_cursor"`
	NextCursor *string                                  `json:"next_cursor"`
	Data       []components.UnifiedHrisEmploymentOutput `json:"data"`
}

func (o *ListHrisEmploymentsResponseBody) GetPrevCursor() *string {
	if o == nil {
		return nil
	}
	return o.PrevCursor
}

func (o *ListHrisEmploymentsResponseBody) GetNextCursor() *string {
	if o == nil {
		return nil
	}
	return o.NextCursor
}

func (o *ListHrisEmploymentsResponseBody) GetData() []components.UnifiedHrisEmploymentOutput {
	if o == nil {
		return []components.UnifiedHrisEmploymentOutput{}
	}
	return o.Data
}

type ListHrisEmploymentsResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	Object   *ListHrisEmploymentsResponseBody

	Next func() (*ListHrisEmploymentsResponse, error)
}

func (o *ListHrisEmploymentsResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *ListHrisEmploymentsResponse) GetObject() *ListHrisEmploymentsResponseBody {
	if o == nil {
		return nil
	}
	return o.Object
}
