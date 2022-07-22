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

// HTTPClientConfig HTTPClientConfig configures an HTTP client.
//
// swagger:model HTTPClientConfig
type HTTPClientConfig struct {

	// authorization
	Authorization *Authorization `json:"authorization,omitempty"`

	// basic auth
	BasicAuth *BasicAuth `json:"basic_auth,omitempty"`

	// bearer token
	BearerToken Secret `json:"bearer_token,omitempty"`

	// The bearer token file for the targets. Deprecated in favour of
	// Authorization.CredentialsFile.
	BearerTokenFile string `json:"bearer_token_file,omitempty"`

	// FollowRedirects specifies whether the client should follow HTTP 3xx redirects.
	// The omitempty flag is not set, because it would be hidden from the
	// marshalled configuration when set to false.
	FollowRedirects bool `json:"follow_redirects,omitempty"`

	// oauth2
	Oauth2 *OAuth2 `json:"oauth2,omitempty"`

	// proxy url
	ProxyURL *URL `json:"proxy_url,omitempty"`

	// tls config
	TLSConfig *TLSConfig `json:"tls_config,omitempty"`
}

// Validate validates this HTTP client config
func (m *HTTPClientConfig) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAuthorization(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateBasicAuth(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateBearerToken(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOauth2(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateProxyURL(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTLSConfig(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *HTTPClientConfig) validateAuthorization(formats strfmt.Registry) error {
	if swag.IsZero(m.Authorization) { // not required
		return nil
	}

	if m.Authorization != nil {
		if err := m.Authorization.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("authorization")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("authorization")
			}
			return err
		}
	}

	return nil
}

func (m *HTTPClientConfig) validateBasicAuth(formats strfmt.Registry) error {
	if swag.IsZero(m.BasicAuth) { // not required
		return nil
	}

	if m.BasicAuth != nil {
		if err := m.BasicAuth.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("basic_auth")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("basic_auth")
			}
			return err
		}
	}

	return nil
}

func (m *HTTPClientConfig) validateBearerToken(formats strfmt.Registry) error {
	if swag.IsZero(m.BearerToken) { // not required
		return nil
	}

	if err := m.BearerToken.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("bearer_token")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("bearer_token")
		}
		return err
	}

	return nil
}

func (m *HTTPClientConfig) validateOauth2(formats strfmt.Registry) error {
	if swag.IsZero(m.Oauth2) { // not required
		return nil
	}

	if m.Oauth2 != nil {
		if err := m.Oauth2.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("oauth2")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("oauth2")
			}
			return err
		}
	}

	return nil
}

func (m *HTTPClientConfig) validateProxyURL(formats strfmt.Registry) error {
	if swag.IsZero(m.ProxyURL) { // not required
		return nil
	}

	if m.ProxyURL != nil {
		if err := m.ProxyURL.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("proxy_url")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("proxy_url")
			}
			return err
		}
	}

	return nil
}

func (m *HTTPClientConfig) validateTLSConfig(formats strfmt.Registry) error {
	if swag.IsZero(m.TLSConfig) { // not required
		return nil
	}

	if m.TLSConfig != nil {
		if err := m.TLSConfig.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("tls_config")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("tls_config")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this HTTP client config based on the context it is used
func (m *HTTPClientConfig) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAuthorization(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateBasicAuth(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateBearerToken(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateOauth2(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateProxyURL(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateTLSConfig(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *HTTPClientConfig) contextValidateAuthorization(ctx context.Context, formats strfmt.Registry) error {

	if m.Authorization != nil {
		if err := m.Authorization.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("authorization")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("authorization")
			}
			return err
		}
	}

	return nil
}

func (m *HTTPClientConfig) contextValidateBasicAuth(ctx context.Context, formats strfmt.Registry) error {

	if m.BasicAuth != nil {
		if err := m.BasicAuth.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("basic_auth")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("basic_auth")
			}
			return err
		}
	}

	return nil
}

func (m *HTTPClientConfig) contextValidateBearerToken(ctx context.Context, formats strfmt.Registry) error {

	if err := m.BearerToken.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("bearer_token")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("bearer_token")
		}
		return err
	}

	return nil
}

func (m *HTTPClientConfig) contextValidateOauth2(ctx context.Context, formats strfmt.Registry) error {

	if m.Oauth2 != nil {
		if err := m.Oauth2.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("oauth2")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("oauth2")
			}
			return err
		}
	}

	return nil
}

func (m *HTTPClientConfig) contextValidateProxyURL(ctx context.Context, formats strfmt.Registry) error {

	if m.ProxyURL != nil {
		if err := m.ProxyURL.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("proxy_url")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("proxy_url")
			}
			return err
		}
	}

	return nil
}

func (m *HTTPClientConfig) contextValidateTLSConfig(ctx context.Context, formats strfmt.Registry) error {

	if m.TLSConfig != nil {
		if err := m.TLSConfig.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("tls_config")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("tls_config")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *HTTPClientConfig) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *HTTPClientConfig) UnmarshalBinary(b []byte) error {
	var res HTTPClientConfig
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
