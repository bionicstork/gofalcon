// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// DomainLastUnscheduledExecution domain last unscheduled execution
//
// swagger:model domain.LastUnscheduledExecution
type DomainLastUnscheduledExecution struct {

	// activity status
	// Required: true
	ActivityStatus *string `json:"activity_status"`

	// id
	// Required: true
	ID *string `json:"id"`

	// last updated ts
	// Required: true
	// Format: date-time
	LastUpdatedTs *strfmt.DateTime `json:"last_updated_ts"`

	// status display
	// Required: true
	StatusDisplay *string `json:"status_display"`
}

// Validate validates this domain last unscheduled execution
func (m *DomainLastUnscheduledExecution) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateActivityStatus(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLastUpdatedTs(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStatusDisplay(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DomainLastUnscheduledExecution) validateActivityStatus(formats strfmt.Registry) error {

	if err := validate.Required("activity_status", "body", m.ActivityStatus); err != nil {
		return err
	}

	return nil
}

func (m *DomainLastUnscheduledExecution) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	return nil
}

func (m *DomainLastUnscheduledExecution) validateLastUpdatedTs(formats strfmt.Registry) error {

	if err := validate.Required("last_updated_ts", "body", m.LastUpdatedTs); err != nil {
		return err
	}

	if err := validate.FormatOf("last_updated_ts", "body", "date-time", m.LastUpdatedTs.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *DomainLastUnscheduledExecution) validateStatusDisplay(formats strfmt.Registry) error {

	if err := validate.Required("status_display", "body", m.StatusDisplay); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this domain last unscheduled execution based on context it is used
func (m *DomainLastUnscheduledExecution) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *DomainLastUnscheduledExecution) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DomainLastUnscheduledExecution) UnmarshalBinary(b []byte) error {
	var res DomainLastUnscheduledExecution
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}