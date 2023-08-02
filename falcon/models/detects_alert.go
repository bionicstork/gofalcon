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

// DetectsAlert detects alert
//
// swagger:model detects.Alert
type DetectsAlert struct {

	// agent id
	AgentID string `json:"agent_id,omitempty"`

	// aggregate id
	AggregateID string `json:"aggregate_id,omitempty"`

	// assigned to name
	AssignedToName string `json:"assigned_to_name,omitempty"`

	// assigned to uid
	AssignedToUID string `json:"assigned_to_uid,omitempty"`

	// assigned to uuid
	AssignedToUUID string `json:"assigned_to_uuid,omitempty"`

	// cid
	Cid string `json:"cid,omitempty"`

	// composite id
	CompositeID string `json:"composite_id,omitempty"`

	// confidence
	Confidence int64 `json:"confidence,omitempty"`

	// crawl edge ids
	CrawlEdgeIds map[string][]string `json:"crawl_edge_ids,omitempty"`

	// crawl traversal
	CrawlTraversal []*ThreatgraphCrawlEdgesRequest `json:"crawl_traversal"`

	// crawl vertex ids
	CrawlVertexIds map[string][]string `json:"crawl_vertex_ids,omitempty"`

	// crawled timestamp
	// Format: date-time
	CrawledTimestamp strfmt.DateTime `json:"crawled_timestamp,omitempty"`

	// created timestamp
	// Format: date-time
	CreatedTimestamp strfmt.DateTime `json:"created_timestamp,omitempty"`

	// description
	Description string `json:"description,omitempty"`

	// display name
	DisplayName string `json:"display_name,omitempty"`

	// email sent
	EmailSent bool `json:"email_sent,omitempty"`

	// external
	External bool `json:"external,omitempty"`

	// id
	// Required: true
	ID *string `json:"id"`

	// name
	Name string `json:"name,omitempty"`

	// objective
	Objective string `json:"objective,omitempty"`

	// pattern id
	PatternID int64 `json:"pattern_id,omitempty"`

	// platform
	Platform string `json:"platform,omitempty"`

	// product
	Product string `json:"product,omitempty"`

	// scenario
	Scenario string `json:"scenario,omitempty"`

	// severity
	Severity int64 `json:"severity,omitempty"`

	// show in ui
	ShowInUI bool `json:"show_in_ui,omitempty"`

	// status
	Status string `json:"status,omitempty"`

	// tactic
	Tactic string `json:"tactic,omitempty"`

	// tactic id
	TacticID string `json:"tactic_id,omitempty"`

	// tags
	Tags []string `json:"tags"`

	// technique
	Technique string `json:"technique,omitempty"`

	// technique id
	TechniqueID string `json:"technique_id,omitempty"`

	// timestamp
	// Format: date-time
	Timestamp strfmt.DateTime `json:"timestamp,omitempty"`

	// type
	// Required: true
	Type *string `json:"type"`

	// updated timestamp
	// Format: date-time
	UpdatedTimestamp strfmt.DateTime `json:"updated_timestamp,omitempty"`
}

// Validate validates this detects alert
func (m *DetectsAlert) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCrawlTraversal(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCrawledTimestamp(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCreatedTimestamp(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTimestamp(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUpdatedTimestamp(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DetectsAlert) validateCrawlTraversal(formats strfmt.Registry) error {
	if swag.IsZero(m.CrawlTraversal) { // not required
		return nil
	}

	for i := 0; i < len(m.CrawlTraversal); i++ {
		if swag.IsZero(m.CrawlTraversal[i]) { // not required
			continue
		}

		if m.CrawlTraversal[i] != nil {
			if err := m.CrawlTraversal[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("crawl_traversal" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("crawl_traversal" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *DetectsAlert) validateCrawledTimestamp(formats strfmt.Registry) error {
	if swag.IsZero(m.CrawledTimestamp) { // not required
		return nil
	}

	if err := validate.FormatOf("crawled_timestamp", "body", "date-time", m.CrawledTimestamp.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *DetectsAlert) validateCreatedTimestamp(formats strfmt.Registry) error {
	if swag.IsZero(m.CreatedTimestamp) { // not required
		return nil
	}

	if err := validate.FormatOf("created_timestamp", "body", "date-time", m.CreatedTimestamp.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *DetectsAlert) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	return nil
}

func (m *DetectsAlert) validateTimestamp(formats strfmt.Registry) error {
	if swag.IsZero(m.Timestamp) { // not required
		return nil
	}

	if err := validate.FormatOf("timestamp", "body", "date-time", m.Timestamp.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *DetectsAlert) validateType(formats strfmt.Registry) error {

	if err := validate.Required("type", "body", m.Type); err != nil {
		return err
	}

	return nil
}

func (m *DetectsAlert) validateUpdatedTimestamp(formats strfmt.Registry) error {
	if swag.IsZero(m.UpdatedTimestamp) { // not required
		return nil
	}

	if err := validate.FormatOf("updated_timestamp", "body", "date-time", m.UpdatedTimestamp.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this detects alert based on the context it is used
func (m *DetectsAlert) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateCrawlTraversal(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DetectsAlert) contextValidateCrawlTraversal(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.CrawlTraversal); i++ {

		if m.CrawlTraversal[i] != nil {

			if swag.IsZero(m.CrawlTraversal[i]) { // not required
				return nil
			}

			if err := m.CrawlTraversal[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("crawl_traversal" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("crawl_traversal" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *DetectsAlert) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DetectsAlert) UnmarshalBinary(b []byte) error {
	var res DetectsAlert
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}