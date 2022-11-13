// Code generated by go-swagger; DO NOT EDIT.

package files

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
)

// NewListFilesParams creates a new ListFilesParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewListFilesParams() *ListFilesParams {
	return &ListFilesParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewListFilesParamsWithTimeout creates a new ListFilesParams object
// with the ability to set a timeout on a request.
func NewListFilesParamsWithTimeout(timeout time.Duration) *ListFilesParams {
	return &ListFilesParams{
		timeout: timeout,
	}
}

// NewListFilesParamsWithContext creates a new ListFilesParams object
// with the ability to set a context for a request.
func NewListFilesParamsWithContext(ctx context.Context) *ListFilesParams {
	return &ListFilesParams{
		Context: ctx,
	}
}

// NewListFilesParamsWithHTTPClient creates a new ListFilesParams object
// with the ability to set a custom HTTPClient for a request.
func NewListFilesParamsWithHTTPClient(client *http.Client) *ListFilesParams {
	return &ListFilesParams{
		HTTPClient: client,
	}
}

/* ListFilesParams contains all the parameters to send to the API endpoint
   for the list files operation.

   Typically these are written to a http.Request.
*/
type ListFilesParams struct {

	/* Continue.

	   The last object ID of the previous page
	*/
	Continue *string

	/* ExcludeOperator.

	   Whether excludeTags requires all tags to match, or just one

	   Default: "all"
	*/
	ExcludeOperator *string

	/* ExcludeTags.

	   Tags to exclude in this query, referenced by tag ID
	*/
	ExcludeTags []int64

	/* HasBeenTagged.

	   Whether to filter to tags that have or have not been tagged
	*/
	HasBeenTagged *bool

	/* IncludeOperator.

	   Whether includeTags requires all tags to match, or just one

	   Default: "all"
	*/
	IncludeOperator *string

	/* IncludeTags.

	   Tags to include in this query, referenced by tag ID
	*/
	IncludeTags []int64

	/* Limit.

	   The count of results to return (aka page size)

	   Default: 5
	*/
	Limit *int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the list files params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ListFilesParams) WithDefaults() *ListFilesParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the list files params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ListFilesParams) SetDefaults() {
	var (
		excludeOperatorDefault = string("all")

		includeOperatorDefault = string("all")

		limitDefault = int64(5)
	)

	val := ListFilesParams{
		ExcludeOperator: &excludeOperatorDefault,
		IncludeOperator: &includeOperatorDefault,
		Limit:           &limitDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the list files params
func (o *ListFilesParams) WithTimeout(timeout time.Duration) *ListFilesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the list files params
func (o *ListFilesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the list files params
func (o *ListFilesParams) WithContext(ctx context.Context) *ListFilesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the list files params
func (o *ListFilesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the list files params
func (o *ListFilesParams) WithHTTPClient(client *http.Client) *ListFilesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the list files params
func (o *ListFilesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithContinue adds the continueVar to the list files params
func (o *ListFilesParams) WithContinue(continueVar *string) *ListFilesParams {
	o.SetContinue(continueVar)
	return o
}

// SetContinue adds the continue to the list files params
func (o *ListFilesParams) SetContinue(continueVar *string) {
	o.Continue = continueVar
}

// WithExcludeOperator adds the excludeOperator to the list files params
func (o *ListFilesParams) WithExcludeOperator(excludeOperator *string) *ListFilesParams {
	o.SetExcludeOperator(excludeOperator)
	return o
}

// SetExcludeOperator adds the excludeOperator to the list files params
func (o *ListFilesParams) SetExcludeOperator(excludeOperator *string) {
	o.ExcludeOperator = excludeOperator
}

// WithExcludeTags adds the excludeTags to the list files params
func (o *ListFilesParams) WithExcludeTags(excludeTags []int64) *ListFilesParams {
	o.SetExcludeTags(excludeTags)
	return o
}

// SetExcludeTags adds the excludeTags to the list files params
func (o *ListFilesParams) SetExcludeTags(excludeTags []int64) {
	o.ExcludeTags = excludeTags
}

// WithHasBeenTagged adds the hasBeenTagged to the list files params
func (o *ListFilesParams) WithHasBeenTagged(hasBeenTagged *bool) *ListFilesParams {
	o.SetHasBeenTagged(hasBeenTagged)
	return o
}

// SetHasBeenTagged adds the hasBeenTagged to the list files params
func (o *ListFilesParams) SetHasBeenTagged(hasBeenTagged *bool) {
	o.HasBeenTagged = hasBeenTagged
}

// WithIncludeOperator adds the includeOperator to the list files params
func (o *ListFilesParams) WithIncludeOperator(includeOperator *string) *ListFilesParams {
	o.SetIncludeOperator(includeOperator)
	return o
}

// SetIncludeOperator adds the includeOperator to the list files params
func (o *ListFilesParams) SetIncludeOperator(includeOperator *string) {
	o.IncludeOperator = includeOperator
}

// WithIncludeTags adds the includeTags to the list files params
func (o *ListFilesParams) WithIncludeTags(includeTags []int64) *ListFilesParams {
	o.SetIncludeTags(includeTags)
	return o
}

// SetIncludeTags adds the includeTags to the list files params
func (o *ListFilesParams) SetIncludeTags(includeTags []int64) {
	o.IncludeTags = includeTags
}

// WithLimit adds the limit to the list files params
func (o *ListFilesParams) WithLimit(limit *int64) *ListFilesParams {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the list files params
func (o *ListFilesParams) SetLimit(limit *int64) {
	o.Limit = limit
}

// WriteToRequest writes these params to a swagger request
func (o *ListFilesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Continue != nil {

		// query param continue
		var qrContinue string

		if o.Continue != nil {
			qrContinue = *o.Continue
		}
		qContinue := qrContinue
		if qContinue != "" {

			if err := r.SetQueryParam("continue", qContinue); err != nil {
				return err
			}
		}
	}

	if o.ExcludeOperator != nil {

		// query param excludeOperator
		var qrExcludeOperator string

		if o.ExcludeOperator != nil {
			qrExcludeOperator = *o.ExcludeOperator
		}
		qExcludeOperator := qrExcludeOperator
		if qExcludeOperator != "" {

			if err := r.SetQueryParam("excludeOperator", qExcludeOperator); err != nil {
				return err
			}
		}
	}

	if o.ExcludeTags != nil {

		// binding items for excludeTags
		joinedExcludeTags := o.bindParamExcludeTags(reg)

		// query array param excludeTags
		if err := r.SetQueryParam("excludeTags", joinedExcludeTags...); err != nil {
			return err
		}
	}

	if o.HasBeenTagged != nil {

		// query param hasBeenTagged
		var qrHasBeenTagged bool

		if o.HasBeenTagged != nil {
			qrHasBeenTagged = *o.HasBeenTagged
		}
		qHasBeenTagged := swag.FormatBool(qrHasBeenTagged)
		if qHasBeenTagged != "" {

			if err := r.SetQueryParam("hasBeenTagged", qHasBeenTagged); err != nil {
				return err
			}
		}
	}

	if o.IncludeOperator != nil {

		// query param includeOperator
		var qrIncludeOperator string

		if o.IncludeOperator != nil {
			qrIncludeOperator = *o.IncludeOperator
		}
		qIncludeOperator := qrIncludeOperator
		if qIncludeOperator != "" {

			if err := r.SetQueryParam("includeOperator", qIncludeOperator); err != nil {
				return err
			}
		}
	}

	if o.IncludeTags != nil {

		// binding items for includeTags
		joinedIncludeTags := o.bindParamIncludeTags(reg)

		// query array param includeTags
		if err := r.SetQueryParam("includeTags", joinedIncludeTags...); err != nil {
			return err
		}
	}

	if o.Limit != nil {

		// query param limit
		var qrLimit int64

		if o.Limit != nil {
			qrLimit = *o.Limit
		}
		qLimit := swag.FormatInt64(qrLimit)
		if qLimit != "" {

			if err := r.SetQueryParam("limit", qLimit); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindParamListFiles binds the parameter excludeTags
func (o *ListFilesParams) bindParamExcludeTags(formats strfmt.Registry) []string {
	excludeTagsIR := o.ExcludeTags

	var excludeTagsIC []string
	for _, excludeTagsIIR := range excludeTagsIR { // explode []int64

		excludeTagsIIV := swag.FormatInt64(excludeTagsIIR) // int64 as string
		excludeTagsIC = append(excludeTagsIC, excludeTagsIIV)
	}

	// items.CollectionFormat: ""
	excludeTagsIS := swag.JoinByFormat(excludeTagsIC, "")

	return excludeTagsIS
}

// bindParamListFiles binds the parameter includeTags
func (o *ListFilesParams) bindParamIncludeTags(formats strfmt.Registry) []string {
	includeTagsIR := o.IncludeTags

	var includeTagsIC []string
	for _, includeTagsIIR := range includeTagsIR { // explode []int64

		includeTagsIIV := swag.FormatInt64(includeTagsIIR) // int64 as string
		includeTagsIC = append(includeTagsIC, includeTagsIIV)
	}

	// items.CollectionFormat: ""
	includeTagsIS := swag.JoinByFormat(includeTagsIC, "")

	return includeTagsIS
}