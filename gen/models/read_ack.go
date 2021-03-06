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

// ReadAck read ack
//
// swagger:model ReadAck
type ReadAck struct {

	// message Id
	// Required: true
	MessageID *string `json:"messageId"`

	// number
	// Required: true
	Number *string `json:"number"`

	// session Id
	// Required: true
	// Format: uuid4
	SessionID *strfmt.UUID4 `json:"sessionId"`
}

// Validate validates this read ack
func (m *ReadAck) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateMessageID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNumber(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSessionID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ReadAck) validateMessageID(formats strfmt.Registry) error {

	if err := validate.Required("messageId", "body", m.MessageID); err != nil {
		return err
	}

	return nil
}

func (m *ReadAck) validateNumber(formats strfmt.Registry) error {

	if err := validate.Required("number", "body", m.Number); err != nil {
		return err
	}

	return nil
}

func (m *ReadAck) validateSessionID(formats strfmt.Registry) error {

	if err := validate.Required("sessionId", "body", m.SessionID); err != nil {
		return err
	}

	if err := validate.FormatOf("sessionId", "body", "uuid4", m.SessionID.String(), formats); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ReadAck) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ReadAck) UnmarshalBinary(b []byte) error {
	var res ReadAck
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
