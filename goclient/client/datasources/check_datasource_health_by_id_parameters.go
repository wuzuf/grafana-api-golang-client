// Code generated by go-swagger; DO NOT EDIT.

package datasources

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

// NewCheckDatasourceHealthByIDParams creates a new CheckDatasourceHealthByIDParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCheckDatasourceHealthByIDParams() *CheckDatasourceHealthByIDParams {
	return &CheckDatasourceHealthByIDParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCheckDatasourceHealthByIDParamsWithTimeout creates a new CheckDatasourceHealthByIDParams object
// with the ability to set a timeout on a request.
func NewCheckDatasourceHealthByIDParamsWithTimeout(timeout time.Duration) *CheckDatasourceHealthByIDParams {
	return &CheckDatasourceHealthByIDParams{
		timeout: timeout,
	}
}

// NewCheckDatasourceHealthByIDParamsWithContext creates a new CheckDatasourceHealthByIDParams object
// with the ability to set a context for a request.
func NewCheckDatasourceHealthByIDParamsWithContext(ctx context.Context) *CheckDatasourceHealthByIDParams {
	return &CheckDatasourceHealthByIDParams{
		Context: ctx,
	}
}

// NewCheckDatasourceHealthByIDParamsWithHTTPClient creates a new CheckDatasourceHealthByIDParams object
// with the ability to set a custom HTTPClient for a request.
func NewCheckDatasourceHealthByIDParamsWithHTTPClient(client *http.Client) *CheckDatasourceHealthByIDParams {
	return &CheckDatasourceHealthByIDParams{
		HTTPClient: client,
	}
}

/* CheckDatasourceHealthByIDParams contains all the parameters to send to the API endpoint
   for the check datasource health by ID operation.

   Typically these are written to a http.Request.
*/
type CheckDatasourceHealthByIDParams struct {

	// ID.
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the check datasource health by ID params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CheckDatasourceHealthByIDParams) WithDefaults() *CheckDatasourceHealthByIDParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the check datasource health by ID params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CheckDatasourceHealthByIDParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the check datasource health by ID params
func (o *CheckDatasourceHealthByIDParams) WithTimeout(timeout time.Duration) *CheckDatasourceHealthByIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the check datasource health by ID params
func (o *CheckDatasourceHealthByIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the check datasource health by ID params
func (o *CheckDatasourceHealthByIDParams) WithContext(ctx context.Context) *CheckDatasourceHealthByIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the check datasource health by ID params
func (o *CheckDatasourceHealthByIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the check datasource health by ID params
func (o *CheckDatasourceHealthByIDParams) WithHTTPClient(client *http.Client) *CheckDatasourceHealthByIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the check datasource health by ID params
func (o *CheckDatasourceHealthByIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the check datasource health by ID params
func (o *CheckDatasourceHealthByIDParams) WithID(id string) *CheckDatasourceHealthByIDParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the check datasource health by ID params
func (o *CheckDatasourceHealthByIDParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *CheckDatasourceHealthByIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
