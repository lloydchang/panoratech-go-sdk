// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/panoratech/go-sdk/models/components"
)

type ListCrmStagesRequest struct {
	// The connection token
	XConnectionToken string `header:"style=simple,explode=false,name=x-connection-token"`
	// Set to true to include data from the original software.
	RemoteData *bool `queryParam:"style=form,explode=true,name=remote_data"`
	// Set to get the number of records.
	Limit *float64 `queryParam:"style=form,explode=true,name=limit"`
	// Set to get the number of records after this cursor.
	Cursor *string `queryParam:"style=form,explode=true,name=cursor"`
}

func (o *ListCrmStagesRequest) GetXConnectionToken() string {
	if o == nil {
		return ""
	}
	return o.XConnectionToken
}

func (o *ListCrmStagesRequest) GetRemoteData() *bool {
	if o == nil {
		return nil
	}
	return o.RemoteData
}

func (o *ListCrmStagesRequest) GetLimit() *float64 {
	if o == nil {
		return nil
	}
	return o.Limit
}

func (o *ListCrmStagesRequest) GetCursor() *string {
	if o == nil {
		return nil
	}
	return o.Cursor
}

type ListCrmStagesResponseBody struct {
	PrevCursor *string                            `json:"prev_cursor"`
	NextCursor *string                            `json:"next_cursor"`
	Data       []components.UnifiedCrmStageOutput `json:"data"`
}

func (o *ListCrmStagesResponseBody) GetPrevCursor() *string {
	if o == nil {
		return nil
	}
	return o.PrevCursor
}

func (o *ListCrmStagesResponseBody) GetNextCursor() *string {
	if o == nil {
		return nil
	}
	return o.NextCursor
}

func (o *ListCrmStagesResponseBody) GetData() []components.UnifiedCrmStageOutput {
	if o == nil {
		return []components.UnifiedCrmStageOutput{}
	}
	return o.Data
}

type ListCrmStagesResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	Object   *ListCrmStagesResponseBody

	Next func() (*ListCrmStagesResponse, error)
}

func (o *ListCrmStagesResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *ListCrmStagesResponse) GetObject() *ListCrmStagesResponseBody {
	if o == nil {
		return nil
	}
	return o.Object
}
