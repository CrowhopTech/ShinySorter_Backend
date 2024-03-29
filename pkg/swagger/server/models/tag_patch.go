// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// TagPatch tag patch
//
// swagger:model tagPatch
type TagPatch struct {

	// description
	// Example: This image contains a Tulip
	Description string `json:"description,omitempty"`

	// user friendly name
	// Example: Tulip
	UserFriendlyName string `json:"userFriendlyName,omitempty"`
}

// Validate validates this tag patch
func (m *TagPatch) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this tag patch based on context it is used
func (m *TagPatch) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *TagPatch) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TagPatch) UnmarshalBinary(b []byte) error {
	var res TagPatch
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
