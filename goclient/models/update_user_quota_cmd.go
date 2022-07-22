// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// UpdateUserQuotaCmd update user quota cmd
//
// swagger:model UpdateUserQuotaCmd
type UpdateUserQuotaCmd struct {

	// limit
	Limit int64 `json:"limit,omitempty"`

	// target
	Target string `json:"target,omitempty"`
}

// Validate validates this update user quota cmd
func (m *UpdateUserQuotaCmd) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this update user quota cmd based on context it is used
func (m *UpdateUserQuotaCmd) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *UpdateUserQuotaCmd) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *UpdateUserQuotaCmd) UnmarshalBinary(b []byte) error {
	var res UpdateUserQuotaCmd
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
