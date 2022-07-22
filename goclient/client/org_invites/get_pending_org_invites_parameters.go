// Code generated by go-swagger; DO NOT EDIT.

package org_invites

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

// NewGetPendingOrgInvitesParams creates a new GetPendingOrgInvitesParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetPendingOrgInvitesParams() *GetPendingOrgInvitesParams {
	return &GetPendingOrgInvitesParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetPendingOrgInvitesParamsWithTimeout creates a new GetPendingOrgInvitesParams object
// with the ability to set a timeout on a request.
func NewGetPendingOrgInvitesParamsWithTimeout(timeout time.Duration) *GetPendingOrgInvitesParams {
	return &GetPendingOrgInvitesParams{
		timeout: timeout,
	}
}

// NewGetPendingOrgInvitesParamsWithContext creates a new GetPendingOrgInvitesParams object
// with the ability to set a context for a request.
func NewGetPendingOrgInvitesParamsWithContext(ctx context.Context) *GetPendingOrgInvitesParams {
	return &GetPendingOrgInvitesParams{
		Context: ctx,
	}
}

// NewGetPendingOrgInvitesParamsWithHTTPClient creates a new GetPendingOrgInvitesParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetPendingOrgInvitesParamsWithHTTPClient(client *http.Client) *GetPendingOrgInvitesParams {
	return &GetPendingOrgInvitesParams{
		HTTPClient: client,
	}
}

/* GetPendingOrgInvitesParams contains all the parameters to send to the API endpoint
   for the get pending org invites operation.

   Typically these are written to a http.Request.
*/
type GetPendingOrgInvitesParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get pending org invites params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetPendingOrgInvitesParams) WithDefaults() *GetPendingOrgInvitesParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get pending org invites params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetPendingOrgInvitesParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get pending org invites params
func (o *GetPendingOrgInvitesParams) WithTimeout(timeout time.Duration) *GetPendingOrgInvitesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get pending org invites params
func (o *GetPendingOrgInvitesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get pending org invites params
func (o *GetPendingOrgInvitesParams) WithContext(ctx context.Context) *GetPendingOrgInvitesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get pending org invites params
func (o *GetPendingOrgInvitesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get pending org invites params
func (o *GetPendingOrgInvitesParams) WithHTTPClient(client *http.Client) *GetPendingOrgInvitesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get pending org invites params
func (o *GetPendingOrgInvitesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *GetPendingOrgInvitesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
