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

// DomainAPIMitreMitigationV1 domain API mitre mitigation v1
//
// swagger:model domain.APIMitreMitigationV1
type DomainAPIMitreMitigationV1 struct {

	// id
	// Required: true
	ID *string `json:"id"`

	// link
	Link string `json:"link,omitempty"`
}

// Validate validates this domain API mitre mitigation v1
func (m *DomainAPIMitreMitigationV1) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DomainAPIMitreMitigationV1) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this domain API mitre mitigation v1 based on context it is used
func (m *DomainAPIMitreMitigationV1) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *DomainAPIMitreMitigationV1) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DomainAPIMitreMitigationV1) UnmarshalBinary(b []byte) error {
	var res DomainAPIMitreMitigationV1
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
