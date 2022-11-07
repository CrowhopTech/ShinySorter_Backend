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

// NewCreateQuestionParams creates a new CreateQuestionParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCreateQuestionParams() *CreateQuestionParams {
	return &CreateQuestionParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCreateQuestionParamsWithTimeout creates a new CreateQuestionParams object
// with the ability to set a timeout on a request.
func NewCreateQuestionParamsWithTimeout(timeout time.Duration) *CreateQuestionParams {
	return &CreateQuestionParams{
		timeout: timeout,
	}
}

// NewCreateQuestionParamsWithContext creates a new CreateQuestionParams object
// with the ability to set a context for a request.
func NewCreateQuestionParamsWithContext(ctx context.Context) *CreateQuestionParams {
	return &CreateQuestionParams{
		Context: ctx,
	}
}

// NewCreateQuestionParamsWithHTTPClient creates a new CreateQuestionParams object
// with the ability to set a custom HTTPClient for a request.
func NewCreateQuestionParamsWithHTTPClient(client *http.Client) *CreateQuestionParams {
	return &CreateQuestionParams{
		HTTPClient: client,
	}
}

/* CreateQuestionParams contains all the parameters to send to the API endpoint
   for the create question operation.

   Typically these are written to a http.Request.
*/
type CreateQuestionParams struct {

	/* NewQuestion.

	   The new question to create
	*/
	NewQuestion *models.QuestionCreate

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the create question params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateQuestionParams) WithDefaults() *CreateQuestionParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the create question params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateQuestionParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the create question params
func (o *CreateQuestionParams) WithTimeout(timeout time.Duration) *CreateQuestionParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create question params
func (o *CreateQuestionParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create question params
func (o *CreateQuestionParams) WithContext(ctx context.Context) *CreateQuestionParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create question params
func (o *CreateQuestionParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create question params
func (o *CreateQuestionParams) WithHTTPClient(client *http.Client) *CreateQuestionParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create question params
func (o *CreateQuestionParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithNewQuestion adds the newQuestion to the create question params
func (o *CreateQuestionParams) WithNewQuestion(newQuestion *models.QuestionCreate) *CreateQuestionParams {
	o.SetNewQuestion(newQuestion)
	return o
}

// SetNewQuestion adds the newQuestion to the create question params
func (o *CreateQuestionParams) SetNewQuestion(newQuestion *models.QuestionCreate) {
	o.NewQuestion = newQuestion
}

// WriteToRequest writes these params to a swagger request
func (o *CreateQuestionParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.NewQuestion != nil {
		if err := r.SetBodyParam(o.NewQuestion); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
