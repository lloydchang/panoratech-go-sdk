// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/panoratech/go-sdk/internal/utils"
	"github.com/panoratech/go-sdk/models/components"
)

type ListTicketingCommentsRequest struct {
	// The connection token
	XConnectionToken string `header:"style=simple,explode=false,name=x-connection-token"`
	// Set to true to include data from the original software.
	RemoteData *bool `queryParam:"style=form,explode=true,name=remote_data"`
	// Set to get the number of records.
	Limit *float64 `default:"30" queryParam:"style=form,explode=true,name=limit"`
	// Set to get the number of records after this cursor.
	Cursor *string `queryParam:"style=form,explode=true,name=cursor"`
}

func (l ListTicketingCommentsRequest) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(l, "", false)
}

func (l *ListTicketingCommentsRequest) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &l, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *ListTicketingCommentsRequest) GetXConnectionToken() string {
	if o == nil {
		return ""
	}
	return o.XConnectionToken
}

func (o *ListTicketingCommentsRequest) GetRemoteData() *bool {
	if o == nil {
		return nil
	}
	return o.RemoteData
}

func (o *ListTicketingCommentsRequest) GetLimit() *float64 {
	if o == nil {
		return nil
	}
	return o.Limit
}

func (o *ListTicketingCommentsRequest) GetCursor() *string {
	if o == nil {
		return nil
	}
	return o.Cursor
}

type ListTicketingCommentsResponseBody struct {
	PrevCursor *string                                    `json:"prev_cursor"`
	NextCursor *string                                    `json:"next_cursor"`
	Data       []components.UnifiedTicketingCommentOutput `json:"data"`
}

func (o *ListTicketingCommentsResponseBody) GetPrevCursor() *string {
	if o == nil {
		return nil
	}
	return o.PrevCursor
}

func (o *ListTicketingCommentsResponseBody) GetNextCursor() *string {
	if o == nil {
		return nil
	}
	return o.NextCursor
}

func (o *ListTicketingCommentsResponseBody) GetData() []components.UnifiedTicketingCommentOutput {
	if o == nil {
		return []components.UnifiedTicketingCommentOutput{}
	}
	return o.Data
}

type ListTicketingCommentsResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	Object   *ListTicketingCommentsResponseBody

	Next func() (*ListTicketingCommentsResponse, error)
}

func (o *ListTicketingCommentsResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *ListTicketingCommentsResponse) GetObject() *ListTicketingCommentsResponseBody {
	if o == nil {
		return nil
	}
	return o.Object
}
