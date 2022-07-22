// Code generated by go-swagger; DO NOT EDIT.

package playlists

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

// NewGetPlaylistDashboardsParams creates a new GetPlaylistDashboardsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetPlaylistDashboardsParams() *GetPlaylistDashboardsParams {
	return &GetPlaylistDashboardsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetPlaylistDashboardsParamsWithTimeout creates a new GetPlaylistDashboardsParams object
// with the ability to set a timeout on a request.
func NewGetPlaylistDashboardsParamsWithTimeout(timeout time.Duration) *GetPlaylistDashboardsParams {
	return &GetPlaylistDashboardsParams{
		timeout: timeout,
	}
}

// NewGetPlaylistDashboardsParamsWithContext creates a new GetPlaylistDashboardsParams object
// with the ability to set a context for a request.
func NewGetPlaylistDashboardsParamsWithContext(ctx context.Context) *GetPlaylistDashboardsParams {
	return &GetPlaylistDashboardsParams{
		Context: ctx,
	}
}

// NewGetPlaylistDashboardsParamsWithHTTPClient creates a new GetPlaylistDashboardsParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetPlaylistDashboardsParamsWithHTTPClient(client *http.Client) *GetPlaylistDashboardsParams {
	return &GetPlaylistDashboardsParams{
		HTTPClient: client,
	}
}

/* GetPlaylistDashboardsParams contains all the parameters to send to the API endpoint
   for the get playlist dashboards operation.

   Typically these are written to a http.Request.
*/
type GetPlaylistDashboardsParams struct {

	// UID.
	UID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get playlist dashboards params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetPlaylistDashboardsParams) WithDefaults() *GetPlaylistDashboardsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get playlist dashboards params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetPlaylistDashboardsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get playlist dashboards params
func (o *GetPlaylistDashboardsParams) WithTimeout(timeout time.Duration) *GetPlaylistDashboardsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get playlist dashboards params
func (o *GetPlaylistDashboardsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get playlist dashboards params
func (o *GetPlaylistDashboardsParams) WithContext(ctx context.Context) *GetPlaylistDashboardsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get playlist dashboards params
func (o *GetPlaylistDashboardsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get playlist dashboards params
func (o *GetPlaylistDashboardsParams) WithHTTPClient(client *http.Client) *GetPlaylistDashboardsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get playlist dashboards params
func (o *GetPlaylistDashboardsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithUID adds the uid to the get playlist dashboards params
func (o *GetPlaylistDashboardsParams) WithUID(uid string) *GetPlaylistDashboardsParams {
	o.SetUID(uid)
	return o
}

// SetUID adds the uid to the get playlist dashboards params
func (o *GetPlaylistDashboardsParams) SetUID(uid string) {
	o.UID = uid
}

// WriteToRequest writes these params to a swagger request
func (o *GetPlaylistDashboardsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
