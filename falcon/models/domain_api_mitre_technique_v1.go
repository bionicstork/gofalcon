// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// DomainAPIMITRETechniqueV1 domain API Mitre technique v1
//
// swagger:model domain.APIMitreTechniqueV1
type DomainAPIMITRETechniqueV1 struct {

	// id
	// Required: true
	ID *string `json:"id"`

	// link
	Link string `json:"link,omitempty"`

	// mitigations
	Mitigations []*DomainAPIMITREMitigationV1 `json:"mitigations"`
}

// Validate validates this domain API Mitre technique v1
func (m *DomainAPIMITRETechniqueV1) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMitigations(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DomainAPIMITRETechniqueV1) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	return nil
}

func (m *DomainAPIMITRETechniqueV1) validateMitigations(formats strfmt.Registry) error {
	if swag.IsZero(m.Mitigations) { // not required
		return nil
	}

	for i := 0; i < len(m.Mitigations); i++ {
		if swag.IsZero(m.Mitigations[i]) { // not required
			continue
		}

		if m.Mitigations[i] != nil {
			if err := m.Mitigations[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("mitigations" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("mitigations" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this domain API Mitre technique v1 based on the context it is used
func (m *DomainAPIMITRETechniqueV1) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateMitigations(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DomainAPIMITRETechniqueV1) contextValidateMitigations(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Mitigations); i++ {

		if m.Mitigations[i] != nil {

			if swag.IsZero(m.Mitigations[i]) { // not required
				return nil
			}

			if err := m.Mitigations[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("mitigations" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("mitigations" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *DomainAPIMITRETechniqueV1) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DomainAPIMITRETechniqueV1) UnmarshalBinary(b []byte) error {
	var res DomainAPIMITRETechniqueV1
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
