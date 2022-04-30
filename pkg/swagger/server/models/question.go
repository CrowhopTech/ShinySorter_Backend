// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// Question question
// Example: {"orderingID":500,"questionID":5,"questionText":"What kinds of flowers are present in this picture?","requiresQuestion":4,"tagOptions":[{"optionText":"Tulips","tagID":5},{"optionText":"Roses","tagID":6},{"optionText":"Violets","tagID":7},{"optionText":"Daisies","tagID":8}]}
//
// swagger:model question
type Question struct {

	// ordering ID
	OrderingID int64 `json:"orderingID,omitempty"`

	// question ID
	// Required: true
	QuestionID *int64 `json:"questionID"`

	// question text
	// Required: true
	QuestionText *string `json:"questionText"`

	// requires question
	RequiresQuestion int64 `json:"requiresQuestion,omitempty"`

	// tag options
	// Required: true
	TagOptions []*QuestionTagOptionsItems0 `json:"tagOptions"`
}

// Validate validates this question
func (m *Question) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateQuestionID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateQuestionText(formats); err != nil {
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

func (m *Question) validateQuestionID(formats strfmt.Registry) error {

	if err := validate.Required("questionID", "body", m.QuestionID); err != nil {
		return err
	}

	return nil
}

func (m *Question) validateQuestionText(formats strfmt.Registry) error {

	if err := validate.Required("questionText", "body", m.QuestionText); err != nil {
		return err
	}

	return nil
}

func (m *Question) validateTagOptions(formats strfmt.Registry) error {

	if err := validate.Required("tagOptions", "body", m.TagOptions); err != nil {
		return err
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

// ContextValidate validate this question based on the context it is used
func (m *Question) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateTagOptions(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Question) contextValidateTagOptions(ctx context.Context, formats strfmt.Registry) error {

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
func (m *Question) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Question) UnmarshalBinary(b []byte) error {
	var res Question
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// QuestionTagOptionsItems0 question tag options items0
//
// swagger:model QuestionTagOptionsItems0
type QuestionTagOptionsItems0 struct {

	// option text
	// Required: true
	OptionText *string `json:"optionText"`

	// tag ID
	// Required: true
	TagID *int64 `json:"tagID"`
}

// Validate validates this question tag options items0
func (m *QuestionTagOptionsItems0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateOptionText(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTagID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *QuestionTagOptionsItems0) validateOptionText(formats strfmt.Registry) error {

	if err := validate.Required("optionText", "body", m.OptionText); err != nil {
		return err
	}

	return nil
}

func (m *QuestionTagOptionsItems0) validateTagID(formats strfmt.Registry) error {

	if err := validate.Required("tagID", "body", m.TagID); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this question tag options items0 based on context it is used
func (m *QuestionTagOptionsItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *QuestionTagOptionsItems0) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *QuestionTagOptionsItems0) UnmarshalBinary(b []byte) error {
	var res QuestionTagOptionsItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
