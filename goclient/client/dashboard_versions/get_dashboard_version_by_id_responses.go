// Code generated by go-swagger; DO NOT EDIT.

package dashboard_versions

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/grafana/grafana-api-golang-client/goclient/models"
)

// GetDashboardVersionByIDReader is a Reader for the GetDashboardVersionByID structure.
type GetDashboardVersionByIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetDashboardVersionByIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetDashboardVersionByIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewGetDashboardVersionByIDUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetDashboardVersionByIDForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetDashboardVersionByIDNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetDashboardVersionByIDInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetDashboardVersionByIDOK creates a GetDashboardVersionByIDOK with default headers values
func NewGetDashboardVersionByIDOK() *GetDashboardVersionByIDOK {
	return &GetDashboardVersionByIDOK{}
}

/* GetDashboardVersionByIDOK describes a response with status code 200, with default header values.

(empty)
*/
type GetDashboardVersionByIDOK struct {
	Payload *models.DashboardVersionMeta
}

func (o *GetDashboardVersionByIDOK) Error() string {
	return fmt.Sprintf("[GET /dashboards/id/{DashboardID}/versions/{DashboardVersionID}][%d] getDashboardVersionByIdOK  %+v", 200, o.Payload)
}
func (o *GetDashboardVersionByIDOK) GetPayload() *models.DashboardVersionMeta {
	return o.Payload
}

func (o *GetDashboardVersionByIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DashboardVersionMeta)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetDashboardVersionByIDUnauthorized creates a GetDashboardVersionByIDUnauthorized with default headers values
func NewGetDashboardVersionByIDUnauthorized() *GetDashboardVersionByIDUnauthorized {
	return &GetDashboardVersionByIDUnauthorized{}
}

/* GetDashboardVersionByIDUnauthorized describes a response with status code 401, with default header values.

UnauthorizedError is returned when the request is not authenticated.
*/
type GetDashboardVersionByIDUnauthorized struct {
	Payload *models.ErrorResponseBody
}

func (o *GetDashboardVersionByIDUnauthorized) Error() string {
	return fmt.Sprintf("[GET /dashboards/id/{DashboardID}/versions/{DashboardVersionID}][%d] getDashboardVersionByIdUnauthorized  %+v", 401, o.Payload)
}
func (o *GetDashboardVersionByIDUnauthorized) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *GetDashboardVersionByIDUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetDashboardVersionByIDForbidden creates a GetDashboardVersionByIDForbidden with default headers values
func NewGetDashboardVersionByIDForbidden() *GetDashboardVersionByIDForbidden {
	return &GetDashboardVersionByIDForbidden{}
}

/* GetDashboardVersionByIDForbidden describes a response with status code 403, with default header values.

ForbiddenError is returned if the user/token has insufficient permissions to access the requested resource.
*/
type GetDashboardVersionByIDForbidden struct {
	Payload *models.ErrorResponseBody
}

func (o *GetDashboardVersionByIDForbidden) Error() string {
	return fmt.Sprintf("[GET /dashboards/id/{DashboardID}/versions/{DashboardVersionID}][%d] getDashboardVersionByIdForbidden  %+v", 403, o.Payload)
}
func (o *GetDashboardVersionByIDForbidden) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *GetDashboardVersionByIDForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetDashboardVersionByIDNotFound creates a GetDashboardVersionByIDNotFound with default headers values
func NewGetDashboardVersionByIDNotFound() *GetDashboardVersionByIDNotFound {
	return &GetDashboardVersionByIDNotFound{}
}

/* GetDashboardVersionByIDNotFound describes a response with status code 404, with default header values.

NotFoundError is returned when the requested resource was not found.
*/
type GetDashboardVersionByIDNotFound struct {
	Payload *models.ErrorResponseBody
}

func (o *GetDashboardVersionByIDNotFound) Error() string {
	return fmt.Sprintf("[GET /dashboards/id/{DashboardID}/versions/{DashboardVersionID}][%d] getDashboardVersionByIdNotFound  %+v", 404, o.Payload)
}
func (o *GetDashboardVersionByIDNotFound) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *GetDashboardVersionByIDNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetDashboardVersionByIDInternalServerError creates a GetDashboardVersionByIDInternalServerError with default headers values
func NewGetDashboardVersionByIDInternalServerError() *GetDashboardVersionByIDInternalServerError {
	return &GetDashboardVersionByIDInternalServerError{}
}

/* GetDashboardVersionByIDInternalServerError describes a response with status code 500, with default header values.

InternalServerError is a general error indicating something went wrong internally.
*/
type GetDashboardVersionByIDInternalServerError struct {
	Payload *models.ErrorResponseBody
}

func (o *GetDashboardVersionByIDInternalServerError) Error() string {
	return fmt.Sprintf("[GET /dashboards/id/{DashboardID}/versions/{DashboardVersionID}][%d] getDashboardVersionByIdInternalServerError  %+v", 500, o.Payload)
}
func (o *GetDashboardVersionByIDInternalServerError) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *GetDashboardVersionByIDInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
