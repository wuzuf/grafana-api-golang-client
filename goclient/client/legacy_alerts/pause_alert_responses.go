// Code generated by go-swagger; DO NOT EDIT.

package legacy_alerts

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/grafana/grafana-api-golang-client/goclient/models"
)

// PauseAlertReader is a Reader for the PauseAlert structure.
type PauseAlertReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PauseAlertReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPauseAlertOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewPauseAlertUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPauseAlertForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPauseAlertNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPauseAlertInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPauseAlertOK creates a PauseAlertOK with default headers values
func NewPauseAlertOK() *PauseAlertOK {
	return &PauseAlertOK{}
}

/* PauseAlertOK describes a response with status code 200, with default header values.

(empty)
*/
type PauseAlertOK struct {
	Payload *models.PauseAlertOKBody
}

func (o *PauseAlertOK) Error() string {
	return fmt.Sprintf("[POST /alerts/{alert_id}/pause][%d] pauseAlertOK  %+v", 200, o.Payload)
}
func (o *PauseAlertOK) GetPayload() *models.PauseAlertOKBody {
	return o.Payload
}

func (o *PauseAlertOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PauseAlertOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPauseAlertUnauthorized creates a PauseAlertUnauthorized with default headers values
func NewPauseAlertUnauthorized() *PauseAlertUnauthorized {
	return &PauseAlertUnauthorized{}
}

/* PauseAlertUnauthorized describes a response with status code 401, with default header values.

UnauthorizedError is returned when the request is not authenticated.
*/
type PauseAlertUnauthorized struct {
	Payload *models.ErrorResponseBody
}

func (o *PauseAlertUnauthorized) Error() string {
	return fmt.Sprintf("[POST /alerts/{alert_id}/pause][%d] pauseAlertUnauthorized  %+v", 401, o.Payload)
}
func (o *PauseAlertUnauthorized) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *PauseAlertUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPauseAlertForbidden creates a PauseAlertForbidden with default headers values
func NewPauseAlertForbidden() *PauseAlertForbidden {
	return &PauseAlertForbidden{}
}

/* PauseAlertForbidden describes a response with status code 403, with default header values.

ForbiddenError is returned if the user/token has insufficient permissions to access the requested resource.
*/
type PauseAlertForbidden struct {
	Payload *models.ErrorResponseBody
}

func (o *PauseAlertForbidden) Error() string {
	return fmt.Sprintf("[POST /alerts/{alert_id}/pause][%d] pauseAlertForbidden  %+v", 403, o.Payload)
}
func (o *PauseAlertForbidden) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *PauseAlertForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPauseAlertNotFound creates a PauseAlertNotFound with default headers values
func NewPauseAlertNotFound() *PauseAlertNotFound {
	return &PauseAlertNotFound{}
}

/* PauseAlertNotFound describes a response with status code 404, with default header values.

NotFoundError is returned when the requested resource was not found.
*/
type PauseAlertNotFound struct {
	Payload *models.ErrorResponseBody
}

func (o *PauseAlertNotFound) Error() string {
	return fmt.Sprintf("[POST /alerts/{alert_id}/pause][%d] pauseAlertNotFound  %+v", 404, o.Payload)
}
func (o *PauseAlertNotFound) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *PauseAlertNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPauseAlertInternalServerError creates a PauseAlertInternalServerError with default headers values
func NewPauseAlertInternalServerError() *PauseAlertInternalServerError {
	return &PauseAlertInternalServerError{}
}

/* PauseAlertInternalServerError describes a response with status code 500, with default header values.

InternalServerError is a general error indicating something went wrong internally.
*/
type PauseAlertInternalServerError struct {
	Payload *models.ErrorResponseBody
}

func (o *PauseAlertInternalServerError) Error() string {
	return fmt.Sprintf("[POST /alerts/{alert_id}/pause][%d] pauseAlertInternalServerError  %+v", 500, o.Payload)
}
func (o *PauseAlertInternalServerError) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *PauseAlertInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
