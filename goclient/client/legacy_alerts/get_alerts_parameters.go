// Code generated by go-swagger; DO NOT EDIT.

package legacy_alerts

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

// NewGetAlertsParams creates a new GetAlertsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetAlertsParams() *GetAlertsParams {
	return &GetAlertsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetAlertsParamsWithTimeout creates a new GetAlertsParams object
// with the ability to set a timeout on a request.
func NewGetAlertsParamsWithTimeout(timeout time.Duration) *GetAlertsParams {
	return &GetAlertsParams{
		timeout: timeout,
	}
}

// NewGetAlertsParamsWithContext creates a new GetAlertsParams object
// with the ability to set a context for a request.
func NewGetAlertsParamsWithContext(ctx context.Context) *GetAlertsParams {
	return &GetAlertsParams{
		Context: ctx,
	}
}

// NewGetAlertsParamsWithHTTPClient creates a new GetAlertsParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetAlertsParamsWithHTTPClient(client *http.Client) *GetAlertsParams {
	return &GetAlertsParams{
		HTTPClient: client,
	}
}

/* GetAlertsParams contains all the parameters to send to the API endpoint
   for the get alerts operation.

   Typically these are written to a http.Request.
*/
type GetAlertsParams struct {

	/* DashboardID.

	   Limit response to alerts in specified dashboard(s). You can specify multiple dashboards.
	*/
	DashboardID []string

	/* DashboardQuery.

	   Limit response to alerts having a dashboard name like this value./ Limit response to alerts having a dashboard name like this value.
	*/
	DashboardQuery *string

	/* DashboardTag.

	   Limit response to alerts of dashboards with specified tags. To do an “AND” filtering with multiple tags, specify the tags parameter multiple times
	*/
	DashboardTag []string

	/* FolderID.

	   Limit response to alerts of dashboards in specified folder(s). You can specify multiple folders
	*/
	FolderID []string

	/* Limit.

	   Limit response to X number of alerts.

	   Format: int64
	*/
	Limit *int64

	/* PanelID.

	   Limit response to alert for a specified panel on a dashboard.

	   Format: int64
	*/
	PanelID *int64

	/* Query.

	   Limit response to alerts having a name like this value.
	*/
	Query *string

	/* State.

	   Return alerts with one or more of the following alert states
	*/
	State *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get alerts params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetAlertsParams) WithDefaults() *GetAlertsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get alerts params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetAlertsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get alerts params
func (o *GetAlertsParams) WithTimeout(timeout time.Duration) *GetAlertsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get alerts params
func (o *GetAlertsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get alerts params
func (o *GetAlertsParams) WithContext(ctx context.Context) *GetAlertsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get alerts params
func (o *GetAlertsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get alerts params
func (o *GetAlertsParams) WithHTTPClient(client *http.Client) *GetAlertsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get alerts params
func (o *GetAlertsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDashboardID adds the dashboardID to the get alerts params
func (o *GetAlertsParams) WithDashboardID(dashboardID []string) *GetAlertsParams {
	o.SetDashboardID(dashboardID)
	return o
}

// SetDashboardID adds the dashboardId to the get alerts params
func (o *GetAlertsParams) SetDashboardID(dashboardID []string) {
	o.DashboardID = dashboardID
}

// WithDashboardQuery adds the dashboardQuery to the get alerts params
func (o *GetAlertsParams) WithDashboardQuery(dashboardQuery *string) *GetAlertsParams {
	o.SetDashboardQuery(dashboardQuery)
	return o
}

// SetDashboardQuery adds the dashboardQuery to the get alerts params
func (o *GetAlertsParams) SetDashboardQuery(dashboardQuery *string) {
	o.DashboardQuery = dashboardQuery
}

// WithDashboardTag adds the dashboardTag to the get alerts params
func (o *GetAlertsParams) WithDashboardTag(dashboardTag []string) *GetAlertsParams {
	o.SetDashboardTag(dashboardTag)
	return o
}

// SetDashboardTag adds the dashboardTag to the get alerts params
func (o *GetAlertsParams) SetDashboardTag(dashboardTag []string) {
	o.DashboardTag = dashboardTag
}

// WithFolderID adds the folderID to the get alerts params
func (o *GetAlertsParams) WithFolderID(folderID []string) *GetAlertsParams {
	o.SetFolderID(folderID)
	return o
}

// SetFolderID adds the folderId to the get alerts params
func (o *GetAlertsParams) SetFolderID(folderID []string) {
	o.FolderID = folderID
}

// WithLimit adds the limit to the get alerts params
func (o *GetAlertsParams) WithLimit(limit *int64) *GetAlertsParams {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the get alerts params
func (o *GetAlertsParams) SetLimit(limit *int64) {
	o.Limit = limit
}

// WithPanelID adds the panelID to the get alerts params
func (o *GetAlertsParams) WithPanelID(panelID *int64) *GetAlertsParams {
	o.SetPanelID(panelID)
	return o
}

// SetPanelID adds the panelId to the get alerts params
func (o *GetAlertsParams) SetPanelID(panelID *int64) {
	o.PanelID = panelID
}

// WithQuery adds the query to the get alerts params
func (o *GetAlertsParams) WithQuery(query *string) *GetAlertsParams {
	o.SetQuery(query)
	return o
}

// SetQuery adds the query to the get alerts params
func (o *GetAlertsParams) SetQuery(query *string) {
	o.Query = query
}

// WithState adds the state to the get alerts params
func (o *GetAlertsParams) WithState(state *string) *GetAlertsParams {
	o.SetState(state)
	return o
}

// SetState adds the state to the get alerts params
func (o *GetAlertsParams) SetState(state *string) {
	o.State = state
}

// WriteToRequest writes these params to a swagger request
func (o *GetAlertsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.DashboardID != nil {

		// binding items for dashboardId
		joinedDashboardID := o.bindParamDashboardID(reg)

		// query array param dashboardId
		if err := r.SetQueryParam("dashboardId", joinedDashboardID...); err != nil {
			return err
		}
	}

	if o.DashboardQuery != nil {

		// query param dashboardQuery
		var qrDashboardQuery string

		if o.DashboardQuery != nil {
			qrDashboardQuery = *o.DashboardQuery
		}
		qDashboardQuery := qrDashboardQuery
		if qDashboardQuery != "" {

			if err := r.SetQueryParam("dashboardQuery", qDashboardQuery); err != nil {
				return err
			}
		}
	}

	if o.DashboardTag != nil {

		// binding items for dashboardTag
		joinedDashboardTag := o.bindParamDashboardTag(reg)

		// query array param dashboardTag
		if err := r.SetQueryParam("dashboardTag", joinedDashboardTag...); err != nil {
			return err
		}
	}

	if o.FolderID != nil {

		// binding items for folderId
		joinedFolderID := o.bindParamFolderID(reg)

		// query array param folderId
		if err := r.SetQueryParam("folderId", joinedFolderID...); err != nil {
			return err
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

	if o.PanelID != nil {

		// query param panelId
		var qrPanelID int64

		if o.PanelID != nil {
			qrPanelID = *o.PanelID
		}
		qPanelID := swag.FormatInt64(qrPanelID)
		if qPanelID != "" {

			if err := r.SetQueryParam("panelId", qPanelID); err != nil {
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

	if o.State != nil {

		// query param state
		var qrState string

		if o.State != nil {
			qrState = *o.State
		}
		qState := qrState
		if qState != "" {

			if err := r.SetQueryParam("state", qState); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindParamGetAlerts binds the parameter dashboardId
func (o *GetAlertsParams) bindParamDashboardID(formats strfmt.Registry) []string {
	dashboardIDIR := o.DashboardID

	var dashboardIDIC []string
	for _, dashboardIDIIR := range dashboardIDIR { // explode []string

		dashboardIDIIV := dashboardIDIIR // string as string
		dashboardIDIC = append(dashboardIDIC, dashboardIDIIV)
	}

	// items.CollectionFormat: ""
	dashboardIDIS := swag.JoinByFormat(dashboardIDIC, "")

	return dashboardIDIS
}

// bindParamGetAlerts binds the parameter dashboardTag
func (o *GetAlertsParams) bindParamDashboardTag(formats strfmt.Registry) []string {
	dashboardTagIR := o.DashboardTag

	var dashboardTagIC []string
	for _, dashboardTagIIR := range dashboardTagIR { // explode []string

		dashboardTagIIV := dashboardTagIIR // string as string
		dashboardTagIC = append(dashboardTagIC, dashboardTagIIV)
	}

	// items.CollectionFormat: "multi"
	dashboardTagIS := swag.JoinByFormat(dashboardTagIC, "multi")

	return dashboardTagIS
}

// bindParamGetAlerts binds the parameter folderId
func (o *GetAlertsParams) bindParamFolderID(formats strfmt.Registry) []string {
	folderIDIR := o.FolderID

	var folderIDIC []string
	for _, folderIDIIR := range folderIDIR { // explode []string

		folderIDIIV := folderIDIIR // string as string
		folderIDIC = append(folderIDIC, folderIDIIV)
	}

	// items.CollectionFormat: "multi"
	folderIDIS := swag.JoinByFormat(folderIDIC, "multi")

	return folderIDIS
}
