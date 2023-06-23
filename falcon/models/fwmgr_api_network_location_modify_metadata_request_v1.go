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

// FwmgrAPINetworkLocationModifyMetadataRequestV1 fwmgr api network location modify metadata request v1
//
// swagger:model fwmgr.api.NetworkLocationModifyMetadataRequestV1
type FwmgrAPINetworkLocationModifyMetadataRequestV1 struct {

	// cid
	// Required: true
	CID *string `json:"cid"`

	// dns resolution targets polling interval
	// Required: true
	DNSResolutionTargetsPollingInterval *int32 `json:"dns_resolution_targets_polling_interval"`

	// https reachable hosts polling interval
	// Required: true
	HTTPSReachableHostsPollingInterval *int32 `json:"https_reachable_hosts_polling_interval"`

	// icmp request targets polling interval
	// Required: true
	ICMPRequestTargetsPollingInterval *int32 `json:"icmp_request_targets_polling_interval"`

	// location precedence
	// Required: true
	LocationPrecedence []string `json:"location_precedence"`
}

// Validate validates this fwmgr api network location modify metadata request v1
func (m *FwmgrAPINetworkLocationModifyMetadataRequestV1) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDNSResolutionTargetsPollingInterval(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateHTTPSReachableHostsPollingInterval(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateICMPRequestTargetsPollingInterval(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLocationPrecedence(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *FwmgrAPINetworkLocationModifyMetadataRequestV1) validateCID(formats strfmt.Registry) error {

	if err := validate.Required("cid", "body", m.CID); err != nil {
		return err
	}

	return nil
}

func (m *FwmgrAPINetworkLocationModifyMetadataRequestV1) validateDNSResolutionTargetsPollingInterval(formats strfmt.Registry) error {

	if err := validate.Required("dns_resolution_targets_polling_interval", "body", m.DNSResolutionTargetsPollingInterval); err != nil {
		return err
	}

	return nil
}

func (m *FwmgrAPINetworkLocationModifyMetadataRequestV1) validateHTTPSReachableHostsPollingInterval(formats strfmt.Registry) error {

	if err := validate.Required("https_reachable_hosts_polling_interval", "body", m.HTTPSReachableHostsPollingInterval); err != nil {
		return err
	}

	return nil
}

func (m *FwmgrAPINetworkLocationModifyMetadataRequestV1) validateICMPRequestTargetsPollingInterval(formats strfmt.Registry) error {

	if err := validate.Required("icmp_request_targets_polling_interval", "body", m.ICMPRequestTargetsPollingInterval); err != nil {
		return err
	}

	return nil
}

func (m *FwmgrAPINetworkLocationModifyMetadataRequestV1) validateLocationPrecedence(formats strfmt.Registry) error {

	if err := validate.Required("location_precedence", "body", m.LocationPrecedence); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this fwmgr api network location modify metadata request v1 based on context it is used
func (m *FwmgrAPINetworkLocationModifyMetadataRequestV1) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *FwmgrAPINetworkLocationModifyMetadataRequestV1) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *FwmgrAPINetworkLocationModifyMetadataRequestV1) UnmarshalBinary(b []byte) error {
	var res FwmgrAPINetworkLocationModifyMetadataRequestV1
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
