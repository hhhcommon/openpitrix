// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// OpenpitrixValidateEmailServiceResponse openpitrix validate email service response
// swagger:model openpitrixValidateEmailServiceResponse
type OpenpitrixValidateEmailServiceResponse struct {

	// validate email service ok or not
	IsSucc bool `json:"is_succ,omitempty"`
}

// Validate validates this openpitrix validate email service response
func (m *OpenpitrixValidateEmailServiceResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *OpenpitrixValidateEmailServiceResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *OpenpitrixValidateEmailServiceResponse) UnmarshalBinary(b []byte) error {
	var res OpenpitrixValidateEmailServiceResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
