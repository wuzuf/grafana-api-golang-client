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

// NewDeletePlaylistParams creates a new DeletePlaylistParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDeletePlaylistParams() *DeletePlaylistParams {
	return &DeletePlaylistParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDeletePlaylistParamsWithTimeout creates a new DeletePlaylistParams object
// with the ability to set a timeout on a request.
func NewDeletePlaylistParamsWithTimeout(timeout time.Duration) *DeletePlaylistParams {
	return &DeletePlaylistParams{
		timeout: timeout,
	}
}

// NewDeletePlaylistParamsWithContext creates a new DeletePlaylistParams object
// with the ability to set a context for a request.
func NewDeletePlaylistParamsWithContext(ctx context.Context) *DeletePlaylistParams {
	return &DeletePlaylistParams{
		Context: ctx,
	}
}

// NewDeletePlaylistParamsWithHTTPClient creates a new DeletePlaylistParams object
// with the ability to set a custom HTTPClient for a request.
func NewDeletePlaylistParamsWithHTTPClient(client *http.Client) *DeletePlaylistParams {
	return &DeletePlaylistParams{
		HTTPClient: client,
	}
}

/* DeletePlaylistParams contains all the parameters to send to the API endpoint
   for the delete playlist operation.

   Typically these are written to a http.Request.
*/
type DeletePlaylistParams struct {

	// UID.
	UID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the delete playlist params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeletePlaylistParams) WithDefaults() *DeletePlaylistParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the delete playlist params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeletePlaylistParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the delete playlist params
func (o *DeletePlaylistParams) WithTimeout(timeout time.Duration) *DeletePlaylistParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete playlist params
func (o *DeletePlaylistParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete playlist params
func (o *DeletePlaylistParams) WithContext(ctx context.Context) *DeletePlaylistParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete playlist params
func (o *DeletePlaylistParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete playlist params
func (o *DeletePlaylistParams) WithHTTPClient(client *http.Client) *DeletePlaylistParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete playlist params
func (o *DeletePlaylistParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithUID adds the uid to the delete playlist params
func (o *DeletePlaylistParams) WithUID(uid string) *DeletePlaylistParams {
	o.SetUID(uid)
	return o
}

// SetUID adds the uid to the delete playlist params
func (o *DeletePlaylistParams) SetUID(uid string) {
	o.UID = uid
}

// WriteToRequest writes these params to a swagger request
func (o *DeletePlaylistParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
