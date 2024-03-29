// Code generated by go-swagger; DO NOT EDIT.

package questions

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/CrowhopTech/shinysorter/backend/pkg/swagger/server/models"
)

// CreateQuestionCreatedCode is the HTTP code returned for type CreateQuestionCreated
const CreateQuestionCreatedCode int = 201

/*CreateQuestionCreated Question was created successfully

swagger:response createQuestionCreated
*/
type CreateQuestionCreated struct {

	/*
	  In: Body
	*/
	Payload *models.QuestionEntry `json:"body,omitempty"`
}

// NewCreateQuestionCreated creates CreateQuestionCreated with default headers values
func NewCreateQuestionCreated() *CreateQuestionCreated {

	return &CreateQuestionCreated{}
}

// WithPayload adds the payload to the create question created response
func (o *CreateQuestionCreated) WithPayload(payload *models.QuestionEntry) *CreateQuestionCreated {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create question created response
func (o *CreateQuestionCreated) SetPayload(payload *models.QuestionEntry) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreateQuestionCreated) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(201)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// CreateQuestionBadRequestCode is the HTTP code returned for type CreateQuestionBadRequest
const CreateQuestionBadRequestCode int = 400

/*CreateQuestionBadRequest Some part of the request was invalid. More information will be included in the error string

swagger:response createQuestionBadRequest
*/
type CreateQuestionBadRequest struct {

	/*
	  In: Body
	*/
	Payload string `json:"body,omitempty"`
}

// NewCreateQuestionBadRequest creates CreateQuestionBadRequest with default headers values
func NewCreateQuestionBadRequest() *CreateQuestionBadRequest {

	return &CreateQuestionBadRequest{}
}

// WithPayload adds the payload to the create question bad request response
func (o *CreateQuestionBadRequest) WithPayload(payload string) *CreateQuestionBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create question bad request response
func (o *CreateQuestionBadRequest) SetPayload(payload string) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreateQuestionBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	payload := o.Payload
	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

// CreateQuestionInternalServerErrorCode is the HTTP code returned for type CreateQuestionInternalServerError
const CreateQuestionInternalServerErrorCode int = 500

/*CreateQuestionInternalServerError Something else went wrong during the request

swagger:response createQuestionInternalServerError
*/
type CreateQuestionInternalServerError struct {

	/*
	  In: Body
	*/
	Payload string `json:"body,omitempty"`
}

// NewCreateQuestionInternalServerError creates CreateQuestionInternalServerError with default headers values
func NewCreateQuestionInternalServerError() *CreateQuestionInternalServerError {

	return &CreateQuestionInternalServerError{}
}

// WithPayload adds the payload to the create question internal server error response
func (o *CreateQuestionInternalServerError) WithPayload(payload string) *CreateQuestionInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create question internal server error response
func (o *CreateQuestionInternalServerError) SetPayload(payload string) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreateQuestionInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	payload := o.Payload
	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}
