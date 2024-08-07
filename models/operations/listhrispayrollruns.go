// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/panoratech/go-sdk/internal/utils"
	"github.com/panoratech/go-sdk/models/components"
)

type ListHrisPayrollRunsRequest struct {
	// The connection token
	XConnectionToken string `header:"style=simple,explode=false,name=x-connection-token"`
	// Set to true to include data from the original software.
	RemoteData *bool `queryParam:"style=form,explode=true,name=remote_data"`
	// Set to get the number of records.
	Limit *float64 `default:"30" queryParam:"style=form,explode=true,name=limit"`
	// Set to get the number of records after this cursor.
	Cursor *string `queryParam:"style=form,explode=true,name=cursor"`
}

func (l ListHrisPayrollRunsRequest) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(l, "", false)
}

func (l *ListHrisPayrollRunsRequest) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &l, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *ListHrisPayrollRunsRequest) GetXConnectionToken() string {
	if o == nil {
		return ""
	}
	return o.XConnectionToken
}

func (o *ListHrisPayrollRunsRequest) GetRemoteData() *bool {
	if o == nil {
		return nil
	}
	return o.RemoteData
}

func (o *ListHrisPayrollRunsRequest) GetLimit() *float64 {
	if o == nil {
		return nil
	}
	return o.Limit
}

func (o *ListHrisPayrollRunsRequest) GetCursor() *string {
	if o == nil {
		return nil
	}
	return o.Cursor
}

type ListHrisPayrollRunsResponseBody struct {
	PrevCursor *string                                  `json:"prev_cursor"`
	NextCursor *string                                  `json:"next_cursor"`
	Data       []components.UnifiedHrisPayrollrunOutput `json:"data"`
}

func (o *ListHrisPayrollRunsResponseBody) GetPrevCursor() *string {
	if o == nil {
		return nil
	}
	return o.PrevCursor
}

func (o *ListHrisPayrollRunsResponseBody) GetNextCursor() *string {
	if o == nil {
		return nil
	}
	return o.NextCursor
}

func (o *ListHrisPayrollRunsResponseBody) GetData() []components.UnifiedHrisPayrollrunOutput {
	if o == nil {
		return []components.UnifiedHrisPayrollrunOutput{}
	}
	return o.Data
}

type ListHrisPayrollRunsResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	Object   *ListHrisPayrollRunsResponseBody

	Next func() (*ListHrisPayrollRunsResponse, error)
}

func (o *ListHrisPayrollRunsResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *ListHrisPayrollRunsResponse) GetObject() *ListHrisPayrollRunsResponseBody {
	if o == nil {
		return nil
	}
	return o.Object
}
