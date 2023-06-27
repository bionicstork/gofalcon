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

// FalconxMetaInfo falconx meta info
//
// swagger:model falconx.MetaInfo
type FalconxMetaInfo struct {

	// meta info
	// Required: true
	MetaInfo *MsaMetaInfo `json:"MetaInfo"`

	// quota
	Quota *FalconxQuota `json:"quota,omitempty"`
}

// Validate validates this falconx meta info
func (m *FalconxMetaInfo) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateMetaInfo(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateQuota(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *FalconxMetaInfo) validateMetaInfo(formats strfmt.Registry) error {

	if err := validate.Required("MetaInfo", "body", m.MetaInfo); err != nil {
		return err
	}

	if m.MetaInfo != nil {
		if err := m.MetaInfo.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("MetaInfo")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("MetaInfo")
			}
			return err
		}
	}

	return nil
}

func (m *FalconxMetaInfo) validateQuota(formats strfmt.Registry) error {
	if swag.IsZero(m.Quota) { // not required
		return nil
	}

	if m.Quota != nil {
		if err := m.Quota.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("quota")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("quota")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this falconx meta info based on the context it is used
func (m *FalconxMetaInfo) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateMetaInfo(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateQuota(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *FalconxMetaInfo) contextValidateMetaInfo(ctx context.Context, formats strfmt.Registry) error {

	if m.MetaInfo != nil {

		if err := m.MetaInfo.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("MetaInfo")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("MetaInfo")
			}
			return err
		}
	}

	return nil
}

func (m *FalconxMetaInfo) contextValidateQuota(ctx context.Context, formats strfmt.Registry) error {

	if m.Quota != nil {

		if swag.IsZero(m.Quota) { // not required
			return nil
		}

		if err := m.Quota.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("quota")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("quota")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *FalconxMetaInfo) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *FalconxMetaInfo) UnmarshalBinary(b []byte) error {
	var res FalconxMetaInfo
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
