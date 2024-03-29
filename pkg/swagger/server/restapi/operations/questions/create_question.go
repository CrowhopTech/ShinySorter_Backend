// Code generated by go-swagger; DO NOT EDIT.

package questions

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// CreateQuestionHandlerFunc turns a function with the right signature into a create question handler
type CreateQuestionHandlerFunc func(CreateQuestionParams) middleware.Responder

// Handle executing the request and returning a response
func (fn CreateQuestionHandlerFunc) Handle(params CreateQuestionParams) middleware.Responder {
	return fn(params)
}

// CreateQuestionHandler interface for that can handle valid create question params
type CreateQuestionHandler interface {
	Handle(CreateQuestionParams) middleware.Responder
}

// NewCreateQuestion creates a new http.Handler for the create question operation
func NewCreateQuestion(ctx *middleware.Context, handler CreateQuestionHandler) *CreateQuestion {
	return &CreateQuestion{Context: ctx, Handler: handler}
}

/* CreateQuestion swagger:route POST /questions questions createQuestion

Creates a new question

*/
type CreateQuestion struct {
	Context *middleware.Context
	Handler CreateQuestionHandler
}

func (o *CreateQuestion) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewCreateQuestionParams()
	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}
