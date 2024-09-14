// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

import (
	"github.com/panoratech/go-sdk/internal/utils"
	"time"
)

// UnifiedHrisEmployeeOutputFieldMappings - The custom field mappings of the object between the remote 3rd party & Panora
type UnifiedHrisEmployeeOutputFieldMappings struct {
}

// UnifiedHrisEmployeeOutputRemoteData - The remote data of the employee in the context of the 3rd Party
type UnifiedHrisEmployeeOutputRemoteData struct {
}

type UnifiedHrisEmployeeOutput struct {
	// The groups the employee belongs to
	Groups []string `json:"groups,omitempty"`
	// UUIDs of the of the Location associated with the company
	Locations []string `json:"locations,omitempty"`
	// The employee number
	EmployeeNumber *string `json:"employee_number,omitempty"`
	// The UUID of the associated company
	CompanyID *string `json:"company_id,omitempty"`
	// The first name of the employee
	FirstName *string `json:"first_name,omitempty"`
	// The last name of the employee
	LastName *string `json:"last_name,omitempty"`
	// The preferred name of the employee
	PreferredName *string `json:"preferred_name,omitempty"`
	// The full display name of the employee
	DisplayFullName *string `json:"display_full_name,omitempty"`
	// The username of the employee
	Username *string `json:"username,omitempty"`
	// The work email of the employee
	WorkEmail *string `json:"work_email,omitempty"`
	// The personal email of the employee
	PersonalEmail *string `json:"personal_email,omitempty"`
	// The mobile phone number of the employee
	MobilePhoneNumber *string `json:"mobile_phone_number,omitempty"`
	// The employments of the employee
	Employments []string `json:"employments,omitempty"`
	// The Social Security Number of the employee
	Ssn *string `json:"ssn,omitempty"`
	// The gender of the employee
	Gender *string `json:"gender,omitempty"`
	// The ethnicity of the employee
	Ethnicity *string `json:"ethnicity,omitempty"`
	// The marital status of the employee
	MaritalStatus *string `json:"marital_status,omitempty"`
	// The date of birth of the employee
	DateOfBirth *time.Time `json:"date_of_birth,omitempty"`
	// The start date of the employee
	StartDate *time.Time `json:"start_date,omitempty"`
	// The employment status of the employee
	EmploymentStatus *string `json:"employment_status,omitempty"`
	// The termination date of the employee
	TerminationDate *time.Time `json:"termination_date,omitempty"`
	// The URL of the employee's avatar
	AvatarURL *string `json:"avatar_url,omitempty"`
	// UUID of the manager (employee) of the employee
	ManagerID *string `json:"manager_id,omitempty"`
	// The custom field mappings of the object between the remote 3rd party & Panora
	FieldMappings *UnifiedHrisEmployeeOutputFieldMappings `json:"field_mappings,omitempty"`
	// The UUID of the employee record
	ID *string `json:"id,omitempty"`
	// The remote ID of the employee in the context of the 3rd Party
	RemoteID *string `json:"remote_id,omitempty"`
	// The remote data of the employee in the context of the 3rd Party
	RemoteData *UnifiedHrisEmployeeOutputRemoteData `json:"remote_data,omitempty"`
	// The date when the employee was created in the 3rd party system
	RemoteCreatedAt *time.Time `json:"remote_created_at,omitempty"`
	// The created date of the employee record
	CreatedAt *time.Time `json:"created_at,omitempty"`
	// The last modified date of the employee record
	ModifiedAt *time.Time `json:"modified_at,omitempty"`
	// Indicates if the employee was deleted in the remote system
	RemoteWasDeleted *bool `json:"remote_was_deleted,omitempty"`
}

func (u UnifiedHrisEmployeeOutput) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(u, "", false)
}

func (u *UnifiedHrisEmployeeOutput) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &u, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *UnifiedHrisEmployeeOutput) GetGroups() []string {
	if o == nil {
		return nil
	}
	return o.Groups
}

func (o *UnifiedHrisEmployeeOutput) GetLocations() []string {
	if o == nil {
		return nil
	}
	return o.Locations
}

func (o *UnifiedHrisEmployeeOutput) GetEmployeeNumber() *string {
	if o == nil {
		return nil
	}
	return o.EmployeeNumber
}

func (o *UnifiedHrisEmployeeOutput) GetCompanyID() *string {
	if o == nil {
		return nil
	}
	return o.CompanyID
}

func (o *UnifiedHrisEmployeeOutput) GetFirstName() *string {
	if o == nil {
		return nil
	}
	return o.FirstName
}

func (o *UnifiedHrisEmployeeOutput) GetLastName() *string {
	if o == nil {
		return nil
	}
	return o.LastName
}

func (o *UnifiedHrisEmployeeOutput) GetPreferredName() *string {
	if o == nil {
		return nil
	}
	return o.PreferredName
}

func (o *UnifiedHrisEmployeeOutput) GetDisplayFullName() *string {
	if o == nil {
		return nil
	}
	return o.DisplayFullName
}

func (o *UnifiedHrisEmployeeOutput) GetUsername() *string {
	if o == nil {
		return nil
	}
	return o.Username
}

func (o *UnifiedHrisEmployeeOutput) GetWorkEmail() *string {
	if o == nil {
		return nil
	}
	return o.WorkEmail
}

func (o *UnifiedHrisEmployeeOutput) GetPersonalEmail() *string {
	if o == nil {
		return nil
	}
	return o.PersonalEmail
}

func (o *UnifiedHrisEmployeeOutput) GetMobilePhoneNumber() *string {
	if o == nil {
		return nil
	}
	return o.MobilePhoneNumber
}

func (o *UnifiedHrisEmployeeOutput) GetEmployments() []string {
	if o == nil {
		return nil
	}
	return o.Employments
}

func (o *UnifiedHrisEmployeeOutput) GetSsn() *string {
	if o == nil {
		return nil
	}
	return o.Ssn
}

func (o *UnifiedHrisEmployeeOutput) GetGender() *string {
	if o == nil {
		return nil
	}
	return o.Gender
}

func (o *UnifiedHrisEmployeeOutput) GetEthnicity() *string {
	if o == nil {
		return nil
	}
	return o.Ethnicity
}

func (o *UnifiedHrisEmployeeOutput) GetMaritalStatus() *string {
	if o == nil {
		return nil
	}
	return o.MaritalStatus
}

func (o *UnifiedHrisEmployeeOutput) GetDateOfBirth() *time.Time {
	if o == nil {
		return nil
	}
	return o.DateOfBirth
}

func (o *UnifiedHrisEmployeeOutput) GetStartDate() *time.Time {
	if o == nil {
		return nil
	}
	return o.StartDate
}

func (o *UnifiedHrisEmployeeOutput) GetEmploymentStatus() *string {
	if o == nil {
		return nil
	}
	return o.EmploymentStatus
}

func (o *UnifiedHrisEmployeeOutput) GetTerminationDate() *time.Time {
	if o == nil {
		return nil
	}
	return o.TerminationDate
}

func (o *UnifiedHrisEmployeeOutput) GetAvatarURL() *string {
	if o == nil {
		return nil
	}
	return o.AvatarURL
}

func (o *UnifiedHrisEmployeeOutput) GetManagerID() *string {
	if o == nil {
		return nil
	}
	return o.ManagerID
}

func (o *UnifiedHrisEmployeeOutput) GetFieldMappings() *UnifiedHrisEmployeeOutputFieldMappings {
	if o == nil {
		return nil
	}
	return o.FieldMappings
}

func (o *UnifiedHrisEmployeeOutput) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *UnifiedHrisEmployeeOutput) GetRemoteID() *string {
	if o == nil {
		return nil
	}
	return o.RemoteID
}

func (o *UnifiedHrisEmployeeOutput) GetRemoteData() *UnifiedHrisEmployeeOutputRemoteData {
	if o == nil {
		return nil
	}
	return o.RemoteData
}

func (o *UnifiedHrisEmployeeOutput) GetRemoteCreatedAt() *time.Time {
	if o == nil {
		return nil
	}
	return o.RemoteCreatedAt
}

func (o *UnifiedHrisEmployeeOutput) GetCreatedAt() *time.Time {
	if o == nil {
		return nil
	}
	return o.CreatedAt
}

func (o *UnifiedHrisEmployeeOutput) GetModifiedAt() *time.Time {
	if o == nil {
		return nil
	}
	return o.ModifiedAt
}

func (o *UnifiedHrisEmployeeOutput) GetRemoteWasDeleted() *bool {
	if o == nil {
		return nil
	}
	return o.RemoteWasDeleted
}
