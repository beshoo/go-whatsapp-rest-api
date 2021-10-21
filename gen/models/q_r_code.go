// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// QRCode q r code
//
// swagger:model QRCode
type QRCode struct {

	// base64
	Base64 string `json:"base64,omitempty"`
}

// Validate validates this q r code
func (m *QRCode) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *QRCode) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *QRCode) UnmarshalBinary(b []byte) error {
	var res QRCode
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}