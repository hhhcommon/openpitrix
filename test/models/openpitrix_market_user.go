// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// OpenpitrixMarketUser openpitrix market user
// swagger:model openpitrixMarketUser
type OpenpitrixMarketUser struct {

	// create time
	CreateTime strfmt.DateTime `json:"create_time,omitempty"`

	// market id
	MarketID string `json:"market_id,omitempty"`

	// owner path
	OwnerPath string `json:"owner_path,omitempty"`

	// user id
	UserID string `json:"user_id,omitempty"`
}

// Validate validates this openpitrix market user
func (m *OpenpitrixMarketUser) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *OpenpitrixMarketUser) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *OpenpitrixMarketUser) UnmarshalBinary(b []byte) error {
	var res OpenpitrixMarketUser
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
