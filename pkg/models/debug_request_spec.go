// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// DebugRequestSpec debug request spec
// swagger:model DebugRequestSpec
type DebugRequestSpec struct {

	// debugger
	// Required: true
	Debugger *string `json:"debugger"`

	// image
	// Required: true
	Image *string `json:"image"`
}

// Validate validates this debug request spec
func (m *DebugRequestSpec) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDebugger(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateImage(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DebugRequestSpec) validateDebugger(formats strfmt.Registry) error {

	if err := validate.Required("debugger", "body", m.Debugger); err != nil {
		return err
	}

	return nil
}

func (m *DebugRequestSpec) validateImage(formats strfmt.Registry) error {

	if err := validate.Required("image", "body", m.Image); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *DebugRequestSpec) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DebugRequestSpec) UnmarshalBinary(b []byte) error {
	var res DebugRequestSpec
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
