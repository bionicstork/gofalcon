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

// DomainAPIHostInfoFacetV1 domain API host info facet v1
//
// swagger:model domain.APIHostInfoFacetV1
type DomainAPIHostInfoFacetV1 struct {

	// build number
	BuildNumber string `json:"build_number,omitempty"`

	// groups
	// Required: true
	Groups []*DomainAPIHostGroup `json:"groups"`

	// host hidden status
	HostHiddenStatus string `json:"host_hidden_status,omitempty"`

	// id
	ID string `json:"id,omitempty"`

	// last seen timestamp
	LastSeenTimestamp string `json:"last_seen_timestamp,omitempty"`

	// local ip
	LocalIP string `json:"local_ip,omitempty"`

	// mac address
	MacAddress string `json:"mac_address,omitempty"`

	// machine domain
	MachineDomain string `json:"machine_domain,omitempty"`

	// name
	// Required: true
	Name *string `json:"name"`

	// os version
	OsVersion string `json:"os_version,omitempty"`

	// ou
	Ou string `json:"ou,omitempty"`

	// platform name
	PlatformName string `json:"platform_name,omitempty"`

	// product type desc
	ProductTypeDesc string `json:"product_type_desc,omitempty"`

	// serial number
	SerialNumber string `json:"serial_number,omitempty"`

	// site name
	SiteName string `json:"site_name,omitempty"`

	// system manufacturer
	SystemManufacturer string `json:"system_manufacturer,omitempty"`

	// system product name
	SystemProductName string `json:"system_product_name,omitempty"`

	// tags
	Tags []string `json:"tags"`
}

// Validate validates this domain API host info facet v1
func (m *DomainAPIHostInfoFacetV1) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateGroups(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DomainAPIHostInfoFacetV1) validateGroups(formats strfmt.Registry) error {

	if err := validate.Required("groups", "body", m.Groups); err != nil {
		return err
	}

	for i := 0; i < len(m.Groups); i++ {
		if swag.IsZero(m.Groups[i]) { // not required
			continue
		}

		if m.Groups[i] != nil {
			if err := m.Groups[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("groups" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("groups" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *DomainAPIHostInfoFacetV1) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this domain API host info facet v1 based on the context it is used
func (m *DomainAPIHostInfoFacetV1) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateGroups(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DomainAPIHostInfoFacetV1) contextValidateGroups(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Groups); i++ {

		if m.Groups[i] != nil {

			if swag.IsZero(m.Groups[i]) { // not required
				return nil
			}

			if err := m.Groups[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("groups" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("groups" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *DomainAPIHostInfoFacetV1) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DomainAPIHostInfoFacetV1) UnmarshalBinary(b []byte) error {
	var res DomainAPIHostInfoFacetV1
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}