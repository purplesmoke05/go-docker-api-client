// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// Node node
// swagger:model Node
type Node struct {

	// Date and time at which the node was added to the swarm in
	// [RFC 3339](https://www.ietf.org/rfc/rfc3339.txt) format with nano-seconds.
	//
	CreatedAt string `json:"CreatedAt,omitempty"`

	// description
	Description *NodeDescription `json:"Description,omitempty"`

	// ID
	ID string `json:"ID,omitempty"`

	// manager status
	ManagerStatus *ManagerStatus `json:"ManagerStatus,omitempty"`

	// spec
	Spec *NodeSpec `json:"Spec,omitempty"`

	// status
	Status *NodeStatus `json:"Status,omitempty"`

	// Date and time at which the node was last updated in
	// [RFC 3339](https://www.ietf.org/rfc/rfc3339.txt) format with nano-seconds.
	//
	UpdatedAt string `json:"UpdatedAt,omitempty"`

	// version
	Version *ObjectVersion `json:"Version,omitempty"`
}

// Validate validates this node
func (m *Node) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDescription(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateManagerStatus(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateSpec(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateStatus(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateVersion(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Node) validateDescription(formats strfmt.Registry) error {

	if swag.IsZero(m.Description) { // not required
		return nil
	}

	if m.Description != nil {

		if err := m.Description.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("Description")
			}
			return err
		}

	}

	return nil
}

func (m *Node) validateManagerStatus(formats strfmt.Registry) error {

	if swag.IsZero(m.ManagerStatus) { // not required
		return nil
	}

	if m.ManagerStatus != nil {

		if err := m.ManagerStatus.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("ManagerStatus")
			}
			return err
		}

	}

	return nil
}

func (m *Node) validateSpec(formats strfmt.Registry) error {

	if swag.IsZero(m.Spec) { // not required
		return nil
	}

	if m.Spec != nil {

		if err := m.Spec.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("Spec")
			}
			return err
		}

	}

	return nil
}

func (m *Node) validateStatus(formats strfmt.Registry) error {

	if swag.IsZero(m.Status) { // not required
		return nil
	}

	if m.Status != nil {

		if err := m.Status.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("Status")
			}
			return err
		}

	}

	return nil
}

func (m *Node) validateVersion(formats strfmt.Registry) error {

	if swag.IsZero(m.Version) { // not required
		return nil
	}

	if m.Version != nil {

		if err := m.Version.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("Version")
			}
			return err
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *Node) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Node) UnmarshalBinary(b []byte) error {
	var res Node
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
