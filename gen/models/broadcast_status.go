// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// BroadcastStatus broadcast status
//
// swagger:model BroadcastStatus
type BroadcastStatus struct {

	// broadcast Id
	BroadcastID string `json:"broadcastId,omitempty"`

	// status
	// Enum: [processing sent]
	Status string `json:"status,omitempty"`
}

// Validate validates this broadcast status
func (m *BroadcastStatus) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var broadcastStatusTypeStatusPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["processing","sent"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		broadcastStatusTypeStatusPropEnum = append(broadcastStatusTypeStatusPropEnum, v)
	}
}

const (

	// BroadcastStatusStatusProcessing captures enum value "processing"
	BroadcastStatusStatusProcessing string = "processing"

	// BroadcastStatusStatusSent captures enum value "sent"
	BroadcastStatusStatusSent string = "sent"
)

// prop value enum
func (m *BroadcastStatus) validateStatusEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, broadcastStatusTypeStatusPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *BroadcastStatus) validateStatus(formats strfmt.Registry) error {

	if swag.IsZero(m.Status) { // not required
		return nil
	}

	// value enum
	if err := m.validateStatusEnum("status", "body", m.Status); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *BroadcastStatus) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *BroadcastStatus) UnmarshalBinary(b []byte) error {
	var res BroadcastStatus
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
