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
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// NewGetChatsParams creates a new GetChatsParams object
// with the default values initialized.
func NewGetChatsParams() GetChatsParams {

	var (
		// initialize parameters with default values

		fromMeDefault = bool(false)
	)

	return GetChatsParams{
		FromMe: &fromMeDefault,
	}
}

// GetChatsParams contains all the bound params for the get chats operation
// typically these are obtained from a http.Request
//
// swagger:parameters getChats
type GetChatsParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*suppy this to load messages before this messageId
	  In: query
	*/
	BeforeMessageID *string
	/*fromMe needs to be supplied if beforeMessagId is given
	  In: query
	  Default: false
	*/
	FromMe *bool
	/*the number of messages in one query, max 300
	  Required: true
	  In: query
	*/
	NumberOfMessages int64
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
// To ensure default values, the struct must have been initialized with NewGetChatsParams() beforehand.
func (o *GetChatsParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	qs := runtime.Values(r.URL.Query())

	qBeforeMessageID, qhkBeforeMessageID, _ := qs.GetOK("beforeMessageId")
	if err := o.bindBeforeMessageID(qBeforeMessageID, qhkBeforeMessageID, route.Formats); err != nil {
		res = append(res, err)
	}

	qFromMe, qhkFromMe, _ := qs.GetOK("fromMe")
	if err := o.bindFromMe(qFromMe, qhkFromMe, route.Formats); err != nil {
		res = append(res, err)
	}

	qNumberOfMessages, qhkNumberOfMessages, _ := qs.GetOK("numberOfMessages")
	if err := o.bindNumberOfMessages(qNumberOfMessages, qhkNumberOfMessages, route.Formats); err != nil {
		res = append(res, err)
	}

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

// bindBeforeMessageID binds and validates parameter BeforeMessageID from query.
func (o *GetChatsParams) bindBeforeMessageID(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: false
	// AllowEmptyValue: false
	if raw == "" { // empty values pass all other validations
		return nil
	}

	o.BeforeMessageID = &raw

	return nil
}

// bindFromMe binds and validates parameter FromMe from query.
func (o *GetChatsParams) bindFromMe(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: false
	// AllowEmptyValue: false
	if raw == "" { // empty values pass all other validations
		// Default values have been previously initialized by NewGetChatsParams()
		return nil
	}

	value, err := swag.ConvertBool(raw)
	if err != nil {
		return errors.InvalidType("fromMe", "query", "bool", raw)
	}
	o.FromMe = &value

	return nil
}

// bindNumberOfMessages binds and validates parameter NumberOfMessages from query.
func (o *GetChatsParams) bindNumberOfMessages(rawData []string, hasKey bool, formats strfmt.Registry) error {
	if !hasKey {
		return errors.Required("numberOfMessages", "query", rawData)
	}
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// AllowEmptyValue: false
	if err := validate.RequiredString("numberOfMessages", "query", raw); err != nil {
		return err
	}

	value, err := swag.ConvertInt64(raw)
	if err != nil {
		return errors.InvalidType("numberOfMessages", "query", "int64", raw)
	}
	o.NumberOfMessages = value

	return nil
}

// bindPhoneNumber binds and validates parameter PhoneNumber from path.
func (o *GetChatsParams) bindPhoneNumber(rawData []string, hasKey bool, formats strfmt.Registry) error {
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
func (o *GetChatsParams) bindSessionID(rawData []string, hasKey bool, formats strfmt.Registry) error {
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
func (o *GetChatsParams) validateSessionID(formats strfmt.Registry) error {

	if err := validate.FormatOf("sessionId", "query", "uuid4", o.SessionID.String(), formats); err != nil {
		return err
	}
	return nil
}
