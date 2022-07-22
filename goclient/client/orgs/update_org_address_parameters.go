// Code generated by go-swagger; DO NOT EDIT.

package orgs

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
	"github.com/go-openapi/swag"

	"github.com/grafana/grafana-api-golang-client/goclient/models"
)

// NewUpdateOrgAddressParams creates a new UpdateOrgAddressParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewUpdateOrgAddressParams() *UpdateOrgAddressParams {
	return &UpdateOrgAddressParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateOrgAddressParamsWithTimeout creates a new UpdateOrgAddressParams object
// with the ability to set a timeout on a request.
func NewUpdateOrgAddressParamsWithTimeout(timeout time.Duration) *UpdateOrgAddressParams {
	return &UpdateOrgAddressParams{
		timeout: timeout,
	}
}

// NewUpdateOrgAddressParamsWithContext creates a new UpdateOrgAddressParams object
// with the ability to set a context for a request.
func NewUpdateOrgAddressParamsWithContext(ctx context.Context) *UpdateOrgAddressParams {
	return &UpdateOrgAddressParams{
		Context: ctx,
	}
}

// NewUpdateOrgAddressParamsWithHTTPClient creates a new UpdateOrgAddressParams object
// with the ability to set a custom HTTPClient for a request.
func NewUpdateOrgAddressParamsWithHTTPClient(client *http.Client) *UpdateOrgAddressParams {
	return &UpdateOrgAddressParams{
		HTTPClient: client,
	}
}

/* UpdateOrgAddressParams contains all the parameters to send to the API endpoint
   for the update org address operation.

   Typically these are written to a http.Request.
*/
type UpdateOrgAddressParams struct {

	// Body.
	Body *models.UpdateOrgAddressForm

	// OrgID.
	//
	// Format: int64
	OrgID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the update org address params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateOrgAddressParams) WithDefaults() *UpdateOrgAddressParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the update org address params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateOrgAddressParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the update org address params
func (o *UpdateOrgAddressParams) WithTimeout(timeout time.Duration) *UpdateOrgAddressParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update org address params
func (o *UpdateOrgAddressParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update org address params
func (o *UpdateOrgAddressParams) WithContext(ctx context.Context) *UpdateOrgAddressParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update org address params
func (o *UpdateOrgAddressParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update org address params
func (o *UpdateOrgAddressParams) WithHTTPClient(client *http.Client) *UpdateOrgAddressParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update org address params
func (o *UpdateOrgAddressParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the update org address params
func (o *UpdateOrgAddressParams) WithBody(body *models.UpdateOrgAddressForm) *UpdateOrgAddressParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the update org address params
func (o *UpdateOrgAddressParams) SetBody(body *models.UpdateOrgAddressForm) {
	o.Body = body
}

// WithOrgID adds the orgID to the update org address params
func (o *UpdateOrgAddressParams) WithOrgID(orgID int64) *UpdateOrgAddressParams {
	o.SetOrgID(orgID)
	return o
}

// SetOrgID adds the orgId to the update org address params
func (o *UpdateOrgAddressParams) SetOrgID(orgID int64) {
	o.OrgID = orgID
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateOrgAddressParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param org_id
	if err := r.SetPathParam("org_id", swag.FormatInt64(o.OrgID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
