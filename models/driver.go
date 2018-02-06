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

// Driver Driver represents a driver (network, logging, secrets).
// swagger:model Driver
type Driver struct {

	// Name of the driver.
	// Required: true
	Name string `json:"Name"`

	// Key/value map of driver-specific options.
	Options map[string]string `json:"Options,omitempty"`
}

// Validate validates this driver
func (m *Driver) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateName(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Driver) validateName(formats strfmt.Registry) error {

	if err := validate.RequiredString("Name", "body", string(m.Name)); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Driver) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Driver) UnmarshalBinary(b []byte) error {
	var res Driver
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
