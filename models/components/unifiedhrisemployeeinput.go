// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

import (
	"github.com/panoratech/go-sdk/internal/utils"
	"time"
)

// UnifiedHrisEmployeeInputFieldMappings - The custom field mappings of the object between the remote 3rd party & Panora
type UnifiedHrisEmployeeInputFieldMappings struct {
}

type UnifiedHrisEmployeeInput struct {
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
	FieldMappings *UnifiedHrisEmployeeInputFieldMappings `json:"field_mappings,omitempty"`
}

func (u UnifiedHrisEmployeeInput) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(u, "", false)
}

func (u *UnifiedHrisEmployeeInput) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &u, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *UnifiedHrisEmployeeInput) GetGroups() []string {
	if o == nil {
		return nil
	}
	return o.Groups
}

func (o *UnifiedHrisEmployeeInput) GetLocations() []string {
	if o == nil {
		return nil
	}
	return o.Locations
}

func (o *UnifiedHrisEmployeeInput) GetEmployeeNumber() *string {
	if o == nil {
		return nil
	}
	return o.EmployeeNumber
}

func (o *UnifiedHrisEmployeeInput) GetCompanyID() *string {
	if o == nil {
		return nil
	}
	return o.CompanyID
}

func (o *UnifiedHrisEmployeeInput) GetFirstName() *string {
	if o == nil {
		return nil
	}
	return o.FirstName
}

func (o *UnifiedHrisEmployeeInput) GetLastName() *string {
	if o == nil {
		return nil
	}
	return o.LastName
}

func (o *UnifiedHrisEmployeeInput) GetPreferredName() *string {
	if o == nil {
		return nil
	}
	return o.PreferredName
}

func (o *UnifiedHrisEmployeeInput) GetDisplayFullName() *string {
	if o == nil {
		return nil
	}
	return o.DisplayFullName
}

func (o *UnifiedHrisEmployeeInput) GetUsername() *string {
	if o == nil {
		return nil
	}
	return o.Username
}

func (o *UnifiedHrisEmployeeInput) GetWorkEmail() *string {
	if o == nil {
		return nil
	}
	return o.WorkEmail
}

func (o *UnifiedHrisEmployeeInput) GetPersonalEmail() *string {
	if o == nil {
		return nil
	}
	return o.PersonalEmail
}

func (o *UnifiedHrisEmployeeInput) GetMobilePhoneNumber() *string {
	if o == nil {
		return nil
	}
	return o.MobilePhoneNumber
}

func (o *UnifiedHrisEmployeeInput) GetEmployments() []string {
	if o == nil {
		return nil
	}
	return o.Employments
}

func (o *UnifiedHrisEmployeeInput) GetSsn() *string {
	if o == nil {
		return nil
	}
	return o.Ssn
}

func (o *UnifiedHrisEmployeeInput) GetGender() *string {
	if o == nil {
		return nil
	}
	return o.Gender
}

func (o *UnifiedHrisEmployeeInput) GetEthnicity() *string {
	if o == nil {
		return nil
	}
	return o.Ethnicity
}

func (o *UnifiedHrisEmployeeInput) GetMaritalStatus() *string {
	if o == nil {
		return nil
	}
	return o.MaritalStatus
}

func (o *UnifiedHrisEmployeeInput) GetDateOfBirth() *time.Time {
	if o == nil {
		return nil
	}
	return o.DateOfBirth
}

func (o *UnifiedHrisEmployeeInput) GetStartDate() *time.Time {
	if o == nil {
		return nil
	}
	return o.StartDate
}

func (o *UnifiedHrisEmployeeInput) GetEmploymentStatus() *string {
	if o == nil {
		return nil
	}
	return o.EmploymentStatus
}

func (o *UnifiedHrisEmployeeInput) GetTerminationDate() *time.Time {
	if o == nil {
		return nil
	}
	return o.TerminationDate
}

func (o *UnifiedHrisEmployeeInput) GetAvatarURL() *string {
	if o == nil {
		return nil
	}
	return o.AvatarURL
}

func (o *UnifiedHrisEmployeeInput) GetManagerID() *string {
	if o == nil {
		return nil
	}
	return o.ManagerID
}

func (o *UnifiedHrisEmployeeInput) GetFieldMappings() *UnifiedHrisEmployeeInputFieldMappings {
	if o == nil {
		return nil
	}
	return o.FieldMappings
}
