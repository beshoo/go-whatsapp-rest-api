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

// MessageItem message item
//
// swagger:model MessageItem
type MessageItem struct {

	// from me
	// Required: true
	FromMe *bool `json:"fromMe"`

	// id
	// Required: true
	ID *string `json:"id"`
}

// Validate validates this message item
func (m *MessageItem) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateFromMe(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *MessageItem) validateFromMe(formats strfmt.Registry) error {

	if err := validate.Required("fromMe", "body", m.FromMe); err != nil {
		return err
	}

	return nil
}

func (m *MessageItem) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *MessageItem) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *MessageItem) UnmarshalBinary(b []byte) error {
	var res MessageItem
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
