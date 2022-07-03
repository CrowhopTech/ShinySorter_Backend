// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// DeleteTagHandlerFunc turns a function with the right signature into a delete tag handler
type DeleteTagHandlerFunc func(DeleteTagParams) middleware.Responder

// Handle executing the request and returning a response
func (fn DeleteTagHandlerFunc) Handle(params DeleteTagParams) middleware.Responder {
	return fn(params)
}

// DeleteTagHandler interface for that can handle valid delete tag params
type DeleteTagHandler interface {
	Handle(DeleteTagParams) middleware.Responder
}

// NewDeleteTag creates a new http.Handler for the delete tag operation
func NewDeleteTag(ctx *middleware.Context, handler DeleteTagHandler) *DeleteTag {
	return &DeleteTag{Context: ctx, Handler: handler}
}

/* DeleteTag swagger:route DELETE /tags/{id} deleteTag

Deletes a tag. Should also remove it from all files that use it.

*/
type DeleteTag struct {
	Context *middleware.Context
	Handler DeleteTagHandler
}

func (o *DeleteTag) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewDeleteTagParams()
	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}
