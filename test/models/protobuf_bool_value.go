// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// ProtobufBoolValue Wrapper message for `bool`.
//
// The JSON representation for `BoolValue` is JSON `true` and `false`.
// swagger:model protobufBoolValue
type ProtobufBoolValue struct {

	// The bool value.
	Value bool `json:"value,omitempty"`
}

// Validate validates this protobuf bool value
func (m *ProtobufBoolValue) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *ProtobufBoolValue) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ProtobufBoolValue) UnmarshalBinary(b []byte) error {
	var res ProtobufBoolValue
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}