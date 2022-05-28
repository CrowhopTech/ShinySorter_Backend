// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/CrowhopTech/shinysorter/backend/pkg/swagger/server/models"
)

// ListQuestionsOKCode is the HTTP code returned for type ListQuestionsOK
const ListQuestionsOKCode int = 200

/*ListQuestionsOK Questions were listed successfully (array may be empty if no questions are registered)

swagger:response listQuestionsOK
*/
type ListQuestionsOK struct {

	/*
	  In: Body
	*/
	Payload []*models.Question `json:"body,omitempty"`
}

// NewListQuestionsOK creates ListQuestionsOK with default headers values
func NewListQuestionsOK() *ListQuestionsOK {

	return &ListQuestionsOK{}
}

// WithPayload adds the payload to the list questions o k response
func (o *ListQuestionsOK) WithPayload(payload []*models.Question) *ListQuestionsOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the list questions o k response
func (o *ListQuestionsOK) SetPayload(payload []*models.Question) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ListQuestionsOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	payload := o.Payload
	if payload == nil {
		// return empty array
		payload = make([]*models.Question, 0, 50)
	}

	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

// ListQuestionsInternalServerErrorCode is the HTTP code returned for type ListQuestionsInternalServerError
const ListQuestionsInternalServerErrorCode int = 500

/*ListQuestionsInternalServerError Something else went wrong during the request

swagger:response listQuestionsInternalServerError
*/
type ListQuestionsInternalServerError struct {

	/*
	  In: Body
	*/
	Payload string `json:"body,omitempty"`
}

// NewListQuestionsInternalServerError creates ListQuestionsInternalServerError with default headers values
func NewListQuestionsInternalServerError() *ListQuestionsInternalServerError {

	return &ListQuestionsInternalServerError{}
}

// WithPayload adds the payload to the list questions internal server error response
func (o *ListQuestionsInternalServerError) WithPayload(payload string) *ListQuestionsInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the list questions internal server error response
func (o *ListQuestionsInternalServerError) SetPayload(payload string) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ListQuestionsInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	payload := o.Payload
	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}