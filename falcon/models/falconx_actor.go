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
)

// FalconxActor falconx actor
//
// swagger:model falconx.Actor
type FalconxActor struct {

	// created timestamp
	CreatedTimestamp string `json:"created_timestamp,omitempty"`

	// description
	Description string `json:"description,omitempty"`

	// first activity timestamp
	FirstActivityTimestamp string `json:"first_activity_timestamp,omitempty"`

	// id
	ID int32 `json:"id,omitempty"`

	// image artifact id
	ImageArtifactID string `json:"image_artifact_id,omitempty"`

	// known as
	KnownAs string `json:"known_as,omitempty"`

	// last activity timestamp
	LastActivityTimestamp string `json:"last_activity_timestamp,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// origins
	Origins []*FalconxEntity `json:"origins"`

	// short description
	ShortDescription string `json:"short_description,omitempty"`

	// slug
	Slug string `json:"slug,omitempty"`

	// target countries
	TargetCountries []*FalconxEntity `json:"target_countries"`

	// target industries
	TargetIndustries []*FalconxEntity `json:"target_industries"`

	// thumbnail artifact id
	ThumbnailArtifactID string `json:"thumbnail_artifact_id,omitempty"`
}

// Validate validates this falconx actor
func (m *FalconxActor) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateOrigins(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTargetCountries(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTargetIndustries(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *FalconxActor) validateOrigins(formats strfmt.Registry) error {
	if swag.IsZero(m.Origins) { // not required
		return nil
	}

	for i := 0; i < len(m.Origins); i++ {
		if swag.IsZero(m.Origins[i]) { // not required
			continue
		}

		if m.Origins[i] != nil {
			if err := m.Origins[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("origins" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("origins" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *FalconxActor) validateTargetCountries(formats strfmt.Registry) error {
	if swag.IsZero(m.TargetCountries) { // not required
		return nil
	}

	for i := 0; i < len(m.TargetCountries); i++ {
		if swag.IsZero(m.TargetCountries[i]) { // not required
			continue
		}

		if m.TargetCountries[i] != nil {
			if err := m.TargetCountries[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("target_countries" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("target_countries" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *FalconxActor) validateTargetIndustries(formats strfmt.Registry) error {
	if swag.IsZero(m.TargetIndustries) { // not required
		return nil
	}

	for i := 0; i < len(m.TargetIndustries); i++ {
		if swag.IsZero(m.TargetIndustries[i]) { // not required
			continue
		}

		if m.TargetIndustries[i] != nil {
			if err := m.TargetIndustries[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("target_industries" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("target_industries" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this falconx actor based on the context it is used
func (m *FalconxActor) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateOrigins(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateTargetCountries(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateTargetIndustries(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *FalconxActor) contextValidateOrigins(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Origins); i++ {

		if m.Origins[i] != nil {

			if swag.IsZero(m.Origins[i]) { // not required
				return nil
			}

			if err := m.Origins[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("origins" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("origins" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *FalconxActor) contextValidateTargetCountries(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.TargetCountries); i++ {

		if m.TargetCountries[i] != nil {

			if swag.IsZero(m.TargetCountries[i]) { // not required
				return nil
			}

			if err := m.TargetCountries[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("target_countries" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("target_countries" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *FalconxActor) contextValidateTargetIndustries(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.TargetIndustries); i++ {

		if m.TargetIndustries[i] != nil {

			if swag.IsZero(m.TargetIndustries[i]) { // not required
				return nil
			}

			if err := m.TargetIndustries[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("target_industries" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("target_industries" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *FalconxActor) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *FalconxActor) UnmarshalBinary(b []byte) error {
	var res FalconxActor
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
