// Code generated by go-swagger; DO NOT EDIT.

package tags

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/CrowhopTech/shinysorter/backend/pkg/swagger/models"
)

// CreateTagReader is a Reader for the CreateTag structure.
type CreateTagReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateTagReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewCreateTagCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewCreateTagBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewCreateTagInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewCreateTagCreated creates a CreateTagCreated with default headers values
func NewCreateTagCreated() *CreateTagCreated {
	return &CreateTagCreated{}
}

/* CreateTagCreated describes a response with status code 201, with default header values.

Tag was created successfully
*/
type CreateTagCreated struct {
	Payload *models.TagEntry
}

func (o *CreateTagCreated) Error() string {
	return fmt.Sprintf("[POST /tags][%d] createTagCreated  %+v", 201, o.Payload)
}
func (o *CreateTagCreated) GetPayload() *models.TagEntry {
	return o.Payload
}

func (o *CreateTagCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.TagEntry)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateTagBadRequest creates a CreateTagBadRequest with default headers values
func NewCreateTagBadRequest() *CreateTagBadRequest {
	return &CreateTagBadRequest{}
}

/* CreateTagBadRequest describes a response with status code 400, with default header values.

Some part of the request was invalid. More information will be included in the error string
*/
type CreateTagBadRequest struct {
	Payload string
}

func (o *CreateTagBadRequest) Error() string {
	return fmt.Sprintf("[POST /tags][%d] createTagBadRequest  %+v", 400, o.Payload)
}
func (o *CreateTagBadRequest) GetPayload() string {
	return o.Payload
}

func (o *CreateTagBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateTagInternalServerError creates a CreateTagInternalServerError with default headers values
func NewCreateTagInternalServerError() *CreateTagInternalServerError {
	return &CreateTagInternalServerError{}
}

/* CreateTagInternalServerError describes a response with status code 500, with default header values.

Something else went wrong during the request
*/
type CreateTagInternalServerError struct {
	Payload string
}

func (o *CreateTagInternalServerError) Error() string {
	return fmt.Sprintf("[POST /tags][%d] createTagInternalServerError  %+v", 500, o.Payload)
}
func (o *CreateTagInternalServerError) GetPayload() string {
	return o.Payload
}

func (o *CreateTagInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
