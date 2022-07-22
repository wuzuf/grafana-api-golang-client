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

// SetUserRolesReader is a Reader for the SetUserRoles structure.
type SetUserRolesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SetUserRolesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewSetUserRolesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewSetUserRolesBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewSetUserRolesForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewSetUserRolesNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewSetUserRolesInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewSetUserRolesOK creates a SetUserRolesOK with default headers values
func NewSetUserRolesOK() *SetUserRolesOK {
	return &SetUserRolesOK{}
}

/* SetUserRolesOK describes a response with status code 200, with default header values.

An OKResponse is returned if the request was successful.
*/
type SetUserRolesOK struct {
	Payload *models.SuccessResponseBody
}

func (o *SetUserRolesOK) Error() string {
	return fmt.Sprintf("[PUT /access-control/users/{userId}/roles][%d] setUserRolesOK  %+v", 200, o.Payload)
}
func (o *SetUserRolesOK) GetPayload() *models.SuccessResponseBody {
	return o.Payload
}

func (o *SetUserRolesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SuccessResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSetUserRolesBadRequest creates a SetUserRolesBadRequest with default headers values
func NewSetUserRolesBadRequest() *SetUserRolesBadRequest {
	return &SetUserRolesBadRequest{}
}

/* SetUserRolesBadRequest describes a response with status code 400, with default header values.

BadRequestError is returned when the request is invalid and it cannot be processed.
*/
type SetUserRolesBadRequest struct {
	Payload *models.ErrorResponseBody
}

func (o *SetUserRolesBadRequest) Error() string {
	return fmt.Sprintf("[PUT /access-control/users/{userId}/roles][%d] setUserRolesBadRequest  %+v", 400, o.Payload)
}
func (o *SetUserRolesBadRequest) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *SetUserRolesBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSetUserRolesForbidden creates a SetUserRolesForbidden with default headers values
func NewSetUserRolesForbidden() *SetUserRolesForbidden {
	return &SetUserRolesForbidden{}
}

/* SetUserRolesForbidden describes a response with status code 403, with default header values.

ForbiddenError is returned if the user/token has insufficient permissions to access the requested resource.
*/
type SetUserRolesForbidden struct {
	Payload *models.ErrorResponseBody
}

func (o *SetUserRolesForbidden) Error() string {
	return fmt.Sprintf("[PUT /access-control/users/{userId}/roles][%d] setUserRolesForbidden  %+v", 403, o.Payload)
}
func (o *SetUserRolesForbidden) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *SetUserRolesForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSetUserRolesNotFound creates a SetUserRolesNotFound with default headers values
func NewSetUserRolesNotFound() *SetUserRolesNotFound {
	return &SetUserRolesNotFound{}
}

/* SetUserRolesNotFound describes a response with status code 404, with default header values.

NotFoundError is returned when the requested resource was not found.
*/
type SetUserRolesNotFound struct {
	Payload *models.ErrorResponseBody
}

func (o *SetUserRolesNotFound) Error() string {
	return fmt.Sprintf("[PUT /access-control/users/{userId}/roles][%d] setUserRolesNotFound  %+v", 404, o.Payload)
}
func (o *SetUserRolesNotFound) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *SetUserRolesNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSetUserRolesInternalServerError creates a SetUserRolesInternalServerError with default headers values
func NewSetUserRolesInternalServerError() *SetUserRolesInternalServerError {
	return &SetUserRolesInternalServerError{}
}

/* SetUserRolesInternalServerError describes a response with status code 500, with default header values.

InternalServerError is a general error indicating something went wrong internally.
*/
type SetUserRolesInternalServerError struct {
	Payload *models.ErrorResponseBody
}

func (o *SetUserRolesInternalServerError) Error() string {
	return fmt.Sprintf("[PUT /access-control/users/{userId}/roles][%d] setUserRolesInternalServerError  %+v", 500, o.Payload)
}
func (o *SetUserRolesInternalServerError) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *SetUserRolesInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
