// Code generated by go-swagger; DO NOT EDIT.

package access_control_provisioning

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

// NewAdminProvisioningReloadAccessControlParams creates a new AdminProvisioningReloadAccessControlParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewAdminProvisioningReloadAccessControlParams() *AdminProvisioningReloadAccessControlParams {
	return &AdminProvisioningReloadAccessControlParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewAdminProvisioningReloadAccessControlParamsWithTimeout creates a new AdminProvisioningReloadAccessControlParams object
// with the ability to set a timeout on a request.
func NewAdminProvisioningReloadAccessControlParamsWithTimeout(timeout time.Duration) *AdminProvisioningReloadAccessControlParams {
	return &AdminProvisioningReloadAccessControlParams{
		timeout: timeout,
	}
}

// NewAdminProvisioningReloadAccessControlParamsWithContext creates a new AdminProvisioningReloadAccessControlParams object
// with the ability to set a context for a request.
func NewAdminProvisioningReloadAccessControlParamsWithContext(ctx context.Context) *AdminProvisioningReloadAccessControlParams {
	return &AdminProvisioningReloadAccessControlParams{
		Context: ctx,
	}
}

// NewAdminProvisioningReloadAccessControlParamsWithHTTPClient creates a new AdminProvisioningReloadAccessControlParams object
// with the ability to set a custom HTTPClient for a request.
func NewAdminProvisioningReloadAccessControlParamsWithHTTPClient(client *http.Client) *AdminProvisioningReloadAccessControlParams {
	return &AdminProvisioningReloadAccessControlParams{
		HTTPClient: client,
	}
}

/* AdminProvisioningReloadAccessControlParams contains all the parameters to send to the API endpoint
   for the admin provisioning reload access control operation.

   Typically these are written to a http.Request.
*/
type AdminProvisioningReloadAccessControlParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the admin provisioning reload access control params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AdminProvisioningReloadAccessControlParams) WithDefaults() *AdminProvisioningReloadAccessControlParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the admin provisioning reload access control params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AdminProvisioningReloadAccessControlParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the admin provisioning reload access control params
func (o *AdminProvisioningReloadAccessControlParams) WithTimeout(timeout time.Duration) *AdminProvisioningReloadAccessControlParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the admin provisioning reload access control params
func (o *AdminProvisioningReloadAccessControlParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the admin provisioning reload access control params
func (o *AdminProvisioningReloadAccessControlParams) WithContext(ctx context.Context) *AdminProvisioningReloadAccessControlParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the admin provisioning reload access control params
func (o *AdminProvisioningReloadAccessControlParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the admin provisioning reload access control params
func (o *AdminProvisioningReloadAccessControlParams) WithHTTPClient(client *http.Client) *AdminProvisioningReloadAccessControlParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the admin provisioning reload access control params
func (o *AdminProvisioningReloadAccessControlParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *AdminProvisioningReloadAccessControlParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
