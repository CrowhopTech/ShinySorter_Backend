// Code generated by go-swagger; DO NOT EDIT.

package tags

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
	"github.com/go-openapi/swag"

	"github.com/CrowhopTech/shinysorter/backend/pkg/swagger/models"
)

// NewPatchTagByIDParams creates a new PatchTagByIDParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPatchTagByIDParams() *PatchTagByIDParams {
	return &PatchTagByIDParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPatchTagByIDParamsWithTimeout creates a new PatchTagByIDParams object
// with the ability to set a timeout on a request.
func NewPatchTagByIDParamsWithTimeout(timeout time.Duration) *PatchTagByIDParams {
	return &PatchTagByIDParams{
		timeout: timeout,
	}
}

// NewPatchTagByIDParamsWithContext creates a new PatchTagByIDParams object
// with the ability to set a context for a request.
func NewPatchTagByIDParamsWithContext(ctx context.Context) *PatchTagByIDParams {
	return &PatchTagByIDParams{
		Context: ctx,
	}
}

// NewPatchTagByIDParamsWithHTTPClient creates a new PatchTagByIDParams object
// with the ability to set a custom HTTPClient for a request.
func NewPatchTagByIDParamsWithHTTPClient(client *http.Client) *PatchTagByIDParams {
	return &PatchTagByIDParams{
		HTTPClient: client,
	}
}

/* PatchTagByIDParams contains all the parameters to send to the API endpoint
   for the patch tag by ID operation.

   Typically these are written to a http.Request.
*/
type PatchTagByIDParams struct {

	/* ID.

	   ID of the tag to modify
	*/
	ID int64

	/* Patch.

	   Patch modifications for the tag
	*/
	Patch *models.TagPatch

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the patch tag by ID params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PatchTagByIDParams) WithDefaults() *PatchTagByIDParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the patch tag by ID params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PatchTagByIDParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the patch tag by ID params
func (o *PatchTagByIDParams) WithTimeout(timeout time.Duration) *PatchTagByIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the patch tag by ID params
func (o *PatchTagByIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the patch tag by ID params
func (o *PatchTagByIDParams) WithContext(ctx context.Context) *PatchTagByIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the patch tag by ID params
func (o *PatchTagByIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the patch tag by ID params
func (o *PatchTagByIDParams) WithHTTPClient(client *http.Client) *PatchTagByIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the patch tag by ID params
func (o *PatchTagByIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the patch tag by ID params
func (o *PatchTagByIDParams) WithID(id int64) *PatchTagByIDParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the patch tag by ID params
func (o *PatchTagByIDParams) SetID(id int64) {
	o.ID = id
}

// WithPatch adds the patch to the patch tag by ID params
func (o *PatchTagByIDParams) WithPatch(patch *models.TagPatch) *PatchTagByIDParams {
	o.SetPatch(patch)
	return o
}

// SetPatch adds the patch to the patch tag by ID params
func (o *PatchTagByIDParams) SetPatch(patch *models.TagPatch) {
	o.Patch = patch
}

// WriteToRequest writes these params to a swagger request
func (o *PatchTagByIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", swag.FormatInt64(o.ID)); err != nil {
		return err
	}
	if o.Patch != nil {
		if err := r.SetBodyParam(o.Patch); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
