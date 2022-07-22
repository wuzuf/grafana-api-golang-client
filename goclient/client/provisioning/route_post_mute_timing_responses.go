// Code generated by go-swagger; DO NOT EDIT.

package provisioning

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/grafana/grafana-api-golang-client/goclient/models"
)

// RoutePostMuteTimingReader is a Reader for the RoutePostMuteTiming structure.
type RoutePostMuteTimingReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *RoutePostMuteTimingReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewRoutePostMuteTimingCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewRoutePostMuteTimingBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewRoutePostMuteTimingCreated creates a RoutePostMuteTimingCreated with default headers values
func NewRoutePostMuteTimingCreated() *RoutePostMuteTimingCreated {
	return &RoutePostMuteTimingCreated{}
}

/* RoutePostMuteTimingCreated describes a response with status code 201, with default header values.

MuteTimeInterval
*/
type RoutePostMuteTimingCreated struct {
	Payload *models.MuteTimeInterval
}

func (o *RoutePostMuteTimingCreated) Error() string {
	return fmt.Sprintf("[POST /api/v1/provisioning/mute-timings][%d] routePostMuteTimingCreated  %+v", 201, o.Payload)
}
func (o *RoutePostMuteTimingCreated) GetPayload() *models.MuteTimeInterval {
	return o.Payload
}

func (o *RoutePostMuteTimingCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.MuteTimeInterval)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRoutePostMuteTimingBadRequest creates a RoutePostMuteTimingBadRequest with default headers values
func NewRoutePostMuteTimingBadRequest() *RoutePostMuteTimingBadRequest {
	return &RoutePostMuteTimingBadRequest{}
}

/* RoutePostMuteTimingBadRequest describes a response with status code 400, with default header values.

ValidationError
*/
type RoutePostMuteTimingBadRequest struct {
	Payload *models.ValidationError
}

func (o *RoutePostMuteTimingBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v1/provisioning/mute-timings][%d] routePostMuteTimingBadRequest  %+v", 400, o.Payload)
}
func (o *RoutePostMuteTimingBadRequest) GetPayload() *models.ValidationError {
	return o.Payload
}

func (o *RoutePostMuteTimingBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ValidationError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
