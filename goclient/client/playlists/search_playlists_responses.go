// Code generated by go-swagger; DO NOT EDIT.

package playlists

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/grafana/grafana-api-golang-client/goclient/models"
)

// SearchPlaylistsReader is a Reader for the SearchPlaylists structure.
type SearchPlaylistsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SearchPlaylistsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewSearchPlaylistsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 500:
		result := NewSearchPlaylistsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewSearchPlaylistsOK creates a SearchPlaylistsOK with default headers values
func NewSearchPlaylistsOK() *SearchPlaylistsOK {
	return &SearchPlaylistsOK{}
}

/* SearchPlaylistsOK describes a response with status code 200, with default header values.

(empty)
*/
type SearchPlaylistsOK struct {
	Payload models.Playlists
}

func (o *SearchPlaylistsOK) Error() string {
	return fmt.Sprintf("[GET /playlists][%d] searchPlaylistsOK  %+v", 200, o.Payload)
}
func (o *SearchPlaylistsOK) GetPayload() models.Playlists {
	return o.Payload
}

func (o *SearchPlaylistsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSearchPlaylistsInternalServerError creates a SearchPlaylistsInternalServerError with default headers values
func NewSearchPlaylistsInternalServerError() *SearchPlaylistsInternalServerError {
	return &SearchPlaylistsInternalServerError{}
}

/* SearchPlaylistsInternalServerError describes a response with status code 500, with default header values.

InternalServerError is a general error indicating something went wrong internally.
*/
type SearchPlaylistsInternalServerError struct {
	Payload *models.ErrorResponseBody
}

func (o *SearchPlaylistsInternalServerError) Error() string {
	return fmt.Sprintf("[GET /playlists][%d] searchPlaylistsInternalServerError  %+v", 500, o.Payload)
}
func (o *SearchPlaylistsInternalServerError) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *SearchPlaylistsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
