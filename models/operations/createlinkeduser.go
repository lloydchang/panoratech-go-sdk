// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/panoratech/go-sdk/models/components"
)

type CreateLinkedUserResponse struct {
	HTTPMeta           components.HTTPMetadata `json:"-"`
	LinkedUserResponse *components.LinkedUserResponse
}

func (o *CreateLinkedUserResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *CreateLinkedUserResponse) GetLinkedUserResponse() *components.LinkedUserResponse {
	if o == nil {
		return nil
	}
	return o.LinkedUserResponse
}
