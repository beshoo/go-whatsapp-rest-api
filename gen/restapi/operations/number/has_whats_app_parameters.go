// Code generated by go-swagger; DO NOT EDIT.

package number

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

// NewHasWhatsAppParams creates a new HasWhatsAppParams object
// no default values defined in spec.
func NewHasWhatsAppParams() HasWhatsAppParams {

	return HasWhatsAppParams{}
}

// HasWhatsAppParams contains all the bound params for the has whats app operation
// typically these are obtained from a http.Request
//
// swagger:parameters hasWhatsApp
type HasWhatsAppParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*
	  Required: true
	  In: path
	*/
	PhoneNumber string
	/*Session Id used while succesfull scan
	  Required: true
	  In: query
	*/
	SessionID strfmt.UUID4
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewHasWhatsAppParams() beforehand.
func (o *HasWhatsAppParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	qs := runtime.Values(r.URL.Query())

	rPhoneNumber, rhkPhoneNumber, _ := route.Params.GetOK("phoneNumber")
	if err := o.bindPhoneNumber(rPhoneNumber, rhkPhoneNumber, route.Formats); err != nil {
		res = append(res, err)
	}

	qSessionID, qhkSessionID, _ := qs.GetOK("sessionId")
	if err := o.bindSessionID(qSessionID, qhkSessionID, route.Formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindPhoneNumber binds and validates parameter PhoneNumber from path.
func (o *HasWhatsAppParams) bindPhoneNumber(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// Parameter is provided by construction from the route

	o.PhoneNumber = raw

	return nil
}

// bindSessionID binds and validates parameter SessionID from query.
func (o *HasWhatsAppParams) bindSessionID(rawData []string, hasKey bool, formats strfmt.Registry) error {
	if !hasKey {
		return errors.Required("sessionId", "query", rawData)
	}
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// AllowEmptyValue: false
	if err := validate.RequiredString("sessionId", "query", raw); err != nil {
		return err
	}

	// Format: uuid4
	value, err := formats.Parse("uuid4", raw)
	if err != nil {
		return errors.InvalidType("sessionId", "query", "strfmt.UUID4", raw)
	}
	o.SessionID = *(value.(*strfmt.UUID4))

	if err := o.validateSessionID(formats); err != nil {
		return err
	}

	return nil
}

// validateSessionID carries on validations for parameter SessionID
func (o *HasWhatsAppParams) validateSessionID(formats strfmt.Registry) error {

	if err := validate.FormatOf("sessionId", "query", "uuid4", o.SessionID.String(), formats); err != nil {
		return err
	}
	return nil
}
