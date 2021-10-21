// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// NotifyLogout notify logout
//
// swagger:model NotifyLogout
type NotifyLogout struct {

	// number
	// Required: true
	Number *string `json:"number"`

	// session Id
	// Required: true
	// Format: uuid4
	SessionID *strfmt.UUID4 `json:"sessionId"`

	// timestamp
	// Required: true
	// Format: date-time
	Timestamp *strfmt.DateTime `json:"timestamp"`
}

// Validate validates this notify logout
func (m *NotifyLogout) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateNumber(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSessionID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTimestamp(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *NotifyLogout) validateNumber(formats strfmt.Registry) error {

	if err := validate.Required("number", "body", m.Number); err != nil {
		return err
	}

	return nil
}

func (m *NotifyLogout) validateSessionID(formats strfmt.Registry) error {

	if err := validate.Required("sessionId", "body", m.SessionID); err != nil {
		return err
	}

	if err := validate.FormatOf("sessionId", "body", "uuid4", m.SessionID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *NotifyLogout) validateTimestamp(formats strfmt.Registry) error {

	if err := validate.Required("timestamp", "body", m.Timestamp); err != nil {
		return err
	}

	if err := validate.FormatOf("timestamp", "body", "date-time", m.Timestamp.String(), formats); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *NotifyLogout) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *NotifyLogout) UnmarshalBinary(b []byte) error {
	var res NotifyLogout
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}