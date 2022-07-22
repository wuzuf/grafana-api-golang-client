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

// AlertingRule adapted from cortex
//
// swagger:model AlertingRule
type AlertingRule struct {

	// alerts
	// Required: true
	Alerts []*Alert `json:"alerts"`

	// annotations
	// Required: true
	Annotations OverrideLabels `json:"annotations"`

	// duration
	Duration float64 `json:"duration,omitempty"`

	// evaluation time
	EvaluationTime float64 `json:"evaluationTime,omitempty"`

	// health
	// Required: true
	Health *string `json:"health"`

	// labels
	Labels OverrideLabels `json:"labels,omitempty"`

	// last error
	LastError string `json:"lastError,omitempty"`

	// last evaluation
	// Format: date-time
	LastEvaluation strfmt.DateTime `json:"lastEvaluation,omitempty"`

	// name
	// Required: true
	Name *string `json:"name"`

	// query
	// Required: true
	Query *string `json:"query"`

	// State can be "pending", "firing", "inactive".
	// Required: true
	State *string `json:"state"`

	// type
	// Required: true
	Type *RuleType `json:"type"`
}

// Validate validates this alerting rule
func (m *AlertingRule) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAlerts(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAnnotations(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateHealth(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLabels(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLastEvaluation(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateQuery(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateState(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AlertingRule) validateAlerts(formats strfmt.Registry) error {

	if err := validate.Required("alerts", "body", m.Alerts); err != nil {
		return err
	}

	for i := 0; i < len(m.Alerts); i++ {
		if swag.IsZero(m.Alerts[i]) { // not required
			continue
		}

		if m.Alerts[i] != nil {
			if err := m.Alerts[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("alerts" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("alerts" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *AlertingRule) validateAnnotations(formats strfmt.Registry) error {

	if err := validate.Required("annotations", "body", m.Annotations); err != nil {
		return err
	}

	if m.Annotations != nil {
		if err := m.Annotations.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("annotations")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("annotations")
			}
			return err
		}
	}

	return nil
}

func (m *AlertingRule) validateHealth(formats strfmt.Registry) error {

	if err := validate.Required("health", "body", m.Health); err != nil {
		return err
	}

	return nil
}

func (m *AlertingRule) validateLabels(formats strfmt.Registry) error {
	if swag.IsZero(m.Labels) { // not required
		return nil
	}

	if m.Labels != nil {
		if err := m.Labels.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("labels")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("labels")
			}
			return err
		}
	}

	return nil
}

func (m *AlertingRule) validateLastEvaluation(formats strfmt.Registry) error {
	if swag.IsZero(m.LastEvaluation) { // not required
		return nil
	}

	if err := validate.FormatOf("lastEvaluation", "body", "date-time", m.LastEvaluation.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *AlertingRule) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

func (m *AlertingRule) validateQuery(formats strfmt.Registry) error {

	if err := validate.Required("query", "body", m.Query); err != nil {
		return err
	}

	return nil
}

func (m *AlertingRule) validateState(formats strfmt.Registry) error {

	if err := validate.Required("state", "body", m.State); err != nil {
		return err
	}

	return nil
}

func (m *AlertingRule) validateType(formats strfmt.Registry) error {

	if err := validate.Required("type", "body", m.Type); err != nil {
		return err
	}

	if err := validate.Required("type", "body", m.Type); err != nil {
		return err
	}

	if m.Type != nil {
		if err := m.Type.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("type")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("type")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this alerting rule based on the context it is used
func (m *AlertingRule) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAlerts(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateAnnotations(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateLabels(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateType(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AlertingRule) contextValidateAlerts(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Alerts); i++ {

		if m.Alerts[i] != nil {
			if err := m.Alerts[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("alerts" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("alerts" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *AlertingRule) contextValidateAnnotations(ctx context.Context, formats strfmt.Registry) error {

	if err := m.Annotations.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("annotations")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("annotations")
		}
		return err
	}

	return nil
}

func (m *AlertingRule) contextValidateLabels(ctx context.Context, formats strfmt.Registry) error {

	if err := m.Labels.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("labels")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("labels")
		}
		return err
	}

	return nil
}

func (m *AlertingRule) contextValidateType(ctx context.Context, formats strfmt.Registry) error {

	if m.Type != nil {
		if err := m.Type.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("type")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("type")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *AlertingRule) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AlertingRule) UnmarshalBinary(b []byte) error {
	var res AlertingRule
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
