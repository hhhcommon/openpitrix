// Code generated by go-swagger; DO NOT EDIT.

package cluster

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime/middleware"

	strfmt "github.com/go-openapi/strfmt"
)

// NewDeleteClustersClusterIDParams creates a new DeleteClustersClusterIDParams object
// with the default values initialized.
func NewDeleteClustersClusterIDParams() DeleteClustersClusterIDParams {
	var ()
	return DeleteClustersClusterIDParams{}
}

// DeleteClustersClusterIDParams contains all the bound params for the delete clusters cluster ID operation
// typically these are obtained from a http.Request
//
// swagger:parameters DeleteClustersClusterID
type DeleteClustersClusterIDParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request

	/*The cluster's id
	  Required: true
	  In: path
	*/
	ClusterID string
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls
func (o *DeleteClustersClusterIDParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error
	o.HTTPRequest = r

	rClusterID, rhkClusterID, _ := route.Params.GetOK("clusterId")
	if err := o.bindClusterID(rClusterID, rhkClusterID, route.Formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *DeleteClustersClusterIDParams) bindClusterID(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	o.ClusterID = raw

	return nil
}
