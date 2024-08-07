// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/panoratech/go-sdk/internal/utils"
	"github.com/panoratech/go-sdk/models/components"
)

type ListAccountingAttachmentsRequest struct {
	// The connection token
	XConnectionToken string `header:"style=simple,explode=false,name=x-connection-token"`
	// Set to true to include data from the original software.
	RemoteData *bool `queryParam:"style=form,explode=true,name=remote_data"`
	// Set to get the number of records.
	Limit *float64 `default:"30" queryParam:"style=form,explode=true,name=limit"`
	// Set to get the number of records after this cursor.
	Cursor *string `queryParam:"style=form,explode=true,name=cursor"`
}

func (l ListAccountingAttachmentsRequest) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(l, "", false)
}

func (l *ListAccountingAttachmentsRequest) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &l, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *ListAccountingAttachmentsRequest) GetXConnectionToken() string {
	if o == nil {
		return ""
	}
	return o.XConnectionToken
}

func (o *ListAccountingAttachmentsRequest) GetRemoteData() *bool {
	if o == nil {
		return nil
	}
	return o.RemoteData
}

func (o *ListAccountingAttachmentsRequest) GetLimit() *float64 {
	if o == nil {
		return nil
	}
	return o.Limit
}

func (o *ListAccountingAttachmentsRequest) GetCursor() *string {
	if o == nil {
		return nil
	}
	return o.Cursor
}

type ListAccountingAttachmentsResponseBody struct {
	PrevCursor *string                                        `json:"prev_cursor"`
	NextCursor *string                                        `json:"next_cursor"`
	Data       []components.UnifiedAccountingAttachmentOutput `json:"data"`
}

func (o *ListAccountingAttachmentsResponseBody) GetPrevCursor() *string {
	if o == nil {
		return nil
	}
	return o.PrevCursor
}

func (o *ListAccountingAttachmentsResponseBody) GetNextCursor() *string {
	if o == nil {
		return nil
	}
	return o.NextCursor
}

func (o *ListAccountingAttachmentsResponseBody) GetData() []components.UnifiedAccountingAttachmentOutput {
	if o == nil {
		return []components.UnifiedAccountingAttachmentOutput{}
	}
	return o.Data
}

type ListAccountingAttachmentsResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	Object   *ListAccountingAttachmentsResponseBody

	Next func() (*ListAccountingAttachmentsResponse, error)
}

func (o *ListAccountingAttachmentsResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *ListAccountingAttachmentsResponse) GetObject() *ListAccountingAttachmentsResponseBody {
	if o == nil {
		return nil
	}
	return o.Object
}
