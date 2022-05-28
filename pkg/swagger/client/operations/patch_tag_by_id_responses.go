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

// PatchTagByIDReader is a Reader for the PatchTagByID structure.
type PatchTagByIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PatchTagByIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPatchTagByIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPatchTagByIDBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPatchTagByIDInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPatchTagByIDOK creates a PatchTagByIDOK with default headers values
func NewPatchTagByIDOK() *PatchTagByIDOK {
	return &PatchTagByIDOK{}
}

/* PatchTagByIDOK describes a response with status code 200, with default header values.

Tag was modified successfully
*/
type PatchTagByIDOK struct {
	Payload *models.Tag
}

func (o *PatchTagByIDOK) Error() string {
	return fmt.Sprintf("[PATCH /tags/{id}][%d] patchTagByIdOK  %+v", 200, o.Payload)
}
func (o *PatchTagByIDOK) GetPayload() *models.Tag {
	return o.Payload
}

func (o *PatchTagByIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Tag)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchTagByIDBadRequest creates a PatchTagByIDBadRequest with default headers values
func NewPatchTagByIDBadRequest() *PatchTagByIDBadRequest {
	return &PatchTagByIDBadRequest{}
}

/* PatchTagByIDBadRequest describes a response with status code 400, with default header values.

Some part of the request was invalid. More information will be included in the error string
*/
type PatchTagByIDBadRequest struct {
	Payload string
}

func (o *PatchTagByIDBadRequest) Error() string {
	return fmt.Sprintf("[PATCH /tags/{id}][%d] patchTagByIdBadRequest  %+v", 400, o.Payload)
}
func (o *PatchTagByIDBadRequest) GetPayload() string {
	return o.Payload
}

func (o *PatchTagByIDBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchTagByIDInternalServerError creates a PatchTagByIDInternalServerError with default headers values
func NewPatchTagByIDInternalServerError() *PatchTagByIDInternalServerError {
	return &PatchTagByIDInternalServerError{}
}

/* PatchTagByIDInternalServerError describes a response with status code 500, with default header values.

Something else went wrong during the request
*/
type PatchTagByIDInternalServerError struct {
	Payload string
}

func (o *PatchTagByIDInternalServerError) Error() string {
	return fmt.Sprintf("[PATCH /tags/{id}][%d] patchTagByIdInternalServerError  %+v", 500, o.Payload)
}
func (o *PatchTagByIDInternalServerError) GetPayload() string {
	return o.Payload
}

func (o *PatchTagByIDInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}