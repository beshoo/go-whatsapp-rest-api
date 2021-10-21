// Code generated by go-swagger; DO NOT EDIT.

package hooks

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"

	"bitbucket.org/rockyOO7/wa-api/gen/models"
)

// NewPostMessageDocParams creates a new PostMessageDocParams object
// no default values defined in spec.
func NewPostMessageDocParams() PostMessageDocParams {

	return PostMessageDocParams{}
}

// PostMessageDocParams contains all the bound params for the post message doc operation
// typically these are obtained from a http.Request
//
// swagger:parameters PostMessageDoc
type PostMessageDocParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*Doc message body
	  In: body
	*/
	Data *models.MessageDoc
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewPostMessageDocParams() beforehand.
func (o *PostMessageDocParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	if runtime.HasBody(r) {
		defer r.Body.Close()
		var body models.MessageDoc
		if err := route.Consumer.Consume(r.Body, &body); err != nil {
			res = append(res, errors.NewParseError("data", "body", "", err))
		} else {
			// validate body object
			if err := body.Validate(route.Formats); err != nil {
				res = append(res, err)
			}

			if len(res) == 0 {
				o.Data = &body
			}
		}
	}
	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}