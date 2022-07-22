// Code generated by go-swagger; DO NOT EDIT.

package access_control

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new access control API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for access control API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	AddBuiltinRole(params *AddBuiltinRoleParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*AddBuiltinRoleOK, error)

	AddTeamRole(params *AddTeamRoleParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*AddTeamRoleOK, error)

	AddUserRole(params *AddUserRoleParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*AddUserRoleOK, error)

	CreateRole(params *CreateRoleParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*CreateRoleCreated, error)

	DeleteCustomRole(params *DeleteCustomRoleParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DeleteCustomRoleOK, error)

	GetAccessControlStatus(params *GetAccessControlStatusParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetAccessControlStatusOK, error)

	GetRole(params *GetRoleParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetRoleOK, error)

	ListBuiltinRoles(params *ListBuiltinRolesParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ListBuiltinRolesOK, error)

	ListRoles(params *ListRolesParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ListRolesOK, error)

	ListTeamRoles(params *ListTeamRolesParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ListTeamRolesOK, error)

	ListUserRoles(params *ListUserRolesParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ListUserRolesOK, error)

	RemoveBuiltinRole(params *RemoveBuiltinRoleParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*RemoveBuiltinRoleOK, error)

	RemoveTeamRole(params *RemoveTeamRoleParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*RemoveTeamRoleOK, error)

	RemoveUserRole(params *RemoveUserRoleParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*RemoveUserRoleOK, error)

	SetTeamRoles(params *SetTeamRolesParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*SetTeamRolesOK, error)

	SetUserRoles(params *SetUserRolesParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*SetUserRolesOK, error)

	UpdateRoleWithPermissions(params *UpdateRoleWithPermissionsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*UpdateRoleWithPermissionsOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  AddBuiltinRole creates a built in role assignment

  You need to have a permission with action `roles.builtin:add` and scope `permissions:type:delegate`. `permissions:type:delegate` scope ensures that users can only create built-in role assignments with the roles which have same, or a subset of permissions which the user has. For example, if a user does not have required permissions for creating users, they won’t be able to create a built-in role assignment which will allow to do that. This is done to prevent escalation of privileges.
*/
func (a *Client) AddBuiltinRole(params *AddBuiltinRoleParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*AddBuiltinRoleOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewAddBuiltinRoleParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "addBuiltinRole",
		Method:             "POST",
		PathPattern:        "/access-control/builtin-roles",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &AddBuiltinRoleReader{formats: a.formats},
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
	success, ok := result.(*AddBuiltinRoleOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for addBuiltinRole: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  AddTeamRole adds team role

  You need to have a permission with action `teams.roles:add` and scope `permissions:type:delegate`.
*/
func (a *Client) AddTeamRole(params *AddTeamRoleParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*AddTeamRoleOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewAddTeamRoleParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "addTeamRole",
		Method:             "POST",
		PathPattern:        "/access-control/teams/{teamId}/roles",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &AddTeamRoleReader{formats: a.formats},
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
	success, ok := result.(*AddTeamRoleOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for addTeamRole: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  AddUserRole adds a user role assignment

  Assign a role to a specific user. For bulk updates consider Set user role assignments.

You need to have a permission with action `users.roles:add` and scope `permissions:type:delegate`. `permissions:type:delegate` scope ensures that users can only assign roles which have same, or a subset of permissions which the user has. For example, if a user does not have required permissions for creating users, they won’t be able to assign a role which will allow to do that. This is done to prevent escalation of privileges.
*/
func (a *Client) AddUserRole(params *AddUserRoleParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*AddUserRoleOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewAddUserRoleParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "addUserRole",
		Method:             "POST",
		PathPattern:        "/access-control/users/{userId}/roles",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &AddUserRoleReader{formats: a.formats},
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
	success, ok := result.(*AddUserRoleOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for addUserRole: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  CreateRole creates a new custom role

  Creates a new custom role and maps given permissions to that role. Note that roles with the same prefix as Fixed Roles can’t be created.

You need to have a permission with action `roles:write` and scope `permissions:type:delegate`. `permissions:type:delegate`` scope ensures that users can only create custom roles with the same, or a subset of permissions which the user has.
For example, if a user does not have required permissions for creating users, they won’t be able to create a custom role which allows to do that. This is done to prevent escalation of privileges.
*/
func (a *Client) CreateRole(params *CreateRoleParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*CreateRoleCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateRoleParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "createRole",
		Method:             "POST",
		PathPattern:        "/access-control/roles",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &CreateRoleReader{formats: a.formats},
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
	success, ok := result.(*CreateRoleCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for createRole: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  DeleteCustomRole deletes a custom role

  Delete a role with the given UID, and it’s permissions. If the role is assigned to a built-in role, the deletion operation will fail, unless force query param is set to true, and in that case all assignments will also be deleted.

You need to have a permission with action `roles:delete` and scope `permissions:type:delegate`. `permissions:type:delegate` scope ensures that users can only delete a custom role with the same, or a subset of permissions which the user has. For example, if a user does not have required permissions for creating users, they won’t be able to delete a custom role which allows to do that.
*/
func (a *Client) DeleteCustomRole(params *DeleteCustomRoleParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DeleteCustomRoleOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteCustomRoleParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "deleteCustomRole",
		Method:             "DELETE",
		PathPattern:        "/access-control/roles/{roleUID}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &DeleteCustomRoleReader{formats: a.formats},
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
	success, ok := result.(*DeleteCustomRoleOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for deleteCustomRole: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetAccessControlStatus gets status

  Returns an indicator to check if fine-grained access control is enabled or not.

You need to have a permission with action `status:accesscontrol` and scope `services:accesscontrol`.
*/
func (a *Client) GetAccessControlStatus(params *GetAccessControlStatusParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetAccessControlStatusOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetAccessControlStatusParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getAccessControlStatus",
		Method:             "GET",
		PathPattern:        "/access-control/status",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetAccessControlStatusReader{formats: a.formats},
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
	success, ok := result.(*GetAccessControlStatusOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getAccessControlStatus: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetRole gets a role

  Get a role for the given UID.

You need to have a permission with action `roles:read` and scope `roles:*`.
*/
func (a *Client) GetRole(params *GetRoleParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetRoleOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetRoleParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getRole",
		Method:             "GET",
		PathPattern:        "/access-control/roles/{roleUID}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetRoleReader{formats: a.formats},
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
	success, ok := result.(*GetRoleOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getRole: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  ListBuiltinRoles gets all built in role assignments

  You need to have a permission with action `roles.builtin:list` with scope `roles:*`.
*/
func (a *Client) ListBuiltinRoles(params *ListBuiltinRolesParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ListBuiltinRolesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListBuiltinRolesParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "listBuiltinRoles",
		Method:             "GET",
		PathPattern:        "/access-control/builtin-roles",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &ListBuiltinRolesReader{formats: a.formats},
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
	success, ok := result.(*ListBuiltinRolesOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for listBuiltinRoles: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  ListRoles gets all roles

  Gets all existing roles. The response contains all global and organization local roles, for the organization which user is signed in.

You need to have a permission with action `roles:list` and scope `roles:*`.
*/
func (a *Client) ListRoles(params *ListRolesParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ListRolesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListRolesParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "listRoles",
		Method:             "GET",
		PathPattern:        "/access-control/roles",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &ListRolesReader{formats: a.formats},
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
	success, ok := result.(*ListRolesOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for listRoles: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  ListTeamRoles gets team roles

  You need to have a permission with action `teams.roles:list` and scope `teams:id:<team ID>`.
*/
func (a *Client) ListTeamRoles(params *ListTeamRolesParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ListTeamRolesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListTeamRolesParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "listTeamRoles",
		Method:             "GET",
		PathPattern:        "/access-control/teams/{teamId}/roles",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &ListTeamRolesReader{formats: a.formats},
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
	success, ok := result.(*ListTeamRolesOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for listTeamRoles: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  ListUserRoles lists roles assigned to a user

  Lists the roles that have been directly assigned to a given user. The list does not include built-in roles (Viewer, Editor, Admin or Grafana Admin), and it does not include roles that have been inherited from a team.

You need to have a permission with action `users.roles:list` and scope `users:id:<user ID>`.
*/
func (a *Client) ListUserRoles(params *ListUserRolesParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ListUserRolesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListUserRolesParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "listUserRoles",
		Method:             "GET",
		PathPattern:        "/access-control/users/{userId}/roles",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &ListUserRolesReader{formats: a.formats},
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
	success, ok := result.(*ListUserRolesOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for listUserRoles: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  RemoveBuiltinRole removes a built in role assignment

  Deletes a built-in role assignment (for one of Viewer, Editor, Admin, or Grafana Admin) to the role with the provided UID.

You need to have a permission with action `roles.builtin:remove` and scope `permissions:type:delegate`. `permissions:type:delegate` scope ensures that users can only remove built-in role assignments with the roles which have same, or a subset of permissions which the user has. For example, if a user does not have required permissions for creating users, they won’t be able to remove a built-in role assignment which allows to do that.
*/
func (a *Client) RemoveBuiltinRole(params *RemoveBuiltinRoleParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*RemoveBuiltinRoleOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewRemoveBuiltinRoleParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "removeBuiltinRole",
		Method:             "DELETE",
		PathPattern:        "/access-control/builtin-roles/{builtinRole}/roles/{roleUID}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &RemoveBuiltinRoleReader{formats: a.formats},
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
	success, ok := result.(*RemoveBuiltinRoleOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for removeBuiltinRole: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  RemoveTeamRole removes team role

  You need to have a permission with action `teams.roles:remove` and scope `permissions:type:delegate`.
*/
func (a *Client) RemoveTeamRole(params *RemoveTeamRoleParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*RemoveTeamRoleOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewRemoveTeamRoleParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "removeTeamRole",
		Method:             "DELETE",
		PathPattern:        "/access-control/teams/{teamId}/roles/{roleUID}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &RemoveTeamRoleReader{formats: a.formats},
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
	success, ok := result.(*RemoveTeamRoleOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for removeTeamRole: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  RemoveUserRole removes a user role assignment

  Revoke a role from a user. For bulk updates consider Set user role assignments.

You need to have a permission with action `users.roles:remove` and scope `permissions:type:delegate`. `permissions:type:delegate` scope ensures that users can only unassign roles which have same, or a subset of permissions which the user has. For example, if a user does not have required permissions for creating users, they won’t be able to unassign a role which will allow to do that. This is done to prevent escalation of privileges.
*/
func (a *Client) RemoveUserRole(params *RemoveUserRoleParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*RemoveUserRoleOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewRemoveUserRoleParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "removeUserRole",
		Method:             "DELETE",
		PathPattern:        "/access-control/users/{userId}/roles/{roleUID}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &RemoveUserRoleReader{formats: a.formats},
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
	success, ok := result.(*RemoveUserRoleOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for removeUserRole: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  SetTeamRoles updates team role

  You need to have a permission with action `teams.roles:add` and `teams.roles:remove` and scope `permissions:type:delegate` for each.
*/
func (a *Client) SetTeamRoles(params *SetTeamRolesParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*SetTeamRolesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewSetTeamRolesParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "setTeamRoles",
		Method:             "PUT",
		PathPattern:        "/access-control/teams/{teamId}/roles",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &SetTeamRolesReader{formats: a.formats},
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
	success, ok := result.(*SetTeamRolesOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for setTeamRoles: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  SetUserRoles sets user role assignments

  Update the user’s role assignments to match the provided set of UIDs. This will remove any assigned roles that aren’t in the request and add roles that are in the set but are not already assigned to the user.
If you want to add or remove a single role, consider using Add a user role assignment or Remove a user role assignment instead.

You need to have a permission with action `users.roles:add` and `users.roles:remove` and scope `permissions:type:delegate` for each. `permissions:type:delegate`  scope ensures that users can only assign or unassign roles which have same, or a subset of permissions which the user has. For example, if a user does not have required permissions for creating users, they won’t be able to assign or unassign a role which will allow to do that. This is done to prevent escalation of privileges.
*/
func (a *Client) SetUserRoles(params *SetUserRolesParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*SetUserRolesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewSetUserRolesParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "setUserRoles",
		Method:             "PUT",
		PathPattern:        "/access-control/users/{userId}/roles",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &SetUserRolesReader{formats: a.formats},
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
	success, ok := result.(*SetUserRolesOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for setUserRoles: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  UpdateRoleWithPermissions updates a custom role

  You need to have a permission with action `roles:write` and scope `permissions:type:delegate`. `permissions:type:delegate`` scope ensures that users can only create custom roles with the same, or a subset of permissions which the user has.
*/
func (a *Client) UpdateRoleWithPermissions(params *UpdateRoleWithPermissionsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*UpdateRoleWithPermissionsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateRoleWithPermissionsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "updateRoleWithPermissions",
		Method:             "PUT",
		PathPattern:        "/access-control/roles/{roleUID}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &UpdateRoleWithPermissionsReader{formats: a.formats},
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
	success, ok := result.(*UpdateRoleWithPermissionsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for updateRoleWithPermissions: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
