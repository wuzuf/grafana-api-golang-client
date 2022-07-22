// Code generated by go-swagger; DO NOT EDIT.

package signed_in_user

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"

	"github.com/grafana/grafana-api-golang-client/goclient/models"
)

// NewRevokeUserAuthTokenParams creates a new RevokeUserAuthTokenParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewRevokeUserAuthTokenParams() *RevokeUserAuthTokenParams {
	return &RevokeUserAuthTokenParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewRevokeUserAuthTokenParamsWithTimeout creates a new RevokeUserAuthTokenParams object
// with the ability to set a timeout on a request.
func NewRevokeUserAuthTokenParamsWithTimeout(timeout time.Duration) *RevokeUserAuthTokenParams {
	return &RevokeUserAuthTokenParams{
		timeout: timeout,
	}
}

// NewRevokeUserAuthTokenParamsWithContext creates a new RevokeUserAuthTokenParams object
// with the ability to set a context for a request.
func NewRevokeUserAuthTokenParamsWithContext(ctx context.Context) *RevokeUserAuthTokenParams {
	return &RevokeUserAuthTokenParams{
		Context: ctx,
	}
}

// NewRevokeUserAuthTokenParamsWithHTTPClient creates a new RevokeUserAuthTokenParams object
// with the ability to set a custom HTTPClient for a request.
func NewRevokeUserAuthTokenParamsWithHTTPClient(client *http.Client) *RevokeUserAuthTokenParams {
	return &RevokeUserAuthTokenParams{
		HTTPClient: client,
	}
}

/* RevokeUserAuthTokenParams contains all the parameters to send to the API endpoint
   for the revoke user auth token operation.

   Typically these are written to a http.Request.
*/
type RevokeUserAuthTokenParams struct {

	// Body.
	Body *models.RevokeAuthTokenCmd

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the revoke user auth token params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *RevokeUserAuthTokenParams) WithDefaults() *RevokeUserAuthTokenParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the revoke user auth token params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *RevokeUserAuthTokenParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the revoke user auth token params
func (o *RevokeUserAuthTokenParams) WithTimeout(timeout time.Duration) *RevokeUserAuthTokenParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the revoke user auth token params
func (o *RevokeUserAuthTokenParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the revoke user auth token params
func (o *RevokeUserAuthTokenParams) WithContext(ctx context.Context) *RevokeUserAuthTokenParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the revoke user auth token params
func (o *RevokeUserAuthTokenParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the revoke user auth token params
func (o *RevokeUserAuthTokenParams) WithHTTPClient(client *http.Client) *RevokeUserAuthTokenParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the revoke user auth token params
func (o *RevokeUserAuthTokenParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the revoke user auth token params
func (o *RevokeUserAuthTokenParams) WithBody(body *models.RevokeAuthTokenCmd) *RevokeUserAuthTokenParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the revoke user auth token params
func (o *RevokeUserAuthTokenParams) SetBody(body *models.RevokeAuthTokenCmd) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *RevokeUserAuthTokenParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
