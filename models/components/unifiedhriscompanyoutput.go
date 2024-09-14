// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

import (
	"github.com/panoratech/go-sdk/internal/utils"
	"time"
)

// UnifiedHrisCompanyOutputFieldMappings - The custom field mappings of the object between the remote 3rd party & Panora
type UnifiedHrisCompanyOutputFieldMappings struct {
}

// UnifiedHrisCompanyOutputRemoteData - The remote data of the company in the context of the 3rd Party
type UnifiedHrisCompanyOutputRemoteData struct {
}

type UnifiedHrisCompanyOutput struct {
	// The legal name of the company
	LegalName *string `json:"legal_name,omitempty"`
	// UUIDs of the of the Location associated with the company
	Locations []string `json:"locations,omitempty"`
	// The display name of the company
	DisplayName *string `json:"display_name,omitempty"`
	// The Employer Identification Numbers (EINs) of the company
	Eins []string `json:"eins,omitempty"`
	// The custom field mappings of the object between the remote 3rd party & Panora
	FieldMappings *UnifiedHrisCompanyOutputFieldMappings `json:"field_mappings,omitempty"`
	// The UUID of the company record
	ID *string `json:"id,omitempty"`
	// The remote ID of the company in the context of the 3rd Party
	RemoteID *string `json:"remote_id,omitempty"`
	// The remote data of the company in the context of the 3rd Party
	RemoteData *UnifiedHrisCompanyOutputRemoteData `json:"remote_data,omitempty"`
	// The date when the company was created in the 3rd party system
	RemoteCreatedAt *time.Time `json:"remote_created_at,omitempty"`
	// The created date of the company record
	CreatedAt *time.Time `json:"created_at,omitempty"`
	// The last modified date of the company record
	ModifiedAt *time.Time `json:"modified_at,omitempty"`
	// Indicates if the company was deleted in the remote system
	RemoteWasDeleted *bool `json:"remote_was_deleted,omitempty"`
}

func (u UnifiedHrisCompanyOutput) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(u, "", false)
}

func (u *UnifiedHrisCompanyOutput) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &u, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *UnifiedHrisCompanyOutput) GetLegalName() *string {
	if o == nil {
		return nil
	}
	return o.LegalName
}

func (o *UnifiedHrisCompanyOutput) GetLocations() []string {
	if o == nil {
		return nil
	}
	return o.Locations
}

func (o *UnifiedHrisCompanyOutput) GetDisplayName() *string {
	if o == nil {
		return nil
	}
	return o.DisplayName
}

func (o *UnifiedHrisCompanyOutput) GetEins() []string {
	if o == nil {
		return nil
	}
	return o.Eins
}

func (o *UnifiedHrisCompanyOutput) GetFieldMappings() *UnifiedHrisCompanyOutputFieldMappings {
	if o == nil {
		return nil
	}
	return o.FieldMappings
}

func (o *UnifiedHrisCompanyOutput) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *UnifiedHrisCompanyOutput) GetRemoteID() *string {
	if o == nil {
		return nil
	}
	return o.RemoteID
}

func (o *UnifiedHrisCompanyOutput) GetRemoteData() *UnifiedHrisCompanyOutputRemoteData {
	if o == nil {
		return nil
	}
	return o.RemoteData
}

func (o *UnifiedHrisCompanyOutput) GetRemoteCreatedAt() *time.Time {
	if o == nil {
		return nil
	}
	return o.RemoteCreatedAt
}

func (o *UnifiedHrisCompanyOutput) GetCreatedAt() *time.Time {
	if o == nil {
		return nil
	}
	return o.CreatedAt
}

func (o *UnifiedHrisCompanyOutput) GetModifiedAt() *time.Time {
	if o == nil {
		return nil
	}
	return o.ModifiedAt
}

func (o *UnifiedHrisCompanyOutput) GetRemoteWasDeleted() *bool {
	if o == nil {
		return nil
	}
	return o.RemoteWasDeleted
}
