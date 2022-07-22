// Code generated by go-swagger; DO NOT EDIT.

package recording_rules

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

// NewUpdateRecordingRuleParams creates a new UpdateRecordingRuleParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewUpdateRecordingRuleParams() *UpdateRecordingRuleParams {
	return &UpdateRecordingRuleParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateRecordingRuleParamsWithTimeout creates a new UpdateRecordingRuleParams object
// with the ability to set a timeout on a request.
func NewUpdateRecordingRuleParamsWithTimeout(timeout time.Duration) *UpdateRecordingRuleParams {
	return &UpdateRecordingRuleParams{
		timeout: timeout,
	}
}

// NewUpdateRecordingRuleParamsWithContext creates a new UpdateRecordingRuleParams object
// with the ability to set a context for a request.
func NewUpdateRecordingRuleParamsWithContext(ctx context.Context) *UpdateRecordingRuleParams {
	return &UpdateRecordingRuleParams{
		Context: ctx,
	}
}

// NewUpdateRecordingRuleParamsWithHTTPClient creates a new UpdateRecordingRuleParams object
// with the ability to set a custom HTTPClient for a request.
func NewUpdateRecordingRuleParamsWithHTTPClient(client *http.Client) *UpdateRecordingRuleParams {
	return &UpdateRecordingRuleParams{
		HTTPClient: client,
	}
}

/* UpdateRecordingRuleParams contains all the parameters to send to the API endpoint
   for the update recording rule operation.

   Typically these are written to a http.Request.
*/
type UpdateRecordingRuleParams struct {

	// Body.
	Body *models.RecordingRuleJSON

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the update recording rule params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateRecordingRuleParams) WithDefaults() *UpdateRecordingRuleParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the update recording rule params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateRecordingRuleParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the update recording rule params
func (o *UpdateRecordingRuleParams) WithTimeout(timeout time.Duration) *UpdateRecordingRuleParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update recording rule params
func (o *UpdateRecordingRuleParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update recording rule params
func (o *UpdateRecordingRuleParams) WithContext(ctx context.Context) *UpdateRecordingRuleParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update recording rule params
func (o *UpdateRecordingRuleParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update recording rule params
func (o *UpdateRecordingRuleParams) WithHTTPClient(client *http.Client) *UpdateRecordingRuleParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update recording rule params
func (o *UpdateRecordingRuleParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the update recording rule params
func (o *UpdateRecordingRuleParams) WithBody(body *models.RecordingRuleJSON) *UpdateRecordingRuleParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the update recording rule params
func (o *UpdateRecordingRuleParams) SetBody(body *models.RecordingRuleJSON) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateRecordingRuleParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
