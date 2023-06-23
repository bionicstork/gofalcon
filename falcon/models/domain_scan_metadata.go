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

// DomainScanMetadata domain scan metadata
//
// swagger:model domain.ScanMetadata
type DomainScanMetadata struct {

	// completed on
	// Format: date-time
	CompletedOn strfmt.DateTime `json:"completed_on,omitempty"`

	// filecount
	Filecount *DomainFileCount `json:"filecount,omitempty"`

	// host id
	// Required: true
	HostID *string `json:"host_id"`

	// host scan id
	HostScanID string `json:"host_scan_id,omitempty"`

	// last updated
	// Format: date-time
	LastUpdated strfmt.DateTime `json:"last_updated,omitempty"`

	// scan control reason
	ScanControlReason string `json:"scan_control_reason,omitempty"`

	// scan host metadata id
	ScanHostMetadataID string `json:"scan_host_metadata_id,omitempty"`

	// severity
	Severity int64 `json:"severity,omitempty"`

	// started on
	// Format: date-time
	StartedOn strfmt.DateTime `json:"started_on,omitempty"`

	// status
	Status string `json:"status,omitempty"`
}

// Validate validates this domain scan metadata
func (m *DomainScanMetadata) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCompletedOn(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFilecount(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateHostID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLastUpdated(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStartedOn(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DomainScanMetadata) validateCompletedOn(formats strfmt.Registry) error {
	if swag.IsZero(m.CompletedOn) { // not required
		return nil
	}

	if err := validate.FormatOf("completed_on", "body", "date-time", m.CompletedOn.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *DomainScanMetadata) validateFilecount(formats strfmt.Registry) error {
	if swag.IsZero(m.Filecount) { // not required
		return nil
	}

	if m.Filecount != nil {
		if err := m.Filecount.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("filecount")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("filecount")
			}
			return err
		}
	}

	return nil
}

func (m *DomainScanMetadata) validateHostID(formats strfmt.Registry) error {

	if err := validate.Required("host_id", "body", m.HostID); err != nil {
		return err
	}

	return nil
}

func (m *DomainScanMetadata) validateLastUpdated(formats strfmt.Registry) error {
	if swag.IsZero(m.LastUpdated) { // not required
		return nil
	}

	if err := validate.FormatOf("last_updated", "body", "date-time", m.LastUpdated.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *DomainScanMetadata) validateStartedOn(formats strfmt.Registry) error {
	if swag.IsZero(m.StartedOn) { // not required
		return nil
	}

	if err := validate.FormatOf("started_on", "body", "date-time", m.StartedOn.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this domain scan metadata based on the context it is used
func (m *DomainScanMetadata) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateFilecount(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DomainScanMetadata) contextValidateFilecount(ctx context.Context, formats strfmt.Registry) error {

	if m.Filecount != nil {

		if swag.IsZero(m.Filecount) { // not required
			return nil
		}

		if err := m.Filecount.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("filecount")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("filecount")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *DomainScanMetadata) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DomainScanMetadata) UnmarshalBinary(b []byte) error {
	var res DomainScanMetadata
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
