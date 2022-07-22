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

// RouteGetTemplateReader is a Reader for the RouteGetTemplate structure.
type RouteGetTemplateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *RouteGetTemplateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewRouteGetTemplateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewRouteGetTemplateNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewRouteGetTemplateOK creates a RouteGetTemplateOK with default headers values
func NewRouteGetTemplateOK() *RouteGetTemplateOK {
	return &RouteGetTemplateOK{}
}

/* RouteGetTemplateOK describes a response with status code 200, with default header values.

MessageTemplate
*/
type RouteGetTemplateOK struct {
	Payload *models.MessageTemplate
}

func (o *RouteGetTemplateOK) Error() string {
	return fmt.Sprintf("[GET /api/v1/provisioning/templates/{name}][%d] routeGetTemplateOK  %+v", 200, o.Payload)
}
func (o *RouteGetTemplateOK) GetPayload() *models.MessageTemplate {
	return o.Payload
}

func (o *RouteGetTemplateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.MessageTemplate)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRouteGetTemplateNotFound creates a RouteGetTemplateNotFound with default headers values
func NewRouteGetTemplateNotFound() *RouteGetTemplateNotFound {
	return &RouteGetTemplateNotFound{}
}

/* RouteGetTemplateNotFound describes a response with status code 404, with default header values.

Not found.
*/
type RouteGetTemplateNotFound struct {
}

func (o *RouteGetTemplateNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v1/provisioning/templates/{name}][%d] routeGetTemplateNotFound ", 404)
}

func (o *RouteGetTemplateNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
