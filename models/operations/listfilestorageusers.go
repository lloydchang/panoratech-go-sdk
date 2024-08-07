// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/panoratech/go-sdk/internal/utils"
	"github.com/panoratech/go-sdk/models/components"
)

type ListFilestorageUsersRequest struct {
	// The connection token
	XConnectionToken string `header:"style=simple,explode=false,name=x-connection-token"`
	// Set to true to include data from the original software.
	RemoteData *bool `queryParam:"style=form,explode=true,name=remote_data"`
	// Set to get the number of records.
	Limit *float64 `default:"30" queryParam:"style=form,explode=true,name=limit"`
	// Set to get the number of records after this cursor.
	Cursor *string `queryParam:"style=form,explode=true,name=cursor"`
}

func (l ListFilestorageUsersRequest) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(l, "", false)
}

func (l *ListFilestorageUsersRequest) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &l, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *ListFilestorageUsersRequest) GetXConnectionToken() string {
	if o == nil {
		return ""
	}
	return o.XConnectionToken
}

func (o *ListFilestorageUsersRequest) GetRemoteData() *bool {
	if o == nil {
		return nil
	}
	return o.RemoteData
}

func (o *ListFilestorageUsersRequest) GetLimit() *float64 {
	if o == nil {
		return nil
	}
	return o.Limit
}

func (o *ListFilestorageUsersRequest) GetCursor() *string {
	if o == nil {
		return nil
	}
	return o.Cursor
}

type ListFilestorageUsersResponseBody struct {
	PrevCursor *string                                   `json:"prev_cursor"`
	NextCursor *string                                   `json:"next_cursor"`
	Data       []components.UnifiedFilestorageUserOutput `json:"data"`
}

func (o *ListFilestorageUsersResponseBody) GetPrevCursor() *string {
	if o == nil {
		return nil
	}
	return o.PrevCursor
}

func (o *ListFilestorageUsersResponseBody) GetNextCursor() *string {
	if o == nil {
		return nil
	}
	return o.NextCursor
}

func (o *ListFilestorageUsersResponseBody) GetData() []components.UnifiedFilestorageUserOutput {
	if o == nil {
		return []components.UnifiedFilestorageUserOutput{}
	}
	return o.Data
}

type ListFilestorageUsersResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	Object   *ListFilestorageUsersResponseBody

	Next func() (*ListFilestorageUsersResponse, error)
}

func (o *ListFilestorageUsersResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *ListFilestorageUsersResponse) GetObject() *ListFilestorageUsersResponseBody {
	if o == nil {
		return nil
	}
	return o.Object
}
