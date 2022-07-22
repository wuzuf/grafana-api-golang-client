// Code generated by go-swagger; DO NOT EDIT.

package recording_rules

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/grafana/grafana-api-golang-client/goclient/models"
)

// UpdateRecordingRuleReader is a Reader for the UpdateRecordingRule structure.
type UpdateRecordingRuleReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateRecordingRuleReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateRecordingRuleOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewUpdateRecordingRuleUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewUpdateRecordingRuleForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewUpdateRecordingRuleNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewUpdateRecordingRuleInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewUpdateRecordingRuleOK creates a UpdateRecordingRuleOK with default headers values
func NewUpdateRecordingRuleOK() *UpdateRecordingRuleOK {
	return &UpdateRecordingRuleOK{}
}

/* UpdateRecordingRuleOK describes a response with status code 200, with default header values.

(empty)
*/
type UpdateRecordingRuleOK struct {
	Payload *models.RecordingRuleJSON
}

func (o *UpdateRecordingRuleOK) Error() string {
	return fmt.Sprintf("[PUT /recording-rules][%d] updateRecordingRuleOK  %+v", 200, o.Payload)
}
func (o *UpdateRecordingRuleOK) GetPayload() *models.RecordingRuleJSON {
	return o.Payload
}

func (o *UpdateRecordingRuleOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RecordingRuleJSON)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateRecordingRuleUnauthorized creates a UpdateRecordingRuleUnauthorized with default headers values
func NewUpdateRecordingRuleUnauthorized() *UpdateRecordingRuleUnauthorized {
	return &UpdateRecordingRuleUnauthorized{}
}

/* UpdateRecordingRuleUnauthorized describes a response with status code 401, with default header values.

UnauthorizedError is returned when the request is not authenticated.
*/
type UpdateRecordingRuleUnauthorized struct {
	Payload *models.ErrorResponseBody
}

func (o *UpdateRecordingRuleUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /recording-rules][%d] updateRecordingRuleUnauthorized  %+v", 401, o.Payload)
}
func (o *UpdateRecordingRuleUnauthorized) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *UpdateRecordingRuleUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateRecordingRuleForbidden creates a UpdateRecordingRuleForbidden with default headers values
func NewUpdateRecordingRuleForbidden() *UpdateRecordingRuleForbidden {
	return &UpdateRecordingRuleForbidden{}
}

/* UpdateRecordingRuleForbidden describes a response with status code 403, with default header values.

ForbiddenError is returned if the user/token has insufficient permissions to access the requested resource.
*/
type UpdateRecordingRuleForbidden struct {
	Payload *models.ErrorResponseBody
}

func (o *UpdateRecordingRuleForbidden) Error() string {
	return fmt.Sprintf("[PUT /recording-rules][%d] updateRecordingRuleForbidden  %+v", 403, o.Payload)
}
func (o *UpdateRecordingRuleForbidden) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *UpdateRecordingRuleForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateRecordingRuleNotFound creates a UpdateRecordingRuleNotFound with default headers values
func NewUpdateRecordingRuleNotFound() *UpdateRecordingRuleNotFound {
	return &UpdateRecordingRuleNotFound{}
}

/* UpdateRecordingRuleNotFound describes a response with status code 404, with default header values.

NotFoundError is returned when the requested resource was not found.
*/
type UpdateRecordingRuleNotFound struct {
	Payload *models.ErrorResponseBody
}

func (o *UpdateRecordingRuleNotFound) Error() string {
	return fmt.Sprintf("[PUT /recording-rules][%d] updateRecordingRuleNotFound  %+v", 404, o.Payload)
}
func (o *UpdateRecordingRuleNotFound) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *UpdateRecordingRuleNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateRecordingRuleInternalServerError creates a UpdateRecordingRuleInternalServerError with default headers values
func NewUpdateRecordingRuleInternalServerError() *UpdateRecordingRuleInternalServerError {
	return &UpdateRecordingRuleInternalServerError{}
}

/* UpdateRecordingRuleInternalServerError describes a response with status code 500, with default header values.

InternalServerError is a general error indicating something went wrong internally.
*/
type UpdateRecordingRuleInternalServerError struct {
	Payload *models.ErrorResponseBody
}

func (o *UpdateRecordingRuleInternalServerError) Error() string {
	return fmt.Sprintf("[PUT /recording-rules][%d] updateRecordingRuleInternalServerError  %+v", 500, o.Payload)
}
func (o *UpdateRecordingRuleInternalServerError) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *UpdateRecordingRuleInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
