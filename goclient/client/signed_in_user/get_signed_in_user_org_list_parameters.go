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
)

// NewGetSignedInUserOrgListParams creates a new GetSignedInUserOrgListParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetSignedInUserOrgListParams() *GetSignedInUserOrgListParams {
	return &GetSignedInUserOrgListParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetSignedInUserOrgListParamsWithTimeout creates a new GetSignedInUserOrgListParams object
// with the ability to set a timeout on a request.
func NewGetSignedInUserOrgListParamsWithTimeout(timeout time.Duration) *GetSignedInUserOrgListParams {
	return &GetSignedInUserOrgListParams{
		timeout: timeout,
	}
}

// NewGetSignedInUserOrgListParamsWithContext creates a new GetSignedInUserOrgListParams object
// with the ability to set a context for a request.
func NewGetSignedInUserOrgListParamsWithContext(ctx context.Context) *GetSignedInUserOrgListParams {
	return &GetSignedInUserOrgListParams{
		Context: ctx,
	}
}

// NewGetSignedInUserOrgListParamsWithHTTPClient creates a new GetSignedInUserOrgListParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetSignedInUserOrgListParamsWithHTTPClient(client *http.Client) *GetSignedInUserOrgListParams {
	return &GetSignedInUserOrgListParams{
		HTTPClient: client,
	}
}

/* GetSignedInUserOrgListParams contains all the parameters to send to the API endpoint
   for the get signed in user org list operation.

   Typically these are written to a http.Request.
*/
type GetSignedInUserOrgListParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get signed in user org list params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetSignedInUserOrgListParams) WithDefaults() *GetSignedInUserOrgListParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get signed in user org list params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetSignedInUserOrgListParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get signed in user org list params
func (o *GetSignedInUserOrgListParams) WithTimeout(timeout time.Duration) *GetSignedInUserOrgListParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get signed in user org list params
func (o *GetSignedInUserOrgListParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get signed in user org list params
func (o *GetSignedInUserOrgListParams) WithContext(ctx context.Context) *GetSignedInUserOrgListParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get signed in user org list params
func (o *GetSignedInUserOrgListParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get signed in user org list params
func (o *GetSignedInUserOrgListParams) WithHTTPClient(client *http.Client) *GetSignedInUserOrgListParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get signed in user org list params
func (o *GetSignedInUserOrgListParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *GetSignedInUserOrgListParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
