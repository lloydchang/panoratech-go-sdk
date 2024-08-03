// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

import (
	"github.com/panoratech/go-sdk/internal/utils"
	"time"
)

type ResyncStatusDto struct {
	Timestamp *time.Time `json:"timestamp"`
	Vertical  *string    `json:"vertical"`
	Provider  *string    `json:"provider"`
	Status    *string    `json:"status"`
}

func (r ResyncStatusDto) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(r, "", false)
}

func (r *ResyncStatusDto) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &r, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *ResyncStatusDto) GetTimestamp() *time.Time {
	if o == nil {
		return nil
	}
	return o.Timestamp
}

func (o *ResyncStatusDto) GetVertical() *string {
	if o == nil {
		return nil
	}
	return o.Vertical
}

func (o *ResyncStatusDto) GetProvider() *string {
	if o == nil {
		return nil
	}
	return o.Provider
}

func (o *ResyncStatusDto) GetStatus() *string {
	if o == nil {
		return nil
	}
	return o.Status
}
