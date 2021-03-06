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

// MessageImage message image
//
// swagger:model MessageImage
type MessageImage struct {

	// caption
	Caption string `json:"caption,omitempty"`

	// context info
	// Required: true
	ContextInfo *MessageContext `json:"contextInfo"`

	// image
	// Required: true
	Image *string `json:"image"`

	// message info
	// Required: true
	MessageInfo *MessageInfo `json:"messageInfo"`

	// thumbnail
	// Required: true
	Thumbnail *string `json:"thumbnail"`
}

// Validate validates this message image
func (m *MessageImage) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateContextInfo(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateImage(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMessageInfo(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateThumbnail(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *MessageImage) validateContextInfo(formats strfmt.Registry) error {

	if err := validate.Required("contextInfo", "body", m.ContextInfo); err != nil {
		return err
	}

	if m.ContextInfo != nil {
		if err := m.ContextInfo.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("contextInfo")
			}
			return err
		}
	}

	return nil
}

func (m *MessageImage) validateImage(formats strfmt.Registry) error {

	if err := validate.Required("image", "body", m.Image); err != nil {
		return err
	}

	return nil
}

func (m *MessageImage) validateMessageInfo(formats strfmt.Registry) error {

	if err := validate.Required("messageInfo", "body", m.MessageInfo); err != nil {
		return err
	}

	if m.MessageInfo != nil {
		if err := m.MessageInfo.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("messageInfo")
			}
			return err
		}
	}

	return nil
}

func (m *MessageImage) validateThumbnail(formats strfmt.Registry) error {

	if err := validate.Required("thumbnail", "body", m.Thumbnail); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *MessageImage) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *MessageImage) UnmarshalBinary(b []byte) error {
	var res MessageImage
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
