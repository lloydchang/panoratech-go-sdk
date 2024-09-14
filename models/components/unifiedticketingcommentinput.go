// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

import (
	"errors"
	"fmt"
	"github.com/panoratech/go-sdk/internal/utils"
)

type UnifiedTicketingCommentInputAttachmentsType string

const (
	UnifiedTicketingCommentInputAttachmentsTypeStr                              UnifiedTicketingCommentInputAttachmentsType = "str"
	UnifiedTicketingCommentInputAttachmentsTypeUnifiedTicketingAttachmentOutput UnifiedTicketingCommentInputAttachmentsType = "UnifiedTicketingAttachmentOutput"
)

type UnifiedTicketingCommentInputAttachments struct {
	Str                              *string
	UnifiedTicketingAttachmentOutput *UnifiedTicketingAttachmentOutput

	Type UnifiedTicketingCommentInputAttachmentsType
}

func CreateUnifiedTicketingCommentInputAttachmentsStr(str string) UnifiedTicketingCommentInputAttachments {
	typ := UnifiedTicketingCommentInputAttachmentsTypeStr

	return UnifiedTicketingCommentInputAttachments{
		Str:  &str,
		Type: typ,
	}
}

func CreateUnifiedTicketingCommentInputAttachmentsUnifiedTicketingAttachmentOutput(unifiedTicketingAttachmentOutput UnifiedTicketingAttachmentOutput) UnifiedTicketingCommentInputAttachments {
	typ := UnifiedTicketingCommentInputAttachmentsTypeUnifiedTicketingAttachmentOutput

	return UnifiedTicketingCommentInputAttachments{
		UnifiedTicketingAttachmentOutput: &unifiedTicketingAttachmentOutput,
		Type:                             typ,
	}
}

func (u *UnifiedTicketingCommentInputAttachments) UnmarshalJSON(data []byte) error {

	var unifiedTicketingAttachmentOutput UnifiedTicketingAttachmentOutput = UnifiedTicketingAttachmentOutput{}
	if err := utils.UnmarshalJSON(data, &unifiedTicketingAttachmentOutput, "", true, true); err == nil {
		u.UnifiedTicketingAttachmentOutput = &unifiedTicketingAttachmentOutput
		u.Type = UnifiedTicketingCommentInputAttachmentsTypeUnifiedTicketingAttachmentOutput
		return nil
	}

	var str string = ""
	if err := utils.UnmarshalJSON(data, &str, "", true, true); err == nil {
		u.Str = &str
		u.Type = UnifiedTicketingCommentInputAttachmentsTypeStr
		return nil
	}

	return fmt.Errorf("could not unmarshal `%s` into any supported union types for UnifiedTicketingCommentInputAttachments", string(data))
}

func (u UnifiedTicketingCommentInputAttachments) MarshalJSON() ([]byte, error) {
	if u.Str != nil {
		return utils.MarshalJSON(u.Str, "", true)
	}

	if u.UnifiedTicketingAttachmentOutput != nil {
		return utils.MarshalJSON(u.UnifiedTicketingAttachmentOutput, "", true)
	}

	return nil, errors.New("could not marshal union type UnifiedTicketingCommentInputAttachments: all fields are null")
}

type UnifiedTicketingCommentInput struct {
	// The body of the comment
	Body *string `json:"body"`
	// The html body of the comment
	HTMLBody *string `json:"html_body,omitempty"`
	// The public status of the comment
	IsPrivate *bool `json:"is_private,omitempty"`
	// The creator type of the comment. Authorized values are either USER or CONTACT
	CreatorType *string `json:"creator_type,omitempty"`
	// The UUID of the ticket the comment is tied to
	TicketID *string `json:"ticket_id,omitempty"`
	// The UUID of the contact which the comment belongs to (if no user_id specified)
	ContactID *string `json:"contact_id,omitempty"`
	// The UUID of the user which the comment belongs to (if no contact_id specified)
	UserID *string `json:"user_id,omitempty"`
	// The attachements UUIDs tied to the comment
	Attachments []UnifiedTicketingCommentInputAttachments `json:"attachments,omitempty"`
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

func (o *UnifiedTicketingCommentInput) GetCreatorType() *string {
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

func (o *UnifiedTicketingCommentInput) GetAttachments() []UnifiedTicketingCommentInputAttachments {
	if o == nil {
		return nil
	}
	return o.Attachments
}
