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
	"github.com/go-openapi/swag"
)

// NewListImagesParams creates a new ListImagesParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewListImagesParams() *ListImagesParams {
	return &ListImagesParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewListImagesParamsWithTimeout creates a new ListImagesParams object
// with the ability to set a timeout on a request.
func NewListImagesParamsWithTimeout(timeout time.Duration) *ListImagesParams {
	return &ListImagesParams{
		timeout: timeout,
	}
}

// NewListImagesParamsWithContext creates a new ListImagesParams object
// with the ability to set a context for a request.
func NewListImagesParamsWithContext(ctx context.Context) *ListImagesParams {
	return &ListImagesParams{
		Context: ctx,
	}
}

// NewListImagesParamsWithHTTPClient creates a new ListImagesParams object
// with the ability to set a custom HTTPClient for a request.
func NewListImagesParamsWithHTTPClient(client *http.Client) *ListImagesParams {
	return &ListImagesParams{
		HTTPClient: client,
	}
}

/* ListImagesParams contains all the parameters to send to the API endpoint
   for the list images operation.

   Typically these are written to a http.Request.
*/
type ListImagesParams struct {

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

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the list images params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ListImagesParams) WithDefaults() *ListImagesParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the list images params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ListImagesParams) SetDefaults() {
	var (
		excludeOperatorDefault = string("all")

		includeOperatorDefault = string("all")
	)

	val := ListImagesParams{
		ExcludeOperator: &excludeOperatorDefault,
		IncludeOperator: &includeOperatorDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the list images params
func (o *ListImagesParams) WithTimeout(timeout time.Duration) *ListImagesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the list images params
func (o *ListImagesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the list images params
func (o *ListImagesParams) WithContext(ctx context.Context) *ListImagesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the list images params
func (o *ListImagesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the list images params
func (o *ListImagesParams) WithHTTPClient(client *http.Client) *ListImagesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the list images params
func (o *ListImagesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithExcludeOperator adds the excludeOperator to the list images params
func (o *ListImagesParams) WithExcludeOperator(excludeOperator *string) *ListImagesParams {
	o.SetExcludeOperator(excludeOperator)
	return o
}

// SetExcludeOperator adds the excludeOperator to the list images params
func (o *ListImagesParams) SetExcludeOperator(excludeOperator *string) {
	o.ExcludeOperator = excludeOperator
}

// WithExcludeTags adds the excludeTags to the list images params
func (o *ListImagesParams) WithExcludeTags(excludeTags []int64) *ListImagesParams {
	o.SetExcludeTags(excludeTags)
	return o
}

// SetExcludeTags adds the excludeTags to the list images params
func (o *ListImagesParams) SetExcludeTags(excludeTags []int64) {
	o.ExcludeTags = excludeTags
}

// WithHasBeenTagged adds the hasBeenTagged to the list images params
func (o *ListImagesParams) WithHasBeenTagged(hasBeenTagged *bool) *ListImagesParams {
	o.SetHasBeenTagged(hasBeenTagged)
	return o
}

// SetHasBeenTagged adds the hasBeenTagged to the list images params
func (o *ListImagesParams) SetHasBeenTagged(hasBeenTagged *bool) {
	o.HasBeenTagged = hasBeenTagged
}

// WithIncludeOperator adds the includeOperator to the list images params
func (o *ListImagesParams) WithIncludeOperator(includeOperator *string) *ListImagesParams {
	o.SetIncludeOperator(includeOperator)
	return o
}

// SetIncludeOperator adds the includeOperator to the list images params
func (o *ListImagesParams) SetIncludeOperator(includeOperator *string) {
	o.IncludeOperator = includeOperator
}

// WithIncludeTags adds the includeTags to the list images params
func (o *ListImagesParams) WithIncludeTags(includeTags []int64) *ListImagesParams {
	o.SetIncludeTags(includeTags)
	return o
}

// SetIncludeTags adds the includeTags to the list images params
func (o *ListImagesParams) SetIncludeTags(includeTags []int64) {
	o.IncludeTags = includeTags
}

// WriteToRequest writes these params to a swagger request
func (o *ListImagesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

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

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindParamListImages binds the parameter excludeTags
func (o *ListImagesParams) bindParamExcludeTags(formats strfmt.Registry) []string {
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

// bindParamListImages binds the parameter includeTags
func (o *ListImagesParams) bindParamIncludeTags(formats strfmt.Registry) []string {
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