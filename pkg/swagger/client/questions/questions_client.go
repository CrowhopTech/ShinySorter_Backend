// Code generated by go-swagger; DO NOT EDIT.

package questions

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new questions API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for questions API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	CreateQuestion(params *CreateQuestionParams, opts ...ClientOption) (*CreateQuestionCreated, error)

	DeleteQuestion(params *DeleteQuestionParams, opts ...ClientOption) (*DeleteQuestionOK, error)

	ListQuestions(params *ListQuestionsParams, opts ...ClientOption) (*ListQuestionsOK, error)

	PatchQuestionByID(params *PatchQuestionByIDParams, opts ...ClientOption) (*PatchQuestionByIDOK, error)

	ReorderQuestions(params *ReorderQuestionsParams, opts ...ClientOption) (*ReorderQuestionsOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  CreateQuestion Creates a new question
*/
func (a *Client) CreateQuestion(params *CreateQuestionParams, opts ...ClientOption) (*CreateQuestionCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateQuestionParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "createQuestion",
		Method:             "POST",
		PathPattern:        "/questions",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json", "multipart/form-data"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &CreateQuestionReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*CreateQuestionCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for createQuestion: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  DeleteQuestion Deletes a question.
*/
func (a *Client) DeleteQuestion(params *DeleteQuestionParams, opts ...ClientOption) (*DeleteQuestionOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteQuestionParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "deleteQuestion",
		Method:             "DELETE",
		PathPattern:        "/questions/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json", "multipart/form-data"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &DeleteQuestionReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*DeleteQuestionOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for deleteQuestion: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  ListQuestions Lists questions
*/
func (a *Client) ListQuestions(params *ListQuestionsParams, opts ...ClientOption) (*ListQuestionsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListQuestionsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "listQuestions",
		Method:             "GET",
		PathPattern:        "/questions",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json", "multipart/form-data"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &ListQuestionsReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ListQuestionsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for listQuestions: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  PatchQuestionByID Modifies question metadata
*/
func (a *Client) PatchQuestionByID(params *PatchQuestionByIDParams, opts ...ClientOption) (*PatchQuestionByIDOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPatchQuestionByIDParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "patchQuestionByID",
		Method:             "PATCH",
		PathPattern:        "/questions/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json", "multipart/form-data"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &PatchQuestionByIDReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*PatchQuestionByIDOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for patchQuestionByID: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  ReorderQuestions Reorders all questions (requires all question IDs to be passed in, e.g. a complete order)
*/
func (a *Client) ReorderQuestions(params *ReorderQuestionsParams, opts ...ClientOption) (*ReorderQuestionsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewReorderQuestionsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "reorderQuestions",
		Method:             "POST",
		PathPattern:        "/questions/reorder",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json", "multipart/form-data"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &ReorderQuestionsReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ReorderQuestionsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for reorderQuestions: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
