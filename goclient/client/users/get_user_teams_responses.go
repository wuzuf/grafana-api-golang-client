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

// GetUserTeamsReader is a Reader for the GetUserTeams structure.
type GetUserTeamsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetUserTeamsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetUserTeamsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewGetUserTeamsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetUserTeamsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetUserTeamsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetUserTeamsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetUserTeamsOK creates a GetUserTeamsOK with default headers values
func NewGetUserTeamsOK() *GetUserTeamsOK {
	return &GetUserTeamsOK{}
}

/* GetUserTeamsOK describes a response with status code 200, with default header values.

(empty)
*/
type GetUserTeamsOK struct {
	Payload []*models.TeamDTO
}

func (o *GetUserTeamsOK) Error() string {
	return fmt.Sprintf("[GET /users/{user_id}/teams][%d] getUserTeamsOK  %+v", 200, o.Payload)
}
func (o *GetUserTeamsOK) GetPayload() []*models.TeamDTO {
	return o.Payload
}

func (o *GetUserTeamsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUserTeamsUnauthorized creates a GetUserTeamsUnauthorized with default headers values
func NewGetUserTeamsUnauthorized() *GetUserTeamsUnauthorized {
	return &GetUserTeamsUnauthorized{}
}

/* GetUserTeamsUnauthorized describes a response with status code 401, with default header values.

UnauthorizedError is returned when the request is not authenticated.
*/
type GetUserTeamsUnauthorized struct {
	Payload *models.ErrorResponseBody
}

func (o *GetUserTeamsUnauthorized) Error() string {
	return fmt.Sprintf("[GET /users/{user_id}/teams][%d] getUserTeamsUnauthorized  %+v", 401, o.Payload)
}
func (o *GetUserTeamsUnauthorized) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *GetUserTeamsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUserTeamsForbidden creates a GetUserTeamsForbidden with default headers values
func NewGetUserTeamsForbidden() *GetUserTeamsForbidden {
	return &GetUserTeamsForbidden{}
}

/* GetUserTeamsForbidden describes a response with status code 403, with default header values.

ForbiddenError is returned if the user/token has insufficient permissions to access the requested resource.
*/
type GetUserTeamsForbidden struct {
	Payload *models.ErrorResponseBody
}

func (o *GetUserTeamsForbidden) Error() string {
	return fmt.Sprintf("[GET /users/{user_id}/teams][%d] getUserTeamsForbidden  %+v", 403, o.Payload)
}
func (o *GetUserTeamsForbidden) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *GetUserTeamsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUserTeamsNotFound creates a GetUserTeamsNotFound with default headers values
func NewGetUserTeamsNotFound() *GetUserTeamsNotFound {
	return &GetUserTeamsNotFound{}
}

/* GetUserTeamsNotFound describes a response with status code 404, with default header values.

NotFoundError is returned when the requested resource was not found.
*/
type GetUserTeamsNotFound struct {
	Payload *models.ErrorResponseBody
}

func (o *GetUserTeamsNotFound) Error() string {
	return fmt.Sprintf("[GET /users/{user_id}/teams][%d] getUserTeamsNotFound  %+v", 404, o.Payload)
}
func (o *GetUserTeamsNotFound) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *GetUserTeamsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUserTeamsInternalServerError creates a GetUserTeamsInternalServerError with default headers values
func NewGetUserTeamsInternalServerError() *GetUserTeamsInternalServerError {
	return &GetUserTeamsInternalServerError{}
}

/* GetUserTeamsInternalServerError describes a response with status code 500, with default header values.

InternalServerError is a general error indicating something went wrong internally.
*/
type GetUserTeamsInternalServerError struct {
	Payload *models.ErrorResponseBody
}

func (o *GetUserTeamsInternalServerError) Error() string {
	return fmt.Sprintf("[GET /users/{user_id}/teams][%d] getUserTeamsInternalServerError  %+v", 500, o.Payload)
}
func (o *GetUserTeamsInternalServerError) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *GetUserTeamsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
