// Code generated by go-swagger; DO NOT EDIT.

package org

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/grafana/grafana-api-golang-client/goclient/models"
)

// UpdateOrgUserForCurrentOrgReader is a Reader for the UpdateOrgUserForCurrentOrg structure.
type UpdateOrgUserForCurrentOrgReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateOrgUserForCurrentOrgReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateOrgUserForCurrentOrgOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewUpdateOrgUserForCurrentOrgBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewUpdateOrgUserForCurrentOrgUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewUpdateOrgUserForCurrentOrgForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewUpdateOrgUserForCurrentOrgInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewUpdateOrgUserForCurrentOrgOK creates a UpdateOrgUserForCurrentOrgOK with default headers values
func NewUpdateOrgUserForCurrentOrgOK() *UpdateOrgUserForCurrentOrgOK {
	return &UpdateOrgUserForCurrentOrgOK{}
}

/* UpdateOrgUserForCurrentOrgOK describes a response with status code 200, with default header values.

An OKResponse is returned if the request was successful.
*/
type UpdateOrgUserForCurrentOrgOK struct {
	Payload *models.SuccessResponseBody
}

func (o *UpdateOrgUserForCurrentOrgOK) Error() string {
	return fmt.Sprintf("[PATCH /org/users/{user_id}][%d] updateOrgUserForCurrentOrgOK  %+v", 200, o.Payload)
}
func (o *UpdateOrgUserForCurrentOrgOK) GetPayload() *models.SuccessResponseBody {
	return o.Payload
}

func (o *UpdateOrgUserForCurrentOrgOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SuccessResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateOrgUserForCurrentOrgBadRequest creates a UpdateOrgUserForCurrentOrgBadRequest with default headers values
func NewUpdateOrgUserForCurrentOrgBadRequest() *UpdateOrgUserForCurrentOrgBadRequest {
	return &UpdateOrgUserForCurrentOrgBadRequest{}
}

/* UpdateOrgUserForCurrentOrgBadRequest describes a response with status code 400, with default header values.

BadRequestError is returned when the request is invalid and it cannot be processed.
*/
type UpdateOrgUserForCurrentOrgBadRequest struct {
	Payload *models.ErrorResponseBody
}

func (o *UpdateOrgUserForCurrentOrgBadRequest) Error() string {
	return fmt.Sprintf("[PATCH /org/users/{user_id}][%d] updateOrgUserForCurrentOrgBadRequest  %+v", 400, o.Payload)
}
func (o *UpdateOrgUserForCurrentOrgBadRequest) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *UpdateOrgUserForCurrentOrgBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateOrgUserForCurrentOrgUnauthorized creates a UpdateOrgUserForCurrentOrgUnauthorized with default headers values
func NewUpdateOrgUserForCurrentOrgUnauthorized() *UpdateOrgUserForCurrentOrgUnauthorized {
	return &UpdateOrgUserForCurrentOrgUnauthorized{}
}

/* UpdateOrgUserForCurrentOrgUnauthorized describes a response with status code 401, with default header values.

UnauthorizedError is returned when the request is not authenticated.
*/
type UpdateOrgUserForCurrentOrgUnauthorized struct {
	Payload *models.ErrorResponseBody
}

func (o *UpdateOrgUserForCurrentOrgUnauthorized) Error() string {
	return fmt.Sprintf("[PATCH /org/users/{user_id}][%d] updateOrgUserForCurrentOrgUnauthorized  %+v", 401, o.Payload)
}
func (o *UpdateOrgUserForCurrentOrgUnauthorized) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *UpdateOrgUserForCurrentOrgUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateOrgUserForCurrentOrgForbidden creates a UpdateOrgUserForCurrentOrgForbidden with default headers values
func NewUpdateOrgUserForCurrentOrgForbidden() *UpdateOrgUserForCurrentOrgForbidden {
	return &UpdateOrgUserForCurrentOrgForbidden{}
}

/* UpdateOrgUserForCurrentOrgForbidden describes a response with status code 403, with default header values.

ForbiddenError is returned if the user/token has insufficient permissions to access the requested resource.
*/
type UpdateOrgUserForCurrentOrgForbidden struct {
	Payload *models.ErrorResponseBody
}

func (o *UpdateOrgUserForCurrentOrgForbidden) Error() string {
	return fmt.Sprintf("[PATCH /org/users/{user_id}][%d] updateOrgUserForCurrentOrgForbidden  %+v", 403, o.Payload)
}
func (o *UpdateOrgUserForCurrentOrgForbidden) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *UpdateOrgUserForCurrentOrgForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateOrgUserForCurrentOrgInternalServerError creates a UpdateOrgUserForCurrentOrgInternalServerError with default headers values
func NewUpdateOrgUserForCurrentOrgInternalServerError() *UpdateOrgUserForCurrentOrgInternalServerError {
	return &UpdateOrgUserForCurrentOrgInternalServerError{}
}

/* UpdateOrgUserForCurrentOrgInternalServerError describes a response with status code 500, with default header values.

InternalServerError is a general error indicating something went wrong internally.
*/
type UpdateOrgUserForCurrentOrgInternalServerError struct {
	Payload *models.ErrorResponseBody
}

func (o *UpdateOrgUserForCurrentOrgInternalServerError) Error() string {
	return fmt.Sprintf("[PATCH /org/users/{user_id}][%d] updateOrgUserForCurrentOrgInternalServerError  %+v", 500, o.Payload)
}
func (o *UpdateOrgUserForCurrentOrgInternalServerError) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *UpdateOrgUserForCurrentOrgInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
