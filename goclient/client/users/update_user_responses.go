// Code generated by go-swagger; DO NOT EDIT.

package users

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/grafana/grafana-api-golang-client/goclient/models"
)

// UpdateUserReader is a Reader for the UpdateUser structure.
type UpdateUserReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateUserReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateUserOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewUpdateUserUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewUpdateUserForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewUpdateUserNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewUpdateUserInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewUpdateUserOK creates a UpdateUserOK with default headers values
func NewUpdateUserOK() *UpdateUserOK {
	return &UpdateUserOK{}
}

/* UpdateUserOK describes a response with status code 200, with default header values.

An OKResponse is returned if the request was successful.
*/
type UpdateUserOK struct {
	Payload *models.SuccessResponseBody
}

func (o *UpdateUserOK) Error() string {
	return fmt.Sprintf("[PUT /users/{user_id}][%d] updateUserOK  %+v", 200, o.Payload)
}
func (o *UpdateUserOK) GetPayload() *models.SuccessResponseBody {
	return o.Payload
}

func (o *UpdateUserOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SuccessResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateUserUnauthorized creates a UpdateUserUnauthorized with default headers values
func NewUpdateUserUnauthorized() *UpdateUserUnauthorized {
	return &UpdateUserUnauthorized{}
}

/* UpdateUserUnauthorized describes a response with status code 401, with default header values.

UnauthorizedError is returned when the request is not authenticated.
*/
type UpdateUserUnauthorized struct {
	Payload *models.ErrorResponseBody
}

func (o *UpdateUserUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /users/{user_id}][%d] updateUserUnauthorized  %+v", 401, o.Payload)
}
func (o *UpdateUserUnauthorized) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *UpdateUserUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateUserForbidden creates a UpdateUserForbidden with default headers values
func NewUpdateUserForbidden() *UpdateUserForbidden {
	return &UpdateUserForbidden{}
}

/* UpdateUserForbidden describes a response with status code 403, with default header values.

ForbiddenError is returned if the user/token has insufficient permissions to access the requested resource.
*/
type UpdateUserForbidden struct {
	Payload *models.ErrorResponseBody
}

func (o *UpdateUserForbidden) Error() string {
	return fmt.Sprintf("[PUT /users/{user_id}][%d] updateUserForbidden  %+v", 403, o.Payload)
}
func (o *UpdateUserForbidden) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *UpdateUserForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateUserNotFound creates a UpdateUserNotFound with default headers values
func NewUpdateUserNotFound() *UpdateUserNotFound {
	return &UpdateUserNotFound{}
}

/* UpdateUserNotFound describes a response with status code 404, with default header values.

NotFoundError is returned when the requested resource was not found.
*/
type UpdateUserNotFound struct {
	Payload *models.ErrorResponseBody
}

func (o *UpdateUserNotFound) Error() string {
	return fmt.Sprintf("[PUT /users/{user_id}][%d] updateUserNotFound  %+v", 404, o.Payload)
}
func (o *UpdateUserNotFound) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *UpdateUserNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateUserInternalServerError creates a UpdateUserInternalServerError with default headers values
func NewUpdateUserInternalServerError() *UpdateUserInternalServerError {
	return &UpdateUserInternalServerError{}
}

/* UpdateUserInternalServerError describes a response with status code 500, with default header values.

InternalServerError is a general error indicating something went wrong internally.
*/
type UpdateUserInternalServerError struct {
	Payload *models.ErrorResponseBody
}

func (o *UpdateUserInternalServerError) Error() string {
	return fmt.Sprintf("[PUT /users/{user_id}][%d] updateUserInternalServerError  %+v", 500, o.Payload)
}
func (o *UpdateUserInternalServerError) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *UpdateUserInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
