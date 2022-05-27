// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"errors"
	"net/url"
	golangswaggerpaths "path"

	"github.com/go-openapi/swag"
)

// ListImagesURL generates an URL for the list images operation
type ListImagesURL struct {
	ExcludeOperator *string
	ExcludeTags     []int64
	HasBeenTagged   *bool
	IncludeOperator *string
	IncludeTags     []int64

	_basePath string
	// avoid unkeyed usage
	_ struct{}
}

// WithBasePath sets the base path for this url builder, only required when it's different from the
// base path specified in the swagger spec.
// When the value of the base path is an empty string
func (o *ListImagesURL) WithBasePath(bp string) *ListImagesURL {
	o.SetBasePath(bp)
	return o
}

// SetBasePath sets the base path for this url builder, only required when it's different from the
// base path specified in the swagger spec.
// When the value of the base path is an empty string
func (o *ListImagesURL) SetBasePath(bp string) {
	o._basePath = bp
}

// Build a url path and query string
func (o *ListImagesURL) Build() (*url.URL, error) {
	var _result url.URL

	var _path = "/images"

	_basePath := o._basePath
	_result.Path = golangswaggerpaths.Join(_basePath, _path)

	qs := make(url.Values)

	var excludeOperatorQ string
	if o.ExcludeOperator != nil {
		excludeOperatorQ = *o.ExcludeOperator
	}
	if excludeOperatorQ != "" {
		qs.Set("excludeOperator", excludeOperatorQ)
	}

	var excludeTagsIR []string
	for _, excludeTagsI := range o.ExcludeTags {
		excludeTagsIS := swag.FormatInt64(excludeTagsI)
		if excludeTagsIS != "" {
			excludeTagsIR = append(excludeTagsIR, excludeTagsIS)
		}
	}

	excludeTags := swag.JoinByFormat(excludeTagsIR, "")

	if len(excludeTags) > 0 {
		qsv := excludeTags[0]
		if qsv != "" {
			qs.Set("excludeTags", qsv)
		}
	}

	var hasBeenTaggedQ string
	if o.HasBeenTagged != nil {
		hasBeenTaggedQ = swag.FormatBool(*o.HasBeenTagged)
	}
	if hasBeenTaggedQ != "" {
		qs.Set("hasBeenTagged", hasBeenTaggedQ)
	}

	var includeOperatorQ string
	if o.IncludeOperator != nil {
		includeOperatorQ = *o.IncludeOperator
	}
	if includeOperatorQ != "" {
		qs.Set("includeOperator", includeOperatorQ)
	}

	var includeTagsIR []string
	for _, includeTagsI := range o.IncludeTags {
		includeTagsIS := swag.FormatInt64(includeTagsI)
		if includeTagsIS != "" {
			includeTagsIR = append(includeTagsIR, includeTagsIS)
		}
	}

	includeTags := swag.JoinByFormat(includeTagsIR, "")

	if len(includeTags) > 0 {
		qsv := includeTags[0]
		if qsv != "" {
			qs.Set("includeTags", qsv)
		}
	}

	_result.RawQuery = qs.Encode()

	return &_result, nil
}

// Must is a helper function to panic when the url builder returns an error
func (o *ListImagesURL) Must(u *url.URL, err error) *url.URL {
	if err != nil {
		panic(err)
	}
	if u == nil {
		panic("url can't be nil")
	}
	return u
}

// String returns the string representation of the path with query string
func (o *ListImagesURL) String() string {
	return o.Must(o.Build()).String()
}

// BuildFull builds a full url with scheme, host, path and query string
func (o *ListImagesURL) BuildFull(scheme, host string) (*url.URL, error) {
	if scheme == "" {
		return nil, errors.New("scheme is required for a full url on ListImagesURL")
	}
	if host == "" {
		return nil, errors.New("host is required for a full url on ListImagesURL")
	}

	base, err := o.Build()
	if err != nil {
		return nil, err
	}

	base.Scheme = scheme
	base.Host = host
	return base, nil
}

// StringFull returns the string representation of a complete url
func (o *ListImagesURL) StringFull(scheme, host string) string {
	return o.Must(o.BuildFull(scheme, host)).String()
}
