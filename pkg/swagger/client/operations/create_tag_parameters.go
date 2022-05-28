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

	"github.com/CrowhopTech/shinysorter/backend/pkg/swagger/models"
)

// NewCreateTagParams creates a new CreateTagParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCreateTagParams() *CreateTagParams {
	return &CreateTagParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCreateTagParamsWithTimeout creates a new CreateTagParams object
// with the ability to set a timeout on a request.
func NewCreateTagParamsWithTimeout(timeout time.Duration) *CreateTagParams {
	return &CreateTagParams{
		timeout: timeout,
	}
}

// NewCreateTagParamsWithContext creates a new CreateTagParams object
// with the ability to set a context for a request.
func NewCreateTagParamsWithContext(ctx context.Context) *CreateTagParams {
	return &CreateTagParams{
		Context: ctx,
	}
}

// NewCreateTagParamsWithHTTPClient creates a new CreateTagParams object
// with the ability to set a custom HTTPClient for a request.
func NewCreateTagParamsWithHTTPClient(client *http.Client) *CreateTagParams {
	return &CreateTagParams{
		HTTPClient: client,
	}
}

/* CreateTagParams contains all the parameters to send to the API endpoint
   for the create tag operation.

   Typically these are written to a http.Request.
*/
type CreateTagParams struct {

	/* NewTag.

	   The new tag to create
	*/
	NewTag *models.Tag

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the create tag params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateTagParams) WithDefaults() *CreateTagParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the create tag params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateTagParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the create tag params
func (o *CreateTagParams) WithTimeout(timeout time.Duration) *CreateTagParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create tag params
func (o *CreateTagParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create tag params
func (o *CreateTagParams) WithContext(ctx context.Context) *CreateTagParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create tag params
func (o *CreateTagParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create tag params
func (o *CreateTagParams) WithHTTPClient(client *http.Client) *CreateTagParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create tag params
func (o *CreateTagParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithNewTag adds the newTag to the create tag params
func (o *CreateTagParams) WithNewTag(newTag *models.Tag) *CreateTagParams {
	o.SetNewTag(newTag)
	return o
}

// SetNewTag adds the newTag to the create tag params
func (o *CreateTagParams) SetNewTag(newTag *models.Tag) {
	o.NewTag = newTag
}

// WriteToRequest writes these params to a swagger request
func (o *CreateTagParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.NewTag != nil {
		if err := r.SetBodyParam(o.NewTag); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}