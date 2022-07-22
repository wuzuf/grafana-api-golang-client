// Code generated by go-swagger; DO NOT EDIT.

package users

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

// NewSearchUsersWithPagingParams creates a new SearchUsersWithPagingParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewSearchUsersWithPagingParams() *SearchUsersWithPagingParams {
	return &SearchUsersWithPagingParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewSearchUsersWithPagingParamsWithTimeout creates a new SearchUsersWithPagingParams object
// with the ability to set a timeout on a request.
func NewSearchUsersWithPagingParamsWithTimeout(timeout time.Duration) *SearchUsersWithPagingParams {
	return &SearchUsersWithPagingParams{
		timeout: timeout,
	}
}

// NewSearchUsersWithPagingParamsWithContext creates a new SearchUsersWithPagingParams object
// with the ability to set a context for a request.
func NewSearchUsersWithPagingParamsWithContext(ctx context.Context) *SearchUsersWithPagingParams {
	return &SearchUsersWithPagingParams{
		Context: ctx,
	}
}

// NewSearchUsersWithPagingParamsWithHTTPClient creates a new SearchUsersWithPagingParams object
// with the ability to set a custom HTTPClient for a request.
func NewSearchUsersWithPagingParamsWithHTTPClient(client *http.Client) *SearchUsersWithPagingParams {
	return &SearchUsersWithPagingParams{
		HTTPClient: client,
	}
}

/* SearchUsersWithPagingParams contains all the parameters to send to the API endpoint
   for the search users with paging operation.

   Typically these are written to a http.Request.
*/
type SearchUsersWithPagingParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the search users with paging params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SearchUsersWithPagingParams) WithDefaults() *SearchUsersWithPagingParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the search users with paging params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SearchUsersWithPagingParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the search users with paging params
func (o *SearchUsersWithPagingParams) WithTimeout(timeout time.Duration) *SearchUsersWithPagingParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the search users with paging params
func (o *SearchUsersWithPagingParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the search users with paging params
func (o *SearchUsersWithPagingParams) WithContext(ctx context.Context) *SearchUsersWithPagingParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the search users with paging params
func (o *SearchUsersWithPagingParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the search users with paging params
func (o *SearchUsersWithPagingParams) WithHTTPClient(client *http.Client) *SearchUsersWithPagingParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the search users with paging params
func (o *SearchUsersWithPagingParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *SearchUsersWithPagingParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
