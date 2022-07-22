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

// SlackConfig SlackConfig configures notifications via Slack.
//
// swagger:model SlackConfig
type SlackConfig struct {

	// actions
	Actions []*SlackAction `json:"actions"`

	// api url
	APIURL *SecretURL `json:"api_url,omitempty"`

	// api url file
	APIURLFile string `json:"api_url_file,omitempty"`

	// callback id
	CallbackID string `json:"callback_id,omitempty"`

	// Slack channel override, (like #other-channel or @username).
	Channel string `json:"channel,omitempty"`

	// color
	Color string `json:"color,omitempty"`

	// fallback
	Fallback string `json:"fallback,omitempty"`

	// fields
	Fields []*SlackField `json:"fields"`

	// footer
	Footer string `json:"footer,omitempty"`

	// http config
	HTTPConfig *HTTPClientConfig `json:"http_config,omitempty"`

	// icon emoji
	IconEmoji string `json:"icon_emoji,omitempty"`

	// icon url
	IconURL string `json:"icon_url,omitempty"`

	// image url
	ImageURL string `json:"image_url,omitempty"`

	// link names
	LinkNames bool `json:"link_names,omitempty"`

	// mrkdwn in
	MrkdwnIn []string `json:"mrkdwn_in"`

	// pretext
	Pretext string `json:"pretext,omitempty"`

	// send resolved
	SendResolved bool `json:"send_resolved,omitempty"`

	// short fields
	ShortFields bool `json:"short_fields,omitempty"`

	// text
	Text string `json:"text,omitempty"`

	// thumb url
	ThumbURL string `json:"thumb_url,omitempty"`

	// title
	Title string `json:"title,omitempty"`

	// title link
	TitleLink string `json:"title_link,omitempty"`

	// username
	Username string `json:"username,omitempty"`
}

// Validate validates this slack config
func (m *SlackConfig) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateActions(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAPIURL(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFields(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateHTTPConfig(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SlackConfig) validateActions(formats strfmt.Registry) error {
	if swag.IsZero(m.Actions) { // not required
		return nil
	}

	for i := 0; i < len(m.Actions); i++ {
		if swag.IsZero(m.Actions[i]) { // not required
			continue
		}

		if m.Actions[i] != nil {
			if err := m.Actions[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("actions" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("actions" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *SlackConfig) validateAPIURL(formats strfmt.Registry) error {
	if swag.IsZero(m.APIURL) { // not required
		return nil
	}

	if m.APIURL != nil {
		if err := m.APIURL.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("api_url")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("api_url")
			}
			return err
		}
	}

	return nil
}

func (m *SlackConfig) validateFields(formats strfmt.Registry) error {
	if swag.IsZero(m.Fields) { // not required
		return nil
	}

	for i := 0; i < len(m.Fields); i++ {
		if swag.IsZero(m.Fields[i]) { // not required
			continue
		}

		if m.Fields[i] != nil {
			if err := m.Fields[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("fields" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("fields" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *SlackConfig) validateHTTPConfig(formats strfmt.Registry) error {
	if swag.IsZero(m.HTTPConfig) { // not required
		return nil
	}

	if m.HTTPConfig != nil {
		if err := m.HTTPConfig.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("http_config")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("http_config")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this slack config based on the context it is used
func (m *SlackConfig) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateActions(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateAPIURL(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateFields(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateHTTPConfig(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SlackConfig) contextValidateActions(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Actions); i++ {

		if m.Actions[i] != nil {
			if err := m.Actions[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("actions" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("actions" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *SlackConfig) contextValidateAPIURL(ctx context.Context, formats strfmt.Registry) error {

	if m.APIURL != nil {
		if err := m.APIURL.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("api_url")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("api_url")
			}
			return err
		}
	}

	return nil
}

func (m *SlackConfig) contextValidateFields(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Fields); i++ {

		if m.Fields[i] != nil {
			if err := m.Fields[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("fields" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("fields" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *SlackConfig) contextValidateHTTPConfig(ctx context.Context, formats strfmt.Registry) error {

	if m.HTTPConfig != nil {
		if err := m.HTTPConfig.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("http_config")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("http_config")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *SlackConfig) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SlackConfig) UnmarshalBinary(b []byte) error {
	var res SlackConfig
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
