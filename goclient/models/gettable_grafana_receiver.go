// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// GettableGrafanaReceiver gettable grafana receiver
//
// swagger:model GettableGrafanaReceiver
type GettableGrafanaReceiver struct {

	// disable resolve message
	DisableResolveMessage bool `json:"disableResolveMessage,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// provenance
	Provenance Provenance `json:"provenance,omitempty"`

	// secure fields
	SecureFields map[string]bool `json:"secureFields,omitempty"`

	// settings
	Settings JSON `json:"settings,omitempty"`

	// type
	Type string `json:"type,omitempty"`

	// uid
	UID string `json:"uid,omitempty"`
}

// Validate validates this gettable grafana receiver
func (m *GettableGrafanaReceiver) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateProvenance(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GettableGrafanaReceiver) validateProvenance(formats strfmt.Registry) error {
	if swag.IsZero(m.Provenance) { // not required
		return nil
	}

	if err := m.Provenance.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("provenance")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("provenance")
		}
		return err
	}

	return nil
}

// ContextValidate validate this gettable grafana receiver based on the context it is used
func (m *GettableGrafanaReceiver) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateProvenance(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GettableGrafanaReceiver) contextValidateProvenance(ctx context.Context, formats strfmt.Registry) error {

	if err := m.Provenance.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("provenance")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("provenance")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *GettableGrafanaReceiver) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GettableGrafanaReceiver) UnmarshalBinary(b []byte) error {
	var res GettableGrafanaReceiver
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
