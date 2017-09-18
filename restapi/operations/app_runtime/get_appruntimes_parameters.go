// Code generated by go-swagger; DO NOT EDIT.

package app_runtime

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"

	strfmt "github.com/go-openapi/strfmt"
)

// NewGetAppruntimesParams creates a new GetAppruntimesParams object
// with the default values initialized.
func NewGetAppruntimesParams() GetAppruntimesParams {
	var (
		pageNumberDefault = int64(1)
		pageSizeDefault   = int32(20)
	)
	return GetAppruntimesParams{
		PageNumber: &pageNumberDefault,

		PageSize: &pageSizeDefault,
	}
}

// GetAppruntimesParams contains all the bound params for the get appruntimes operation
// typically these are obtained from a http.Request
//
// swagger:parameters GetAppruntimes
type GetAppruntimesParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request

	/*Page number
	  In: query
	  Default: 1
	*/
	PageNumber *int64
	/*Number of clusters returned
	  Minimum: > 0
	  Multiple Of: 10
	  In: query
	  Default: 20
	*/
	PageSize *int32
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls
func (o *GetAppruntimesParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error
	o.HTTPRequest = r

	qs := runtime.Values(r.URL.Query())

	qPageNumber, qhkPageNumber, _ := qs.GetOK("pageNumber")
	if err := o.bindPageNumber(qPageNumber, qhkPageNumber, route.Formats); err != nil {
		res = append(res, err)
	}

	qPageSize, qhkPageSize, _ := qs.GetOK("pageSize")
	if err := o.bindPageSize(qPageSize, qhkPageSize, route.Formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetAppruntimesParams) bindPageNumber(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}
	if raw == "" { // empty values pass all other validations
		var pageNumberDefault int64 = int64(1)
		o.PageNumber = &pageNumberDefault
		return nil
	}

	value, err := swag.ConvertInt64(raw)
	if err != nil {
		return errors.InvalidType("pageNumber", "query", "int64", raw)
	}
	o.PageNumber = &value

	return nil
}

func (o *GetAppruntimesParams) bindPageSize(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}
	if raw == "" { // empty values pass all other validations
		var pageSizeDefault int32 = int32(20)
		o.PageSize = &pageSizeDefault
		return nil
	}

	value, err := swag.ConvertInt32(raw)
	if err != nil {
		return errors.InvalidType("pageSize", "query", "int32", raw)
	}
	o.PageSize = &value

	if err := o.validatePageSize(formats); err != nil {
		return err
	}

	return nil
}

func (o *GetAppruntimesParams) validatePageSize(formats strfmt.Registry) error {

	if err := validate.MinimumInt("pageSize", "query", int64(*o.PageSize), 0, true); err != nil {
		return err
	}

	if err := validate.MultipleOf("pageSize", "query", float64(*o.PageSize), 10); err != nil {
		return err
	}

	return nil
}
