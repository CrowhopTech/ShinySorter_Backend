// Code generated by go-swagger; DO NOT EDIT.

package tags

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/CrowhopTech/shinysorter/backend/pkg/swagger/server/models"
)

// PatchTagByIDOKCode is the HTTP code returned for type PatchTagByIDOK
const PatchTagByIDOKCode int = 200

/*PatchTagByIDOK Tag was modified successfully

swagger:response patchTagByIdOK
*/
type PatchTagByIDOK struct {

	/*
	  In: Body
	*/
	Payload *models.TagEntry `json:"body,omitempty"`
}

// NewPatchTagByIDOK creates PatchTagByIDOK with default headers values
func NewPatchTagByIDOK() *PatchTagByIDOK {

	return &PatchTagByIDOK{}
}

// WithPayload adds the payload to the patch tag by Id o k response
func (o *PatchTagByIDOK) WithPayload(payload *models.TagEntry) *PatchTagByIDOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the patch tag by Id o k response
func (o *PatchTagByIDOK) SetPayload(payload *models.TagEntry) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PatchTagByIDOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// PatchTagByIDBadRequestCode is the HTTP code returned for type PatchTagByIDBadRequest
const PatchTagByIDBadRequestCode int = 400

/*PatchTagByIDBadRequest Some part of the request was invalid. More information will be included in the error string

swagger:response patchTagByIdBadRequest
*/
type PatchTagByIDBadRequest struct {

	/*
	  In: Body
	*/
	Payload string `json:"body,omitempty"`
}

// NewPatchTagByIDBadRequest creates PatchTagByIDBadRequest with default headers values
func NewPatchTagByIDBadRequest() *PatchTagByIDBadRequest {

	return &PatchTagByIDBadRequest{}
}

// WithPayload adds the payload to the patch tag by Id bad request response
func (o *PatchTagByIDBadRequest) WithPayload(payload string) *PatchTagByIDBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the patch tag by Id bad request response
func (o *PatchTagByIDBadRequest) SetPayload(payload string) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PatchTagByIDBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	payload := o.Payload
	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

// PatchTagByIDInternalServerErrorCode is the HTTP code returned for type PatchTagByIDInternalServerError
const PatchTagByIDInternalServerErrorCode int = 500

/*PatchTagByIDInternalServerError Something else went wrong during the request

swagger:response patchTagByIdInternalServerError
*/
type PatchTagByIDInternalServerError struct {

	/*
	  In: Body
	*/
	Payload string `json:"body,omitempty"`
}

// NewPatchTagByIDInternalServerError creates PatchTagByIDInternalServerError with default headers values
func NewPatchTagByIDInternalServerError() *PatchTagByIDInternalServerError {

	return &PatchTagByIDInternalServerError{}
}

// WithPayload adds the payload to the patch tag by Id internal server error response
func (o *PatchTagByIDInternalServerError) WithPayload(payload string) *PatchTagByIDInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the patch tag by Id internal server error response
func (o *PatchTagByIDInternalServerError) SetPayload(payload string) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PatchTagByIDInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	payload := o.Payload
	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}