// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

type CreateLinkedUserDto struct {
	LinkedUserOriginID *string `json:"linked_user_origin_id"`
	Alias              *string `json:"alias"`
}

func (o *CreateLinkedUserDto) GetLinkedUserOriginID() *string {
	if o == nil {
		return nil
	}
	return o.LinkedUserOriginID
}

func (o *CreateLinkedUserDto) GetAlias() *string {
	if o == nil {
		return nil
	}
	return o.Alias
}
