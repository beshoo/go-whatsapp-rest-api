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

// MessageAudio message audio
//
// swagger:model MessageAudio
type MessageAudio struct {

	// audio
	// Required: true
	Audio *string `json:"audio"`

	// audio length
	AudioLength string `json:"audioLength,omitempty"`

	// context info
	// Required: true
	ContextInfo *MessageContext `json:"contextInfo"`

	// message info
	// Required: true
	MessageInfo *MessageInfo `json:"messageInfo"`
}

// Validate validates this message audio
func (m *MessageAudio) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAudio(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateContextInfo(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMessageInfo(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *MessageAudio) validateAudio(formats strfmt.Registry) error {

	if err := validate.Required("audio", "body", m.Audio); err != nil {
		return err
	}

	return nil
}

func (m *MessageAudio) validateContextInfo(formats strfmt.Registry) error {

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

func (m *MessageAudio) validateMessageInfo(formats strfmt.Registry) error {

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

// MarshalBinary interface implementation
func (m *MessageAudio) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *MessageAudio) UnmarshalBinary(b []byte) error {
	var res MessageAudio
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}