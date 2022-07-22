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
)

// NewSearchOrgsParams creates a new SearchOrgsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewSearchOrgsParams() *SearchOrgsParams {
	return &SearchOrgsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewSearchOrgsParamsWithTimeout creates a new SearchOrgsParams object
// with the ability to set a timeout on a request.
func NewSearchOrgsParamsWithTimeout(timeout time.Duration) *SearchOrgsParams {
	return &SearchOrgsParams{
		timeout: timeout,
	}
}

// NewSearchOrgsParamsWithContext creates a new SearchOrgsParams object
// with the ability to set a context for a request.
func NewSearchOrgsParamsWithContext(ctx context.Context) *SearchOrgsParams {
	return &SearchOrgsParams{
		Context: ctx,
	}
}

// NewSearchOrgsParamsWithHTTPClient creates a new SearchOrgsParams object
// with the ability to set a custom HTTPClient for a request.
func NewSearchOrgsParamsWithHTTPClient(client *http.Client) *SearchOrgsParams {
	return &SearchOrgsParams{
		HTTPClient: client,
	}
}

/* SearchOrgsParams contains all the parameters to send to the API endpoint
   for the search orgs operation.

   Typically these are written to a http.Request.
*/
type SearchOrgsParams struct {

	// Name.
	Name *string

	// Page.
	//
	// Format: int64
	// Default: 1
	Page *int64

	/* Perpage.

	     Number of items per page
	The totalCount field in the response can be used for pagination list E.g. if totalCount is equal to 100 teams and the perpage parameter is set to 10 then there are 10 pages of teams.

	     Format: int64
	     Default: 1000
	*/
	Perpage *int64

	/* Query.

	   If set it will return results where the query value is contained in the name field. Query values with spaces need to be URL encoded.
	*/
	Query *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the search orgs params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SearchOrgsParams) WithDefaults() *SearchOrgsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the search orgs params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SearchOrgsParams) SetDefaults() {
	var (
		pageDefault = int64(1)

		perpageDefault = int64(1000)
	)

	val := SearchOrgsParams{
		Page:    &pageDefault,
		Perpage: &perpageDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the search orgs params
func (o *SearchOrgsParams) WithTimeout(timeout time.Duration) *SearchOrgsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the search orgs params
func (o *SearchOrgsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the search orgs params
func (o *SearchOrgsParams) WithContext(ctx context.Context) *SearchOrgsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the search orgs params
func (o *SearchOrgsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the search orgs params
func (o *SearchOrgsParams) WithHTTPClient(client *http.Client) *SearchOrgsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the search orgs params
func (o *SearchOrgsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithName adds the name to the search orgs params
func (o *SearchOrgsParams) WithName(name *string) *SearchOrgsParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the search orgs params
func (o *SearchOrgsParams) SetName(name *string) {
	o.Name = name
}

// WithPage adds the page to the search orgs params
func (o *SearchOrgsParams) WithPage(page *int64) *SearchOrgsParams {
	o.SetPage(page)
	return o
}

// SetPage adds the page to the search orgs params
func (o *SearchOrgsParams) SetPage(page *int64) {
	o.Page = page
}

// WithPerpage adds the perpage to the search orgs params
func (o *SearchOrgsParams) WithPerpage(perpage *int64) *SearchOrgsParams {
	o.SetPerpage(perpage)
	return o
}

// SetPerpage adds the perpage to the search orgs params
func (o *SearchOrgsParams) SetPerpage(perpage *int64) {
	o.Perpage = perpage
}

// WithQuery adds the query to the search orgs params
func (o *SearchOrgsParams) WithQuery(query *string) *SearchOrgsParams {
	o.SetQuery(query)
	return o
}

// SetQuery adds the query to the search orgs params
func (o *SearchOrgsParams) SetQuery(query *string) {
	o.Query = query
}

// WriteToRequest writes these params to a swagger request
func (o *SearchOrgsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Name != nil {

		// query param name
		var qrName string

		if o.Name != nil {
			qrName = *o.Name
		}
		qName := qrName
		if qName != "" {

			if err := r.SetQueryParam("name", qName); err != nil {
				return err
			}
		}
	}

	if o.Page != nil {

		// query param page
		var qrPage int64

		if o.Page != nil {
			qrPage = *o.Page
		}
		qPage := swag.FormatInt64(qrPage)
		if qPage != "" {

			if err := r.SetQueryParam("page", qPage); err != nil {
				return err
			}
		}
	}

	if o.Perpage != nil {

		// query param perpage
		var qrPerpage int64

		if o.Perpage != nil {
			qrPerpage = *o.Perpage
		}
		qPerpage := swag.FormatInt64(qrPerpage)
		if qPerpage != "" {

			if err := r.SetQueryParam("perpage", qPerpage); err != nil {
				return err
			}
		}
	}

	if o.Query != nil {

		// query param query
		var qrQuery string

		if o.Query != nil {
			qrQuery = *o.Query
		}
		qQuery := qrQuery
		if qQuery != "" {

			if err := r.SetQueryParam("query", qQuery); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
