// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// TaskSpecPluginSpecPluginPrivilegeItems Describes a permission accepted by the user upon installing the plugin.
// swagger:model taskSpecPluginSpecPluginPrivilegeItems
type TaskSpecPluginSpecPluginPrivilegeItems struct {

	// description
	Description string `json:"Description,omitempty"`

	// name
	Name string `json:"Name,omitempty"`

	// value
	Value []string `json:"Value"`
}

// Validate validates this task spec plugin spec plugin privilege items
func (m *TaskSpecPluginSpecPluginPrivilegeItems) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateValue(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TaskSpecPluginSpecPluginPrivilegeItems) validateValue(formats strfmt.Registry) error {

	if swag.IsZero(m.Value) { // not required
		return nil
	}

	return nil
}

// MarshalBinary interface implementation
func (m *TaskSpecPluginSpecPluginPrivilegeItems) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TaskSpecPluginSpecPluginPrivilegeItems) UnmarshalBinary(b []byte) error {
	var res TaskSpecPluginSpecPluginPrivilegeItems
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
