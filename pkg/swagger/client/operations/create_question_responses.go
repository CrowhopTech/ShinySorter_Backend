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

// CreateQuestionReader is a Reader for the CreateQuestion structure.
type CreateQuestionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateQuestionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewCreateQuestionCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewCreateQuestionBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewCreateQuestionInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewCreateQuestionCreated creates a CreateQuestionCreated with default headers values
func NewCreateQuestionCreated() *CreateQuestionCreated {
	return &CreateQuestionCreated{}
}

/* CreateQuestionCreated describes a response with status code 201, with default header values.

Question was created successfully
*/
type CreateQuestionCreated struct {
	Payload *models.QuestionEntry
}

func (o *CreateQuestionCreated) Error() string {
	return fmt.Sprintf("[POST /questions][%d] createQuestionCreated  %+v", 201, o.Payload)
}
func (o *CreateQuestionCreated) GetPayload() *models.QuestionEntry {
	return o.Payload
}

func (o *CreateQuestionCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.QuestionEntry)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateQuestionBadRequest creates a CreateQuestionBadRequest with default headers values
func NewCreateQuestionBadRequest() *CreateQuestionBadRequest {
	return &CreateQuestionBadRequest{}
}

/* CreateQuestionBadRequest describes a response with status code 400, with default header values.

Some part of the request was invalid. More information will be included in the error string
*/
type CreateQuestionBadRequest struct {
	Payload string
}

func (o *CreateQuestionBadRequest) Error() string {
	return fmt.Sprintf("[POST /questions][%d] createQuestionBadRequest  %+v", 400, o.Payload)
}
func (o *CreateQuestionBadRequest) GetPayload() string {
	return o.Payload
}

func (o *CreateQuestionBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateQuestionInternalServerError creates a CreateQuestionInternalServerError with default headers values
func NewCreateQuestionInternalServerError() *CreateQuestionInternalServerError {
	return &CreateQuestionInternalServerError{}
}

/* CreateQuestionInternalServerError describes a response with status code 500, with default header values.

Something else went wrong during the request
*/
type CreateQuestionInternalServerError struct {
	Payload string
}

func (o *CreateQuestionInternalServerError) Error() string {
	return fmt.Sprintf("[POST /questions][%d] createQuestionInternalServerError  %+v", 500, o.Payload)
}
func (o *CreateQuestionInternalServerError) GetPayload() string {
	return o.Payload
}

func (o *CreateQuestionInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
