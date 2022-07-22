// Code generated by go-swagger; DO NOT EDIT.

package access_control

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/grafana/grafana-api-golang-client/goclient/models"
)

// AddTeamRoleReader is a Reader for the AddTeamRole structure.
type AddTeamRoleReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AddTeamRoleReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewAddTeamRoleOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewAddTeamRoleBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewAddTeamRoleForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewAddTeamRoleNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewAddTeamRoleInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewAddTeamRoleOK creates a AddTeamRoleOK with default headers values
func NewAddTeamRoleOK() *AddTeamRoleOK {
	return &AddTeamRoleOK{}
}

/* AddTeamRoleOK describes a response with status code 200, with default header values.

An OKResponse is returned if the request was successful.
*/
type AddTeamRoleOK struct {
	Payload *models.SuccessResponseBody
}

func (o *AddTeamRoleOK) Error() string {
	return fmt.Sprintf("[POST /access-control/teams/{teamId}/roles][%d] addTeamRoleOK  %+v", 200, o.Payload)
}
func (o *AddTeamRoleOK) GetPayload() *models.SuccessResponseBody {
	return o.Payload
}

func (o *AddTeamRoleOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SuccessResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAddTeamRoleBadRequest creates a AddTeamRoleBadRequest with default headers values
func NewAddTeamRoleBadRequest() *AddTeamRoleBadRequest {
	return &AddTeamRoleBadRequest{}
}

/* AddTeamRoleBadRequest describes a response with status code 400, with default header values.

BadRequestError is returned when the request is invalid and it cannot be processed.
*/
type AddTeamRoleBadRequest struct {
	Payload *models.ErrorResponseBody
}

func (o *AddTeamRoleBadRequest) Error() string {
	return fmt.Sprintf("[POST /access-control/teams/{teamId}/roles][%d] addTeamRoleBadRequest  %+v", 400, o.Payload)
}
func (o *AddTeamRoleBadRequest) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *AddTeamRoleBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAddTeamRoleForbidden creates a AddTeamRoleForbidden with default headers values
func NewAddTeamRoleForbidden() *AddTeamRoleForbidden {
	return &AddTeamRoleForbidden{}
}

/* AddTeamRoleForbidden describes a response with status code 403, with default header values.

ForbiddenError is returned if the user/token has insufficient permissions to access the requested resource.
*/
type AddTeamRoleForbidden struct {
	Payload *models.ErrorResponseBody
}

func (o *AddTeamRoleForbidden) Error() string {
	return fmt.Sprintf("[POST /access-control/teams/{teamId}/roles][%d] addTeamRoleForbidden  %+v", 403, o.Payload)
}
func (o *AddTeamRoleForbidden) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *AddTeamRoleForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAddTeamRoleNotFound creates a AddTeamRoleNotFound with default headers values
func NewAddTeamRoleNotFound() *AddTeamRoleNotFound {
	return &AddTeamRoleNotFound{}
}

/* AddTeamRoleNotFound describes a response with status code 404, with default header values.

NotFoundError is returned when the requested resource was not found.
*/
type AddTeamRoleNotFound struct {
	Payload *models.ErrorResponseBody
}

func (o *AddTeamRoleNotFound) Error() string {
	return fmt.Sprintf("[POST /access-control/teams/{teamId}/roles][%d] addTeamRoleNotFound  %+v", 404, o.Payload)
}
func (o *AddTeamRoleNotFound) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *AddTeamRoleNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAddTeamRoleInternalServerError creates a AddTeamRoleInternalServerError with default headers values
func NewAddTeamRoleInternalServerError() *AddTeamRoleInternalServerError {
	return &AddTeamRoleInternalServerError{}
}

/* AddTeamRoleInternalServerError describes a response with status code 500, with default header values.

InternalServerError is a general error indicating something went wrong internally.
*/
type AddTeamRoleInternalServerError struct {
	Payload *models.ErrorResponseBody
}

func (o *AddTeamRoleInternalServerError) Error() string {
	return fmt.Sprintf("[POST /access-control/teams/{teamId}/roles][%d] addTeamRoleInternalServerError  %+v", 500, o.Payload)
}
func (o *AddTeamRoleInternalServerError) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *AddTeamRoleInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
