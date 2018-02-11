// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// ContainerTopOKBody container top o k body
// swagger:model containerTopOKBody
type ContainerTopOKBody struct {

	// Each process running in the container, where each is process is an array of values corresponding to the titles
	Processes [][]string `json:"Processes"`

	// The ps column titles
	Titles []string `json:"Titles"`
}

// Validate validates this container top o k body
func (m *ContainerTopOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateProcesses(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateTitles(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ContainerTopOKBody) validateProcesses(formats strfmt.Registry) error {

	if swag.IsZero(m.Processes) { // not required
		return nil
	}

	return nil
}

func (m *ContainerTopOKBody) validateTitles(formats strfmt.Registry) error {

	if swag.IsZero(m.Titles) { // not required
		return nil
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ContainerTopOKBody) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ContainerTopOKBody) UnmarshalBinary(b []byte) error {
	var res ContainerTopOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}