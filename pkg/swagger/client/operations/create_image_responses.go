// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/CrowhopTech/shinysorter/backend/pkg/swagger/models"
)

// CreateImageReader is a Reader for the CreateImage structure.
type CreateImageReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateImageReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewCreateImageCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewCreateImageBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewCreateImageInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewCreateImageCreated creates a CreateImageCreated with default headers values
func NewCreateImageCreated() *CreateImageCreated {
	return &CreateImageCreated{}
}

/* CreateImageCreated describes a response with status code 201, with default header values.

Image was created successfully
*/
type CreateImageCreated struct {
	Payload *models.Image
}

func (o *CreateImageCreated) Error() string {
	return fmt.Sprintf("[POST /images][%d] createImageCreated  %+v", 201, o.Payload)
}
func (o *CreateImageCreated) GetPayload() *models.Image {
	return o.Payload
}

func (o *CreateImageCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Image)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateImageBadRequest creates a CreateImageBadRequest with default headers values
func NewCreateImageBadRequest() *CreateImageBadRequest {
	return &CreateImageBadRequest{}
}

/* CreateImageBadRequest describes a response with status code 400, with default header values.

Some part of the provided Image was invalid.
*/
type CreateImageBadRequest struct {
}

func (o *CreateImageBadRequest) Error() string {
	return fmt.Sprintf("[POST /images][%d] createImageBadRequest ", 400)
}

func (o *CreateImageBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCreateImageInternalServerError creates a CreateImageInternalServerError with default headers values
func NewCreateImageInternalServerError() *CreateImageInternalServerError {
	return &CreateImageInternalServerError{}
}

/* CreateImageInternalServerError describes a response with status code 500, with default header values.

Something else went wrong during the request
*/
type CreateImageInternalServerError struct {
	Payload string
}

func (o *CreateImageInternalServerError) Error() string {
	return fmt.Sprintf("[POST /images][%d] createImageInternalServerError  %+v", 500, o.Payload)
}
func (o *CreateImageInternalServerError) GetPayload() string {
	return o.Payload
}

func (o *CreateImageInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}