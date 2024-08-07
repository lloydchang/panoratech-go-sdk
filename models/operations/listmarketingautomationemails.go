// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/panoratech/go-sdk/internal/utils"
	"github.com/panoratech/go-sdk/models/components"
)

type ListMarketingautomationEmailsRequest struct {
	// The connection token
	XConnectionToken string `header:"style=simple,explode=false,name=x-connection-token"`
	// Set to true to include data from the original software.
	RemoteData *bool `queryParam:"style=form,explode=true,name=remote_data"`
	// Set to get the number of records.
	Limit *float64 `default:"30" queryParam:"style=form,explode=true,name=limit"`
	// Set to get the number of records after this cursor.
	Cursor *string `queryParam:"style=form,explode=true,name=cursor"`
}

func (l ListMarketingautomationEmailsRequest) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(l, "", false)
}

func (l *ListMarketingautomationEmailsRequest) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &l, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *ListMarketingautomationEmailsRequest) GetXConnectionToken() string {
	if o == nil {
		return ""
	}
	return o.XConnectionToken
}

func (o *ListMarketingautomationEmailsRequest) GetRemoteData() *bool {
	if o == nil {
		return nil
	}
	return o.RemoteData
}

func (o *ListMarketingautomationEmailsRequest) GetLimit() *float64 {
	if o == nil {
		return nil
	}
	return o.Limit
}

func (o *ListMarketingautomationEmailsRequest) GetCursor() *string {
	if o == nil {
		return nil
	}
	return o.Cursor
}

type ListMarketingautomationEmailsResponseBody struct {
	PrevCursor *string                                            `json:"prev_cursor"`
	NextCursor *string                                            `json:"next_cursor"`
	Data       []components.UnifiedMarketingautomationEmailOutput `json:"data"`
}

func (o *ListMarketingautomationEmailsResponseBody) GetPrevCursor() *string {
	if o == nil {
		return nil
	}
	return o.PrevCursor
}

func (o *ListMarketingautomationEmailsResponseBody) GetNextCursor() *string {
	if o == nil {
		return nil
	}
	return o.NextCursor
}

func (o *ListMarketingautomationEmailsResponseBody) GetData() []components.UnifiedMarketingautomationEmailOutput {
	if o == nil {
		return []components.UnifiedMarketingautomationEmailOutput{}
	}
	return o.Data
}

type ListMarketingautomationEmailsResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	Object   *ListMarketingautomationEmailsResponseBody

	Next func() (*ListMarketingautomationEmailsResponse, error)
}

func (o *ListMarketingautomationEmailsResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *ListMarketingautomationEmailsResponse) GetObject() *ListMarketingautomationEmailsResponseBody {
	if o == nil {
		return nil
	}
	return o.Object
}
