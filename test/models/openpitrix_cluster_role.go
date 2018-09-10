// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// OpenpitrixClusterRole openpitrix cluster role
// swagger:model openpitrixClusterRole
type OpenpitrixClusterRole struct {

	// api version
	APIVersion string `json:"api_version,omitempty"`

	// cluster id
	ClusterID string `json:"cluster_id,omitempty"`

	// cpu
	CPU int64 `json:"cpu,omitempty"`

	// env
	Env string `json:"env,omitempty"`

	// file system
	FileSystem string `json:"file_system,omitempty"`

	// gpu
	Gpu int64 `json:"gpu,omitempty"`

	// instance size
	InstanceSize int64 `json:"instance_size,omitempty"`

	// memory
	Memory int64 `json:"memory,omitempty"`

	// mount options
	MountOptions string `json:"mount_options,omitempty"`

	// mount point
	MountPoint string `json:"mount_point,omitempty"`

	// ready replicas
	ReadyReplicas int64 `json:"ready_replicas,omitempty"`

	// replicas
	Replicas int64 `json:"replicas,omitempty"`

	// role
	Role string `json:"role,omitempty"`

	// storage size
	StorageSize int64 `json:"storage_size,omitempty"`
}

// Validate validates this openpitrix cluster role
func (m *OpenpitrixClusterRole) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *OpenpitrixClusterRole) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *OpenpitrixClusterRole) UnmarshalBinary(b []byte) error {
	var res OpenpitrixClusterRole
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
