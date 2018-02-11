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

// PluginConfig The config of a plugin.
// swagger:model pluginConfig
type PluginConfig struct {

	// args
	// Required: true
	Args PluginConfigArgs `json:"Args"`

	// description
	// Required: true
	Description string `json:"Description"`

	// Docker Version used to create the plugin
	DockerVersion string `json:"DockerVersion,omitempty"`

	// documentation
	// Required: true
	Documentation string `json:"Documentation"`

	// entrypoint
	// Required: true
	Entrypoint []string `json:"Entrypoint"`

	// env
	// Required: true
	Env PluginConfigEnv `json:"Env"`

	// interface
	// Required: true
	Interface PluginConfigInterface `json:"Interface"`

	// ipc host
	// Required: true
	IpcHost bool `json:"IpcHost"`

	// linux
	// Required: true
	Linux PluginConfigLinux `json:"Linux"`

	// mounts
	// Required: true
	Mounts PluginConfigMounts `json:"Mounts"`

	// network
	// Required: true
	Network PluginConfigNetwork `json:"Network"`

	// pid host
	// Required: true
	PidHost bool `json:"PidHost"`

	// propagated mount
	// Required: true
	PropagatedMount string `json:"PropagatedMount"`

	// user
	User PluginConfigUser `json:"User,omitempty"`

	// work dir
	// Required: true
	WorkDir string `json:"WorkDir"`

	// rootfs
	Rootfs *PluginConfigRootfs `json:"rootfs,omitempty"`
}

// Validate validates this plugin config
func (m *PluginConfig) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateArgs(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateDescription(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateDocumentation(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateEntrypoint(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateEnv(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateInterface(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateIpcHost(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateLinux(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateMounts(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateNetwork(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validatePidHost(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validatePropagatedMount(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateUser(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateWorkDir(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateRootfs(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PluginConfig) validateArgs(formats strfmt.Registry) error {

	if err := m.Args.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("Args")
		}
		return err
	}

	return nil
}

func (m *PluginConfig) validateDescription(formats strfmt.Registry) error {

	if err := validate.RequiredString("Description", "body", string(m.Description)); err != nil {
		return err
	}

	return nil
}

func (m *PluginConfig) validateDocumentation(formats strfmt.Registry) error {

	if err := validate.RequiredString("Documentation", "body", string(m.Documentation)); err != nil {
		return err
	}

	return nil
}

func (m *PluginConfig) validateEntrypoint(formats strfmt.Registry) error {

	if err := validate.Required("Entrypoint", "body", m.Entrypoint); err != nil {
		return err
	}

	return nil
}

func (m *PluginConfig) validateEnv(formats strfmt.Registry) error {

	if err := validate.Required("Env", "body", m.Env); err != nil {
		return err
	}

	if err := m.Env.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("Env")
		}
		return err
	}

	return nil
}

func (m *PluginConfig) validateInterface(formats strfmt.Registry) error {

	if err := m.Interface.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("Interface")
		}
		return err
	}

	return nil
}

func (m *PluginConfig) validateIpcHost(formats strfmt.Registry) error {

	if err := validate.Required("IpcHost", "body", bool(m.IpcHost)); err != nil {
		return err
	}

	return nil
}

func (m *PluginConfig) validateLinux(formats strfmt.Registry) error {

	if err := m.Linux.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("Linux")
		}
		return err
	}

	return nil
}

func (m *PluginConfig) validateMounts(formats strfmt.Registry) error {

	if err := validate.Required("Mounts", "body", m.Mounts); err != nil {
		return err
	}

	if err := m.Mounts.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("Mounts")
		}
		return err
	}

	return nil
}

func (m *PluginConfig) validateNetwork(formats strfmt.Registry) error {

	if err := m.Network.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("Network")
		}
		return err
	}

	return nil
}

func (m *PluginConfig) validatePidHost(formats strfmt.Registry) error {

	if err := validate.Required("PidHost", "body", bool(m.PidHost)); err != nil {
		return err
	}

	return nil
}

func (m *PluginConfig) validatePropagatedMount(formats strfmt.Registry) error {

	if err := validate.RequiredString("PropagatedMount", "body", string(m.PropagatedMount)); err != nil {
		return err
	}

	return nil
}

func (m *PluginConfig) validateUser(formats strfmt.Registry) error {

	if swag.IsZero(m.User) { // not required
		return nil
	}

	if err := m.User.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("User")
		}
		return err
	}

	return nil
}

func (m *PluginConfig) validateWorkDir(formats strfmt.Registry) error {

	if err := validate.RequiredString("WorkDir", "body", string(m.WorkDir)); err != nil {
		return err
	}

	return nil
}

func (m *PluginConfig) validateRootfs(formats strfmt.Registry) error {

	if swag.IsZero(m.Rootfs) { // not required
		return nil
	}

	if m.Rootfs != nil {

		if err := m.Rootfs.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("rootfs")
			}
			return err
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *PluginConfig) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PluginConfig) UnmarshalBinary(b []byte) error {
	var res PluginConfig
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}