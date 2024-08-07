// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/panoratech/go-sdk/internal/utils"
	"github.com/panoratech/go-sdk/models/components"
)

type ListFilestorageFileRequest struct {
	// The connection token
	XConnectionToken string `header:"style=simple,explode=false,name=x-connection-token"`
	// Set to true to include data from the original software.
	RemoteData *bool `queryParam:"style=form,explode=true,name=remote_data"`
	// Set to get the number of records.
	Limit *float64 `default:"30" queryParam:"style=form,explode=true,name=limit"`
	// Set to get the number of records after this cursor.
	Cursor *string `queryParam:"style=form,explode=true,name=cursor"`
}

func (l ListFilestorageFileRequest) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(l, "", false)
}

func (l *ListFilestorageFileRequest) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &l, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *ListFilestorageFileRequest) GetXConnectionToken() string {
	if o == nil {
		return ""
	}
	return o.XConnectionToken
}

func (o *ListFilestorageFileRequest) GetRemoteData() *bool {
	if o == nil {
		return nil
	}
	return o.RemoteData
}

func (o *ListFilestorageFileRequest) GetLimit() *float64 {
	if o == nil {
		return nil
	}
	return o.Limit
}

func (o *ListFilestorageFileRequest) GetCursor() *string {
	if o == nil {
		return nil
	}
	return o.Cursor
}

type ListFilestorageFileResponseBody struct {
	PrevCursor *string                                   `json:"prev_cursor"`
	NextCursor *string                                   `json:"next_cursor"`
	Data       []components.UnifiedFilestorageFileOutput `json:"data"`
}

func (o *ListFilestorageFileResponseBody) GetPrevCursor() *string {
	if o == nil {
		return nil
	}
	return o.PrevCursor
}

func (o *ListFilestorageFileResponseBody) GetNextCursor() *string {
	if o == nil {
		return nil
	}
	return o.NextCursor
}

func (o *ListFilestorageFileResponseBody) GetData() []components.UnifiedFilestorageFileOutput {
	if o == nil {
		return []components.UnifiedFilestorageFileOutput{}
	}
	return o.Data
}

type ListFilestorageFileResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	Object   *ListFilestorageFileResponseBody

	Next func() (*ListFilestorageFileResponse, error)
}

func (o *ListFilestorageFileResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *ListFilestorageFileResponse) GetObject() *ListFilestorageFileResponseBody {
	if o == nil {
		return nil
	}
	return o.Object
}
