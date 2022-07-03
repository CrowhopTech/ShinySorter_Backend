// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// SetFileContentReader is a Reader for the SetFileContent structure.
type SetFileContentReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SetFileContentReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewSetFileContentOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewSetFileContentBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewSetFileContentNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewSetFileContentInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewSetFileContentOK creates a SetFileContentOK with default headers values
func NewSetFileContentOK() *SetFileContentOK {
	return &SetFileContentOK{}
}

/* SetFileContentOK describes a response with status code 200, with default header values.

The file contents were modified successfully
*/
type SetFileContentOK struct {
}

func (o *SetFileContentOK) Error() string {
	return fmt.Sprintf("[PATCH /files/contents/{id}][%d] setFileContentOK ", 200)
}

func (o *SetFileContentOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewSetFileContentBadRequest creates a SetFileContentBadRequest with default headers values
func NewSetFileContentBadRequest() *SetFileContentBadRequest {
	return &SetFileContentBadRequest{}
}

/* SetFileContentBadRequest describes a response with status code 400, with default header values.

Some part of the request was invalid. More information will be included in the error string
*/
type SetFileContentBadRequest struct {
	Payload string
}

func (o *SetFileContentBadRequest) Error() string {
	return fmt.Sprintf("[PATCH /files/contents/{id}][%d] setFileContentBadRequest  %+v", 400, o.Payload)
}
func (o *SetFileContentBadRequest) GetPayload() string {
	return o.Payload
}

func (o *SetFileContentBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSetFileContentNotFound creates a SetFileContentNotFound with default headers values
func NewSetFileContentNotFound() *SetFileContentNotFound {
	return &SetFileContentNotFound{}
}

/* SetFileContentNotFound describes a response with status code 404, with default header values.

The given file was not found.
*/
type SetFileContentNotFound struct {
}

func (o *SetFileContentNotFound) Error() string {
	return fmt.Sprintf("[PATCH /files/contents/{id}][%d] setFileContentNotFound ", 404)
}

func (o *SetFileContentNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewSetFileContentInternalServerError creates a SetFileContentInternalServerError with default headers values
func NewSetFileContentInternalServerError() *SetFileContentInternalServerError {
	return &SetFileContentInternalServerError{}
}

/* SetFileContentInternalServerError describes a response with status code 500, with default header values.

Something else went wrong during the request
*/
type SetFileContentInternalServerError struct {
	Payload string
}

func (o *SetFileContentInternalServerError) Error() string {
	return fmt.Sprintf("[PATCH /files/contents/{id}][%d] setFileContentInternalServerError  %+v", 500, o.Payload)
}
func (o *SetFileContentInternalServerError) GetPayload() string {
	return o.Payload
}

func (o *SetFileContentInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
