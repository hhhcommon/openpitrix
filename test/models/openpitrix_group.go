// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// OpenpitrixGroup openpitrix group
// swagger:model openpitrixGroup
type OpenpitrixGroup struct {

	// the time when user create
	CreateTime strfmt.DateTime `json:"create_time,omitempty"`

	// group description
	Description string `json:"description,omitempty"`

	// group id
	GroupID string `json:"group_id,omitempty"`

	// group path, a concat string gid-xxx.gid-xxx.gid...
	GroupPath string `json:"group_path,omitempty"`

	// group name
	Name string `json:"name,omitempty"`

	// parent group id
	ParentGroupID string `json:"parent_group_id,omitempty"`

	// group status eg.[active|deleted]
	Status string `json:"status,omitempty"`

	// record group status changed time
	StatusTime strfmt.DateTime `json:"status_time,omitempty"`

	// the time when group update
	UpdateTime strfmt.DateTime `json:"update_time,omitempty"`
}

// Validate validates this openpitrix group
func (m *OpenpitrixGroup) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *OpenpitrixGroup) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *OpenpitrixGroup) UnmarshalBinary(b []byte) error {
	var res OpenpitrixGroup
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
