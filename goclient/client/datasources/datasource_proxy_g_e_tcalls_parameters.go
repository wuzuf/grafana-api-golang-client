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

// NewDatasourceProxyGETcallsParams creates a new DatasourceProxyGETcallsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDatasourceProxyGETcallsParams() *DatasourceProxyGETcallsParams {
	return &DatasourceProxyGETcallsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDatasourceProxyGETcallsParamsWithTimeout creates a new DatasourceProxyGETcallsParams object
// with the ability to set a timeout on a request.
func NewDatasourceProxyGETcallsParamsWithTimeout(timeout time.Duration) *DatasourceProxyGETcallsParams {
	return &DatasourceProxyGETcallsParams{
		timeout: timeout,
	}
}

// NewDatasourceProxyGETcallsParamsWithContext creates a new DatasourceProxyGETcallsParams object
// with the ability to set a context for a request.
func NewDatasourceProxyGETcallsParamsWithContext(ctx context.Context) *DatasourceProxyGETcallsParams {
	return &DatasourceProxyGETcallsParams{
		Context: ctx,
	}
}

// NewDatasourceProxyGETcallsParamsWithHTTPClient creates a new DatasourceProxyGETcallsParams object
// with the ability to set a custom HTTPClient for a request.
func NewDatasourceProxyGETcallsParamsWithHTTPClient(client *http.Client) *DatasourceProxyGETcallsParams {
	return &DatasourceProxyGETcallsParams{
		HTTPClient: client,
	}
}

/* DatasourceProxyGETcallsParams contains all the parameters to send to the API endpoint
   for the datasource proxy g e tcalls operation.

   Typically these are written to a http.Request.
*/
type DatasourceProxyGETcallsParams struct {

	// DatasourceProxyRoute.
	DatasourceProxyRoute string

	// ID.
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the datasource proxy g e tcalls params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DatasourceProxyGETcallsParams) WithDefaults() *DatasourceProxyGETcallsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the datasource proxy g e tcalls params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DatasourceProxyGETcallsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the datasource proxy g e tcalls params
func (o *DatasourceProxyGETcallsParams) WithTimeout(timeout time.Duration) *DatasourceProxyGETcallsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the datasource proxy g e tcalls params
func (o *DatasourceProxyGETcallsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the datasource proxy g e tcalls params
func (o *DatasourceProxyGETcallsParams) WithContext(ctx context.Context) *DatasourceProxyGETcallsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the datasource proxy g e tcalls params
func (o *DatasourceProxyGETcallsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the datasource proxy g e tcalls params
func (o *DatasourceProxyGETcallsParams) WithHTTPClient(client *http.Client) *DatasourceProxyGETcallsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the datasource proxy g e tcalls params
func (o *DatasourceProxyGETcallsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDatasourceProxyRoute adds the datasourceProxyRoute to the datasource proxy g e tcalls params
func (o *DatasourceProxyGETcallsParams) WithDatasourceProxyRoute(datasourceProxyRoute string) *DatasourceProxyGETcallsParams {
	o.SetDatasourceProxyRoute(datasourceProxyRoute)
	return o
}

// SetDatasourceProxyRoute adds the datasourceProxyRoute to the datasource proxy g e tcalls params
func (o *DatasourceProxyGETcallsParams) SetDatasourceProxyRoute(datasourceProxyRoute string) {
	o.DatasourceProxyRoute = datasourceProxyRoute
}

// WithID adds the id to the datasource proxy g e tcalls params
func (o *DatasourceProxyGETcallsParams) WithID(id string) *DatasourceProxyGETcallsParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the datasource proxy g e tcalls params
func (o *DatasourceProxyGETcallsParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *DatasourceProxyGETcallsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param datasource_proxy_route
	if err := r.SetPathParam("datasource_proxy_route", o.DatasourceProxyRoute); err != nil {
		return err
	}

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
