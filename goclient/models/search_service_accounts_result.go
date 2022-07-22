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

// SearchServiceAccountsResult swagger: model
//
// swagger:model SearchServiceAccountsResult
type SearchServiceAccountsResult struct {

	// page
	Page int64 `json:"page,omitempty"`

	// per page
	PerPage int64 `json:"perPage,omitempty"`

	// service accounts
	ServiceAccounts []*ServiceAccountDTO `json:"serviceAccounts"`

	// It can be used for pagination of the user list
	// E.g. if totalCount is equal to 100 users and
	// the perpage parameter is set to 10 then there are 10 pages of users.
	TotalCount int64 `json:"totalCount,omitempty"`
}

// Validate validates this search service accounts result
func (m *SearchServiceAccountsResult) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateServiceAccounts(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SearchServiceAccountsResult) validateServiceAccounts(formats strfmt.Registry) error {
	if swag.IsZero(m.ServiceAccounts) { // not required
		return nil
	}

	for i := 0; i < len(m.ServiceAccounts); i++ {
		if swag.IsZero(m.ServiceAccounts[i]) { // not required
			continue
		}

		if m.ServiceAccounts[i] != nil {
			if err := m.ServiceAccounts[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("serviceAccounts" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("serviceAccounts" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this search service accounts result based on the context it is used
func (m *SearchServiceAccountsResult) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateServiceAccounts(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SearchServiceAccountsResult) contextValidateServiceAccounts(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.ServiceAccounts); i++ {

		if m.ServiceAccounts[i] != nil {
			if err := m.ServiceAccounts[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("serviceAccounts" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("serviceAccounts" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *SearchServiceAccountsResult) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SearchServiceAccountsResult) UnmarshalBinary(b []byte) error {
	var res SearchServiceAccountsResult
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
