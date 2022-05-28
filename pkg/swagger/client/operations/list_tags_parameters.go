// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

// NewListTagsParams creates a new ListTagsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewListTagsParams() *ListTagsParams {
	return &ListTagsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewListTagsParamsWithTimeout creates a new ListTagsParams object
// with the ability to set a timeout on a request.
func NewListTagsParamsWithTimeout(timeout time.Duration) *ListTagsParams {
	return &ListTagsParams{
		timeout: timeout,
	}
}

// NewListTagsParamsWithContext creates a new ListTagsParams object
// with the ability to set a context for a request.
func NewListTagsParamsWithContext(ctx context.Context) *ListTagsParams {
	return &ListTagsParams{
		Context: ctx,
	}
}

// NewListTagsParamsWithHTTPClient creates a new ListTagsParams object
// with the ability to set a custom HTTPClient for a request.
func NewListTagsParamsWithHTTPClient(client *http.Client) *ListTagsParams {
	return &ListTagsParams{
		HTTPClient: client,
	}
}

/* ListTagsParams contains all the parameters to send to the API endpoint
   for the list tags operation.

   Typically these are written to a http.Request.
*/
type ListTagsParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the list tags params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ListTagsParams) WithDefaults() *ListTagsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the list tags params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ListTagsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the list tags params
func (o *ListTagsParams) WithTimeout(timeout time.Duration) *ListTagsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the list tags params
func (o *ListTagsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the list tags params
func (o *ListTagsParams) WithContext(ctx context.Context) *ListTagsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the list tags params
func (o *ListTagsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the list tags params
func (o *ListTagsParams) WithHTTPClient(client *http.Client) *ListTagsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the list tags params
func (o *ListTagsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *ListTagsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
