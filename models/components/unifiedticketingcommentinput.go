// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

import (
	"encoding/json"
	"fmt"
)

// UnifiedTicketingCommentInputCreatorType - The creator type of the comment. Authorized values are either USER or CONTACT
type UnifiedTicketingCommentInputCreatorType string

const (
	UnifiedTicketingCommentInputCreatorTypeUser    UnifiedTicketingCommentInputCreatorType = "USER"
	UnifiedTicketingCommentInputCreatorTypeContact UnifiedTicketingCommentInputCreatorType = "CONTACT"
)

func (e UnifiedTicketingCommentInputCreatorType) ToPointer() *UnifiedTicketingCommentInputCreatorType {
	return &e
}
func (e *UnifiedTicketingCommentInputCreatorType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "USER":
		fallthrough
	case "CONTACT":
		*e = UnifiedTicketingCommentInputCreatorType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for UnifiedTicketingCommentInputCreatorType: %v", v)
	}
}

type UnifiedTicketingCommentInput struct {
	// The body of the comment
	Body *string `json:"body"`
	// The html body of the comment
	HTMLBody *string `json:"html_body,omitempty"`
	// The public status of the comment
	IsPrivate *bool `json:"is_private,omitempty"`
	// The creator type of the comment. Authorized values are either USER or CONTACT
	CreatorType *UnifiedTicketingCommentInputCreatorType `json:"creator_type,omitempty"`
	// The UUID of the ticket the comment is tied to
	TicketID *string `json:"ticket_id,omitempty"`
	// The UUID of the contact which the comment belongs to (if no user_id specified)
	ContactID *string `json:"contact_id,omitempty"`
	// The UUID of the user which the comment belongs to (if no contact_id specified)
	UserID *string `json:"user_id,omitempty"`
	// The attachements UUIDs tied to the comment
	Attachments []string `json:"attachments,omitempty"`
}

func (o *UnifiedTicketingCommentInput) GetBody() *string {
	if o == nil {
		return nil
	}
	return o.Body
}

func (o *UnifiedTicketingCommentInput) GetHTMLBody() *string {
	if o == nil {
		return nil
	}
	return o.HTMLBody
}

func (o *UnifiedTicketingCommentInput) GetIsPrivate() *bool {
	if o == nil {
		return nil
	}
	return o.IsPrivate
}

func (o *UnifiedTicketingCommentInput) GetCreatorType() *UnifiedTicketingCommentInputCreatorType {
	if o == nil {
		return nil
	}
	return o.CreatorType
}

func (o *UnifiedTicketingCommentInput) GetTicketID() *string {
	if o == nil {
		return nil
	}
	return o.TicketID
}

func (o *UnifiedTicketingCommentInput) GetContactID() *string {
	if o == nil {
		return nil
	}
	return o.ContactID
}

func (o *UnifiedTicketingCommentInput) GetUserID() *string {
	if o == nil {
		return nil
	}
	return o.UserID
}

func (o *UnifiedTicketingCommentInput) GetAttachments() []string {
	if o == nil {
		return nil
	}
	return o.Attachments
}
