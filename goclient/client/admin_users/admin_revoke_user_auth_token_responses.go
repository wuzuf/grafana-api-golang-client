// Code generated by go-swagger; DO NOT EDIT.

package admin_users

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/grafana/grafana-api-golang-client/goclient/models"
)

// AdminRevokeUserAuthTokenReader is a Reader for the AdminRevokeUserAuthToken structure.
type AdminRevokeUserAuthTokenReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AdminRevokeUserAuthTokenReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewAdminRevokeUserAuthTokenOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewAdminRevokeUserAuthTokenBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewAdminRevokeUserAuthTokenUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewAdminRevokeUserAuthTokenForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewAdminRevokeUserAuthTokenNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewAdminRevokeUserAuthTokenInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewAdminRevokeUserAuthTokenOK creates a AdminRevokeUserAuthTokenOK with default headers values
func NewAdminRevokeUserAuthTokenOK() *AdminRevokeUserAuthTokenOK {
	return &AdminRevokeUserAuthTokenOK{}
}

/* AdminRevokeUserAuthTokenOK describes a response with status code 200, with default header values.

An OKResponse is returned if the request was successful.
*/
type AdminRevokeUserAuthTokenOK struct {
	Payload *models.SuccessResponseBody
}

func (o *AdminRevokeUserAuthTokenOK) Error() string {
	return fmt.Sprintf("[POST /admin/users/{user_id}/revoke-auth-token][%d] adminRevokeUserAuthTokenOK  %+v", 200, o.Payload)
}
func (o *AdminRevokeUserAuthTokenOK) GetPayload() *models.SuccessResponseBody {
	return o.Payload
}

func (o *AdminRevokeUserAuthTokenOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SuccessResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAdminRevokeUserAuthTokenBadRequest creates a AdminRevokeUserAuthTokenBadRequest with default headers values
func NewAdminRevokeUserAuthTokenBadRequest() *AdminRevokeUserAuthTokenBadRequest {
	return &AdminRevokeUserAuthTokenBadRequest{}
}

/* AdminRevokeUserAuthTokenBadRequest describes a response with status code 400, with default header values.

BadRequestError is returned when the request is invalid and it cannot be processed.
*/
type AdminRevokeUserAuthTokenBadRequest struct {
	Payload *models.ErrorResponseBody
}

func (o *AdminRevokeUserAuthTokenBadRequest) Error() string {
	return fmt.Sprintf("[POST /admin/users/{user_id}/revoke-auth-token][%d] adminRevokeUserAuthTokenBadRequest  %+v", 400, o.Payload)
}
func (o *AdminRevokeUserAuthTokenBadRequest) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *AdminRevokeUserAuthTokenBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAdminRevokeUserAuthTokenUnauthorized creates a AdminRevokeUserAuthTokenUnauthorized with default headers values
func NewAdminRevokeUserAuthTokenUnauthorized() *AdminRevokeUserAuthTokenUnauthorized {
	return &AdminRevokeUserAuthTokenUnauthorized{}
}

/* AdminRevokeUserAuthTokenUnauthorized describes a response with status code 401, with default header values.

UnauthorizedError is returned when the request is not authenticated.
*/
type AdminRevokeUserAuthTokenUnauthorized struct {
	Payload *models.ErrorResponseBody
}

func (o *AdminRevokeUserAuthTokenUnauthorized) Error() string {
	return fmt.Sprintf("[POST /admin/users/{user_id}/revoke-auth-token][%d] adminRevokeUserAuthTokenUnauthorized  %+v", 401, o.Payload)
}
func (o *AdminRevokeUserAuthTokenUnauthorized) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *AdminRevokeUserAuthTokenUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAdminRevokeUserAuthTokenForbidden creates a AdminRevokeUserAuthTokenForbidden with default headers values
func NewAdminRevokeUserAuthTokenForbidden() *AdminRevokeUserAuthTokenForbidden {
	return &AdminRevokeUserAuthTokenForbidden{}
}

/* AdminRevokeUserAuthTokenForbidden describes a response with status code 403, with default header values.

ForbiddenError is returned if the user/token has insufficient permissions to access the requested resource.
*/
type AdminRevokeUserAuthTokenForbidden struct {
	Payload *models.ErrorResponseBody
}

func (o *AdminRevokeUserAuthTokenForbidden) Error() string {
	return fmt.Sprintf("[POST /admin/users/{user_id}/revoke-auth-token][%d] adminRevokeUserAuthTokenForbidden  %+v", 403, o.Payload)
}
func (o *AdminRevokeUserAuthTokenForbidden) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *AdminRevokeUserAuthTokenForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAdminRevokeUserAuthTokenNotFound creates a AdminRevokeUserAuthTokenNotFound with default headers values
func NewAdminRevokeUserAuthTokenNotFound() *AdminRevokeUserAuthTokenNotFound {
	return &AdminRevokeUserAuthTokenNotFound{}
}

/* AdminRevokeUserAuthTokenNotFound describes a response with status code 404, with default header values.

NotFoundError is returned when the requested resource was not found.
*/
type AdminRevokeUserAuthTokenNotFound struct {
	Payload *models.ErrorResponseBody
}

func (o *AdminRevokeUserAuthTokenNotFound) Error() string {
	return fmt.Sprintf("[POST /admin/users/{user_id}/revoke-auth-token][%d] adminRevokeUserAuthTokenNotFound  %+v", 404, o.Payload)
}
func (o *AdminRevokeUserAuthTokenNotFound) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *AdminRevokeUserAuthTokenNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAdminRevokeUserAuthTokenInternalServerError creates a AdminRevokeUserAuthTokenInternalServerError with default headers values
func NewAdminRevokeUserAuthTokenInternalServerError() *AdminRevokeUserAuthTokenInternalServerError {
	return &AdminRevokeUserAuthTokenInternalServerError{}
}

/* AdminRevokeUserAuthTokenInternalServerError describes a response with status code 500, with default header values.

InternalServerError is a general error indicating something went wrong internally.
*/
type AdminRevokeUserAuthTokenInternalServerError struct {
	Payload *models.ErrorResponseBody
}

func (o *AdminRevokeUserAuthTokenInternalServerError) Error() string {
	return fmt.Sprintf("[POST /admin/users/{user_id}/revoke-auth-token][%d] adminRevokeUserAuthTokenInternalServerError  %+v", 500, o.Payload)
}
func (o *AdminRevokeUserAuthTokenInternalServerError) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *AdminRevokeUserAuthTokenInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
