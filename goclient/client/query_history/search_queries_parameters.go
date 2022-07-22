// Code generated by go-swagger; DO NOT EDIT.

package query_history

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

// NewSearchQueriesParams creates a new SearchQueriesParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewSearchQueriesParams() *SearchQueriesParams {
	return &SearchQueriesParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewSearchQueriesParamsWithTimeout creates a new SearchQueriesParams object
// with the ability to set a timeout on a request.
func NewSearchQueriesParamsWithTimeout(timeout time.Duration) *SearchQueriesParams {
	return &SearchQueriesParams{
		timeout: timeout,
	}
}

// NewSearchQueriesParamsWithContext creates a new SearchQueriesParams object
// with the ability to set a context for a request.
func NewSearchQueriesParamsWithContext(ctx context.Context) *SearchQueriesParams {
	return &SearchQueriesParams{
		Context: ctx,
	}
}

// NewSearchQueriesParamsWithHTTPClient creates a new SearchQueriesParams object
// with the ability to set a custom HTTPClient for a request.
func NewSearchQueriesParamsWithHTTPClient(client *http.Client) *SearchQueriesParams {
	return &SearchQueriesParams{
		HTTPClient: client,
	}
}

/* SearchQueriesParams contains all the parameters to send to the API endpoint
   for the search queries operation.

   Typically these are written to a http.Request.
*/
type SearchQueriesParams struct {

	/* DatasourceUID.

	   List of data source UIDs to search for
	*/
	DatasourceUID []string

	/* From.

	   From range for the query history search

	   Format: int64
	*/
	From *int64

	/* Limit.

	   Limit the number of returned results

	   Format: int64
	*/
	Limit *int64

	/* OnlyStarred.

	   Flag indicating if only starred queries should be returned
	*/
	OnlyStarred *bool

	/* Page.

	   Use this parameter to access hits beyond limit. Numbering starts at 1. limit param acts as page size.

	   Format: int64
	*/
	Page *int64

	/* SearchString.

	   Text inside query or comments that is searched for
	*/
	SearchString *string

	/* Sort.

	   Sort method

	   Default: "time-desc"
	*/
	Sort *string

	/* To.

	   To range for the query history search

	   Format: int64
	*/
	To *int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the search queries params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SearchQueriesParams) WithDefaults() *SearchQueriesParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the search queries params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SearchQueriesParams) SetDefaults() {
	var (
		sortDefault = string("time-desc")
	)

	val := SearchQueriesParams{
		Sort: &sortDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the search queries params
func (o *SearchQueriesParams) WithTimeout(timeout time.Duration) *SearchQueriesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the search queries params
func (o *SearchQueriesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the search queries params
func (o *SearchQueriesParams) WithContext(ctx context.Context) *SearchQueriesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the search queries params
func (o *SearchQueriesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the search queries params
func (o *SearchQueriesParams) WithHTTPClient(client *http.Client) *SearchQueriesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the search queries params
func (o *SearchQueriesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDatasourceUID adds the datasourceUID to the search queries params
func (o *SearchQueriesParams) WithDatasourceUID(datasourceUID []string) *SearchQueriesParams {
	o.SetDatasourceUID(datasourceUID)
	return o
}

// SetDatasourceUID adds the datasourceUid to the search queries params
func (o *SearchQueriesParams) SetDatasourceUID(datasourceUID []string) {
	o.DatasourceUID = datasourceUID
}

// WithFrom adds the from to the search queries params
func (o *SearchQueriesParams) WithFrom(from *int64) *SearchQueriesParams {
	o.SetFrom(from)
	return o
}

// SetFrom adds the from to the search queries params
func (o *SearchQueriesParams) SetFrom(from *int64) {
	o.From = from
}

// WithLimit adds the limit to the search queries params
func (o *SearchQueriesParams) WithLimit(limit *int64) *SearchQueriesParams {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the search queries params
func (o *SearchQueriesParams) SetLimit(limit *int64) {
	o.Limit = limit
}

// WithOnlyStarred adds the onlyStarred to the search queries params
func (o *SearchQueriesParams) WithOnlyStarred(onlyStarred *bool) *SearchQueriesParams {
	o.SetOnlyStarred(onlyStarred)
	return o
}

// SetOnlyStarred adds the onlyStarred to the search queries params
func (o *SearchQueriesParams) SetOnlyStarred(onlyStarred *bool) {
	o.OnlyStarred = onlyStarred
}

// WithPage adds the page to the search queries params
func (o *SearchQueriesParams) WithPage(page *int64) *SearchQueriesParams {
	o.SetPage(page)
	return o
}

// SetPage adds the page to the search queries params
func (o *SearchQueriesParams) SetPage(page *int64) {
	o.Page = page
}

// WithSearchString adds the searchString to the search queries params
func (o *SearchQueriesParams) WithSearchString(searchString *string) *SearchQueriesParams {
	o.SetSearchString(searchString)
	return o
}

// SetSearchString adds the searchString to the search queries params
func (o *SearchQueriesParams) SetSearchString(searchString *string) {
	o.SearchString = searchString
}

// WithSort adds the sort to the search queries params
func (o *SearchQueriesParams) WithSort(sort *string) *SearchQueriesParams {
	o.SetSort(sort)
	return o
}

// SetSort adds the sort to the search queries params
func (o *SearchQueriesParams) SetSort(sort *string) {
	o.Sort = sort
}

// WithTo adds the to to the search queries params
func (o *SearchQueriesParams) WithTo(to *int64) *SearchQueriesParams {
	o.SetTo(to)
	return o
}

// SetTo adds the to to the search queries params
func (o *SearchQueriesParams) SetTo(to *int64) {
	o.To = to
}

// WriteToRequest writes these params to a swagger request
func (o *SearchQueriesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.DatasourceUID != nil {

		// binding items for datasourceUid
		joinedDatasourceUID := o.bindParamDatasourceUID(reg)

		// query array param datasourceUid
		if err := r.SetQueryParam("datasourceUid", joinedDatasourceUID...); err != nil {
			return err
		}
	}

	if o.From != nil {

		// query param from
		var qrFrom int64

		if o.From != nil {
			qrFrom = *o.From
		}
		qFrom := swag.FormatInt64(qrFrom)
		if qFrom != "" {

			if err := r.SetQueryParam("from", qFrom); err != nil {
				return err
			}
		}
	}

	if o.Limit != nil {

		// query param limit
		var qrLimit int64

		if o.Limit != nil {
			qrLimit = *o.Limit
		}
		qLimit := swag.FormatInt64(qrLimit)
		if qLimit != "" {

			if err := r.SetQueryParam("limit", qLimit); err != nil {
				return err
			}
		}
	}

	if o.OnlyStarred != nil {

		// query param onlyStarred
		var qrOnlyStarred bool

		if o.OnlyStarred != nil {
			qrOnlyStarred = *o.OnlyStarred
		}
		qOnlyStarred := swag.FormatBool(qrOnlyStarred)
		if qOnlyStarred != "" {

			if err := r.SetQueryParam("onlyStarred", qOnlyStarred); err != nil {
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

	if o.SearchString != nil {

		// query param searchString
		var qrSearchString string

		if o.SearchString != nil {
			qrSearchString = *o.SearchString
		}
		qSearchString := qrSearchString
		if qSearchString != "" {

			if err := r.SetQueryParam("searchString", qSearchString); err != nil {
				return err
			}
		}
	}

	if o.Sort != nil {

		// query param sort
		var qrSort string

		if o.Sort != nil {
			qrSort = *o.Sort
		}
		qSort := qrSort
		if qSort != "" {

			if err := r.SetQueryParam("sort", qSort); err != nil {
				return err
			}
		}
	}

	if o.To != nil {

		// query param to
		var qrTo int64

		if o.To != nil {
			qrTo = *o.To
		}
		qTo := swag.FormatInt64(qrTo)
		if qTo != "" {

			if err := r.SetQueryParam("to", qTo); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindParamSearchQueries binds the parameter datasourceUid
func (o *SearchQueriesParams) bindParamDatasourceUID(formats strfmt.Registry) []string {
	datasourceUIDIR := o.DatasourceUID

	var datasourceUIDIC []string
	for _, datasourceUIDIIR := range datasourceUIDIR { // explode []string

		datasourceUIDIIV := datasourceUIDIIR // string as string
		datasourceUIDIC = append(datasourceUIDIC, datasourceUIDIIV)
	}

	// items.CollectionFormat: "multi"
	datasourceUIDIS := swag.JoinByFormat(datasourceUIDIC, "multi")

	return datasourceUIDIS
}
