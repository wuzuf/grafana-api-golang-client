// Code generated by go-swagger; DO NOT EDIT.

package datasource_permissions

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

// NewDisablePermissionsParams creates a new DisablePermissionsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDisablePermissionsParams() *DisablePermissionsParams {
	return &DisablePermissionsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDisablePermissionsParamsWithTimeout creates a new DisablePermissionsParams object
// with the ability to set a timeout on a request.
func NewDisablePermissionsParamsWithTimeout(timeout time.Duration) *DisablePermissionsParams {
	return &DisablePermissionsParams{
		timeout: timeout,
	}
}

// NewDisablePermissionsParamsWithContext creates a new DisablePermissionsParams object
// with the ability to set a context for a request.
func NewDisablePermissionsParamsWithContext(ctx context.Context) *DisablePermissionsParams {
	return &DisablePermissionsParams{
		Context: ctx,
	}
}

// NewDisablePermissionsParamsWithHTTPClient creates a new DisablePermissionsParams object
// with the ability to set a custom HTTPClient for a request.
func NewDisablePermissionsParamsWithHTTPClient(client *http.Client) *DisablePermissionsParams {
	return &DisablePermissionsParams{
		HTTPClient: client,
	}
}

/* DisablePermissionsParams contains all the parameters to send to the API endpoint
   for the disable permissions operation.

   Typically these are written to a http.Request.
*/
type DisablePermissionsParams struct {

	// DatasourceID.
	DatasourceID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the disable permissions params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DisablePermissionsParams) WithDefaults() *DisablePermissionsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the disable permissions params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DisablePermissionsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the disable permissions params
func (o *DisablePermissionsParams) WithTimeout(timeout time.Duration) *DisablePermissionsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the disable permissions params
func (o *DisablePermissionsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the disable permissions params
func (o *DisablePermissionsParams) WithContext(ctx context.Context) *DisablePermissionsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the disable permissions params
func (o *DisablePermissionsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the disable permissions params
func (o *DisablePermissionsParams) WithHTTPClient(client *http.Client) *DisablePermissionsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the disable permissions params
func (o *DisablePermissionsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDatasourceID adds the datasourceID to the disable permissions params
func (o *DisablePermissionsParams) WithDatasourceID(datasourceID string) *DisablePermissionsParams {
	o.SetDatasourceID(datasourceID)
	return o
}

// SetDatasourceID adds the datasourceId to the disable permissions params
func (o *DisablePermissionsParams) SetDatasourceID(datasourceID string) {
	o.DatasourceID = datasourceID
}

// WriteToRequest writes these params to a swagger request
func (o *DisablePermissionsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param datasourceId
	if err := r.SetPathParam("datasourceId", o.DatasourceID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
