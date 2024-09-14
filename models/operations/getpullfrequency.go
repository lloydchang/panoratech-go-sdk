// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/panoratech/go-sdk/models/components"
)

type GetPullFrequencyResponse struct {
	HTTPMeta               components.HTTPMetadata `json:"-"`
	UpdatePullFrequencyDto *components.UpdatePullFrequencyDto
}

func (o *GetPullFrequencyResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *GetPullFrequencyResponse) GetUpdatePullFrequencyDto() *components.UpdatePullFrequencyDto {
	if o == nil {
		return nil
	}
	return o.UpdatePullFrequencyDto
}
