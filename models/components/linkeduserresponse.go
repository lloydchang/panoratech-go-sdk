// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

type LinkedUserResponse struct {
	IDLinkedUser       *string `json:"id_linked_user"`
	LinkedUserOriginID *string `json:"linked_user_origin_id"`
	Alias              *string `json:"alias"`
	IDProject          *string `json:"id_project"`
}

func (o *LinkedUserResponse) GetIDLinkedUser() *string {
	if o == nil {
		return nil
	}
	return o.IDLinkedUser
}

func (o *LinkedUserResponse) GetLinkedUserOriginID() *string {
	if o == nil {
		return nil
	}
	return o.LinkedUserOriginID
}

func (o *LinkedUserResponse) GetAlias() *string {
	if o == nil {
		return nil
	}
	return o.Alias
}

func (o *LinkedUserResponse) GetIDProject() *string {
	if o == nil {
		return nil
	}
	return o.IDProject
}