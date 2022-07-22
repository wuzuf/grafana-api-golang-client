// Code generated by go-swagger; DO NOT EDIT.

package admin_users

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
)

// NewAdminDeleteUserParams creates a new AdminDeleteUserParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewAdminDeleteUserParams() *AdminDeleteUserParams {
	return &AdminDeleteUserParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewAdminDeleteUserParamsWithTimeout creates a new AdminDeleteUserParams object
// with the ability to set a timeout on a request.
func NewAdminDeleteUserParamsWithTimeout(timeout time.Duration) *AdminDeleteUserParams {
	return &AdminDeleteUserParams{
		timeout: timeout,
	}
}

// NewAdminDeleteUserParamsWithContext creates a new AdminDeleteUserParams object
// with the ability to set a context for a request.
func NewAdminDeleteUserParamsWithContext(ctx context.Context) *AdminDeleteUserParams {
	return &AdminDeleteUserParams{
		Context: ctx,
	}
}

// NewAdminDeleteUserParamsWithHTTPClient creates a new AdminDeleteUserParams object
// with the ability to set a custom HTTPClient for a request.
func NewAdminDeleteUserParamsWithHTTPClient(client *http.Client) *AdminDeleteUserParams {
	return &AdminDeleteUserParams{
		HTTPClient: client,
	}
}

/* AdminDeleteUserParams contains all the parameters to send to the API endpoint
   for the admin delete user operation.

   Typically these are written to a http.Request.
*/
type AdminDeleteUserParams struct {

	// UserID.
	//
	// Format: int64
	UserID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the admin delete user params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AdminDeleteUserParams) WithDefaults() *AdminDeleteUserParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the admin delete user params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AdminDeleteUserParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the admin delete user params
func (o *AdminDeleteUserParams) WithTimeout(timeout time.Duration) *AdminDeleteUserParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the admin delete user params
func (o *AdminDeleteUserParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the admin delete user params
func (o *AdminDeleteUserParams) WithContext(ctx context.Context) *AdminDeleteUserParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the admin delete user params
func (o *AdminDeleteUserParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the admin delete user params
func (o *AdminDeleteUserParams) WithHTTPClient(client *http.Client) *AdminDeleteUserParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the admin delete user params
func (o *AdminDeleteUserParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithUserID adds the userID to the admin delete user params
func (o *AdminDeleteUserParams) WithUserID(userID int64) *AdminDeleteUserParams {
	o.SetUserID(userID)
	return o
}

// SetUserID adds the userId to the admin delete user params
func (o *AdminDeleteUserParams) SetUserID(userID int64) {
	o.UserID = userID
}

// WriteToRequest writes these params to a swagger request
func (o *AdminDeleteUserParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param user_id
	if err := r.SetPathParam("user_id", swag.FormatInt64(o.UserID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
