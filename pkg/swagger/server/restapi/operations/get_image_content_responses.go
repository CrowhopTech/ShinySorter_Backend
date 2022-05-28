// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"io"
	"net/http"

	"github.com/go-openapi/runtime"
)

// GetImageContentOKCode is the HTTP code returned for type GetImageContentOK
const GetImageContentOKCode int = 200

/*GetImageContentOK Returns the image contents

swagger:response getImageContentOK
*/
type GetImageContentOK struct {
	/*

	 */
	ContentType string `json:"Content-Type"`

	/*
	  In: Body
	*/
	Payload io.ReadCloser `json:"body,omitempty"`
}

// NewGetImageContentOK creates GetImageContentOK with default headers values
func NewGetImageContentOK() *GetImageContentOK {

	return &GetImageContentOK{}
}

// WithContentType adds the contentType to the get image content o k response
func (o *GetImageContentOK) WithContentType(contentType string) *GetImageContentOK {
	o.ContentType = contentType
	return o
}

// SetContentType sets the contentType to the get image content o k response
func (o *GetImageContentOK) SetContentType(contentType string) {
	o.ContentType = contentType
}

// WithPayload adds the payload to the get image content o k response
func (o *GetImageContentOK) WithPayload(payload io.ReadCloser) *GetImageContentOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get image content o k response
func (o *GetImageContentOK) SetPayload(payload io.ReadCloser) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetImageContentOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	// response header Content-Type

	contentType := o.ContentType
	if contentType != "" {
		rw.Header().Set("Content-Type", contentType)
	}

	rw.WriteHeader(200)
	payload := o.Payload
	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

// GetImageContentNotFoundCode is the HTTP code returned for type GetImageContentNotFound
const GetImageContentNotFoundCode int = 404

/*GetImageContentNotFound The given image was not found.

swagger:response getImageContentNotFound
*/
type GetImageContentNotFound struct {
}

// NewGetImageContentNotFound creates GetImageContentNotFound with default headers values
func NewGetImageContentNotFound() *GetImageContentNotFound {

	return &GetImageContentNotFound{}
}

// WriteResponse to the client
func (o *GetImageContentNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(404)
}

// GetImageContentInternalServerErrorCode is the HTTP code returned for type GetImageContentInternalServerError
const GetImageContentInternalServerErrorCode int = 500

/*GetImageContentInternalServerError Something else went wrong during the request

swagger:response getImageContentInternalServerError
*/
type GetImageContentInternalServerError struct {

	/*
	  In: Body
	*/
	Payload string `json:"body,omitempty"`
}

// NewGetImageContentInternalServerError creates GetImageContentInternalServerError with default headers values
func NewGetImageContentInternalServerError() *GetImageContentInternalServerError {

	return &GetImageContentInternalServerError{}
}

// WithPayload adds the payload to the get image content internal server error response
func (o *GetImageContentInternalServerError) WithPayload(payload string) *GetImageContentInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get image content internal server error response
func (o *GetImageContentInternalServerError) SetPayload(payload string) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetImageContentInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	payload := o.Payload
	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}