// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
)

// DsPermissionType Datasource permission
// Description:
// `0` - No Access
// `1` - Query
// Enum: 0,1
//
// swagger:model DsPermissionType
type DsPermissionType int64

// Validate validates this ds permission type
func (m DsPermissionType) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this ds permission type based on context it is used
func (m DsPermissionType) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}
