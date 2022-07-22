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

// RestoreDashboardVersionByUIDOKBody restore dashboard version by Uid o k body
//
// swagger:model restoreDashboardVersionByUidOKBody
type RestoreDashboardVersionByUIDOKBody struct {

	// ID The unique identifier (id) of the created/updated dashboard.
	// Example: 1
	// Required: true
	ID *string `json:"id"`

	// Status status of the response.
	// Example: success
	// Required: true
	Status *string `json:"status"`

	// Slug The slug of the dashboard.
	// Example: my-dashboard
	// Required: true
	Title *string `json:"title"`

	// UID The unique identifier (uid) of the created/updated dashboard.
	// Example: nHz3SXiiz
	// Required: true
	UID *string `json:"uid"`

	// URL The relative URL for accessing the created/updated dashboard.
	// Example: /d/nHz3SXiiz/my-dashboard
	// Required: true
	URL *string `json:"url"`

	// Version The version of the dashboard.
	// Example: 2
	// Required: true
	Version *int64 `json:"version"`
}

// Validate validates this restore dashboard version by Uid o k body
func (m *RestoreDashboardVersionByUIDOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTitle(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateURL(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVersion(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RestoreDashboardVersionByUIDOKBody) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	return nil
}

func (m *RestoreDashboardVersionByUIDOKBody) validateStatus(formats strfmt.Registry) error {

	if err := validate.Required("status", "body", m.Status); err != nil {
		return err
	}

	return nil
}

func (m *RestoreDashboardVersionByUIDOKBody) validateTitle(formats strfmt.Registry) error {

	if err := validate.Required("title", "body", m.Title); err != nil {
		return err
	}

	return nil
}

func (m *RestoreDashboardVersionByUIDOKBody) validateUID(formats strfmt.Registry) error {

	if err := validate.Required("uid", "body", m.UID); err != nil {
		return err
	}

	return nil
}

func (m *RestoreDashboardVersionByUIDOKBody) validateURL(formats strfmt.Registry) error {

	if err := validate.Required("url", "body", m.URL); err != nil {
		return err
	}

	return nil
}

func (m *RestoreDashboardVersionByUIDOKBody) validateVersion(formats strfmt.Registry) error {

	if err := validate.Required("version", "body", m.Version); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this restore dashboard version by Uid o k body based on context it is used
func (m *RestoreDashboardVersionByUIDOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *RestoreDashboardVersionByUIDOKBody) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RestoreDashboardVersionByUIDOKBody) UnmarshalBinary(b []byte) error {
	var res RestoreDashboardVersionByUIDOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
