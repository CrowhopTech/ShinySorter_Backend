// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// QuestionPatch question patch
//
// swagger:model questionPatch
type QuestionPatch struct {

	// Whether this functions as an "and" (true, only one option selected) or an "or" question false, default, can select multiple)
	// Enum: [true false]
	MutuallyExclusive string `json:"mutuallyExclusive,omitempty"`

	// ordering ID
	OrderingID int64 `json:"orderingID,omitempty"`

	// question text
	QuestionText string `json:"questionText,omitempty"`

	// tag options
	TagOptions []*TagOption `json:"tagOptions"`
}

// Validate validates this question patch
func (m *QuestionPatch) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateMutuallyExclusive(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTagOptions(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var questionPatchTypeMutuallyExclusivePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["true","false"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		questionPatchTypeMutuallyExclusivePropEnum = append(questionPatchTypeMutuallyExclusivePropEnum, v)
	}
}

const (

	// QuestionPatchMutuallyExclusiveTrue captures enum value "true"
	QuestionPatchMutuallyExclusiveTrue string = "true"

	// QuestionPatchMutuallyExclusiveFalse captures enum value "false"
	QuestionPatchMutuallyExclusiveFalse string = "false"
)

// prop value enum
func (m *QuestionPatch) validateMutuallyExclusiveEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, questionPatchTypeMutuallyExclusivePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *QuestionPatch) validateMutuallyExclusive(formats strfmt.Registry) error {
	if swag.IsZero(m.MutuallyExclusive) { // not required
		return nil
	}

	// value enum
	if err := m.validateMutuallyExclusiveEnum("mutuallyExclusive", "body", m.MutuallyExclusive); err != nil {
		return err
	}

	return nil
}

func (m *QuestionPatch) validateTagOptions(formats strfmt.Registry) error {
	if swag.IsZero(m.TagOptions) { // not required
		return nil
	}

	for i := 0; i < len(m.TagOptions); i++ {
		if swag.IsZero(m.TagOptions[i]) { // not required
			continue
		}

		if m.TagOptions[i] != nil {
			if err := m.TagOptions[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("tagOptions" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("tagOptions" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this question patch based on the context it is used
func (m *QuestionPatch) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateTagOptions(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *QuestionPatch) contextValidateTagOptions(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.TagOptions); i++ {

		if m.TagOptions[i] != nil {
			if err := m.TagOptions[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("tagOptions" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("tagOptions" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *QuestionPatch) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *QuestionPatch) UnmarshalBinary(b []byte) error {
	var res QuestionPatch
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
