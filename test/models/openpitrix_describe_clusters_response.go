// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// OpenpitrixDescribeClustersResponse openpitrix describe clusters response
// swagger:model openpitrixDescribeClustersResponse
type OpenpitrixDescribeClustersResponse struct {

	// cluster set
	ClusterSet OpenpitrixDescribeClustersResponseClusterSet `json:"cluster_set"`

	// total count of qualified cluster
	TotalCount int64 `json:"total_count,omitempty"`
}

// Validate validates this openpitrix describe clusters response
func (m *OpenpitrixDescribeClustersResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *OpenpitrixDescribeClustersResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *OpenpitrixDescribeClustersResponse) UnmarshalBinary(b []byte) error {
	var res OpenpitrixDescribeClustersResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
