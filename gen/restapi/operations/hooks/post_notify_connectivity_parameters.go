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

// NewPostNotifyConnectivityParams creates a new PostNotifyConnectivityParams object
// no default values defined in spec.
func NewPostNotifyConnectivityParams() PostNotifyConnectivityParams {

	return PostNotifyConnectivityParams{}
}

// PostNotifyConnectivityParams contains all the bound params for the post notify connectivity operation
// typically these are obtained from a http.Request
//
// swagger:parameters PostNotifyConnectivity
type PostNotifyConnectivityParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*when the user logs out
	  In: body
	*/
	Data *models.NotifyConnectivity
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewPostNotifyConnectivityParams() beforehand.
func (o *PostNotifyConnectivityParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	if runtime.HasBody(r) {
		defer r.Body.Close()
		var body models.NotifyConnectivity
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
