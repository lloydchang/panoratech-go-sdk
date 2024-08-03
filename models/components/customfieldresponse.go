// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

import (
	"github.com/panoratech/go-sdk/internal/utils"
	"time"
)

type CustomFieldResponse struct {
	// Attribute Id
	IDAttribute *string `json:"id_attribute"`
	// Attribute Status
	Status *string `json:"status"`
	// Attribute Ressource Owner Type
	RessourceOwnerType *string `json:"ressource_owner_type"`
	// Attribute Slug
	Slug *string `json:"slug"`
	// Attribute Description
	Description *string `json:"description"`
	// Attribute Data Type
	DataType *string `json:"data_type"`
	// Attribute Remote Id
	RemoteID *string `json:"remote_id"`
	// Attribute Source
	Source *string `json:"source"`
	// Attribute Id Entity
	IDEntity *string `json:"id_entity"`
	// Attribute Id Project
	IDProject *string `json:"id_project"`
	// Attribute Scope
	Scope *string `json:"scope"`
	// Attribute Id Consumer
	IDConsumer *string `json:"id_consumer"`
	// Attribute Created Date
	CreatedAt *time.Time `json:"created_at"`
	// Attribute Modified Date
	ModifiedAt *time.Time `json:"modified_at"`
}

func (c CustomFieldResponse) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(c, "", false)
}

func (c *CustomFieldResponse) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &c, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *CustomFieldResponse) GetIDAttribute() *string {
	if o == nil {
		return nil
	}
	return o.IDAttribute
}

func (o *CustomFieldResponse) GetStatus() *string {
	if o == nil {
		return nil
	}
	return o.Status
}

func (o *CustomFieldResponse) GetRessourceOwnerType() *string {
	if o == nil {
		return nil
	}
	return o.RessourceOwnerType
}

func (o *CustomFieldResponse) GetSlug() *string {
	if o == nil {
		return nil
	}
	return o.Slug
}

func (o *CustomFieldResponse) GetDescription() *string {
	if o == nil {
		return nil
	}
	return o.Description
}

func (o *CustomFieldResponse) GetDataType() *string {
	if o == nil {
		return nil
	}
	return o.DataType
}

func (o *CustomFieldResponse) GetRemoteID() *string {
	if o == nil {
		return nil
	}
	return o.RemoteID
}

func (o *CustomFieldResponse) GetSource() *string {
	if o == nil {
		return nil
	}
	return o.Source
}

func (o *CustomFieldResponse) GetIDEntity() *string {
	if o == nil {
		return nil
	}
	return o.IDEntity
}

func (o *CustomFieldResponse) GetIDProject() *string {
	if o == nil {
		return nil
	}
	return o.IDProject
}

func (o *CustomFieldResponse) GetScope() *string {
	if o == nil {
		return nil
	}
	return o.Scope
}

func (o *CustomFieldResponse) GetIDConsumer() *string {
	if o == nil {
		return nil
	}
	return o.IDConsumer
}

func (o *CustomFieldResponse) GetCreatedAt() *time.Time {
	if o == nil {
		return nil
	}
	return o.CreatedAt
}

func (o *CustomFieldResponse) GetModifiedAt() *time.Time {
	if o == nil {
		return nil
	}
	return o.ModifiedAt
}
