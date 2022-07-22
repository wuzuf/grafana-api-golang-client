// Code generated by go-swagger; DO NOT EDIT.

package dashboards

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

// NewGetDashboardByUIDParams creates a new GetDashboardByUIDParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetDashboardByUIDParams() *GetDashboardByUIDParams {
	return &GetDashboardByUIDParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetDashboardByUIDParamsWithTimeout creates a new GetDashboardByUIDParams object
// with the ability to set a timeout on a request.
func NewGetDashboardByUIDParamsWithTimeout(timeout time.Duration) *GetDashboardByUIDParams {
	return &GetDashboardByUIDParams{
		timeout: timeout,
	}
}

// NewGetDashboardByUIDParamsWithContext creates a new GetDashboardByUIDParams object
// with the ability to set a context for a request.
func NewGetDashboardByUIDParamsWithContext(ctx context.Context) *GetDashboardByUIDParams {
	return &GetDashboardByUIDParams{
		Context: ctx,
	}
}

// NewGetDashboardByUIDParamsWithHTTPClient creates a new GetDashboardByUIDParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetDashboardByUIDParamsWithHTTPClient(client *http.Client) *GetDashboardByUIDParams {
	return &GetDashboardByUIDParams{
		HTTPClient: client,
	}
}

/* GetDashboardByUIDParams contains all the parameters to send to the API endpoint
   for the get dashboard by UID operation.

   Typically these are written to a http.Request.
*/
type GetDashboardByUIDParams struct {

	// UID.
	UID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get dashboard by UID params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetDashboardByUIDParams) WithDefaults() *GetDashboardByUIDParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get dashboard by UID params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetDashboardByUIDParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get dashboard by UID params
func (o *GetDashboardByUIDParams) WithTimeout(timeout time.Duration) *GetDashboardByUIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get dashboard by UID params
func (o *GetDashboardByUIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get dashboard by UID params
func (o *GetDashboardByUIDParams) WithContext(ctx context.Context) *GetDashboardByUIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get dashboard by UID params
func (o *GetDashboardByUIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get dashboard by UID params
func (o *GetDashboardByUIDParams) WithHTTPClient(client *http.Client) *GetDashboardByUIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get dashboard by UID params
func (o *GetDashboardByUIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithUID adds the uid to the get dashboard by UID params
func (o *GetDashboardByUIDParams) WithUID(uid string) *GetDashboardByUIDParams {
	o.SetUID(uid)
	return o
}

// SetUID adds the uid to the get dashboard by UID params
func (o *GetDashboardByUIDParams) SetUID(uid string) {
	o.UID = uid
}

// WriteToRequest writes these params to a swagger request
func (o *GetDashboardByUIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param uid
	if err := r.SetPathParam("uid", o.UID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
