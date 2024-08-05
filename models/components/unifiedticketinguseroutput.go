// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

import (
	"github.com/panoratech/go-sdk/internal/utils"
	"time"
)

type UnifiedTicketingUserOutput struct {
	// The name of the user
	Name *string `json:"name"`
	// The email address of the user
	EmailAddress *string `json:"email_address"`
	// The teams whose the user is part of
	Teams []string `json:"teams,omitempty"`
	// The account or organization the user is part of
	AccountID *string `json:"account_id,omitempty"`
	// The custom field mappings of the user between the remote 3rd party & Panora
	FieldMappings map[string]any `json:"field_mappings,omitempty"`
	// The UUID of the user
	ID *string `json:"id,omitempty"`
	// The id of the user in the context of the 3rd Party
	RemoteID *string `json:"remote_id,omitempty"`
	// The remote data of the user in the context of the 3rd Party
	RemoteData map[string]any `json:"remote_data,omitempty"`
	// The created date of the object
	CreatedAt *time.Time `json:"created_at,omitempty"`
	// The modified date of the object
	ModifiedAt *time.Time `json:"modified_at,omitempty"`
}

func (u UnifiedTicketingUserOutput) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(u, "", false)
}

func (u *UnifiedTicketingUserOutput) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &u, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *UnifiedTicketingUserOutput) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

func (o *UnifiedTicketingUserOutput) GetEmailAddress() *string {
	if o == nil {
		return nil
	}
	return o.EmailAddress
}

func (o *UnifiedTicketingUserOutput) GetTeams() []string {
	if o == nil {
		return nil
	}
	return o.Teams
}

func (o *UnifiedTicketingUserOutput) GetAccountID() *string {
	if o == nil {
		return nil
	}
	return o.AccountID
}

func (o *UnifiedTicketingUserOutput) GetFieldMappings() map[string]any {
	if o == nil {
		return nil
	}
	return o.FieldMappings
}

func (o *UnifiedTicketingUserOutput) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *UnifiedTicketingUserOutput) GetRemoteID() *string {
	if o == nil {
		return nil
	}
	return o.RemoteID
}

func (o *UnifiedTicketingUserOutput) GetRemoteData() map[string]any {
	if o == nil {
		return nil
	}
	return o.RemoteData
}

func (o *UnifiedTicketingUserOutput) GetCreatedAt() *time.Time {
	if o == nil {
		return nil
	}
	return o.CreatedAt
}

func (o *UnifiedTicketingUserOutput) GetModifiedAt() *time.Time {
	if o == nil {
		return nil
	}
	return o.ModifiedAt
}
