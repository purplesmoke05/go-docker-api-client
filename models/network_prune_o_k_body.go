// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// NetworkPruneOKBody network prune o k body
// swagger:model networkPruneOKBody
type NetworkPruneOKBody struct {

	// Networks that were deleted
	NetworksDeleted []string `json:"NetworksDeleted"`
}

// Validate validates this network prune o k body
func (m *NetworkPruneOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateNetworksDeleted(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *NetworkPruneOKBody) validateNetworksDeleted(formats strfmt.Registry) error {

	if swag.IsZero(m.NetworksDeleted) { // not required
		return nil
	}

	return nil
}

// MarshalBinary interface implementation
func (m *NetworkPruneOKBody) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *NetworkPruneOKBody) UnmarshalBinary(b []byte) error {
	var res NetworkPruneOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
