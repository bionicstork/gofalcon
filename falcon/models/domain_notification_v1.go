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

// DomainNotificationV1 domain notification v1
//
// swagger:model domain.NotificationV1
type DomainNotificationV1 struct {

	// The email of the user who is assigned to this notification
	AssignedToUID string `json:"assigned_to_uid,omitempty"`

	// The name of the user who is assigned to this notification
	AssignedToUsername string `json:"assigned_to_username,omitempty"`

	// The unique ID of the user who is assigned to this notification
	AssignedToUUID string `json:"assigned_to_uuid,omitempty"`

	// Summary of the data breach which matched the rule
	BreachSummary *DomainMatchedBreachSummaryV1 `json:"breach_summary,omitempty"`

	// cid
	// Required: true
	CID *string `json:"cid"`

	// The date when the notification was generated
	// Required: true
	// Format: date-time
	CreatedDate *strfmt.DateTime `json:"created_date"`

	// Highlighted content based on the rule that generated the notifications. Highlights are surrounded with a `<cs-highlight>` tag
	Highlights []string `json:"highlights"`

	// The ID of the notification
	// Required: true
	ID *string `json:"id"`

	// The author who posted the intelligence item
	ItemAuthor string `json:"item_author,omitempty"`

	// The ID of the author who posted the intelligence item
	ItemAuthorID string `json:"item_author_id,omitempty"`

	// Timestamp when the item is considered to have been created
	// Required: true
	// Format: date-time
	ItemDate *strfmt.DateTime `json:"item_date"`

	// ID of the item which matched the rule
	// Required: true
	ItemID *string `json:"item_id"`

	// The site where the intelligence item was found
	ItemSite string `json:"item_site,omitempty"`

	// The ID of the site where the intelligence item was found
	ItemSiteID string `json:"item_site_id,omitempty"`

	// Type of the item which matched the rule: `post`, `reply`, `botnet_config`, `breach`, etc.
	// Required: true
	ItemType *string `json:"item_type"`

	// notes
	Notes []*SadomainNotificationNote `json:"notes"`

	// ID of the raw intel item that matched the rule
	// Required: true
	RawIntelID *string `json:"raw_intel_id"`

	// The ID of the rule that generated this notification
	// Required: true
	RuleID *string `json:"rule_id"`

	// The name of the rule that generated this notification
	// Required: true
	RuleName *string `json:"rule_name"`

	// The priority of the rule that generated this notification
	// Required: true
	RulePriority *string `json:"rule_priority"`

	// The topic of the rule that generated this notification
	// Required: true
	RuleTopic *string `json:"rule_topic"`

	// Category of the source that generated the notification
	SourceCategory string `json:"source_category,omitempty"`

	// The notification status. This can be one of: `new`, `in-progress`, `closed-false-positive`, `closed-true-positive`.
	// Required: true
	Status *string `json:"status"`

	// Details about the infrastructure component that matched the Typosquatting rule
	Typosquatting *SadomainTyposquattingComponent `json:"typosquatting,omitempty"`

	// The date when the notification was updated
	// Required: true
	// Format: date-time
	UpdatedDate *strfmt.DateTime `json:"updated_date"`
}

// Validate validates this domain notification v1
func (m *DomainNotificationV1) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateBreachSummary(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCreatedDate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateItemDate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateItemID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateItemType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNotes(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRawIntelID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRuleID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRuleName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRulePriority(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRuleTopic(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTyposquatting(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUpdatedDate(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DomainNotificationV1) validateBreachSummary(formats strfmt.Registry) error {
	if swag.IsZero(m.BreachSummary) { // not required
		return nil
	}

	if m.BreachSummary != nil {
		if err := m.BreachSummary.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("breach_summary")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("breach_summary")
			}
			return err
		}
	}

	return nil
}

func (m *DomainNotificationV1) validateCID(formats strfmt.Registry) error {

	if err := validate.Required("cid", "body", m.CID); err != nil {
		return err
	}

	return nil
}

func (m *DomainNotificationV1) validateCreatedDate(formats strfmt.Registry) error {

	if err := validate.Required("created_date", "body", m.CreatedDate); err != nil {
		return err
	}

	if err := validate.FormatOf("created_date", "body", "date-time", m.CreatedDate.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *DomainNotificationV1) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	return nil
}

func (m *DomainNotificationV1) validateItemDate(formats strfmt.Registry) error {

	if err := validate.Required("item_date", "body", m.ItemDate); err != nil {
		return err
	}

	if err := validate.FormatOf("item_date", "body", "date-time", m.ItemDate.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *DomainNotificationV1) validateItemID(formats strfmt.Registry) error {

	if err := validate.Required("item_id", "body", m.ItemID); err != nil {
		return err
	}

	return nil
}

func (m *DomainNotificationV1) validateItemType(formats strfmt.Registry) error {

	if err := validate.Required("item_type", "body", m.ItemType); err != nil {
		return err
	}

	return nil
}

func (m *DomainNotificationV1) validateNotes(formats strfmt.Registry) error {
	if swag.IsZero(m.Notes) { // not required
		return nil
	}

	for i := 0; i < len(m.Notes); i++ {
		if swag.IsZero(m.Notes[i]) { // not required
			continue
		}

		if m.Notes[i] != nil {
			if err := m.Notes[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("notes" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("notes" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *DomainNotificationV1) validateRawIntelID(formats strfmt.Registry) error {

	if err := validate.Required("raw_intel_id", "body", m.RawIntelID); err != nil {
		return err
	}

	return nil
}

func (m *DomainNotificationV1) validateRuleID(formats strfmt.Registry) error {

	if err := validate.Required("rule_id", "body", m.RuleID); err != nil {
		return err
	}

	return nil
}

func (m *DomainNotificationV1) validateRuleName(formats strfmt.Registry) error {

	if err := validate.Required("rule_name", "body", m.RuleName); err != nil {
		return err
	}

	return nil
}

func (m *DomainNotificationV1) validateRulePriority(formats strfmt.Registry) error {

	if err := validate.Required("rule_priority", "body", m.RulePriority); err != nil {
		return err
	}

	return nil
}

func (m *DomainNotificationV1) validateRuleTopic(formats strfmt.Registry) error {

	if err := validate.Required("rule_topic", "body", m.RuleTopic); err != nil {
		return err
	}

	return nil
}

func (m *DomainNotificationV1) validateStatus(formats strfmt.Registry) error {

	if err := validate.Required("status", "body", m.Status); err != nil {
		return err
	}

	return nil
}

func (m *DomainNotificationV1) validateTyposquatting(formats strfmt.Registry) error {
	if swag.IsZero(m.Typosquatting) { // not required
		return nil
	}

	if m.Typosquatting != nil {
		if err := m.Typosquatting.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("typosquatting")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("typosquatting")
			}
			return err
		}
	}

	return nil
}

func (m *DomainNotificationV1) validateUpdatedDate(formats strfmt.Registry) error {

	if err := validate.Required("updated_date", "body", m.UpdatedDate); err != nil {
		return err
	}

	if err := validate.FormatOf("updated_date", "body", "date-time", m.UpdatedDate.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this domain notification v1 based on the context it is used
func (m *DomainNotificationV1) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateBreachSummary(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateNotes(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateTyposquatting(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DomainNotificationV1) contextValidateBreachSummary(ctx context.Context, formats strfmt.Registry) error {

	if m.BreachSummary != nil {

		if swag.IsZero(m.BreachSummary) { // not required
			return nil
		}

		if err := m.BreachSummary.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("breach_summary")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("breach_summary")
			}
			return err
		}
	}

	return nil
}

func (m *DomainNotificationV1) contextValidateNotes(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Notes); i++ {

		if m.Notes[i] != nil {

			if swag.IsZero(m.Notes[i]) { // not required
				return nil
			}

			if err := m.Notes[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("notes" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("notes" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *DomainNotificationV1) contextValidateTyposquatting(ctx context.Context, formats strfmt.Registry) error {

	if m.Typosquatting != nil {

		if swag.IsZero(m.Typosquatting) { // not required
			return nil
		}

		if err := m.Typosquatting.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("typosquatting")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("typosquatting")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *DomainNotificationV1) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DomainNotificationV1) UnmarshalBinary(b []byte) error {
	var res DomainNotificationV1
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
