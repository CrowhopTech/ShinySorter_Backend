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

// CheckHealthReader is a Reader for the CheckHealth structure.
type CheckHealthReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CheckHealthReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCheckHealthOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 503:
		result := NewCheckHealthServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewCheckHealthOK creates a CheckHealthOK with default headers values
func NewCheckHealthOK() *CheckHealthOK {
	return &CheckHealthOK{}
}

/* CheckHealthOK describes a response with status code 200, with default header values.

OK message
*/
type CheckHealthOK struct {
	Payload string
}

func (o *CheckHealthOK) Error() string {
	return fmt.Sprintf("[GET /healthz][%d] checkHealthOK  %+v", 200, o.Payload)
}
func (o *CheckHealthOK) GetPayload() string {
	return o.Payload
}

func (o *CheckHealthOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCheckHealthServiceUnavailable creates a CheckHealthServiceUnavailable with default headers values
func NewCheckHealthServiceUnavailable() *CheckHealthServiceUnavailable {
	return &CheckHealthServiceUnavailable{}
}

/* CheckHealthServiceUnavailable describes a response with status code 503, with default header values.

Server still starting
*/
type CheckHealthServiceUnavailable struct {
	Payload string
}

func (o *CheckHealthServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /healthz][%d] checkHealthServiceUnavailable  %+v", 503, o.Payload)
}
func (o *CheckHealthServiceUnavailable) GetPayload() string {
	return o.Payload
}

func (o *CheckHealthServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
