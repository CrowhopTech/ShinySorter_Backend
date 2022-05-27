// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/CrowhopTech/shinysorter/backend/pkg/swagger/server/models"
)

// CreateImageCreatedCode is the HTTP code returned for type CreateImageCreated
const CreateImageCreatedCode int = 201

/*CreateImageCreated Image was created successfully

swagger:response createImageCreated
*/
type CreateImageCreated struct {

	/*
	  In: Body
	*/
	Payload *models.Image `json:"body,omitempty"`
}

// NewCreateImageCreated creates CreateImageCreated with default headers values
func NewCreateImageCreated() *CreateImageCreated {

	return &CreateImageCreated{}
}

// WithPayload adds the payload to the create image created response
func (o *CreateImageCreated) WithPayload(payload *models.Image) *CreateImageCreated {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create image created response
func (o *CreateImageCreated) SetPayload(payload *models.Image) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreateImageCreated) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(201)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// CreateImageBadRequestCode is the HTTP code returned for type CreateImageBadRequest
const CreateImageBadRequestCode int = 400

/*CreateImageBadRequest Some part of the provided Image was invalid.

swagger:response createImageBadRequest
*/
type CreateImageBadRequest struct {
}

// NewCreateImageBadRequest creates CreateImageBadRequest with default headers values
func NewCreateImageBadRequest() *CreateImageBadRequest {

	return &CreateImageBadRequest{}
}

// WriteResponse to the client
func (o *CreateImageBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(400)
}

// CreateImageInternalServerErrorCode is the HTTP code returned for type CreateImageInternalServerError
const CreateImageInternalServerErrorCode int = 500

/*CreateImageInternalServerError Something else went wrong during the request

swagger:response createImageInternalServerError
*/
type CreateImageInternalServerError struct {

	/*
	  In: Body
	*/
	Payload string `json:"body,omitempty"`
}

// NewCreateImageInternalServerError creates CreateImageInternalServerError with default headers values
func NewCreateImageInternalServerError() *CreateImageInternalServerError {

	return &CreateImageInternalServerError{}
}

// WithPayload adds the payload to the create image internal server error response
func (o *CreateImageInternalServerError) WithPayload(payload string) *CreateImageInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create image internal server error response
func (o *CreateImageInternalServerError) SetPayload(payload string) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreateImageInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	payload := o.Payload
	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}
