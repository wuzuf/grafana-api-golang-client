// Code generated by go-swagger; DO NOT EDIT.

package legacy_alerts_notification_channels

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/grafana/grafana-api-golang-client/goclient/models"
)

// GetAlertNotificationChannelByIDReader is a Reader for the GetAlertNotificationChannelByID structure.
type GetAlertNotificationChannelByIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetAlertNotificationChannelByIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetAlertNotificationChannelByIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewGetAlertNotificationChannelByIDUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetAlertNotificationChannelByIDForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetAlertNotificationChannelByIDNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetAlertNotificationChannelByIDInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetAlertNotificationChannelByIDOK creates a GetAlertNotificationChannelByIDOK with default headers values
func NewGetAlertNotificationChannelByIDOK() *GetAlertNotificationChannelByIDOK {
	return &GetAlertNotificationChannelByIDOK{}
}

/* GetAlertNotificationChannelByIDOK describes a response with status code 200, with default header values.

(empty)
*/
type GetAlertNotificationChannelByIDOK struct {
	Payload *models.AlertNotification
}

func (o *GetAlertNotificationChannelByIDOK) Error() string {
	return fmt.Sprintf("[GET /alert-notifications/{notification_channel_id}][%d] getAlertNotificationChannelByIdOK  %+v", 200, o.Payload)
}
func (o *GetAlertNotificationChannelByIDOK) GetPayload() *models.AlertNotification {
	return o.Payload
}

func (o *GetAlertNotificationChannelByIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.AlertNotification)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAlertNotificationChannelByIDUnauthorized creates a GetAlertNotificationChannelByIDUnauthorized with default headers values
func NewGetAlertNotificationChannelByIDUnauthorized() *GetAlertNotificationChannelByIDUnauthorized {
	return &GetAlertNotificationChannelByIDUnauthorized{}
}

/* GetAlertNotificationChannelByIDUnauthorized describes a response with status code 401, with default header values.

UnauthorizedError is returned when the request is not authenticated.
*/
type GetAlertNotificationChannelByIDUnauthorized struct {
	Payload *models.ErrorResponseBody
}

func (o *GetAlertNotificationChannelByIDUnauthorized) Error() string {
	return fmt.Sprintf("[GET /alert-notifications/{notification_channel_id}][%d] getAlertNotificationChannelByIdUnauthorized  %+v", 401, o.Payload)
}
func (o *GetAlertNotificationChannelByIDUnauthorized) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *GetAlertNotificationChannelByIDUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAlertNotificationChannelByIDForbidden creates a GetAlertNotificationChannelByIDForbidden with default headers values
func NewGetAlertNotificationChannelByIDForbidden() *GetAlertNotificationChannelByIDForbidden {
	return &GetAlertNotificationChannelByIDForbidden{}
}

/* GetAlertNotificationChannelByIDForbidden describes a response with status code 403, with default header values.

ForbiddenError is returned if the user/token has insufficient permissions to access the requested resource.
*/
type GetAlertNotificationChannelByIDForbidden struct {
	Payload *models.ErrorResponseBody
}

func (o *GetAlertNotificationChannelByIDForbidden) Error() string {
	return fmt.Sprintf("[GET /alert-notifications/{notification_channel_id}][%d] getAlertNotificationChannelByIdForbidden  %+v", 403, o.Payload)
}
func (o *GetAlertNotificationChannelByIDForbidden) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *GetAlertNotificationChannelByIDForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAlertNotificationChannelByIDNotFound creates a GetAlertNotificationChannelByIDNotFound with default headers values
func NewGetAlertNotificationChannelByIDNotFound() *GetAlertNotificationChannelByIDNotFound {
	return &GetAlertNotificationChannelByIDNotFound{}
}

/* GetAlertNotificationChannelByIDNotFound describes a response with status code 404, with default header values.

NotFoundError is returned when the requested resource was not found.
*/
type GetAlertNotificationChannelByIDNotFound struct {
	Payload *models.ErrorResponseBody
}

func (o *GetAlertNotificationChannelByIDNotFound) Error() string {
	return fmt.Sprintf("[GET /alert-notifications/{notification_channel_id}][%d] getAlertNotificationChannelByIdNotFound  %+v", 404, o.Payload)
}
func (o *GetAlertNotificationChannelByIDNotFound) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *GetAlertNotificationChannelByIDNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAlertNotificationChannelByIDInternalServerError creates a GetAlertNotificationChannelByIDInternalServerError with default headers values
func NewGetAlertNotificationChannelByIDInternalServerError() *GetAlertNotificationChannelByIDInternalServerError {
	return &GetAlertNotificationChannelByIDInternalServerError{}
}

/* GetAlertNotificationChannelByIDInternalServerError describes a response with status code 500, with default header values.

InternalServerError is a general error indicating something went wrong internally.
*/
type GetAlertNotificationChannelByIDInternalServerError struct {
	Payload *models.ErrorResponseBody
}

func (o *GetAlertNotificationChannelByIDInternalServerError) Error() string {
	return fmt.Sprintf("[GET /alert-notifications/{notification_channel_id}][%d] getAlertNotificationChannelByIdInternalServerError  %+v", 500, o.Payload)
}
func (o *GetAlertNotificationChannelByIDInternalServerError) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *GetAlertNotificationChannelByIDInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
