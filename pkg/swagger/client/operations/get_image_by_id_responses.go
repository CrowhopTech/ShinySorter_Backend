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

// GetImageByIDReader is a Reader for the GetImageByID structure.
type GetImageByIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetImageByIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetImageByIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewGetImageByIDNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetImageByIDInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetImageByIDOK creates a GetImageByIDOK with default headers values
func NewGetImageByIDOK() *GetImageByIDOK {
	return &GetImageByIDOK{}
}

/* GetImageByIDOK describes a response with status code 200, with default header values.

Returns the found image.
*/
type GetImageByIDOK struct {
	Payload *models.Image
}

func (o *GetImageByIDOK) Error() string {
	return fmt.Sprintf("[GET /images/{id}][%d] getImageByIdOK  %+v", 200, o.Payload)
}
func (o *GetImageByIDOK) GetPayload() *models.Image {
	return o.Payload
}

func (o *GetImageByIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Image)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetImageByIDNotFound creates a GetImageByIDNotFound with default headers values
func NewGetImageByIDNotFound() *GetImageByIDNotFound {
	return &GetImageByIDNotFound{}
}

/* GetImageByIDNotFound describes a response with status code 404, with default header values.

The given image was not found.
*/
type GetImageByIDNotFound struct {
}

func (o *GetImageByIDNotFound) Error() string {
	return fmt.Sprintf("[GET /images/{id}][%d] getImageByIdNotFound ", 404)
}

func (o *GetImageByIDNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetImageByIDInternalServerError creates a GetImageByIDInternalServerError with default headers values
func NewGetImageByIDInternalServerError() *GetImageByIDInternalServerError {
	return &GetImageByIDInternalServerError{}
}

/* GetImageByIDInternalServerError describes a response with status code 500, with default header values.

Something else went wrong during the request
*/
type GetImageByIDInternalServerError struct {
	Payload string
}

func (o *GetImageByIDInternalServerError) Error() string {
	return fmt.Sprintf("[GET /images/{id}][%d] getImageByIdInternalServerError  %+v", 500, o.Payload)
}
func (o *GetImageByIDInternalServerError) GetPayload() string {
	return o.Payload
}

func (o *GetImageByIDInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
