// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// ImageMessage image message
//
// swagger:model ImageMessage
type ImageMessage struct {

	// caption
	Caption string `json:"caption,omitempty"`

	// Fully qualified url
	// Required: true
	// Format: uri
	Image *strfmt.URI `json:"image"`

	// number reply ids
	// Required: true
	NumberReplyIds []*NumberReplyIds `json:"numberReplyIds"`

	// session Id
	// Required: true
	// Format: uuid4
	SessionID *strfmt.UUID4 `json:"sessionId"`
}

// Validate validates this image message
func (m *ImageMessage) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateImage(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNumberReplyIds(formats); err != nil {
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

func (m *ImageMessage) validateImage(formats strfmt.Registry) error {

	if err := validate.Required("image", "body", m.Image); err != nil {
		return err
	}

	if err := validate.FormatOf("image", "body", "uri", m.Image.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *ImageMessage) validateNumberReplyIds(formats strfmt.Registry) error {

	if err := validate.Required("numberReplyIds", "body", m.NumberReplyIds); err != nil {
		return err
	}

	for i := 0; i < len(m.NumberReplyIds); i++ {
		if swag.IsZero(m.NumberReplyIds[i]) { // not required
			continue
		}

		if m.NumberReplyIds[i] != nil {
			if err := m.NumberReplyIds[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("numberReplyIds" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ImageMessage) validateSessionID(formats strfmt.Registry) error {

	if err := validate.Required("sessionId", "body", m.SessionID); err != nil {
		return err
	}

	if err := validate.FormatOf("sessionId", "body", "uuid4", m.SessionID.String(), formats); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ImageMessage) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ImageMessage) UnmarshalBinary(b []byte) error {
	var res ImageMessage
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
