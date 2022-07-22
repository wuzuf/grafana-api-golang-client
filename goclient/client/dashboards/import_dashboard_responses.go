// Code generated by go-swagger; DO NOT EDIT.

package dashboards

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/grafana/grafana-api-golang-client/goclient/models"
)

// ImportDashboardReader is a Reader for the ImportDashboard structure.
type ImportDashboardReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ImportDashboardReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewImportDashboardOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewImportDashboardBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewImportDashboardUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 412:
		result := NewImportDashboardPreconditionFailed()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 422:
		result := NewImportDashboardUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewImportDashboardInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewImportDashboardOK creates a ImportDashboardOK with default headers values
func NewImportDashboardOK() *ImportDashboardOK {
	return &ImportDashboardOK{}
}

/* ImportDashboardOK describes a response with status code 200, with default header values.

(empty)
*/
type ImportDashboardOK struct {
	Payload *models.ImportDashboardResponse
}

func (o *ImportDashboardOK) Error() string {
	return fmt.Sprintf("[POST /dashboards/import][%d] importDashboardOK  %+v", 200, o.Payload)
}
func (o *ImportDashboardOK) GetPayload() *models.ImportDashboardResponse {
	return o.Payload
}

func (o *ImportDashboardOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ImportDashboardResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewImportDashboardBadRequest creates a ImportDashboardBadRequest with default headers values
func NewImportDashboardBadRequest() *ImportDashboardBadRequest {
	return &ImportDashboardBadRequest{}
}

/* ImportDashboardBadRequest describes a response with status code 400, with default header values.

BadRequestError is returned when the request is invalid and it cannot be processed.
*/
type ImportDashboardBadRequest struct {
	Payload *models.ErrorResponseBody
}

func (o *ImportDashboardBadRequest) Error() string {
	return fmt.Sprintf("[POST /dashboards/import][%d] importDashboardBadRequest  %+v", 400, o.Payload)
}
func (o *ImportDashboardBadRequest) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *ImportDashboardBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewImportDashboardUnauthorized creates a ImportDashboardUnauthorized with default headers values
func NewImportDashboardUnauthorized() *ImportDashboardUnauthorized {
	return &ImportDashboardUnauthorized{}
}

/* ImportDashboardUnauthorized describes a response with status code 401, with default header values.

UnauthorizedError is returned when the request is not authenticated.
*/
type ImportDashboardUnauthorized struct {
	Payload *models.ErrorResponseBody
}

func (o *ImportDashboardUnauthorized) Error() string {
	return fmt.Sprintf("[POST /dashboards/import][%d] importDashboardUnauthorized  %+v", 401, o.Payload)
}
func (o *ImportDashboardUnauthorized) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *ImportDashboardUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewImportDashboardPreconditionFailed creates a ImportDashboardPreconditionFailed with default headers values
func NewImportDashboardPreconditionFailed() *ImportDashboardPreconditionFailed {
	return &ImportDashboardPreconditionFailed{}
}

/* ImportDashboardPreconditionFailed describes a response with status code 412, with default header values.

PreconditionFailedError
*/
type ImportDashboardPreconditionFailed struct {
	Payload *models.ErrorResponseBody
}

func (o *ImportDashboardPreconditionFailed) Error() string {
	return fmt.Sprintf("[POST /dashboards/import][%d] importDashboardPreconditionFailed  %+v", 412, o.Payload)
}
func (o *ImportDashboardPreconditionFailed) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *ImportDashboardPreconditionFailed) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewImportDashboardUnprocessableEntity creates a ImportDashboardUnprocessableEntity with default headers values
func NewImportDashboardUnprocessableEntity() *ImportDashboardUnprocessableEntity {
	return &ImportDashboardUnprocessableEntity{}
}

/* ImportDashboardUnprocessableEntity describes a response with status code 422, with default header values.

UnprocessableEntityError
*/
type ImportDashboardUnprocessableEntity struct {
	Payload *models.ErrorResponseBody
}

func (o *ImportDashboardUnprocessableEntity) Error() string {
	return fmt.Sprintf("[POST /dashboards/import][%d] importDashboardUnprocessableEntity  %+v", 422, o.Payload)
}
func (o *ImportDashboardUnprocessableEntity) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *ImportDashboardUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewImportDashboardInternalServerError creates a ImportDashboardInternalServerError with default headers values
func NewImportDashboardInternalServerError() *ImportDashboardInternalServerError {
	return &ImportDashboardInternalServerError{}
}

/* ImportDashboardInternalServerError describes a response with status code 500, with default header values.

InternalServerError is a general error indicating something went wrong internally.
*/
type ImportDashboardInternalServerError struct {
	Payload *models.ErrorResponseBody
}

func (o *ImportDashboardInternalServerError) Error() string {
	return fmt.Sprintf("[POST /dashboards/import][%d] importDashboardInternalServerError  %+v", 500, o.Payload)
}
func (o *ImportDashboardInternalServerError) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *ImportDashboardInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
