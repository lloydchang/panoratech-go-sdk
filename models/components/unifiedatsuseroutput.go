// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

import (
	"encoding/json"
	"fmt"
	"github.com/panoratech/go-sdk/internal/utils"
	"time"
)

// AccessRole - The access role of the user
type AccessRole string

const (
	AccessRoleSuperAdmin        AccessRole = "SUPER_ADMIN"
	AccessRoleAdmin             AccessRole = "ADMIN"
	AccessRoleTeamMember        AccessRole = "TEAM_MEMBER"
	AccessRoleLimitedTeamMember AccessRole = "LIMITED_TEAM_MEMBER"
	AccessRoleInterviewer       AccessRole = "INTERVIEWER"
)

func (e AccessRole) ToPointer() *AccessRole {
	return &e
}
func (e *AccessRole) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "SUPER_ADMIN":
		fallthrough
	case "ADMIN":
		fallthrough
	case "TEAM_MEMBER":
		fallthrough
	case "LIMITED_TEAM_MEMBER":
		fallthrough
	case "INTERVIEWER":
		*e = AccessRole(v)
		return nil
	default:
		return fmt.Errorf("invalid value for AccessRole: %v", v)
	}
}

type UnifiedAtsUserOutput struct {
	// The first name of the user
	FirstName *string `json:"first_name,omitempty"`
	// The last name of the user
	LastName *string `json:"last_name,omitempty"`
	// The email of the user
	Email *string `json:"email,omitempty"`
	// Whether the user is disabled
	Disabled *bool `json:"disabled,omitempty"`
	// The access role of the user
	AccessRole *AccessRole `json:"access_role,omitempty"`
	// The remote creation date of the user
	RemoteCreatedAt *time.Time `json:"remote_created_at,omitempty"`
	// The remote modification date of the user
	RemoteModifiedAt *time.Time `json:"remote_modified_at,omitempty"`
	// The custom field mappings of the object between the remote 3rd party & Panora
	FieldMappings map[string]any `json:"field_mappings,omitempty"`
	// The UUID of the user
	ID *string `json:"id,omitempty"`
	// The remote ID of the user in the context of the 3rd Party
	RemoteID *string `json:"remote_id,omitempty"`
	// The remote data of the user in the context of the 3rd Party
	RemoteData map[string]any `json:"remote_data,omitempty"`
	// The created date of the object
	CreatedAt *time.Time `json:"created_at,omitempty"`
	// The modified date of the object
	ModifiedAt *time.Time `json:"modified_at,omitempty"`
}

func (u UnifiedAtsUserOutput) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(u, "", false)
}

func (u *UnifiedAtsUserOutput) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &u, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *UnifiedAtsUserOutput) GetFirstName() *string {
	if o == nil {
		return nil
	}
	return o.FirstName
}

func (o *UnifiedAtsUserOutput) GetLastName() *string {
	if o == nil {
		return nil
	}
	return o.LastName
}

func (o *UnifiedAtsUserOutput) GetEmail() *string {
	if o == nil {
		return nil
	}
	return o.Email
}

func (o *UnifiedAtsUserOutput) GetDisabled() *bool {
	if o == nil {
		return nil
	}
	return o.Disabled
}

func (o *UnifiedAtsUserOutput) GetAccessRole() *AccessRole {
	if o == nil {
		return nil
	}
	return o.AccessRole
}

func (o *UnifiedAtsUserOutput) GetRemoteCreatedAt() *time.Time {
	if o == nil {
		return nil
	}
	return o.RemoteCreatedAt
}

func (o *UnifiedAtsUserOutput) GetRemoteModifiedAt() *time.Time {
	if o == nil {
		return nil
	}
	return o.RemoteModifiedAt
}

func (o *UnifiedAtsUserOutput) GetFieldMappings() map[string]any {
	if o == nil {
		return nil
	}
	return o.FieldMappings
}

func (o *UnifiedAtsUserOutput) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *UnifiedAtsUserOutput) GetRemoteID() *string {
	if o == nil {
		return nil
	}
	return o.RemoteID
}

func (o *UnifiedAtsUserOutput) GetRemoteData() map[string]any {
	if o == nil {
		return nil
	}
	return o.RemoteData
}

func (o *UnifiedAtsUserOutput) GetCreatedAt() *time.Time {
	if o == nil {
		return nil
	}
	return o.CreatedAt
}

func (o *UnifiedAtsUserOutput) GetModifiedAt() *time.Time {
	if o == nil {
		return nil
	}
	return o.ModifiedAt
}
