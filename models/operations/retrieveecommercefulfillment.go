// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/panoratech/go-sdk/models/components"
)

type RetrieveEcommerceFulfillmentRequest struct {
	// The connection token
	XConnectionToken string `header:"style=simple,explode=false,name=x-connection-token"`
	// id of the fulfillment you want to retrieve.
	ID string `pathParam:"style=simple,explode=false,name=id"`
	// Set to true to include data from the original Ats software.
	RemoteData *bool `queryParam:"style=form,explode=true,name=remote_data"`
}

func (o *RetrieveEcommerceFulfillmentRequest) GetXConnectionToken() string {
	if o == nil {
		return ""
	}
	return o.XConnectionToken
}

func (o *RetrieveEcommerceFulfillmentRequest) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *RetrieveEcommerceFulfillmentRequest) GetRemoteData() *bool {
	if o == nil {
		return nil
	}
	return o.RemoteData
}

type RetrieveEcommerceFulfillmentResponse struct {
	HTTPMeta                          components.HTTPMetadata `json:"-"`
	UnifiedEcommerceFulfillmentOutput *components.UnifiedEcommerceFulfillmentOutput
}

func (o *RetrieveEcommerceFulfillmentResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *RetrieveEcommerceFulfillmentResponse) GetUnifiedEcommerceFulfillmentOutput() *components.UnifiedEcommerceFulfillmentOutput {
	if o == nil {
		return nil
	}
	return o.UnifiedEcommerceFulfillmentOutput
}