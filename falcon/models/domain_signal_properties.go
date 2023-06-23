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

// DomainSignalProperties domain signal properties
//
// swagger:model domain.SignalProperties
type DomainSignalProperties struct {

	// aid
	// Required: true
	Aid *string `json:"aid"`

	// assessment
	// Required: true
	Assessment *DomainAssessment `json:"assessment"`

	// assessment items
	// Required: true
	AssessmentItems *DomainAssessmentItems `json:"assessment_items"`

	// cid
	// Required: true
	CID *string `json:"cid"`

	// event platform
	// Required: true
	EventPlatform *string `json:"event_platform"`

	// modified time
	// Required: true
	// Format: date-time
	ModifiedTime *strfmt.DateTime `json:"modified_time"`

	// product type desc
	// Required: true
	ProductTypeDesc *string `json:"product_type_desc"`

	// sensor file status
	// Required: true
	SensorFileStatus *string `json:"sensor_file_status"`

	// system serial number
	// Required: true
	SystemSerialNumber *string `json:"system_serial_number"`
}

// Validate validates this domain signal properties
func (m *DomainSignalProperties) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAid(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAssessment(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAssessmentItems(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEventPlatform(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateModifiedTime(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateProductTypeDesc(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSensorFileStatus(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSystemSerialNumber(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DomainSignalProperties) validateAid(formats strfmt.Registry) error {

	if err := validate.Required("aid", "body", m.Aid); err != nil {
		return err
	}

	return nil
}

func (m *DomainSignalProperties) validateAssessment(formats strfmt.Registry) error {

	if err := validate.Required("assessment", "body", m.Assessment); err != nil {
		return err
	}

	if m.Assessment != nil {
		if err := m.Assessment.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("assessment")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("assessment")
			}
			return err
		}
	}

	return nil
}

func (m *DomainSignalProperties) validateAssessmentItems(formats strfmt.Registry) error {

	if err := validate.Required("assessment_items", "body", m.AssessmentItems); err != nil {
		return err
	}

	if m.AssessmentItems != nil {
		if err := m.AssessmentItems.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("assessment_items")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("assessment_items")
			}
			return err
		}
	}

	return nil
}

func (m *DomainSignalProperties) validateCID(formats strfmt.Registry) error {

	if err := validate.Required("cid", "body", m.CID); err != nil {
		return err
	}

	return nil
}

func (m *DomainSignalProperties) validateEventPlatform(formats strfmt.Registry) error {

	if err := validate.Required("event_platform", "body", m.EventPlatform); err != nil {
		return err
	}

	return nil
}

func (m *DomainSignalProperties) validateModifiedTime(formats strfmt.Registry) error {

	if err := validate.Required("modified_time", "body", m.ModifiedTime); err != nil {
		return err
	}

	if err := validate.FormatOf("modified_time", "body", "date-time", m.ModifiedTime.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *DomainSignalProperties) validateProductTypeDesc(formats strfmt.Registry) error {

	if err := validate.Required("product_type_desc", "body", m.ProductTypeDesc); err != nil {
		return err
	}

	return nil
}

func (m *DomainSignalProperties) validateSensorFileStatus(formats strfmt.Registry) error {

	if err := validate.Required("sensor_file_status", "body", m.SensorFileStatus); err != nil {
		return err
	}

	return nil
}

func (m *DomainSignalProperties) validateSystemSerialNumber(formats strfmt.Registry) error {

	if err := validate.Required("system_serial_number", "body", m.SystemSerialNumber); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this domain signal properties based on the context it is used
func (m *DomainSignalProperties) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAssessment(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateAssessmentItems(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DomainSignalProperties) contextValidateAssessment(ctx context.Context, formats strfmt.Registry) error {

	if m.Assessment != nil {

		if err := m.Assessment.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("assessment")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("assessment")
			}
			return err
		}
	}

	return nil
}

func (m *DomainSignalProperties) contextValidateAssessmentItems(ctx context.Context, formats strfmt.Registry) error {

	if m.AssessmentItems != nil {

		if err := m.AssessmentItems.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("assessment_items")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("assessment_items")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *DomainSignalProperties) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DomainSignalProperties) UnmarshalBinary(b []byte) error {
	var res DomainSignalProperties
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
