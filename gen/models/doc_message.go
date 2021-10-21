// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// DocMessage doc message
//
// swagger:model DocMessage
type DocMessage struct {

	// Fully qualified url
	// Required: true
	// Format: uri
	Doc *strfmt.URI `json:"doc"`

	// doc type
	// Required: true
	// Enum: [DOC DOCX CSV XLS XLSX PDF PPT PPTX GZ ZIP 7z TEXT]
	DocType *string `json:"docType"`

	// number reply ids
	NumberReplyIds []*NumberReplyIds `json:"numberReplyIds"`

	// session Id
	// Required: true
	// Format: uuid4
	SessionID *strfmt.UUID4 `json:"sessionId"`

	// title
	// Required: true
	Title *string `json:"title"`
}

// Validate validates this doc message
func (m *DocMessage) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDoc(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDocType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNumberReplyIds(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSessionID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTitle(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DocMessage) validateDoc(formats strfmt.Registry) error {

	if err := validate.Required("doc", "body", m.Doc); err != nil {
		return err
	}

	if err := validate.FormatOf("doc", "body", "uri", m.Doc.String(), formats); err != nil {
		return err
	}

	return nil
}

var docMessageTypeDocTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["DOC","DOCX","CSV","XLS","XLSX","PDF","PPT","PPTX","GZ","ZIP","7z","TEXT"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		docMessageTypeDocTypePropEnum = append(docMessageTypeDocTypePropEnum, v)
	}
}

const (

	// DocMessageDocTypeDOC captures enum value "DOC"
	DocMessageDocTypeDOC string = "DOC"

	// DocMessageDocTypeDOCX captures enum value "DOCX"
	DocMessageDocTypeDOCX string = "DOCX"

	// DocMessageDocTypeCSV captures enum value "CSV"
	DocMessageDocTypeCSV string = "CSV"

	// DocMessageDocTypeXLS captures enum value "XLS"
	DocMessageDocTypeXLS string = "XLS"

	// DocMessageDocTypeXLSX captures enum value "XLSX"
	DocMessageDocTypeXLSX string = "XLSX"

	// DocMessageDocTypePDF captures enum value "PDF"
	DocMessageDocTypePDF string = "PDF"

	// DocMessageDocTypePPT captures enum value "PPT"
	DocMessageDocTypePPT string = "PPT"

	// DocMessageDocTypePPTX captures enum value "PPTX"
	DocMessageDocTypePPTX string = "PPTX"

	// DocMessageDocTypeGZ captures enum value "GZ"
	DocMessageDocTypeGZ string = "GZ"

	// DocMessageDocTypeZIP captures enum value "ZIP"
	DocMessageDocTypeZIP string = "ZIP"

	// DocMessageDocTypeNr7z captures enum value "7z"
	DocMessageDocTypeNr7z string = "7z"

	// DocMessageDocTypeTEXT captures enum value "TEXT"
	DocMessageDocTypeTEXT string = "TEXT"
)

// prop value enum
func (m *DocMessage) validateDocTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, docMessageTypeDocTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *DocMessage) validateDocType(formats strfmt.Registry) error {

	if err := validate.Required("docType", "body", m.DocType); err != nil {
		return err
	}

	// value enum
	if err := m.validateDocTypeEnum("docType", "body", *m.DocType); err != nil {
		return err
	}

	return nil
}

func (m *DocMessage) validateNumberReplyIds(formats strfmt.Registry) error {

	if swag.IsZero(m.NumberReplyIds) { // not required
		return nil
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

func (m *DocMessage) validateSessionID(formats strfmt.Registry) error {

	if err := validate.Required("sessionId", "body", m.SessionID); err != nil {
		return err
	}

	if err := validate.FormatOf("sessionId", "body", "uuid4", m.SessionID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *DocMessage) validateTitle(formats strfmt.Registry) error {

	if err := validate.Required("title", "body", m.Title); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *DocMessage) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DocMessage) UnmarshalBinary(b []byte) error {
	var res DocMessage
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}