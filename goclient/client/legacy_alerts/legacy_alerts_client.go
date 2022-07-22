// Code generated by go-swagger; DO NOT EDIT.

package legacy_alerts

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new legacy alerts API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for legacy alerts API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	GetAlertByID(params *GetAlertByIDParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetAlertByIDOK, error)

	GetAlerts(params *GetAlertsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetAlertsOK, error)

	GetDashboardStates(params *GetDashboardStatesParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetDashboardStatesOK, error)

	PauseAlert(params *PauseAlertParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PauseAlertOK, error)

	TestAlert(params *TestAlertParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*TestAlertOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  GetAlertByID gets alert by ID

  “evalMatches” data in the response is cached in the db when and only when the state of the alert changes (e.g. transitioning from “ok” to “alerting” state).
If data from one server triggers the alert first and, before that server is seen leaving alerting state, a second server also enters a state that would trigger the alert, the second server will not be visible in “evalMatches” data.
*/
func (a *Client) GetAlertByID(params *GetAlertByIDParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetAlertByIDOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetAlertByIDParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getAlertByID",
		Method:             "GET",
		PathPattern:        "/alerts/{alert_id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetAlertByIDReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetAlertByIDOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getAlertByID: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetAlerts gets legacy alerts
*/
func (a *Client) GetAlerts(params *GetAlertsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetAlertsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetAlertsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getAlerts",
		Method:             "GET",
		PathPattern:        "/alerts",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetAlertsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetAlertsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getAlerts: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetDashboardStates gets alert states for a dashboard
*/
func (a *Client) GetDashboardStates(params *GetDashboardStatesParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetDashboardStatesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetDashboardStatesParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getDashboardStates",
		Method:             "GET",
		PathPattern:        "/alerts/states-for-dashboard",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetDashboardStatesReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetDashboardStatesOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getDashboardStates: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  PauseAlert pauses unpause alert by id
*/
func (a *Client) PauseAlert(params *PauseAlertParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PauseAlertOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPauseAlertParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "pauseAlert",
		Method:             "POST",
		PathPattern:        "/alerts/{alert_id}/pause",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &PauseAlertReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*PauseAlertOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for pauseAlert: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  TestAlert tests alert
*/
func (a *Client) TestAlert(params *TestAlertParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*TestAlertOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewTestAlertParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "testAlert",
		Method:             "POST",
		PathPattern:        "/alerts/test",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &TestAlertReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*TestAlertOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for testAlert: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
