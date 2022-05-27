// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// DeleteQuestionHandlerFunc turns a function with the right signature into a delete question handler
type DeleteQuestionHandlerFunc func(DeleteQuestionParams) middleware.Responder

// Handle executing the request and returning a response
func (fn DeleteQuestionHandlerFunc) Handle(params DeleteQuestionParams) middleware.Responder {
	return fn(params)
}

// DeleteQuestionHandler interface for that can handle valid delete question params
type DeleteQuestionHandler interface {
	Handle(DeleteQuestionParams) middleware.Responder
}

// NewDeleteQuestion creates a new http.Handler for the delete question operation
func NewDeleteQuestion(ctx *middleware.Context, handler DeleteQuestionHandler) *DeleteQuestion {
	return &DeleteQuestion{Context: ctx, Handler: handler}
}

/* DeleteQuestion swagger:route DELETE /questions/{id} deleteQuestion

Deletes a question.

*/
type DeleteQuestion struct {
	Context *middleware.Context
	Handler DeleteQuestionHandler
}

func (o *DeleteQuestion) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewDeleteQuestionParams()
	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}
