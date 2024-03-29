// Code generated by go-swagger; DO NOT EDIT.

package files

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/CrowhopTech/shinysorter/backend/pkg/swagger/server/models"
)

// PatchFileByIDOKCode is the HTTP code returned for type PatchFileByIDOK
const PatchFileByIDOKCode int = 200

/*PatchFileByIDOK Returns the modified file.

swagger:response patchFileByIdOK
*/
type PatchFileByIDOK struct {

	/*
	  In: Body
	*/
	Payload *models.FileEntry `json:"body,omitempty"`
}

// NewPatchFileByIDOK creates PatchFileByIDOK with default headers values
func NewPatchFileByIDOK() *PatchFileByIDOK {

	return &PatchFileByIDOK{}
}

// WithPayload adds the payload to the patch file by Id o k response
func (o *PatchFileByIDOK) WithPayload(payload *models.FileEntry) *PatchFileByIDOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the patch file by Id o k response
func (o *PatchFileByIDOK) SetPayload(payload *models.FileEntry) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PatchFileByIDOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// PatchFileByIDBadRequestCode is the HTTP code returned for type PatchFileByIDBadRequest
const PatchFileByIDBadRequestCode int = 400

/*PatchFileByIDBadRequest Some part of the request was invalid. More information will be included in the error string

swagger:response patchFileByIdBadRequest
*/
type PatchFileByIDBadRequest struct {

	/*
	  In: Body
	*/
	Payload string `json:"body,omitempty"`
}

// NewPatchFileByIDBadRequest creates PatchFileByIDBadRequest with default headers values
func NewPatchFileByIDBadRequest() *PatchFileByIDBadRequest {

	return &PatchFileByIDBadRequest{}
}

// WithPayload adds the payload to the patch file by Id bad request response
func (o *PatchFileByIDBadRequest) WithPayload(payload string) *PatchFileByIDBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the patch file by Id bad request response
func (o *PatchFileByIDBadRequest) SetPayload(payload string) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PatchFileByIDBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	payload := o.Payload
	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

// PatchFileByIDInternalServerErrorCode is the HTTP code returned for type PatchFileByIDInternalServerError
const PatchFileByIDInternalServerErrorCode int = 500

/*PatchFileByIDInternalServerError Something else went wrong during the request

swagger:response patchFileByIdInternalServerError
*/
type PatchFileByIDInternalServerError struct {

	/*
	  In: Body
	*/
	Payload string `json:"body,omitempty"`
}

// NewPatchFileByIDInternalServerError creates PatchFileByIDInternalServerError with default headers values
func NewPatchFileByIDInternalServerError() *PatchFileByIDInternalServerError {

	return &PatchFileByIDInternalServerError{}
}

// WithPayload adds the payload to the patch file by Id internal server error response
func (o *PatchFileByIDInternalServerError) WithPayload(payload string) *PatchFileByIDInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the patch file by Id internal server error response
func (o *PatchFileByIDInternalServerError) SetPayload(payload string) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PatchFileByIDInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	payload := o.Payload
	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}
