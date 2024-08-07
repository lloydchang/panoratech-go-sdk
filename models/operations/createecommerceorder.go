// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/panoratech/go-sdk/models/components"
)

type CreateEcommerceOrderRequest struct {
	// The connection token
	XConnectionToken string `header:"style=simple,explode=false,name=x-connection-token"`
	// Set to true to include data from the original Accounting software.
	RemoteData                 *bool                                 `queryParam:"style=form,explode=true,name=remote_data"`
	UnifiedEcommerceOrderInput components.UnifiedEcommerceOrderInput `request:"mediaType=application/json"`
}

func (o *CreateEcommerceOrderRequest) GetXConnectionToken() string {
	if o == nil {
		return ""
	}
	return o.XConnectionToken
}

func (o *CreateEcommerceOrderRequest) GetRemoteData() *bool {
	if o == nil {
		return nil
	}
	return o.RemoteData
}

func (o *CreateEcommerceOrderRequest) GetUnifiedEcommerceOrderInput() components.UnifiedEcommerceOrderInput {
	if o == nil {
		return components.UnifiedEcommerceOrderInput{}
	}
	return o.UnifiedEcommerceOrderInput
}

type CreateEcommerceOrderResponse struct {
	HTTPMeta                    components.HTTPMetadata `json:"-"`
	UnifiedEcommerceOrderOutput *components.UnifiedEcommerceOrderOutput
}

func (o *CreateEcommerceOrderResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *CreateEcommerceOrderResponse) GetUnifiedEcommerceOrderOutput() *components.UnifiedEcommerceOrderOutput {
	if o == nil {
		return nil
	}
	return o.UnifiedEcommerceOrderOutput
}
