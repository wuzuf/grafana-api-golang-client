// Code generated by go-swagger; DO NOT EDIT.

package admin

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

// NewPauseAllAlertsParams creates a new PauseAllAlertsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPauseAllAlertsParams() *PauseAllAlertsParams {
	return &PauseAllAlertsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPauseAllAlertsParamsWithTimeout creates a new PauseAllAlertsParams object
// with the ability to set a timeout on a request.
func NewPauseAllAlertsParamsWithTimeout(timeout time.Duration) *PauseAllAlertsParams {
	return &PauseAllAlertsParams{
		timeout: timeout,
	}
}

// NewPauseAllAlertsParamsWithContext creates a new PauseAllAlertsParams object
// with the ability to set a context for a request.
func NewPauseAllAlertsParamsWithContext(ctx context.Context) *PauseAllAlertsParams {
	return &PauseAllAlertsParams{
		Context: ctx,
	}
}

// NewPauseAllAlertsParamsWithHTTPClient creates a new PauseAllAlertsParams object
// with the ability to set a custom HTTPClient for a request.
func NewPauseAllAlertsParamsWithHTTPClient(client *http.Client) *PauseAllAlertsParams {
	return &PauseAllAlertsParams{
		HTTPClient: client,
	}
}

/* PauseAllAlertsParams contains all the parameters to send to the API endpoint
   for the pause all alerts operation.

   Typically these are written to a http.Request.
*/
type PauseAllAlertsParams struct {

	// Body.
	Body *models.PauseAllAlertsCommand

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the pause all alerts params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PauseAllAlertsParams) WithDefaults() *PauseAllAlertsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the pause all alerts params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PauseAllAlertsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the pause all alerts params
func (o *PauseAllAlertsParams) WithTimeout(timeout time.Duration) *PauseAllAlertsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the pause all alerts params
func (o *PauseAllAlertsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the pause all alerts params
func (o *PauseAllAlertsParams) WithContext(ctx context.Context) *PauseAllAlertsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the pause all alerts params
func (o *PauseAllAlertsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the pause all alerts params
func (o *PauseAllAlertsParams) WithHTTPClient(client *http.Client) *PauseAllAlertsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the pause all alerts params
func (o *PauseAllAlertsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the pause all alerts params
func (o *PauseAllAlertsParams) WithBody(body *models.PauseAllAlertsCommand) *PauseAllAlertsParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the pause all alerts params
func (o *PauseAllAlertsParams) SetBody(body *models.PauseAllAlertsCommand) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *PauseAllAlertsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
