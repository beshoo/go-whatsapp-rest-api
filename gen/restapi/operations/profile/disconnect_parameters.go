// Code generated by go-swagger; DO NOT EDIT.

package profile

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/validate"
)

// NewDisconnectParams creates a new DisconnectParams object
// no default values defined in spec.
func NewDisconnectParams() DisconnectParams {

	return DisconnectParams{}
}

// DisconnectParams contains all the bound params for the disconnect operation
// typically these are obtained from a http.Request
//
// swagger:parameters disconnect
type DisconnectParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*Session Id used while succesfull scan
	  Required: true
	  In: formData
	*/
	SessionID strfmt.UUID4
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewDisconnectParams() beforehand.
func (o *DisconnectParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	if err := r.ParseMultipartForm(32 << 20); err != nil {
		if err != http.ErrNotMultipart {
			return errors.New(400, "%v", err)
		} else if err := r.ParseForm(); err != nil {
			return errors.New(400, "%v", err)
		}
	}
	fds := runtime.Values(r.Form)

	fdSessionID, fdhkSessionID, _ := fds.GetOK("sessionId")
	if err := o.bindSessionID(fdSessionID, fdhkSessionID, route.Formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindSessionID binds and validates parameter SessionID from formData.
func (o *DisconnectParams) bindSessionID(rawData []string, hasKey bool, formats strfmt.Registry) error {
	if !hasKey {
		return errors.Required("sessionId", "formData", rawData)
	}
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true

	if err := validate.RequiredString("sessionId", "formData", raw); err != nil {
		return err
	}

	// Format: uuid4
	value, err := formats.Parse("uuid4", raw)
	if err != nil {
		return errors.InvalidType("sessionId", "formData", "strfmt.UUID4", raw)
	}
	o.SessionID = *(value.(*strfmt.UUID4))

	if err := o.validateSessionID(formats); err != nil {
		return err
	}

	return nil
}

// validateSessionID carries on validations for parameter SessionID
func (o *DisconnectParams) validateSessionID(formats strfmt.Registry) error {

	if err := validate.FormatOf("sessionId", "formData", "uuid4", o.SessionID.String(), formats); err != nil {
		return err
	}
	return nil
}
