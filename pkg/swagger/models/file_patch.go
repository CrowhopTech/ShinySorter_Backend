// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// FilePatch file patch
//
// swagger:model filePatch
type FilePatch struct {

	// has been tagged
	HasBeenTagged *bool `json:"hasBeenTagged,omitempty"`

	// mime type
	// Example: image/png
	MimeType string `json:"mimeType,omitempty"`

	// tags
	// Example: [5,7,37]
	Tags []int64 `json:"tags"`
}

// Validate validates this file patch
func (m *FilePatch) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this file patch based on context it is used
func (m *FilePatch) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *FilePatch) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *FilePatch) UnmarshalBinary(b []byte) error {
	var res FilePatch
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}