// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// LibraryElementDTOMetaUser LibraryElementDTOMetaUser is the meta information for user that creates/changes the library element.
//
// swagger:model LibraryElementDTOMetaUser
type LibraryElementDTOMetaUser struct {

	// avatar Url
	AvatarURL string `json:"avatarUrl,omitempty"`

	// id
	ID int64 `json:"id,omitempty"`

	// name
	Name string `json:"name,omitempty"`
}

// Validate validates this library element d t o meta user
func (m *LibraryElementDTOMetaUser) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this library element d t o meta user based on context it is used
func (m *LibraryElementDTOMetaUser) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *LibraryElementDTOMetaUser) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *LibraryElementDTOMetaUser) UnmarshalBinary(b []byte) error {
	var res LibraryElementDTOMetaUser
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
