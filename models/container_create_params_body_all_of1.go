// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// ContainerCreateParamsBodyAllOf1 container create params body all of1
// swagger:model containerCreateParamsBodyAllOf1
type ContainerCreateParamsBodyAllOf1 struct {

	// host config
	HostConfig *HostConfig `json:"HostConfig,omitempty"`

	// networking config
	NetworkingConfig *ContainerCreateParamsBodyAllOf1NetworkingConfig `json:"NetworkingConfig,omitempty"`
}

// Validate validates this container create params body all of1
func (m *ContainerCreateParamsBodyAllOf1) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateNetworkingConfig(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ContainerCreateParamsBodyAllOf1) validateNetworkingConfig(formats strfmt.Registry) error {

	if swag.IsZero(m.NetworkingConfig) { // not required
		return nil
	}

	if m.NetworkingConfig != nil {

		if err := m.NetworkingConfig.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("NetworkingConfig")
			}
			return err
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *ContainerCreateParamsBodyAllOf1) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ContainerCreateParamsBodyAllOf1) UnmarshalBinary(b []byte) error {
	var res ContainerCreateParamsBodyAllOf1
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}