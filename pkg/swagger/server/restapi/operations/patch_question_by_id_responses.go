// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/CrowhopTech/shinysorter/backend/pkg/swagger/server/models"
)

// PatchQuestionByIDOKCode is the HTTP code returned for type PatchQuestionByIDOK
const PatchQuestionByIDOKCode int = 200

/*PatchQuestionByIDOK Question was modified successfully

swagger:response patchQuestionByIdOK
*/
type PatchQuestionByIDOK struct {

	/*
	  In: Body
	*/
	Payload *models.Question `json:"body,omitempty"`
}

// NewPatchQuestionByIDOK creates PatchQuestionByIDOK with default headers values
func NewPatchQuestionByIDOK() *PatchQuestionByIDOK {

	return &PatchQuestionByIDOK{}
}

// WithPayload adds the payload to the patch question by Id o k response
func (o *PatchQuestionByIDOK) WithPayload(payload *models.Question) *PatchQuestionByIDOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the patch question by Id o k response
func (o *PatchQuestionByIDOK) SetPayload(payload *models.Question) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PatchQuestionByIDOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// PatchQuestionByIDBadRequestCode is the HTTP code returned for type PatchQuestionByIDBadRequest
const PatchQuestionByIDBadRequestCode int = 400

/*PatchQuestionByIDBadRequest Some part of the request was invalid. More information will be included in the error string

swagger:response patchQuestionByIdBadRequest
*/
type PatchQuestionByIDBadRequest struct {

	/*
	  In: Body
	*/
	Payload string `json:"body,omitempty"`
}

// NewPatchQuestionByIDBadRequest creates PatchQuestionByIDBadRequest with default headers values
func NewPatchQuestionByIDBadRequest() *PatchQuestionByIDBadRequest {

	return &PatchQuestionByIDBadRequest{}
}

// WithPayload adds the payload to the patch question by Id bad request response
func (o *PatchQuestionByIDBadRequest) WithPayload(payload string) *PatchQuestionByIDBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the patch question by Id bad request response
func (o *PatchQuestionByIDBadRequest) SetPayload(payload string) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PatchQuestionByIDBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	payload := o.Payload
	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

// PatchQuestionByIDInternalServerErrorCode is the HTTP code returned for type PatchQuestionByIDInternalServerError
const PatchQuestionByIDInternalServerErrorCode int = 500

/*PatchQuestionByIDInternalServerError Something else went wrong during the request

swagger:response patchQuestionByIdInternalServerError
*/
type PatchQuestionByIDInternalServerError struct {

	/*
	  In: Body
	*/
	Payload string `json:"body,omitempty"`
}

// NewPatchQuestionByIDInternalServerError creates PatchQuestionByIDInternalServerError with default headers values
func NewPatchQuestionByIDInternalServerError() *PatchQuestionByIDInternalServerError {

	return &PatchQuestionByIDInternalServerError{}
}

// WithPayload adds the payload to the patch question by Id internal server error response
func (o *PatchQuestionByIDInternalServerError) WithPayload(payload string) *PatchQuestionByIDInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the patch question by Id internal server error response
func (o *PatchQuestionByIDInternalServerError) SetPayload(payload string) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PatchQuestionByIDInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	payload := o.Payload
	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}
