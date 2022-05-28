// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"io"
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/validate"

	"github.com/CrowhopTech/shinysorter/backend/pkg/swagger/server/models"
)

// NewPatchImageByIDParams creates a new PatchImageByIDParams object
//
// There are no default values defined in the spec.
func NewPatchImageByIDParams() PatchImageByIDParams {

	return PatchImageByIDParams{}
}

// PatchImageByIDParams contains all the bound params for the patch image by Id operation
// typically these are obtained from a http.Request
//
// swagger:parameters patchImageById
type PatchImageByIDParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*Image ID
	  Required: true
	  In: path
	*/
	ID string
	/*Patch modifications for the image
	  Required: true
	  In: body
	*/
	Patch *models.Image
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewPatchImageByIDParams() beforehand.
func (o *PatchImageByIDParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	rID, rhkID, _ := route.Params.GetOK("id")
	if err := o.bindID(rID, rhkID, route.Formats); err != nil {
		res = append(res, err)
	}

	if runtime.HasBody(r) {
		defer r.Body.Close()
		var body models.Image
		if err := route.Consumer.Consume(r.Body, &body); err != nil {
			if err == io.EOF {
				res = append(res, errors.Required("patch", "body", ""))
			} else {
				res = append(res, errors.NewParseError("patch", "body", "", err))
			}
		} else {
			// validate body object
			if err := body.Validate(route.Formats); err != nil {
				res = append(res, err)
			}

			ctx := validate.WithOperationRequest(context.Background())
			if err := body.ContextValidate(ctx, route.Formats); err != nil {
				res = append(res, err)
			}

			if len(res) == 0 {
				o.Patch = &body
			}
		}
	} else {
		res = append(res, errors.Required("patch", "body", ""))
	}
	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindID binds and validates parameter ID from path.
func (o *PatchImageByIDParams) bindID(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// Parameter is provided by construction from the route
	o.ID = raw

	return nil
}