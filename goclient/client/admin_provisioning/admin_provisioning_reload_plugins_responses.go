// Code generated by go-swagger; DO NOT EDIT.

package admin_provisioning

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/grafana/grafana-api-golang-client/goclient/models"
)

// AdminProvisioningReloadPluginsReader is a Reader for the AdminProvisioningReloadPlugins structure.
type AdminProvisioningReloadPluginsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AdminProvisioningReloadPluginsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewAdminProvisioningReloadPluginsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewAdminProvisioningReloadPluginsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewAdminProvisioningReloadPluginsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewAdminProvisioningReloadPluginsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewAdminProvisioningReloadPluginsOK creates a AdminProvisioningReloadPluginsOK with default headers values
func NewAdminProvisioningReloadPluginsOK() *AdminProvisioningReloadPluginsOK {
	return &AdminProvisioningReloadPluginsOK{}
}

/* AdminProvisioningReloadPluginsOK describes a response with status code 200, with default header values.

An OKResponse is returned if the request was successful.
*/
type AdminProvisioningReloadPluginsOK struct {
	Payload *models.SuccessResponseBody
}

func (o *AdminProvisioningReloadPluginsOK) Error() string {
	return fmt.Sprintf("[POST /admin/provisioning/plugins/reload][%d] adminProvisioningReloadPluginsOK  %+v", 200, o.Payload)
}
func (o *AdminProvisioningReloadPluginsOK) GetPayload() *models.SuccessResponseBody {
	return o.Payload
}

func (o *AdminProvisioningReloadPluginsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SuccessResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAdminProvisioningReloadPluginsUnauthorized creates a AdminProvisioningReloadPluginsUnauthorized with default headers values
func NewAdminProvisioningReloadPluginsUnauthorized() *AdminProvisioningReloadPluginsUnauthorized {
	return &AdminProvisioningReloadPluginsUnauthorized{}
}

/* AdminProvisioningReloadPluginsUnauthorized describes a response with status code 401, with default header values.

UnauthorizedError is returned when the request is not authenticated.
*/
type AdminProvisioningReloadPluginsUnauthorized struct {
	Payload *models.ErrorResponseBody
}

func (o *AdminProvisioningReloadPluginsUnauthorized) Error() string {
	return fmt.Sprintf("[POST /admin/provisioning/plugins/reload][%d] adminProvisioningReloadPluginsUnauthorized  %+v", 401, o.Payload)
}
func (o *AdminProvisioningReloadPluginsUnauthorized) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *AdminProvisioningReloadPluginsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAdminProvisioningReloadPluginsForbidden creates a AdminProvisioningReloadPluginsForbidden with default headers values
func NewAdminProvisioningReloadPluginsForbidden() *AdminProvisioningReloadPluginsForbidden {
	return &AdminProvisioningReloadPluginsForbidden{}
}

/* AdminProvisioningReloadPluginsForbidden describes a response with status code 403, with default header values.

ForbiddenError is returned if the user/token has insufficient permissions to access the requested resource.
*/
type AdminProvisioningReloadPluginsForbidden struct {
	Payload *models.ErrorResponseBody
}

func (o *AdminProvisioningReloadPluginsForbidden) Error() string {
	return fmt.Sprintf("[POST /admin/provisioning/plugins/reload][%d] adminProvisioningReloadPluginsForbidden  %+v", 403, o.Payload)
}
func (o *AdminProvisioningReloadPluginsForbidden) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *AdminProvisioningReloadPluginsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAdminProvisioningReloadPluginsInternalServerError creates a AdminProvisioningReloadPluginsInternalServerError with default headers values
func NewAdminProvisioningReloadPluginsInternalServerError() *AdminProvisioningReloadPluginsInternalServerError {
	return &AdminProvisioningReloadPluginsInternalServerError{}
}

/* AdminProvisioningReloadPluginsInternalServerError describes a response with status code 500, with default header values.

InternalServerError is a general error indicating something went wrong internally.
*/
type AdminProvisioningReloadPluginsInternalServerError struct {
	Payload *models.ErrorResponseBody
}

func (o *AdminProvisioningReloadPluginsInternalServerError) Error() string {
	return fmt.Sprintf("[POST /admin/provisioning/plugins/reload][%d] adminProvisioningReloadPluginsInternalServerError  %+v", 500, o.Payload)
}
func (o *AdminProvisioningReloadPluginsInternalServerError) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *AdminProvisioningReloadPluginsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
