// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// SwarmAllOf1 swarm all of1
// swagger:model swarmAllOf1
type SwarmAllOf1 struct {

	// join tokens
	JoinTokens *JoinTokens `json:"JoinTokens,omitempty"`
}

// Validate validates this swarm all of1
func (m *SwarmAllOf1) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateJoinTokens(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SwarmAllOf1) validateJoinTokens(formats strfmt.Registry) error {

	if swag.IsZero(m.JoinTokens) { // not required
		return nil
	}

	if m.JoinTokens != nil {

		if err := m.JoinTokens.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("JoinTokens")
			}
			return err
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *SwarmAllOf1) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SwarmAllOf1) UnmarshalBinary(b []byte) error {
	var res SwarmAllOf1
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}