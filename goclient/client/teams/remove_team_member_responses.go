// Code generated by go-swagger; DO NOT EDIT.

package teams

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/grafana/grafana-api-golang-client/goclient/models"
)

// RemoveTeamMemberReader is a Reader for the RemoveTeamMember structure.
type RemoveTeamMemberReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *RemoveTeamMemberReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewRemoveTeamMemberOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewRemoveTeamMemberUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewRemoveTeamMemberForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewRemoveTeamMemberNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewRemoveTeamMemberInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewRemoveTeamMemberOK creates a RemoveTeamMemberOK with default headers values
func NewRemoveTeamMemberOK() *RemoveTeamMemberOK {
	return &RemoveTeamMemberOK{}
}

/* RemoveTeamMemberOK describes a response with status code 200, with default header values.

An OKResponse is returned if the request was successful.
*/
type RemoveTeamMemberOK struct {
	Payload *models.SuccessResponseBody
}

func (o *RemoveTeamMemberOK) Error() string {
	return fmt.Sprintf("[DELETE /teams/{team_id}/members/{user_id}][%d] removeTeamMemberOK  %+v", 200, o.Payload)
}
func (o *RemoveTeamMemberOK) GetPayload() *models.SuccessResponseBody {
	return o.Payload
}

func (o *RemoveTeamMemberOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SuccessResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRemoveTeamMemberUnauthorized creates a RemoveTeamMemberUnauthorized with default headers values
func NewRemoveTeamMemberUnauthorized() *RemoveTeamMemberUnauthorized {
	return &RemoveTeamMemberUnauthorized{}
}

/* RemoveTeamMemberUnauthorized describes a response with status code 401, with default header values.

UnauthorizedError is returned when the request is not authenticated.
*/
type RemoveTeamMemberUnauthorized struct {
	Payload *models.ErrorResponseBody
}

func (o *RemoveTeamMemberUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /teams/{team_id}/members/{user_id}][%d] removeTeamMemberUnauthorized  %+v", 401, o.Payload)
}
func (o *RemoveTeamMemberUnauthorized) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *RemoveTeamMemberUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRemoveTeamMemberForbidden creates a RemoveTeamMemberForbidden with default headers values
func NewRemoveTeamMemberForbidden() *RemoveTeamMemberForbidden {
	return &RemoveTeamMemberForbidden{}
}

/* RemoveTeamMemberForbidden describes a response with status code 403, with default header values.

ForbiddenError is returned if the user/token has insufficient permissions to access the requested resource.
*/
type RemoveTeamMemberForbidden struct {
	Payload *models.ErrorResponseBody
}

func (o *RemoveTeamMemberForbidden) Error() string {
	return fmt.Sprintf("[DELETE /teams/{team_id}/members/{user_id}][%d] removeTeamMemberForbidden  %+v", 403, o.Payload)
}
func (o *RemoveTeamMemberForbidden) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *RemoveTeamMemberForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRemoveTeamMemberNotFound creates a RemoveTeamMemberNotFound with default headers values
func NewRemoveTeamMemberNotFound() *RemoveTeamMemberNotFound {
	return &RemoveTeamMemberNotFound{}
}

/* RemoveTeamMemberNotFound describes a response with status code 404, with default header values.

NotFoundError is returned when the requested resource was not found.
*/
type RemoveTeamMemberNotFound struct {
	Payload *models.ErrorResponseBody
}

func (o *RemoveTeamMemberNotFound) Error() string {
	return fmt.Sprintf("[DELETE /teams/{team_id}/members/{user_id}][%d] removeTeamMemberNotFound  %+v", 404, o.Payload)
}
func (o *RemoveTeamMemberNotFound) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *RemoveTeamMemberNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRemoveTeamMemberInternalServerError creates a RemoveTeamMemberInternalServerError with default headers values
func NewRemoveTeamMemberInternalServerError() *RemoveTeamMemberInternalServerError {
	return &RemoveTeamMemberInternalServerError{}
}

/* RemoveTeamMemberInternalServerError describes a response with status code 500, with default header values.

InternalServerError is a general error indicating something went wrong internally.
*/
type RemoveTeamMemberInternalServerError struct {
	Payload *models.ErrorResponseBody
}

func (o *RemoveTeamMemberInternalServerError) Error() string {
	return fmt.Sprintf("[DELETE /teams/{team_id}/members/{user_id}][%d] removeTeamMemberInternalServerError  %+v", 500, o.Payload)
}
func (o *RemoveTeamMemberInternalServerError) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *RemoveTeamMemberInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
