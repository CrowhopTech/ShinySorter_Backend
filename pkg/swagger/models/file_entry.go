// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// FileEntry file entry
//
// swagger:model fileEntry
type FileEntry struct {

	// has been tagged
	// Required: true
	HasBeenTagged bool `json:"hasBeenTagged"`

	// id
	// Example: 507f1f77bcf86cd799439011
	// Required: true
	// Max Length: 24
	// Min Length: 24
	ID *string `json:"id"`

	// md5sum
	// Example: 0a8bd0c4863ec1720da0f69d2795d18a
	// Required: true
	Md5sum *string `json:"md5sum"`

	// mime type
	// Example: image/png
	// Required: true
	MimeType *string `json:"mimeType"`

	// name
	// Example: filename.jpg
	// Required: true
	Name *string `json:"name"`

	// tags
	// Example: [5,7,37]
	// Required: true
	Tags []int64 `json:"tags"`
}

// Validate validates this file entry
func (m *FileEntry) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateHasBeenTagged(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMd5sum(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMimeType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTags(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *FileEntry) validateHasBeenTagged(formats strfmt.Registry) error {

	if err := validate.Required("hasBeenTagged", "body", bool(m.HasBeenTagged)); err != nil {
		return err
	}

	return nil
}

func (m *FileEntry) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	if err := validate.MinLength("id", "body", *m.ID, 24); err != nil {
		return err
	}

	if err := validate.MaxLength("id", "body", *m.ID, 24); err != nil {
		return err
	}

	return nil
}

func (m *FileEntry) validateMd5sum(formats strfmt.Registry) error {

	if err := validate.Required("md5sum", "body", m.Md5sum); err != nil {
		return err
	}

	return nil
}

func (m *FileEntry) validateMimeType(formats strfmt.Registry) error {

	if err := validate.Required("mimeType", "body", m.MimeType); err != nil {
		return err
	}

	return nil
}

func (m *FileEntry) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

func (m *FileEntry) validateTags(formats strfmt.Registry) error {

	if err := validate.Required("tags", "body", m.Tags); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this file entry based on context it is used
func (m *FileEntry) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *FileEntry) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *FileEntry) UnmarshalBinary(b []byte) error {
	var res FileEntry
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
