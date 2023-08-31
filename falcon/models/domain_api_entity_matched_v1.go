// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// DomainAPIEntityMatchedV1 domain API entity matched v1
//
// swagger:model domain.APIEntityMatchedV1
type DomainAPIEntityMatchedV1 struct {

	// asset id
	AssetID string `json:"asset_id,omitempty"`

	// data provider
	DataProvider string `json:"data_provider,omitempty"`

	// provider asset id
	ProviderAssetID string `json:"provider_asset_id,omitempty"`
}

// Validate validates this domain API entity matched v1
func (m *DomainAPIEntityMatchedV1) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this domain API entity matched v1 based on context it is used
func (m *DomainAPIEntityMatchedV1) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *DomainAPIEntityMatchedV1) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DomainAPIEntityMatchedV1) UnmarshalBinary(b []byte) error {
	var res DomainAPIEntityMatchedV1
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}